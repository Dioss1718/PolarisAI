package gitops

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

// This function generates infrastructure code based on the diff and decision
func GenerateInfraCode(diff Diff, d Decision) InfraCode {

	// Prepare payload to send to GenAI API
	payload := map[string]interface{}{
		"node":    d.NodeID,
		"changes": diff.ChangeSet,
		"reason":  d.Reason,
	}

	// Convert payload into JSON format
	body, _ := json.Marshal(payload)

	// Create HTTP POST request to GenAI endpoint
	req, _ := http.NewRequest(
		"POST",
		os.Getenv("GENAI_ENDPOINT"),
		bytes.NewBuffer(body),
	)

	// Add authorization and content type headers
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	// Create HTTP client and send request
	client := &http.Client{}
	resp, err := client.Do(req)

	// If request fails, return empty infra code with default format
	if err != nil {
		return InfraCode{Content: "", Format: "terraform"}
	}
	defer resp.Body.Close()

	// Decode response from API
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	// Return generated code and its format
	return InfraCode{
		Content: result["code"],
		Format:  "terraform",
	}
}
