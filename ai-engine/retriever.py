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

    if node_type in {"DATABASE", "COMPUTE", "OBJECT_STORAGE", "LOAD_BALANCER", "IAM_ROLE", "MULTI_SERVICE"}:
        categories.add("sla")

    if "SECURE" in action or "RESTRICT" in action:
        categories.add("security")

    if "TERMINATE" in action or "DOWNSIZE" in action:
        categories.add("architecture")
        categories.add("sla")

    return list(categories)


def retrieve(query: str, node_type: str, action: str, top_k: int = TOP_K):
    query_embedding = model.encode(query).tolist()
    categories = infer_categories(node_type, action)

    results = collection.query(
        query_embeddings=[query_embedding],
        n_results=top_k,
        where={"category": {"$in": categories}},
    )

    docs = results.get("documents", [[]])[0]
    metas = results.get("metadatas", [[]])[0]
    return docs, metas