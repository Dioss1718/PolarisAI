function badge(status) {
  if (status === "APPROVED") return "border-emerald-500/30 bg-emerald-500/10 text-emerald-200";
  if (status === "MODIFIED") return "border-amber-500/30 bg-amber-500/10 text-amber-200";
  return "border-rose-500/30 bg-rose-500/10 text-rose-200";
}

export default function RecommendationList({ recommendations = [], onSelect }) {
  return (
    <div className="min-h-0 overflow-auto rounded-2xl border border-white/10 bg-white/[0.04] p-4">
      <div className="mb-3 flex items-center justify-between">
        <div>
          <div className="text-lg font-semibold">Governance Actions</div>
          <div className="text-sm text-slate-400">
            Detected issue → selected remediation → projected enterprise outcome.
          </div>
        </div>
      </div>

      <div className="space-y-3">
        {recommendations.map((r) => (
          <button
            key={`${r.nodeId}-${r.finalAction}`}
            onClick={() => onSelect(r)}
            className="block w-full rounded-2xl border border-white/10 bg-slate-950/60 p-4 text-left hover:border-sky-500/30"
          >
            <div className="flex flex-wrap items-start justify-between gap-3">
              <div>
                <div className="font-medium">{r.nodeId}</div>
                <div className="text-xs text-slate-500">
                  {r.cloud} · {r.type} · {r.environment}
                </div>
              </div>

              <div className="flex flex-wrap gap-2">
                <span className={`rounded-full border px-2.5 py-1 text-[11px] ${badge(r.status)}`}>
                  {r.status}
                </span>
                <span className="rounded-full border border-white/10 bg-slate-900/70 px-2.5 py-1 text-[11px] text-slate-300">
                  Confidence {Number(r.confidence || 0).toFixed(2)}
                </span>
                <span className="rounded-full border border-white/10 bg-slate-900/70 px-2.5 py-1 text-[11px] text-slate-300">
                  Safety {r.safetyLevel || "—"}
                </span>
              </div>
            </div>

            <div className="mt-3 grid grid-cols-4 gap-3 text-sm">
              <Metric label="Selected Action" value={r.finalAction} />
              <Metric label="Risk" value={Number(r.risk).toFixed(2)} />
              <Metric label="Cost Δ" value={Number(r.costDelta || 0).toFixed(2)} />
              <Metric label="Blast Radius" value={r.blastRadius ?? "—"} />
            </div>

            <div className="mt-3 rounded-xl bg-slate-900/70 p-3 text-sm text-slate-300">
              {r.reason}
            </div>

            <div className="mt-3 grid gap-3 md:grid-cols-2">
              <Metric label="GitOps Path" value={r.gitOpsPath || "—"} />
              <Metric label="Rollback Path" value={r.rollbackPath || "—"} />
            </div>
          </button>
        ))}
      </div>
    </div>
  );
}

function Metric({ label, value }) {
  return (
    <div className="rounded-xl bg-slate-900/70 p-3">
      <div className="text-[11px] uppercase tracking-wider text-slate-500">{label}</div>
      <div className="mt-1 font-medium break-words">{value}</div>
    </div>
  );
}