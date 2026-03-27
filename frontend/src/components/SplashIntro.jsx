import React, { useEffect, useState } from "react";

export default function SplashIntro({ onEnter }) {
  const [visible, setVisible] = useState(false);
  const [expand, setExpand] = useState(false);

  useEffect(() => {
    const t1 = setTimeout(() => setVisible(true), 300);
    const t2 = setTimeout(() => setExpand(true), 1800);
    const t3 = setTimeout(() => onEnter(), 3500);

    return () => {
      clearTimeout(t1);
      clearTimeout(t2);
      clearTimeout(t3);
    };
  }, [onEnter]);

  return (
    <div className="relative flex min-h-screen items-center justify-center overflow-hidden bg-[#020617]">
      <div className="absolute inset-0 bg-[radial-gradient(circle_at_top_left,rgba(56,189,248,0.16),transparent_30%),radial-gradient(circle_at_bottom_right,rgba(168,85,247,0.14),transparent_28%),radial-gradient(circle_at_center,rgba(255,255,255,0.03),transparent_45%)]" />

      <div
        className={`relative transition-all duration-[1400ms] ease-out ${
          expand ? "scale-[1.85] opacity-0 blur-xl" : "scale-100 opacity-100"
        }`}
      >
        <div
          className={`glass-logo-shell transition-all duration-700 ${
            visible ? "translate-y-0 opacity-100" : "translate-y-8 opacity-0"
          }`}
        >
          <div className="logo-orb">
            <div className="logo-ring" />
            <div className="logo-core">
              <svg
                width="90"
                height="90"
                viewBox="0 0 100 100"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M50 10L66 42L50 34L34 42L50 10Z"
                  fill="url(#grad1)"
                />
                <path
                  d="M50 90L34 58L50 66L66 58L50 90Z"
                  fill="url(#grad1)"
                />
                <circle cx="50" cy="50" r="10" fill="url(#grad2)" />
                <defs>
                  <linearGradient id="grad1" x1="20" y1="20" x2="80" y2="80">
                    <stop stopColor="#7dd3fc" />
                    <stop offset="1" stopColor="#a78bfa" />
                  </linearGradient>
                  <linearGradient id="grad2" x1="40" y1="40" x2="60" y2="60">
                    <stop stopColor="#ffffff" />
                    <stop offset="1" stopColor="#7dd3fc" />
                  </linearGradient>
                </defs>
              </svg>
            </div>
          </div>

          <div className="mt-8 text-center">
            <div className="text-xs uppercase tracking-[0.45em] text-sky-200/70">
              Autonomous Cloud Governance
            </div>
            <h1 className="mt-3 text-5xl font-semibold tracking-tight text-white">
              PolarisAI
            </h1>
            <p className="mt-3 text-sm text-slate-300">
              Graph intelligence · policy validation · explainable remediation
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}