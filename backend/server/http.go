package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/diya-suryawanshi/cloud/backend/auth"
	"github.com/diya-suryawanshi/cloud/backend/orchestrator"
	"github.com/diya-suryawanshi/cloud/rbac"
)

type contextKey string

const sessionKey contextKey = "session"

var authService = auth.NewService()
var runtimeState = newRuntimeManager()

type loginPayload struct {
	EmployeeIDAlt1 string `json:"employeeId"`
	EmployeeIDAlt2 string `json:"employeeID"`
	Password       string `json:"password"`
}

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h(w, r)
	}
}

func requireAuth(next func(http.ResponseWriter, *http.Request, auth.Session)) http.HandlerFunc {
	return withCORS(func(w http.ResponseWriter, r *http.Request) {
		token := auth.ExtractBearerToken(r)
		if token == "" {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing bearer token"})
			return
		}

		session, ok := authService.Resolve(token)
		if !ok {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid session"})
			return
		}

		ctx := context.WithValue(r.Context(), sessionKey, session)
		next(w, r.WithContext(ctx), session)
	})
}

func Start(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", withCORS(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"status": "ok",
			"services": []map[string]string{
				{"name": "Governance API", "status": "up"},
				{"name": "Simulation Engine", "status": "up"},
				{"name": "AI Engine", "status": "up"},
				{"name": "Forecast Engine", "status": "up"},
				{"name": "GitOps Engine", "status": "up"},
			},
		})
	}))

	mux.HandleFunc("/api/login", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}

		var req loginPayload
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}

		employeeID := strings.TrimSpace(req.EmployeeIDAlt1)
		if employeeID == "" {
			employeeID = strings.TrimSpace(req.EmployeeIDAlt2)
		}
		employeeID = strings.ToUpper(employeeID)

		password := strings.TrimSpace(req.Password)

		if employeeID == "" || password == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "employee ID and password are required"})
			return
		}

		session, err := authService.Login(employeeID, password)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, session)
	}))

	mux.HandleFunc("/api/me", requireAuth(func(w http.ResponseWriter, r *http.Request, session auth.Session) {
		writeJSON(w, http.StatusOK, session)
	}))

	mux.HandleFunc("/api/state", requireAuth(func(w http.ResponseWriter, r *http.Request, session auth.Session) {
		scenario := r.URL.Query().Get("scenario")
		if scenario == "" {
			scenario = "FULL_CHAOS"
		}

		seed := 42
		if s := r.URL.Query().Get("seed"); s != "" {
			if parsed, err := strconv.Atoi(s); err == nil {
				seed = parsed
			}
		}

		result := runtimeState.latestFor(scenario, seed)
		if result == nil {
			writeJSON(w, http.StatusNoContent, map[string]string{"status": "no-run-yet"})
			return
		}

		overlay := runtimeState.overlayGitOps(result, scenario, seed)
		writeJSON(w, http.StatusOK, filterPipelineForRole(overlay, session.Role))
	}))

	mux.HandleFunc("/api/run", requireAuth(func(w http.ResponseWriter, r *http.Request, session auth.Session) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}

		if !rbac.CanRunGovernance(session.Role) {
			writeJSON(w, http.StatusForbidden, map[string]string{"error": "role not allowed to run governance"})
			return
		}

		var req orchestrator.RunRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}

		if req.ManualData != nil && !rbac.CanAccess(session.Role, rbac.FeatureSimulationStudio) {
			writeJSON(w, http.StatusForbidden, map[string]string{"error": "role not allowed to use simulation studio"})
			return
		}

		result, err := orchestrator.Run(req)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		scenario := req.Scenario
		if scenario == "" {
			scenario = "FULL_CHAOS"
		}

		seed := req.Seed
		if seed == 0 {
			seed = 42
		}

		runtimeState.setLatest(result, scenario, seed)
		runtimeState.seedApprovalsFromPipeline(result, scenario, seed)

		overlay := runtimeState.overlayGitOps(result, scenario, seed)
		writeJSON(w, http.StatusOK, filterPipelineForRole(overlay, session.Role))
	}))

	mux.HandleFunc("/api/copilot", requireAuth(func(w http.ResponseWriter, r *http.Request, session auth.Session) {
		handleCopilot(w, r, session)
	}))

	mux.HandleFunc("/api/gitops/approve", requireAuth(func(w http.ResponseWriter, r *http.Request, session auth.Session) {
		handleGitOpsApprove(w, r, session)
	}))

	mux.HandleFunc("/api/gitops/reject", requireAuth(func(w http.ResponseWriter, r *http.Request, session auth.Session) {
		handleGitOpsReject(w, r, session)
	}))

	mux.HandleFunc("/api/gitops/audit", requireAuth(func(w http.ResponseWriter, r *http.Request, session auth.Session) {
		handleGitOpsAudit(w, r, session)
	}))

	mux.HandleFunc("/api/gitops/refresh-pr", requireAuth(func(w http.ResponseWriter, r *http.Request, session auth.Session) {
		handleGitOpsRefreshPR(w, r, session)
	}))
	mux.HandleFunc("/api/compare", requireAuth(func(w http.ResponseWriter, r *http.Request, session auth.Session) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}

		if !rbac.CanRunGovernance(session.Role) {
			writeJSON(w, http.StatusForbidden, map[string]string{"error": "role not allowed to run what-if comparison"})
			return
		}

		var req compareRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}

		if req.ManualData != nil && !rbac.CanAccess(session.Role, rbac.FeatureSimulationStudio) {
			writeJSON(w, http.StatusForbidden, map[string]string{"error": "role not allowed to use simulation studio"})
			return
		}

		result, err := runWhatIfCompare(req)
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
