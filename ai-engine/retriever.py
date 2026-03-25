import chromadb
from chromadb.config import Settings
from sentence_transformers import SentenceTransformer
from config import EMBED_MODEL, CHROMA_DB_DIR

model = SentenceTransformer(EMBED_MODEL)

client = chromadb.Client(Settings(persist_directory=CHROMA_DB_DIR))
collection = client.get_or_create_collection(name="cloud_docs")


def retrieve(query, node_type):
    query_embedding = model.encode(query).tolist()

    results = collection.query(
        query_embeddings=[query_embedding],
        n_results=3
    )

    return results.get("documents", [[]])[0]