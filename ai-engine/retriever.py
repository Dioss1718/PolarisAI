from sentence_transformers import SentenceTransformer
import chromadb
from config import EMBED_MODEL, CHROMA_DB_DIR

model = SentenceTransformer(EMBED_MODEL)

# ✅ SAME as ingest (CRITICAL)
client = chromadb.PersistentClient(path=CHROMA_DB_DIR)
collection = client.get_or_create_collection(name="cloud_docs")


def infer_categories(node_type: str, action: str):
    categories = ["architecture"]

    node_type = node_type.upper()

    if node_type in {"DATABASE", "COMPUTE", "OBJECT_STORAGE", "LOAD_BALANCER", "IAM_ROLE"}:
        categories.append("sla")

    categories.append("security")
    categories.append("compliance")

    return list(set(categories))


def retrieve(query, node_type, action, top_k=4):
    query_embedding = model.encode(query).tolist()

    categories = infer_categories(node_type, action)

    results = collection.query(
        query_embeddings=[query_embedding],
        n_results=top_k,
        where={"category": {"$in": categories}}
    )

    docs = results.get("documents", [[]])[0]
    metas = results.get("metadatas", [[]])[0]

    return docs, metas