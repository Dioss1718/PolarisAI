import { api, aiApi } from "./client";

export const runPipeline = () => api.get("/run");

export const getGraph = () => api.get("/graph");

export const getRecommendations = () => api.get("/recommendations");

export const getForecast = (node) =>
  api.get(`/forecast?node=${node}`);

export const explain = (payload) =>
  aiApi.post("/explain", payload);

export const generateInfra = (payload) =>
  aiApi.post("/infra", payload);