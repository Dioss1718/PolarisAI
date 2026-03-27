import React from "react";
import { Activity, Play, RefreshCcw, Shield, GitBranch, Sparkles } from "lucide-react";

export default function TopBar({
  onRun,
  loading,
  scenario,
  setScenario,
  seed,
  setSeed,
}) {
  return (
    <div className="sticky top-4 z-40 rounded-2xl border border-white/10 bg-slate-950/65 shadow-glow backdrop-blur-xl">
      <div className="flex flex-col gap-4 p-4 lg:flex-row lg:items-center lg:justify-between">
        <div className="flex items-center gap-4">
          <div className="flex h-14 w-14 items-center justify-center rounded-2xl border border-sky-400/20 bg-white/5 shadow-[0_0_30px_rgba(56,189,248,0.14)]">
            <svg
              width="28"
              height="28"
              viewBox="0 0 100 100"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path d="M50 10L66 42L50 34L34 42L50 10Z" fill="url(#g1)" />
              <path d="M50 90L34 58L50 66L66 58L50 90Z" fill="url(#g1)" />
              <circle cx="50" cy="50" r="10" fill="url(#g2)" />
              <defs>
                <linearGradient id="g1" x1="20" y1="20" x2="80" y2="80">
                  <stop stopColor="#7dd3fc" />
                  <stop offset="1" stopColor="#a78bfa" />
                </linearGradient>
                <linearGradient id="g2" x1="40" y1="40" x2="60" y2="60">
                  <stop stopColor="#ffffff" />
                  <stop offset="1" stopColor="#7dd3fc" />
                </linearGradient>
              </defs>
            </svg>
          </div>

          <div>
            <div className="text-[11px] uppercase tracking-[0.38em] text-sky-300/70">
              PolarisAI Control Plane
            </div>
            <h1 className="mt-1 text-2xl font-semibold tracking-tight text-white">
              Autonomous Cloud Governance Console
            </h1>
            <p className="mt-1 text-sm text-slate-400">
              Operate infrastructure as a connected governance system, not a collection of isolated alerts.
            </p>
          </div>
        </div>

        <div className="grid grid-cols-1 gap-3 sm:grid-cols-4">
          <ControlCard icon={<Activity size={14} />} label="Scenario">
            <select
              value={scenario}
              onChange={(e) => setScenario(e.target.value)}
              className="mt-2 w-full rounded-xl border border-white/10 bg-slate-900/90 px-3 py-2 text-sm text-slate-100 outline-none"
            >
              <option value="FULL_CHAOS">FULL_CHAOS</option>
              <option value="SECURITY_BREACH">SECURITY_BREACH</option>
              <option value="COST_SPIKE">COST_SPIKE</option>
              <option value="COMPLIANCE_DRIFT">COMPLIANCE_DRIFT</option>
              <option value="BILL_SHOCK">BILL_SHOCK</option>
            </select>
          </ControlCard>

          <ControlCard icon={<Shield size={14} />} label="Seed">
            <input
              value={seed}
              onChange={(e) => setSeed(Number(e.target.value || 0))}
              className="mt-2 w-full rounded-xl border border-white/10 bg-slate-900/90 px-3 py-2 text-sm text-slate-100 outline-none"
            />
          </ControlCard>

          <ControlCard icon={<GitBranch size={14} />} label="Mode">
            <div className="mt-2 text-sm font-medium text-slate-100">Operator Console</div>
            <div className="text-xs text-slate-500">Graph-native governance</div>
          </ControlCard>

          <button
            onClick={onRun}
            disabled={loading}
            className="flex min-h-[92px] items-center justify-center rounded-2xl border border-sky-400/20 bg-[linear-gradient(135deg,rgba(56,189,248,0.95),rgba(99,102,241,0.95))] px-4 py-3 text-slate-950 shadow-[0_10px_40px_rgba(56,189,248,0.35)] transition hover:scale-[1.01] disabled:opacity-50"
          >
            <div className="text-center">
              <div className="flex items-center justify-center gap-2 font-semibold">
                {loading ? <RefreshCcw size={16} className="animate-spin" /> : <Play size={16} />}
                {loading ? "Running..." : "Run Governance"}
              </div>
              <div className="mt-1 text-xs text-slate-900/75">
                Execute full decision loop
              </div>
            </div>
          </button>
        </div>
      </div>

      <div className="flex flex-wrap items-center gap-3 border-t border-white/5 px-4 py-3 text-xs text-slate-400">
        <div className="inline-flex items-center gap-2 rounded-full border border-white/10 bg-white/5 px-3 py-1">
          <Sparkles size={12} className="text-sky-300" />
          Premium operator workspace
        </div>
        <div>Graph-aware risk analysis</div>
        <div>Policy-enforced remediation</div>
        <div>Explainable decisions</div>
        <div>GitOps-ready execution</div>
      </div>
    </div>
  );
}

function ControlCard({ icon, label, children }) {
  return (
    <div className="rounded-2xl border border-white/10 bg-white/[0.04] px-4 py-3">
      <div className="flex items-center gap-2 text-[11px] uppercase tracking-[0.24em] text-slate-400">
        {icon}
        {label}
      </div>
      {children}
    </div>
  );
}