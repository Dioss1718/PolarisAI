from sentence_transformers import SentenceTransformer
import chromadb
from config import EMBED_MODEL, CHROMA_DB_DIR, TOP_K

model = SentenceTransformer(EMBED_MODEL)
client = chromadb.PersistentClient(path=CHROMA_DB_DIR)
collection = client.get_or_create_collection(name="cloud_docs")


def infer_categories(node_type: str, action: str):
    categories = {"architecture", "compliance", "security"}

    node_type = (node_type or "").upper()
    action = (action or "").upper()

    if node_type in {
        "DATABASE", "COMPUTE", "OBJECT_STORAGE", "LOAD_BALANCER", "IAM_ROLE", "MULTI_SERVICE",
    }:
        categories.add("sla")

    if "SECURE" in action or "RESTRICT" in action or "PATCH" in action:
        categories.add("security")

    if "TERMINATE" in action or "DOWNSIZE" in action:
        categories.add("architecture")
        categories.add("sla")

    return list(categories)


def infer_sla_sources(node_type: str, action: str, query: str):
    q = (query or "").lower()
    node_type = (node_type or "").upper()

    sources = set()

    if "aws" in q:
        if node_type == "COMPUTE":
            sources.add("aws_ec2.txt")
        elif node_type == "DATABASE":
            sources.add("aws_rds.txt")
        elif node_type == "OBJECT_STORAGE":
            sources.add("aws_s3.txt")

    if "azure" in q:
        sources.add("azure_global.txt")

    if "gcp" in q or "google" in q:
        if node_type == "COMPUTE":
            sources.add("gcp_compute_engine.txt")
        elif node_type == "DATABASE":
            sources.add("gcp_cloud_sql.txt")
        elif node_type == "OBJECT_STORAGE":
            sources.add("gcp_cloud_storage.txt")

    return list(sources)


def build_target_query(query: str, node_type: str, action: str):
    return f"""
    cloud infrastructure decision:
    action={action}
    node_type={node_type}
    query={query}
    need SLA impact, security implications, compliance implications, and recommended remediation
    """.strip()


def retrieve(query: str, node_type: str, action: str, top_k: int = TOP_K):
    target_query = build_target_query(query, node_type, action)
    query_embedding = model.encode(target_query).tolist()

    categories = infer_categories(node_type, action)
    preferred_sources = infer_sla_sources(node_type, action, query)

    docs = []
    metas = []

    # First pass: targeted source retrieval if available
    if preferred_sources:
        results = collection.query(
            query_embeddings=[query_embedding],
            n_results=top_k,
            where={
                "$and": [
                    {"category": {"$in": categories}},
                    {"source": {"$in": preferred_sources}},
                ]
            },
        )
        docs = results.get("documents", [[]])[0]
        metas = results.get("metadatas", [[]])[0]

    # Fallback pass: broader category retrieval
    if len(docs) < top_k:
        results = collection.query(
            query_embeddings=[query_embedding],
            n_results=top_k,
            where={"category": {"$in": categories}},
        )
        docs = results.get("documents", [[]])[0]
        metas = results.get("metadatas", [[]])[0]

    return docs, metas