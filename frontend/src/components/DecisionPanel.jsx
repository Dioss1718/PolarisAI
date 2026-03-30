import RecommendationList from "./RecommendationList";
import ExplainabilityList from "./ExplainabilityList";

export default function DecisionPanel({ recommendations, explanations, onSelectRecommendation, onSelectNode }) {
  return (
    <div className="grid min-h-0 grid-rows-[1fr_1fr] gap-3">
      <RecommendationList recommendations={recommendations} onSelect={onSelectRecommendation} />
      <ExplainabilityList explanations={explanations} onSelect={onSelectNode} />
    </div>
  );
}