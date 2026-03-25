from groq import Groq
import os

client = Groq(api_key=os.getenv("GROQ_API_KEY"))

def call_llm(prompt: str) -> str:
    try:
        completion = client.chat.completions.create(
            model="llama-3.1-8b-instant",  # stable + production ready
            messages=[
                {"role": "system", "content": "You are a senior cloud architect AI."},
                {"role": "user", "content": prompt}
            ],
            temperature=0.2,
            max_tokens=800
        )

        return completion.choices[0].message.content

    except Exception as e:
        raise RuntimeError(f"LLM failure: {str(e)}")