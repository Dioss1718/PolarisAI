import { useMemo, useState } from "react";
import { askCopilot } from "../api/client";

export default function CopilotDrawer({ open, onClose, state, onSelectNode }) {
  const [query, setQuery] = useState("");
  const [answer, setAnswer] = useState(
    "Ask about risk, bill shock, approved actions, attack paths, or a specific node."
  );
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const normalizedNodes = useMemo(() => state?.nodes || [], [state?.nodes]);

  const trySelectMatchedNode = (rawQuery) => {
    const q = String(rawQuery || "").trim().toLowerCase();
    if (!q) return;

    const matchedNode = normalizedNodes.find(
      (n) =>
        q.includes(String(n.id || "").toLowerCase()) ||
        q.includes(String(n.label || "").toLowerCase())
    );

    if (matchedNode) {
      onSelectNode?.(matchedNode.id);
    }
  };

  const handleSubmit = async () => {
    const trimmed = query.trim();
    if (!trimmed) return;

    setLoading(true);
    setError("");

    try {
      trySelectMatchedNode(trimmed);

      const data = await askCopilot({
        query: trimmed,
        scenario: state?.scenario || "FULL_CHAOS",
        seed: state?.seed || 42,
      });

      setAnswer(
        data?.answer ||
          "Copilot could not generate a response from the current pipeline state."
      );
    } catch (err) {
      const backendError =
        err?.response?.data?.error ||
        err?.message ||
        "Copilot request failed";

      setError(backendError);
      setAnswer(
        "Copilot is unavailable right now. Please verify the backend and AI engine are running."
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <div
      className={`fixed right-0 top-0 z-[90] h-screen w-[420px] transform border-l border-white/10 bg-slate-950/96 p-4 backdrop-blur-xl transition-transform duration-300 ${
        open ? "translate-x-0" : "translate-x-full"
      }`}
    >
      <div className="flex items-center justify-between">
        <div>
          <div className="text-lg font-semibold">Copilot</div>
          <div className="text-sm text-slate-400">
            Backed by the AI engine and grounded on the current pipeline state.
          </div>
        </div>
        <button
          onClick={onClose}
          className="rounded-lg border border-white/10 px-3 py-1 text-sm"
        >
          Close
        </button>
      </div>

      <div className="mt-4 flex items-center gap-2">
        <input
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter" && !loading) {
              handleSubmit();
            }
          }}
          placeholder="Ask about risk, bill shock, approved actions..."
          className="flex-1 rounded-xl border border-white/10 bg-slate-900 px-4 py-3 outline-none"
        />
        <button
          onClick={handleSubmit}
          disabled={!query.trim() || loading}
          className="rounded-xl bg-sky-500 px-4 py-3 font-medium text-slate-950 disabled:opacity-50"
        >
          {loading ? "..." : "Send"}
        </button>
      </div>

      {error ? (
        <div className="mt-4 rounded-xl border border-rose-500/30 bg-rose-500/10 p-3 text-sm text-rose-200">
          {error}
        </div>
      ) : null}

      <div className="mt-4 rounded-xl bg-slate-900 p-4 text-sm whitespace-pre-wrap text-slate-200">
        {loading ? "Copilot is reasoning over the current governance state..." : answer}
      </div>
    </div>
  );
}