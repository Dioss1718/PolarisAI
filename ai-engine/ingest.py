import os
from sentence_transformers import SentenceTransformer
import chromadb
from chromadb.config import Settings
from config import DATA_PATH, CHROMA_DB_DIR, EMBED_MODEL

# ✅ Persistent DB (CRITICAL FIX)
client = chromadb.PersistentClient(path=CHROMA_DB_DIR)

collection = client.get_or_create_collection(name="cloud_docs")

model = SentenceTransformer(EMBED_MODEL)


def load_documents():
    docs = []

    for root, _, files in os.walk(DATA_PATH):
        for file in files:
            if file.endswith(".txt"):
                path = os.path.join(root, file)

                with open(path, "r", encoding="utf-8") as f:
                    text = f.read()

                docs.append({
                    "content": text,
                    "source": file,
                    "category": os.path.basename(root)
                })

    return docs


def ingest():
    docs = load_documents()

    print(f"📥 Ingesting {len(docs)} documents...")

    for i, doc in enumerate(docs):
        embedding = model.encode(doc["content"]).tolist()

        collection.add(
            documents=[doc["content"]],
            embeddings=[embedding],
            metadatas=[{
                "source": doc["source"],
                "category": doc["category"]
            }],
            ids=[str(i)]
        )

    print("✅ Documents ingested into persistent vector DB")


if __name__ == "__main__":
    ingest()