import { useMemo, useState } from "react";

export default function CopilotDrawer({ open, onClose, state, onSelectNode }) {
  const [query, setQuery] = useState("");
  const [submittedQuery, setSubmittedQuery] = useState("");

  const answer = useMemo(() => {
    const q = submittedQuery.trim().toLowerCase();
    if (!q) return "Ask about risk, bill shock, approved actions, attack paths, or a specific node.";

    if (q.includes("bill shock")) {
      const shocked = (state.forecasts || []).filter((f) => f.billShock).map((f) => f.nodeId);
      return shocked.length ? `Bill shock risk detected on: ${shocked.join(", ")}` : "No bill shock risk detected.";
    }

    if (q.includes("approved")) {
      const approved = (state.recommendations || [])
        .filter((r) => r.status === "APPROVED" || r.status === "MODIFIED")
        .map((r) => `${r.nodeId}: ${r.finalAction}`);
      return approved.length ? approved.join(" | ") : "No approved actions found.";
    }

    if (q.includes("attack path")) {
      return state.attackPaths?.length
        ? `Detected ${state.attackPaths.length} attack path(s). Select one in the attack path panel for graph highlighting.`
        : "No attack paths are currently present in the loaded state.";
    }

    const matchedNode = (state.nodes || []).find(
      (n) =>
        q.includes(String(n.id).toLowerCase()) ||
        q.includes(String(n.label).toLowerCase())
    );

    if (matchedNode) {
      onSelectNode?.(matchedNode.id);
      const rec = (state.recommendations || []).find((r) => r.nodeId === matchedNode.id);
      return rec
        ? `${matchedNode.id} is ${matchedNode.environment} ${matchedNode.type} on ${matchedNode.cloud}. Risk=${matchedNode.risk}. Final action=${rec.finalAction}.`
        : `${matchedNode.id} selected. No final recommendation available.`;
    }

    if (q.includes("risk")) {
      const sorted = [...(state.nodes || [])].sort((a, b) => b.risk - a.risk).slice(0, 3);
      return sorted.length
        ? `Top risky resources: ${sorted.map((n) => `${n.id} (${n.risk})`).join(", ")}`
        : "No node risk data is currently loaded.";
    }

    return "Try: bill shock, attack paths, approved actions, top risk nodes, why aws_vm1.";
  }, [submittedQuery, state, onSelectNode]);

  const handleSubmit = () => {
    if (!query.trim()) return;
    setSubmittedQuery(query);
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
          <div className="text-sm text-slate-400">Grounded on current pipeline state.</div>
        </div>
        <button onClick={onClose} className="rounded-lg border border-white/10 px-3 py-1 text-sm">
          Close
        </button>
      </div>

      <div className="mt-4 flex items-center gap-2">
        <input
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") handleSubmit();
          }}
          placeholder="Ask about risk, bill shock, approved actions..."
          className="flex-1 rounded-xl border border-white/10 bg-slate-900 px-4 py-3 outline-none"
        />
        <button
          onClick={handleSubmit}
          disabled={!query.trim()}
          className="rounded-xl bg-sky-500 px-4 py-3 font-medium text-slate-950 disabled:opacity-50"
        >
          Send
        </button>
      </div>

      <div className="mt-4 rounded-xl bg-slate-900 p-4 text-sm whitespace-pre-wrap text-slate-200">
        {answer}
      </div>
    </div>
  );
}