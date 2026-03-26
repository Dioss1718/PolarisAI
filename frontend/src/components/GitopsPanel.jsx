import React from "react";

export default function GitOpsPanel({ gitops }) {
  const safeGitops = gitops || {
    status: "idle",
    message: "No GitOps run yet.",
    prs: [],
  };

  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4">
      <h2 className="text-lg font-semibold">GitOps Execution Gate</h2>

      <div className="mt-4 rounded-xl border border-borderSoft bg-slate-950/70 p-4">
        <div className="flex items-center justify-between">
          <div className="text-sm font-medium">Pipeline Status</div>
          <div
            className={`rounded-full px-3 py-1 text-xs uppercase tracking-wider ${
              safeGitops.status === "ready"
                ? "bg-emerald-500/15 text-emerald-300"
                : safeGitops.status === "skipped"
                ? "bg-amber-500/15 text-amber-300"
                : "bg-slate-500/15 text-slate-300"
            }`}
          >
            {safeGitops.status}
          </div>
        </div>

        <div className="mt-3 text-sm text-slate-300">{safeGitops.message}</div>
      </div>

      <div className="mt-4 space-y-3">
        {safeGitops.prs && safeGitops.prs.length > 0 ? (
          safeGitops.prs.map((pr) => (
            <div
              key={`${pr.nodeId || "node"}-${pr.prNumber || "pr"}`}
              className="rounded-xl border border-borderSoft bg-slate-950/70 p-4"
            >
              <div className="flex items-center justify-between">
                <div>
                  <div className="font-medium text-sky-300">
                    {pr.nodeId || "Unknown Node"}
                  </div>
                  <div className="text-xs text-slate-500">
                    {(pr.action || "Unknown Action")} · Branch: {pr.branch || "-"}
                  </div>
                </div>

                <div className="rounded-full bg-emerald-500/15 px-3 py-1 text-xs uppercase tracking-wider text-emerald-300">
                  PR #{pr.prNumber || "-"}
                </div>
              </div>

              <div className="mt-3 text-sm text-slate-400">
                {pr.message || "Pull request created for governed remediation."}
              </div>

              {pr.url ? (
                <a
                  href={pr.url}
                  target="_blank"
                  rel="noreferrer"
                  className="mt-3 inline-block rounded-lg border border-sky-500/30 bg-sky-500/10 px-3 py-2 text-sm text-sky-300 hover:bg-sky-500/20"
                >
                  Open Pull Request
                </a>
              ) : null}
            </div>
          ))
        ) : (
          <div className="rounded-xl border border-borderSoft bg-slate-950/70 p-4 text-sm text-slate-500">
            No pull requests available for this run.
          </div>
        )}
      </div>
    </div>
  );
}