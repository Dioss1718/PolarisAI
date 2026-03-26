import os
import hashlib
from sentence_transformers import SentenceTransformer
import chromadb
from config import DATA_PATH, CHROMA_DB_DIR, EMBED_MODEL

client = chromadb.PersistentClient(path=CHROMA_DB_DIR)
collection = client.get_or_create_collection(name="cloud_docs")
model = SentenceTransformer(EMBED_MODEL)


def chunk_text(text: str, max_chars: int = 1200):
    text = text.strip()
    if len(text) <= max_chars:
        return [text]

    chunks = []
    start = 0
    while start < len(text):
        end = min(start + max_chars, len(text))
        chunks.append(text[start:end].strip())
        start = end
    return [c for c in chunks if c]


def load_documents():
    docs = []

    for root, _, files in os.walk(DATA_PATH):
        for file in files:
            if not file.endswith(".txt"):
                continue

            path = os.path.join(root, file)
            category = os.path.basename(root)

            with open(path, "r", encoding="utf-8") as f:
                text = f.read()

            chunks = chunk_text(text)
            for idx, chunk in enumerate(chunks):
                docs.append({
                    "content": chunk,
                    "source": file,
                    "category": category,
                    "chunk_id": idx
                })

    return docs


def make_id(source: str, chunk_id: int, content: str) -> str:
    digest = hashlib.md5(content.encode("utf-8")).hexdigest()
    return f"{source}-{chunk_id}-{digest}"


def ingest():
    docs = load_documents()
    print(f"Ingesting {len(docs)} chunks...")

    existing_ids = set()
    try:
        existing = collection.get()
        existing_ids = set(existing.get("ids", []))
    except Exception:
        pass

    added = 0

    for doc in docs:
        doc_id = make_id(doc["source"], doc["chunk_id"], doc["content"])
        if doc_id in existing_ids:
            continue

        embedding = model.encode(doc["content"]).tolist()

        collection.add(
            documents=[doc["content"]],
            embeddings=[embedding],
            metadatas=[{
                "source": doc["source"],
                "category": doc["category"],
                "chunk_id": doc["chunk_id"],
            }],
            ids=[doc_id]
        )
        added += 1

    print(f"Added {added} new chunks to ChromaDB")


if __name__ == "__main__":
    ingest()