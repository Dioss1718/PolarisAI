export default function FeedbackPanel({ feedback = {} }) {
  const avgReward = Number(feedback.avgReward || 0).toFixed(2);
  const count = Number(feedback.count || 0);
  const riskWeight = Number(feedback.riskWeight || 0).toFixed(2);
  const costWeight = Number(feedback.costWeight || 0).toFixed(2);
  const penalty = Number(feedback.penalty || 0).toFixed(2);

  return (
    <div className="min-h-0 overflow-auto rounded-2xl border border-white/10 bg-white/[0.04] p-4">
      <div className="flex items-center justify-between">
        <div>
          <div className="text-lg font-semibold">Adaptive Feedback Loop</div>
          <div className="text-sm text-slate-400">Self-improving optimization weights from past decisions.</div>
        </div>
      </div>

      <div className="mt-4 grid grid-cols-2 gap-3">
        <Metric label="Avg Reward" value={avgReward} />
        <Metric label="Decision Count" value={count} />
        <Metric label="Risk Weight" value={riskWeight} />
        <Metric label="Cost Weight" value={costWeight} />
      </div>

      <div className="mt-4 rounded-xl border border-white/10 bg-slate-950/60 p-4 text-sm text-slate-300">
        <div className="text-xs uppercase tracking-wider text-slate-500">Penalty Coefficient</div>
        <div className="mt-2 text-2xl font-semibold text-violet-300">{penalty}</div>
        <div className="mt-3 text-xs text-slate-500">
          This panel demonstrates the closed-loop learning system that updates trade-off behavior over time.
        </div>
      </div>
    </div>
  );
}

function Metric({ label, value }) {
  return (
    <div className="rounded-xl border border-white/10 bg-slate-950/60 p-4">
      <div className="text-xs uppercase tracking-wider text-slate-500">{label}</div>
      <div className="mt-2 text-xl font-semibold">{value}</div>
    </div>
  );
}