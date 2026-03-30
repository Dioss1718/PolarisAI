import axios from "axios";

const TOKEN_KEY = "polaris_token";
const SESSION_KEY = "polaris_session";

export function getToken() {
  return localStorage.getItem(TOKEN_KEY) || "";
}

export function setSession(session) {
  if (!session?.token) {
    throw new Error("Invalid session payload: missing token");
  }
  localStorage.setItem(TOKEN_KEY, session.token);
  localStorage.setItem(SESSION_KEY, JSON.stringify(session));
}

export function getSession() {
  const raw = localStorage.getItem(SESSION_KEY);
  if (!raw) return null;
  try {
    return JSON.parse(raw);
  } catch {
    return null;
  }
}

export function clearSession() {
  localStorage.removeItem(TOKEN_KEY);
  localStorage.removeItem(SESSION_KEY);
}

export const api = axios.create({
  baseURL: "http://127.0.0.1:8080/api",
  timeout: 300000,
});

api.interceptors.request.use((config) => {
  const token = getToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export async function login(employeeId, password) {
  const normalizedEmployeeId = String(employeeId || "").trim().toUpperCase();
  const normalizedPassword = String(password || "").trim();

  const payload = {
    employeeId: normalizedEmployeeId,
    employeeID: normalizedEmployeeId,
    password: normalizedPassword,
  };

  const { data } = await api.post("/login", payload);
  setSession(data);
  return data;
}

export async function getMe() {
  const { data } = await api.get("/me");
  return data;
}

export async function runPipeline(payload) {
  const { data } = await api.post("/run", payload);
  return data;
}

export async function getState({ scenario, seed }) {
  const params = new URLSearchParams();
  if (scenario) params.set("scenario", scenario);
  if (seed !== undefined && seed !== null) params.set("seed", String(seed));
  const { data } = await api.get(`/state?${params.toString()}`);
  return data;
}

export async function getServiceHealth() {
  const { data } = await api.get("/health");
  return data;
}