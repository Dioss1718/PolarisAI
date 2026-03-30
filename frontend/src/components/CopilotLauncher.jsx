import { MessageSquareMore } from "lucide-react";

export default function CopilotLauncher({ onClick }) {
  return (
    <button
      onClick={onClick}
      className="
        group
        fixed bottom-6 right-6 z-[80]
        flex h-14 w-14 items-center justify-center
        rounded-full
        border border-white/10
        bg-[linear-gradient(135deg,rgba(139,92,246,0.95),rgba(56,189,248,0.95))]
        text-white
        shadow-[0_12px_40px_rgba(139,92,246,0.45)]
        backdrop-blur-xl
        transition-all duration-300
        hover:scale-[1.06]
        hover:shadow-[0_18px_60px_rgba(139,92,246,0.65)]
        active:scale-[0.96]
      "
    >
      {/* Glow ring */}
      <div className="absolute inset-0 rounded-full bg-violet-500/20 blur-xl opacity-0 transition group-hover:opacity-100" />

      {/* Icon */}
      <MessageSquareMore
        size={22}
        className="relative z-10 transition-transform duration-300 group-hover:scale-110"
      />

      {/* Pulse indicator (AI active feel) */}
      <span className="absolute right-2 top-2 flex h-3 w-3">
        <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-sky-400 opacity-75"></span>
        <span className="relative inline-flex h-3 w-3 rounded-full bg-sky-500"></span>
      </span>
    </button>
  );
}