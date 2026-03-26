from groq import Groq
import os
from config import LLM_MODEL

# 🔥 Production-grade client
client = Groq(api_key=os.getenv("GROQ_API_KEY"))

def call_llm(prompt: str) -> str:
    try:
        completion = client.chat.completions.create(
            model=LLM_MODEL,
            messages=[
                {
                    "role": "system",
                    "content": "You are a PRINCIPAL cloud architect AI. Be precise, grounded, and non-hallucinating."
                },
                {"role": "user", "content": prompt}
            ],
            temperature=0.2,  # low hallucination
            max_tokens=800
        )

        output = completion.choices[0].message.content

        if not output or len(output.strip()) < 10:
            raise ValueError("Empty LLM response")

        return output

    except Exception as e:
        raise RuntimeError(f"LLM failure: {str(e)}")