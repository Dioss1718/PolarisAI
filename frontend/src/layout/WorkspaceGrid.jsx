export default function WorkspaceGrid({ leftTop, leftBottom, rightTop, rightBottom }) {
  return (
    <div className="grid h-full min-h-0 grid-cols-[1.65fr_0.95fr] gap-3 overflow-hidden">
      <div className="grid min-h-0 grid-rows-[1fr_240px] gap-3">
        <div className="min-h-0">{leftTop}</div>
        <div className="min-h-0">{leftBottom}</div>
      </div>

      <div className="grid min-h-0 grid-rows-[1fr_290px] gap-3">
        <div className="min-h-0">{rightTop}</div>
        <div className="min-h-0">{rightBottom}</div>
      </div>
    </div>
  );
}