export default function GitOpsPanel({ gitops }) {
  const prs = Array.isArray(gitops?.prs) ? gitops.prs : [];
  const status = String(gitops?.status || "idle").toUpperCase();

  return (
    <div className="space-y-4">
      <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
        <div className="flex items-center justify-between gap-3">
          <div>
            <div className="text-lg font-semibold">GitOps Execution Gate</div>
            <div className="mt-1 text-sm text-slate-400">Pipeline status</div>
          </div>

          <span className={`rounded-full px-3 py-1 text-xs font-semibold ${
            status === "READY"
              ? "bg-emerald-500/15 text-emerald-300 border border-emerald-500/20"
              : "bg-white/10 text-slate-300 border border-white/10"
          }`}>
            {status}
          </span>
        </div>

        <div className="mt-4 text-sm text-slate-300">
          {gitops?.message || "No GitOps activity available."}
        </div>
      </div>

      {prs.length === 0 ? (
        <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4 text-sm text-slate-400">
          No pull requests available for this run.
        </div>
      ) : (
        prs.map((pr, index) => {
          const prNumber = pr.prNumber ?? "—";
          const nodeId = pr.nodeId || "Unknown Node";
          const action = pr.action || "Unknown Action";
          const branch = pr.branch || "—";
          const url = pr.url || pr.URL || "";
          const message = pr.message || "Pull request created for governed remediation.";
          const prStatus = String(pr.status || "READY").toUpperCase();

          return (
            <div
              key={`${prNumber}-${index}`}
              className="rounded-2xl border border-white/10 bg-slate-950/60 p-4"
            >
              <div className="flex flex-wrap items-start justify-between gap-3">
                <div>
                  <div className="text-xl font-semibold text-sky-300">{nodeId}</div>
                  <div className="mt-1 text-sm text-slate-500">{action}</div>
                </div>

                <div className="flex items-center gap-2">
                  <span className="rounded-full border border-white/10 bg-slate-900/70 px-3 py-1 text-xs text-slate-300">
                    PR #{prNumber}
                  </span>
                  <span className={`rounded-full px-3 py-1 text-xs font-semibold ${
                    prStatus === "READY"
                      ? "bg-emerald-500/15 text-emerald-300 border border-emerald-500/20"
                      : "bg-white/10 text-slate-300 border border-white/10"
                  }`}>
                    {prStatus}
                  </span>
                </div>
              </div>

              <div className="mt-4 grid gap-3 md:grid-cols-2">
                <MetaField label="Repository" value="polaris-gitops" />
                <MetaField label="Branch" value={branch} />
              </div>

              <div className="mt-4 rounded-xl bg-slate-900/70 p-3 text-sm text-slate-300">
                {message}
              </div>

              <div className="mt-4 flex flex-wrap gap-3">
                {url ? (
                  <a
                    href={url}
                    target="_blank"
                    rel="noreferrer"
                    className="rounded-xl border border-sky-500/20 bg-sky-500/10 px-4 py-2 text-sm text-sky-200 hover:bg-sky-500/15"
                  >
                    Open Pull Request
                  </a>
                ) : (
                  <button
                    disabled
                    className="rounded-xl border border-white/10 bg-white/5 px-4 py-2 text-sm text-slate-500"
                  >
                    Open Pull Request
                  </button>
                )}

                <button
                  disabled
                  className="rounded-xl border border-white/10 bg-white/5 px-4 py-2 text-sm text-slate-500"
                >
                  View Branch
                </button>
              </div>
            </div>
          );
        })
      )}
    </div>
  );
}

function MetaField({ label, value }) {
  return (
    <div className="rounded-xl bg-slate-900/70 p-3">
      <div className="text-[11px] uppercase tracking-wider text-slate-500">{label}</div>
      <div className="mt-1 break-all font-medium text-slate-200">{value}</div>
    </div>
  );
}