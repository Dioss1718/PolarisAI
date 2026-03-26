import os
from groq import Groq
from config import LLM_MODEL

client = Groq(api_key=os.getenv("GROQ_API_KEY"))


def call_llm(prompt: str) -> str:
    try:
        completion = client.chat.completions.create(
            model=LLM_MODEL,
            messages=[
                {
                    "role": "system",
                    "content": "You are a principal cloud architect AI. Stay grounded, precise, and concise."
                },
                {"role": "user", "content": prompt}
            ],
            temperature=0.2,
            max_tokens=800
        )

        output = completion.choices[0].message.content
        if not output or len(output.strip()) < 10:
            raise ValueError("Empty LLM response")
        return output

    except Exception as e:
        raise RuntimeError(f"LLM failure: {str(e)}")