import { CheckCircle2, Loader2, CircleDot, ChevronRight } from "lucide-react";

function normalizeStatus(status) {
  const s = String(status || "").toLowerCase();
  if (s.includes("complete") || s === "ready") return "complete";
  if (s.includes("run") || s.includes("progress")) return "running";
  if (s.includes("skip")) return "skipped";
  return "idle";
}

export default function StageTicker({ loading, stages = [] }) {
  const visibleStages = Array.isArray(stages) ? stages : [];
  const stillRunning =
    loading || visibleStages.some((s) => normalizeStatus(s.status) === "running");

  if (!visibleStages.length && !loading) return null;

  return (
    <div className="mt-3 rounded-2xl border border-white/10 bg-white/[0.04] px-4 py-3">
      <div className="flex flex-wrap items-center gap-2">
        {visibleStages.map((stage, idx) => {
          const status = normalizeStatus(stage.status);

          return (
            <div key={`${stage.name}-${idx}`} className="flex items-center gap-2">
              <div
                className={`flex items-center gap-2 rounded-full border px-3 py-1.5 text-xs ${
                  status === "complete"
                    ? "border-emerald-500/30 bg-emerald-500/10 text-emerald-200"
                    : status === "running"
                    ? "border-sky-500/30 bg-sky-500/10 text-sky-200"
                    : status === "skipped"
                    ? "border-amber-500/30 bg-amber-500/10 text-amber-200"
                    : "border-white/10 bg-slate-950/60 text-slate-300"
                }`}
              >
                {status === "complete" ? (
                  <CheckCircle2 size={14} />
                ) : status === "running" ? (
                  <Loader2 size={14} className="animate-spin" />
                ) : (
                  <CircleDot size={14} />
                )}
                <span>{stage.name}</span>
              </div>
              {idx < visibleStages.length - 1 ? (
                <ChevronRight size={14} className="text-slate-500" />
              ) : null}
            </div>
          );
        })}

        {stillRunning ? (
          <div className="ml-auto text-xs text-sky-300">
            Governance pipeline in progress...
          </div>
        ) : (
          <div className="ml-auto text-xs text-emerald-300">
            Governance pipeline completed.
          </div>
        )}
      </div>
    </div>
  );
}