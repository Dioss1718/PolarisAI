import { useState } from "react";
import {
  Play,
  RefreshCcw,
  ShieldEllipsis,
  LogOut,
  LayoutPanelTop,
  Sparkles,
  Menu,
} from "lucide-react";

export default function HeaderBar({
  scenario,
  setScenario,
  seed,
  setSeed,
  onRun,
  loading,
  onOpenLogin,
  session,
  onLogout,
  onOpenWorkspace,
}) {
  const [menuOpen, setMenuOpen] = useState(false);

  return (
    <div className="grid h-[94px] grid-cols-[minmax(0,1fr)_170px_64px_190px_180px_260px] items-center gap-3 rounded-[24px] border border-white/10 bg-white/[0.05] px-4 backdrop-blur-xl">
      <div className="flex min-w-0 items-center gap-4 overflow-hidden">
        <div className="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl border border-sky-400/20 bg-slate-950/60 shadow-[0_0_35px_rgba(56,189,248,0.18)]">
          <ShieldEllipsis className="text-sky-300" size={22} />
        </div>

        <div className="min-w-0">
          <div className="truncate text-[11px] uppercase tracking-[0.34em] text-sky-300/70">
            PolarisAI
          </div>
          <div className="truncate text-[22px] font-semibold">
            GenAI Cloud Security Copilot Control Plane
          </div>
          <div className="truncate text-xs text-slate-400">
            Misconfig detection · risk prioritization · FinOps optimization · explainable governance
          </div>
        </div>
      </div>

      <div className="rounded-2xl border border-white/10 bg-slate-950/55 px-3 py-3">
        <div className="text-[11px] uppercase tracking-[0.24em] text-slate-400">
          Scenario
        </div>
        <select
          value={scenario}
          onChange={(e) => setScenario(e.target.value)}
          className="mt-2 w-full rounded-xl border border-white/10 bg-slate-900/90 px-3 py-2 text-sm text-slate-100 outline-none"
        >
          <option value="FULL_CHAOS">FULL_CHAOS</option>
          <option value="SECURITY_BREACH">SECURITY_BREACH</option>
          <option value="COST_SPIKE">COST_SPIKE</option>
          <option value="POLICY_DRIFT">POLICY_DRIFT</option>
        </select>
      </div>

      <div className="rounded-2xl border border-white/10 bg-slate-950/55 px-2 py-3">
        <div className="text-[11px] uppercase tracking-[0.24em] text-slate-400">
          Seed
        </div>
        <input
          value={seed}
          onChange={(e) => setSeed(Number(e.target.value || 0))}
          className="mt-2 w-full rounded-xl border border-white/10 bg-slate-900/90 px-2 py-2 text-sm text-slate-100 outline-none"
        />
      </div>

      <button
        onClick={onRun}
        disabled={loading}
        className="flex h-[56px] items-center justify-center rounded-2xl border border-sky-400/20 bg-[linear-gradient(135deg,rgba(56,189,248,0.95),rgba(99,102,241,0.95))] px-4 text-sm font-semibold text-slate-950 shadow-[0_10px_40px_rgba(56,189,248,0.35)] transition hover:scale-[1.01] disabled:opacity-50"
      >
        <div className="flex items-center gap-2">
          {loading ? <RefreshCcw size={15} className="animate-spin" /> : <Play size={15} />}
          {loading ? "Running..." : "Run Governance"}
        </div>
      </button>

      <button
        onClick={() => onOpenWorkspace?.("governance")}
        className="flex h-[56px] items-center justify-center gap-2 rounded-2xl border border-white/10 bg-slate-950/70 px-4 text-sm hover:bg-slate-900"
      >
        <LayoutPanelTop size={16} />
        Workspace
      </button>

      <div className="flex min-w-0 items-center justify-end gap-2 relative">
        <button
          onClick={() => setMenuOpen((v) => !v)}
          className="flex h-[56px] w-[56px] shrink-0 items-center justify-center rounded-2xl border border-white/10 bg-slate-950/70 hover:bg-slate-900"
        >
          <Menu size={18} />
        </button>

        {menuOpen ? (
          <div className="absolute right-[200px] top-[66px] z-50 w-64 rounded-2xl border border-white/10 bg-slate-950/95 p-2 shadow-2xl">
            <button
              onClick={() => {
                onOpenWorkspace?.("explainability");
                setMenuOpen(false);
              }}
              className="w-full rounded-xl px-3 py-2 text-left text-sm text-slate-200 hover:bg-slate-900"
            >
              Explainability
            </button>
            <button
              onClick={() => {
                onOpenWorkspace?.("feedback");
                setMenuOpen(false);
              }}
              className="w-full rounded-xl px-3 py-2 text-left text-sm text-slate-200 hover:bg-slate-900"
            >
              Adaptive Feedback
            </button>
            <button
              onClick={() => {
                onOpenWorkspace?.("negotiation");
                setMenuOpen(false);
              }}
              className="w-full rounded-xl px-3 py-2 text-left text-sm text-slate-200 hover:bg-slate-900"
            >
              Negotiation & Tradeoffs
            </button>
          </div>
        ) : null}

        {session ? (
          <>
            <div className="min-w-0 rounded-2xl border border-white/10 bg-slate-950/70 px-4 py-3 text-sm">
              <div className="flex items-center gap-2 font-medium">
                <Sparkles size={14} className="text-sky-300" />
                <span className="truncate">{session.name}</span>
              </div>
              <div className="truncate text-xs text-slate-500">{session.role}</div>
            </div>
            <button
              onClick={onLogout}
              className="flex h-[56px] w-[56px] shrink-0 items-center justify-center rounded-2xl border border-white/10 bg-slate-950/70 hover:bg-slate-900"
            >
              <LogOut size={16} />
            </button>
          </>
        ) : (
          <button
            onClick={onOpenLogin}
            className="rounded-2xl border border-white/10 bg-slate-950/70 px-4 py-3 text-sm hover:bg-slate-900"
          >
            Employee Login
          </button>
        )}
      </div>
    </div>
  );
}