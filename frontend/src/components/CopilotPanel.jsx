import { useState } from "react";

export default function CopilotPanel({ state, onSelectNode }) {
  const [query, setQuery] = useState("");
  const [answer, setAnswer] = useState("Ask the control plane about risky nodes, bill shock, approved actions, or a specific resource.");

  const runQuery = () => {
    const q = query.trim().toLowerCase();
    if (!q) return;

    if (q.includes("bill shock")) {
      const shocked = state.forecasts.filter((f) => f.billShock).map((f) => f.nodeId);
      setAnswer(shocked.length ? `Bill shock risk detected on: ${shocked.join(", ")}` : "No bill shock risk detected.");
      return;
    }

    if (q.includes("approved")) {
      const approved = state.recommendations
        .filter((r) => r.status === "APPROVED" || r.status === "MODIFIED")
        .map((r) => `${r.nodeId}: ${r.finalAction}`);
      setAnswer(approved.length ? approved.join(" | ") : "No approved actions found.");
      return;
    }

    const matchedNode = state.nodes.find((n) => q.includes(n.id.toLowerCase()) || q.includes(n.label.toLowerCase()));
    if (matchedNode) {
      onSelectNode(matchedNode.id);
      const rec = state.recommendations.find((r) => r.nodeId === matchedNode.id);
      setAnswer(
        rec
          ? `${matchedNode.id} is ${matchedNode.environment} ${matchedNode.type} on ${matchedNode.cloud}. Risk=${matchedNode.risk}. Final action=${rec.finalAction}.`
          : `${matchedNode.id} selected. No final recommendation available.`
      );
      return;
    }

    if (q.includes("risk")) {
      const sorted = [...state.nodes].sort((a, b) => b.risk - a.risk).slice(0, 3);
      setAnswer(`Top risky resources: ${sorted.map((n) => `${n.id} (${n.risk})`).join(", ")}`);
      return;
    }

    setAnswer("Try prompts like: 'show bill shock', 'approved actions', 'why aws_vm1', or 'top risk nodes'.");
  };

  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4">
      <h2 className="text-lg font-semibold">Operator Command Space</h2>
      <p className="mt-1 text-sm text-slate-400">
        Query the current governance state like an autonomous cloud operator.
      </p>

      <div className="mt-4 flex gap-2">
        <input
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          placeholder="show bill shock, why aws_vm1, approved actions..."
          className="flex-1 rounded-xl border border-borderSoft bg-slate-950/70 px-4 py-3 outline-none"
        />
        <button
          onClick={runQuery}
          className="rounded-xl bg-violet-500 px-4 py-3 font-medium text-slate-950 hover:bg-violet-400"
        >
          Ask
        </button>
      </div>

      <div className="mt-4 rounded-xl bg-slate-950/70 p-4 text-sm text-slate-200 min-h-[112px] whitespace-pre-wrap">
        {answer}
      </div>
    </div>
  );
}