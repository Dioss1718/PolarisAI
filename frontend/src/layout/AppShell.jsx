export default function AppShell({ children }) {
  return (
    <div className="min-h-screen w-full bg-[#050816] text-white">
      <div className="mx-auto w-full max-w-[1920px] px-4 py-4">
        {children}
      </div>
    </div>
  );
}