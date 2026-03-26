export default function FeedbackPanel({ feedback }) {
  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4">
      <h2 className="text-lg font-semibold">Adaptive Learning Loop</h2>
      <div className="mt-4 grid grid-cols-2 gap-3 lg:grid-cols-4">
        <Metric label="Avg Reward" value={feedback.avgReward?.toFixed(2)} />
        <Metric label="Count" value={feedback.count} />
        <Metric label="Risk Weight" value={feedback.riskWeight?.toFixed(2)} />
        <Metric label="Cost Weight" value={feedback.costWeight?.toFixed(2)} />
      </div>
      <div className="mt-3 rounded-xl bg-slate-950/70 p-3 text-sm text-slate-400">
        Penalty coefficient: <span className="text-slate-200">{feedback.penalty?.toFixed(2)}</span>
      </div>
    </div>
  );
}

function Metric({ label, value }) {
  return (
    <div className="rounded-xl bg-slate-950/70 p-4">
      <div className="text-xs uppercase tracking-wider text-slate-500">{label}</div>
      <div className="mt-2 text-xl font-semibold">{value}</div>
    </div>
  );
}