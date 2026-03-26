import axios from "axios";

export const api = axios.create({
  baseURL: "http://127.0.0.1:8080/api",
  timeout: 300000,
});

export async function runPipeline(payload) {
  const { data } = await api.post("/run", payload);
  return data;
}

export async function getState() {
  const { data } = await api.get("/state");
  return data;
}

export async function health() {
  const { data } = await api.get("/health");
  return data;
}