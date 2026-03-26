package gitops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type infraRequest struct {
	NodeID  string   `json:"node_id"`
	Action  string   `json:"action"`
	Reason  string   `json:"reason"`
	Changes []string `json:"changes"`
	Format  string   `json:"format"`
}

type infraResponse struct {
	Code     string `json:"code"`
	Format   string `json:"format"`
	Title    string `json:"title,omitempty"`
	Summary  string `json:"summary,omitempty"`
	Grounded bool   `json:"grounded,omitempty"`
}

var infraHTTPClient = &http.Client{
	Timeout: 10 * time.Second,
}

func GenerateInfraCode(diff Diff, d Decision) InfraCode {
	reqBody := infraRequest{
		NodeID:  d.NodeID,
		Action:  d.FinalAction,
		Reason:  d.Reason,
		Changes: diff.ChangeSet,
		Format:  "terraform",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return fallbackInfraCode(d)
	}

	resp, err := infraHTTPClient.Post(
		"http://localhost:8000/infra",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return fallbackInfraCode(d)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fallbackInfraCode(d)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fallbackInfraCode(d)
	}

	var out infraResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return fallbackInfraCode(d)
	}

	if out.Code == "" {
		return fallbackInfraCode(d)
	}

	format := out.Format
	if format == "" {
		format = "terraform"
	}

	return InfraCode{
		Content: out.Code,
		Format:  format,
	}
}

func fallbackInfraCode(d Decision) InfraCode {
	return InfraCode{
		Content: fmt.Sprintf(`
resource "null_resource" "%s" {
  provisioner "local-exec" {
    command = "echo applied %s"
  }
}
`, d.NodeID, d.FinalAction),
		Format: "terraform",
	}
}
