import { Bell } from "lucide-react";

export default function NotificationBell({ count = 0, onClick }) {
  return (
    <button
      onClick={onClick}
      className="relative flex h-11 w-11 items-center justify-center rounded-2xl border border-white/10 bg-slate-950/70 hover:bg-slate-900"
      title={count > 0 ? `${count} pull request notification(s)` : "Open GitOps"}
    >
      <Bell size={18} className="text-slate-200" />
      {count > 0 ? (
        <span className="absolute -right-1 -top-1 rounded-full bg-emerald-500 px-1.5 py-0.5 text-[10px] font-semibold text-slate-950">
          {count}
        </span>
      ) : null}
    </button>
  );
}