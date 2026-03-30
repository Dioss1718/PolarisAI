import { useEffect, useState } from "react";
import { login } from "../api/client";

export default function LoginDialog({ open, onClose, onLoggedIn }) {
  const [employeeId, setEmployeeId] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (open) {
      setError("");
    }
  }, [open]);

  if (!open) return null;

  const submit = async () => {
    const normalizedEmployeeId = employeeId.trim().toUpperCase();
    const normalizedPassword = password.trim();

    if (!normalizedEmployeeId || !normalizedPassword) {
      setError("Employee ID and password are required.");
      return;
    }

    setLoading(true);
    setError("");

    try {
      const session = await login(normalizedEmployeeId, normalizedPassword);
      onLoggedIn?.(session);
    } catch (err) {
      const serverError = err?.response?.data?.error;
      if (err?.response?.status === 401) {
        setError(serverError || "Invalid employee ID or password.");
      } else {
        setError(serverError || err.message || "Login failed");
      }
    } finally {
      setLoading(false);
    }
  };

  const onKeyDown = (e) => {
    if (e.key === "Enter" && !loading) {
      submit();
    }
  };

  return (
    <div className="fixed inset-0 z-[100] flex items-center justify-center bg-black/50 backdrop-blur-sm">
      <div className="w-full max-w-md rounded-2xl border border-white/10 bg-slate-950 p-6 shadow-2xl">
        <h2 className="text-xl font-semibold text-white">Employee Login</h2>
        <p className="mt-1 text-sm text-slate-400">
          Authenticate to unlock role-aware governance controls.
        </p>

        <div className="mt-5 space-y-3">
          <input
            value={employeeId}
            onChange={(e) => setEmployeeId(e.target.value)}
            onKeyDown={onKeyDown}
            placeholder="Employee ID (e.g. AT001)"
            className="w-full rounded-xl border border-white/10 bg-slate-900 px-4 py-3 text-white outline-none"
          />
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            onKeyDown={onKeyDown}
            placeholder="Password"
            className="w-full rounded-xl border border-white/10 bg-slate-900 px-4 py-3 text-white outline-none"
          />
        </div>

        <div className="mt-3 rounded-xl border border-white/10 bg-slate-900/60 p-3 text-xs text-slate-400">
          Demo credentials:
          <div className="mt-2 space-y-1">
            <div>AT001 / admin123</div>
            <div>AT002 / devops123</div>
            <div>AT003 / security123</div>
          </div>
        </div>

        {error ? (
          <div className="mt-3 rounded-xl border border-rose-500/30 bg-rose-500/10 p-3 text-sm text-rose-200">
            {error}
          </div>
        ) : null}

        <div className="mt-5 flex justify-end gap-2">
          <button
            onClick={onClose}
            className="rounded-xl border border-white/10 px-4 py-2 text-white"
          >
            Close
          </button>
          <button
            onClick={submit}
            disabled={loading}
            className="rounded-xl bg-sky-500 px-4 py-2 font-medium text-slate-950 disabled:opacity-50"
          >
            {loading ? "Signing in..." : "Sign In"}
          </button>
        </div>
      </div>
    </div>
  );
}