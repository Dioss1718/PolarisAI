import { useState } from "react";
import { FileJson, Wand2 } from "lucide-react";

const sampleManual = {
  nodes: [
    {
      id: "manual_vm1",
      type: "COMPUTE",
      name: "Manual VM",
      cloud_provider: "AWS",
      region: "ap-south-1",
      environment: "PROD",
      cost: 120,
      utilization: 18,
      exposure: "PUBLIC",
      criticality: 8,
      compliance_flags: ["PCI"]
    }
  ],
  edges: []
};

export default function SimulationStudio({ onRunManual, loading }) {
  const [manualJSON, setManualJSON] = useState(JSON.stringify(sampleManual, null, 2));

  const submitManual = () => {
    try {
      const parsed = JSON.parse(manualJSON);
      onRunManual(parsed);
    } catch {
      alert("Manual simulation JSON is invalid.");
    }
  };

  return (
    <div className="rounded-2xl border border-borderSoft bg-panel/90 shadow-glow p-4">
      <div className="flex items-center gap-2 text-sky-300">
        <Wand2 size={18} />
        <h2 className="text-lg font-semibold">Simulation Studio</h2>
      </div>
      <p className="mt-2 text-sm text-slate-400">
        Run scenario-based synthetic simulation or manually inject a custom cloud graph for operator what-if testing.
      </p>

      <div className="mt-4">
        <label className="mb-2 flex items-center gap-2 text-sm text-slate-300">
          <FileJson size={16} />
          Manual synthetic data
        </label>
        <textarea
          value={manualJSON}
          onChange={(e) => setManualJSON(e.target.value)}
          className="h-64 w-full rounded-xl border border-borderSoft bg-slate-950/70 p-4 text-xs text-slate-200 outline-none"
        />
        <button
          onClick={submitManual}
          disabled={loading}
          className="mt-3 rounded-xl border border-sky-500/40 bg-sky-500/10 px-4 py-2 text-sm text-sky-300 hover:bg-sky-500/20 disabled:opacity-50"
        >
          Run Manual Simulation
        </button>
      </div>
    </div>
  );
}