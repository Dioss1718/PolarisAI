package server

import (
	"encoding/json"
	"net/http"

	"github.com/diya-suryawanshi/cloud/backend/orchestrator"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h(w, r)
	}
}

func Start(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", withCORS(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}))

	mux.HandleFunc("/api/state", withCORS(func(w http.ResponseWriter, r *http.Request) {
		state := orchestrator.GetLatestState()
		if state == nil {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "no pipeline run yet"})
			return
		}
		writeJSON(w, http.StatusOK, state)
	}))

	mux.HandleFunc("/api/run", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}

		var req orchestrator.RunRequest
		_ = json.NewDecoder(r.Body).Decode(&req)

		result, err := orchestrator.Run(req)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, result)
	}))

	return http.ListenAndServe(addr, mux)
}

func writeJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}
