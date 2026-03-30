export default function ServicePill({ name, status }) {
  return (
    <div className={`rounded-full px-3 py-1 text-xs ${status === "up" ? "border border-emerald-500/20 bg-emerald-500/10 text-emerald-300" : "border border-rose-500/20 bg-rose-500/10 text-rose-300"}`}>
      {name}: {String(status).toUpperCase()}
    </div>
  );
}