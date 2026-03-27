import React from "react";

export default function ServiceStatusBar({ services = [], generatedAt, scenario, seed }) {
  return (
    <div className="rounded-2xl border border-white/10 bg-slate-950/55 p-4 shadow-glow backdrop-blur-xl">
      <div className="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
        <div>
          <h2 className="text-lg font-semibold text-white">System Readiness</h2>
          <p className="mt-1 text-sm text-slate-400">
            Live service health and run context for the current governance session.
          </p>
        </div>

        <div className="flex flex-wrap gap-2">
          {services.map((svc) => (
            <div
              key={svc.name}
              className={`rounded-full px-3 py-1 text-xs font-medium tracking-wide ${
                svc.status === "up"
                  ? "border border-emerald-500/25 bg-emerald-500/15 text-emerald-300"
                  : "border border-rose-500/25 bg-rose-500/15 text-rose-300"
              }`}
            >
              {svc.name}: {svc.status.toUpperCase()}
            </div>
          ))}
        </div>
      </div>

      <div className="mt-4 grid grid-cols-1 gap-3 md:grid-cols-3">
        <InfoCard label="Latest Scenario" value={scenario || "-"} />
        <InfoCard label="Seed" value={seed ?? "-"} />
        <InfoCard
          label="Last Run"
          value={generatedAt ? new Date(generatedAt).toLocaleString() : "No run yet"}
        />
      </div>
    </div>
  );
}

function InfoCard({ label, value }) {
  return (
    <div className="rounded-xl border border-white/8 bg-white/[0.03] p-3">
      <div className="text-xs uppercase tracking-wider text-slate-500">{label}</div>
      <div className="mt-1 text-sm text-slate-100">{String(value)}</div>
    </div>
  );
}