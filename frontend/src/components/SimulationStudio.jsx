import { useState } from "react";

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
      compliance_flags: ["PCI"],
    },
  ],
  edges: [],
};

export default function SimulationStudio({ onRunManual, loading, allowed = true }) {
  const [manualJSON, setManualJSON] = useState(JSON.stringify(sampleManual, null, 2));

  const submit = () => {
    try {
      const parsed = JSON.parse(manualJSON);
      onRunManual(parsed);
    } catch {
      window.alert("Manual simulation JSON is invalid.");
    }
  };

  if (!allowed) {
    return (
      <div className="rounded-2xl border border-white/10 bg-white/[0.04] p-4">
        <div className="text-lg font-semibold">Simulation Studio</div>
        <div className="mt-2 text-sm text-slate-400">
          Your role does not have access to simulation studio.
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-0 overflow-hidden rounded-2xl border border-white/10 bg-white/[0.04] p-4">
      <div className="text-lg font-semibold">Simulation Studio</div>
      <p className="mt-1 text-sm text-slate-400">
        Inject manual cloud graph state directly into the pipeline.
      </p>

      <textarea
        value={manualJSON}
        onChange={(e) => setManualJSON(e.target.value)}
        className="mt-4 h-[calc(100%-110px)] w-full rounded-xl border border-white/10 bg-slate-950/80 p-4 text-xs outline-none"
      />

      <button
        onClick={submit}
        disabled={loading}
        className="mt-3 rounded-xl border border-sky-500/35 bg-sky-500/10 px-4 py-2 text-sm text-sky-300 disabled:opacity-50"
      >
        Run Manual Simulation
      </button>
    </div>
  );
}