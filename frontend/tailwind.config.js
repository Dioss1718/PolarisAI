/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,jsx}"],
  theme: {
    extend: {
      colors: {
        panel: "#0f172a",
        panelSoft: "#111827",
        borderSoft: "#1f2937",
        accent: "#38bdf8"
      },
      boxShadow: {
        glow: "0 0 0 1px rgba(56,189,248,0.15), 0 8px 32px rgba(2,6,23,0.45)"
      }
    }
  },
  plugins: []
}