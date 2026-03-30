export default function AppShell({ children }) {
  return (
    <div className="bg-app h-screen w-screen overflow-hidden text-slate-100">
      <div className="mx-auto h-full max-w-[1800px] px-4 py-4">
        {children}
      </div>
    </div>
  );
}