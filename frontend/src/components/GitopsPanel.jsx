import { useEffect, useMemo, useState } from "react";
import {
  approveGitOps,
  rejectGitOps,
  refreshGitOpsPR,
  getGitOpsAudit,
} from "../api/client";

export default function GitOpsPanel({
  gitops,
  session,
  scenario,
  seed,
  onRefreshState,
}) {
  const prs = Array.isArray(gitops?.prs) ? gitops.prs : [];
  const status = String(gitops?.status || "idle").toUpperCase();
  const canReview = session?.features?.GITOPS_MERGE === "FULL";

  const [auditItems, setAuditItems] = useState([]);
  const [busyId, setBusyId] = useState("");
  const [comments, setComments] = useState({});
  const [panelError, setPanelError] = useState("");

  const loadAudit = async () => {
    try {
      const data = await getGitOpsAudit({ scenario, seed });
      setAuditItems(Array.isArray(data?.items) ? data.items : []);
    } catch (err) {
      setPanelError(
        err?.response?.data?.error || err?.message || "Failed to load GitOps audit"
      );
    }
  };

  useEffect(() => {
    loadAudit();
  }, [scenario, seed, gitops?.prs?.length]);

  const summary = useMemo(() => {
    const counts = {
      pending: 0,
      rejected: 0,
      created: 0,
      merged: 0,
      blocked: 0,
      failed: 0,
    };

    prs.forEach((item) => {
      const s = String(item?.status || "").toUpperCase();
      if (s === "PENDING_APPROVAL") counts.pending += 1;
      else if (s === "REJECTED") counts.rejected += 1;
      else if (s === "PR_CREATED") counts.created += 1;
      else if (s === "MERGED") counts.merged += 1;
      else if (s === "BLOCKED") counts.blocked += 1;
      else if (s === "FAILED") counts.failed += 1;
    });

    return counts;
  }, [prs]);

  const getComment = (approvalId) => comments[approvalId] || "";

  const setComment = (approvalId, value) => {
    setComments((prev) => ({ ...prev, [approvalId]: value }));
  };

  const runAction = async (approvalId, action) => {
    setPanelError("");
    setBusyId(approvalId);

    try {
      const payload = {
        scenario,
        seed,
        approvalId,
        comment: getComment(approvalId),
      };

      if (action === "approve") {
        await approveGitOps(payload);
      } else if (action === "reject") {
        await rejectGitOps(payload);
      } else if (action === "refresh") {
        await refreshGitOpsPR(payload);
      }

      await onRefreshState?.();
      await loadAudit();
    } catch (err) {
      setPanelError(
        err?.response?.data?.error || err?.message || "GitOps action failed"
      );
    } finally {
      setBusyId("");
    }
  };

  return (
    <div className="space-y-4">
      <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
        <div className="flex items-center justify-between gap-3">
          <div>
            <div className="text-lg font-semibold">GitOps Approval Gate</div>
            <div className="mt-1 text-sm text-slate-400">
              Validated remediations wait here until an authorized reviewer approves or rejects them.
            </div>
          </div>

          <span
            className={`rounded-full px-3 py-1 text-xs font-semibold ${
              status === "PENDING_APPROVAL"
                ? "border border-amber-500/20 bg-amber-500/10 text-amber-300"
                : "border border-white/10 bg-white/10 text-slate-300"
            }`}
          >
            {status}
          </span>
        </div>

        <div className="mt-4 grid gap-3 md:grid-cols-3 xl:grid-cols-6">
          <Metric label="Pending" value={summary.pending} />
          <Metric label="Rejected" value={summary.rejected} />
          <Metric label="PR Created" value={summary.created} />
          <Metric label="Merged" value={summary.merged} />
          <Metric label="Blocked" value={summary.blocked} />
          <Metric label="Failed" value={summary.failed} />
        </div>

        <div className="mt-4 text-sm text-slate-300">
          {gitops?.message || "No GitOps activity available."}
        </div>

        <div className="mt-3 text-xs text-slate-500">
          Reviewer access: {canReview ? "Enabled" : "Read-only"}
        </div>

        {panelError ? (
          <div className="mt-4 rounded-xl border border-rose-500/30 bg-rose-500/10 p-3 text-sm text-rose-200">
            {panelError}
          </div>
        ) : null}
      </div>

      {prs.length === 0 ? (
        <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4 text-sm text-slate-400">
          No approval requests available for this run.
        </div>
      ) : (
        prs.map((pr, index) => {
          const approvalId = pr.approvalId || `${pr.nodeId}-${index}`;
          const prStatus = String(pr.status || "PENDING_APPROVAL").toUpperCase();
          const isPending = prStatus === "PENDING_APPROVAL";
          const isCreated = prStatus === "PR_CREATED";
          const isBusy = busyId === approvalId;

          return (
            <div
              key={approvalId}
              className="rounded-2xl border border-white/10 bg-slate-950/60 p-4"
            >
              <div className="flex flex-wrap items-start justify-between gap-3">
                <div>
                  <div className="text-xl font-semibold text-sky-300">
                    {pr.nodeId || "Unknown Node"}
                  </div>
                  <div className="mt-1 text-sm text-slate-400">
                    {pr.action || "Unknown Action"}
                  </div>
                </div>

                <div className="flex items-center gap-2">
                  <span className="rounded-full border border-white/10 bg-slate-900/70 px-3 py-1 text-xs text-slate-300">
                    {approvalId}
                  </span>
                  <span
                    className={`rounded-full px-3 py-1 text-xs font-semibold ${
                      prStatus === "PENDING_APPROVAL"
                        ? "border border-amber-500/20 bg-amber-500/10 text-amber-300"
                        : prStatus === "REJECTED"
                        ? "border border-rose-500/20 bg-rose-500/10 text-rose-300"
                        : prStatus === "PR_CREATED" || prStatus === "MERGED"
                        ? "border border-emerald-500/20 bg-emerald-500/10 text-emerald-300"
                        : "border border-white/10 bg-white/10 text-slate-300"
                    }`}
                  >
                    {prStatus}
                  </span>
                </div>
              </div>

              <div className="mt-4 grid gap-3 md:grid-cols-2">
                <MetaField label="Requested At" value={pr.requestedAt || "—"} />
                <MetaField label="Reviewed By" value={pr.reviewedBy || "—"} />
                <MetaField label="Reviewed At" value={pr.reviewedAt || "—"} />
                <MetaField label="Branch" value={pr.branch || "—"} />
              </div>

              <div className="mt-4 rounded-xl bg-slate-900/70 p-3 text-sm text-slate-300">
                {pr.message || "Awaiting reviewer decision."}
              </div>

              <div className="mt-4">
                <div className="text-xs uppercase tracking-wide text-slate-500">
                  Review Comment
                </div>
                <textarea
                  value={getComment(approvalId)}
                  onChange={(e) => setComment(approvalId, e.target.value)}
                  placeholder="Add a reviewer note..."
                  className="mt-2 min-h-[88px] w-full rounded-xl border border-white/10 bg-slate-900/70 p-3 text-sm outline-none"
                  disabled={!canReview || !isPending || isBusy}
                />
              </div>

              <div className="mt-4 flex flex-wrap gap-3">
                {canReview && isPending ? (
                  <>
                    <button
                      onClick={() => runAction(approvalId, "approve")}
                      disabled={isBusy}
                      className="rounded-xl border border-emerald-500/20 bg-emerald-500/10 px-4 py-2 text-sm text-emerald-200 disabled:opacity-50"
                    >
                      {isBusy ? "Working..." : "Approve and Create PR"}
                    </button>

                    <button
                      onClick={() => runAction(approvalId, "reject")}
                      disabled={isBusy}
                      className="rounded-xl border border-rose-500/20 bg-rose-500/10 px-4 py-2 text-sm text-rose-200 disabled:opacity-50"
                    >
                      {isBusy ? "Working..." : "Reject"}
                    </button>
                  </>
                ) : null}

                {isCreated || prStatus === "MERGED" ? (
                  <button
                    onClick={() => runAction(approvalId, "refresh")}
                    disabled={isBusy}
                    className="rounded-xl border border-white/10 bg-white/5 px-4 py-2 text-sm text-slate-200 disabled:opacity-50"
                  >
                    Refresh PR Status
                  </button>
                ) : null}

                {pr.url ? (
                  <a
                    href={pr.url}
                    target="_blank"
                    rel="noreferrer"
                    className="rounded-xl border border-sky-500/20 bg-sky-500/10 px-4 py-2 text-sm text-sky-200 hover:bg-sky-500/15"
                  >
                    Open Pull Request
                  </a>
                ) : null}
              </div>
            </div>
          );
        })
      )}

      <div className="rounded-2xl border border-white/10 bg-slate-950/60 p-4">
        <div className="text-lg font-semibold">Approval Audit Trail</div>
        <div className="mt-1 text-sm text-slate-400">
          Every review action is recorded with actor, timestamp, status, and comment.
        </div>

        <div className="mt-4 space-y-3">
          {auditItems.length === 0 ? (
            <div className="text-sm text-slate-500">No audit events recorded yet.</div>
          ) : (
            auditItems.map((item, index) => (
              <div
                key={`${item.timestamp}-${item.approvalId}-${index}`}
                className="rounded-xl border border-white/10 bg-slate-900/70 p-3"
              >
                <div className="flex flex-wrap items-center justify-between gap-2">
                  <div className="text-sm font-medium text-slate-200">
                    {item.action} · {item.nodeId} · {item.finalAction}
                  </div>
                  <div className="text-xs text-slate-500">{item.timestamp}</div>
                </div>
                <div className="mt-2 text-sm text-slate-300">
                  Actor: {item.actor} · Status: {item.status}
                </div>
                {item.reviewComment ? (
                  <div className="mt-2 text-sm text-slate-400">
                    Comment: {item.reviewComment}
                  </div>
                ) : null}
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
}

function Metric({ label, value }) {
  return (
    <div className="rounded-xl bg-slate-900/70 p-3">
      <div className="text-[11px] uppercase tracking-wider text-slate-500">{label}</div>
      <div className="mt-1 text-lg font-semibold text-slate-200">{value}</div>
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