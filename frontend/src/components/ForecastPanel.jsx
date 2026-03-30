import BillShockTable from "./BillShockTable";

export default function ForecastPanel({ forecasts, onSelect }) {
  return <BillShockTable items={forecasts} onSelect={onSelect} />;
}