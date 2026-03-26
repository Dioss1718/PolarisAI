import os
from groq import Groq
from dotenv import load_dotenv
from config import LLM_MODEL

# Load environment variables from .env if present
load_dotenv()

_client = None


def get_client() -> Groq:
    global _client

    if _client is not None:
        return _client

    api_key = os.getenv("GROQ_API_KEY")
    if not api_key:
        raise RuntimeError(
            "GROQ_API_KEY is missing. Set it in your environment or .env file."
        )

    _client = Groq(api_key=api_key)
    return _client


def call_llm(prompt: str) -> str:
    try:
        client = get_client()

        completion = client.chat.completions.create(
            model=LLM_MODEL,
            messages=[
                {
                    "role": "system",
                    "content": (
                        "You are a principal cloud architect AI. "
                        "Stay grounded, precise, concise, and production-oriented."
                    ),
                },
                {"role": "user", "content": prompt},
            ],
            temperature=0.2,
            max_tokens=800,
        )

        output = completion.choices[0].message.content
        if not output or len(output.strip()) < 10:
            raise ValueError("Empty LLM response")

        return output.strip()

    except Exception as e:
        raise RuntimeError(f"LLM failure: {str(e)}")