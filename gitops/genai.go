package gitops

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func GenerateInfraCode(diff Diff, d Decision) InfraCode {

	payload := map[string]interface{}{
		"node":    d.NodeID,
		"changes": diff.ChangeSet,
		"reason":  d.Reason,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"POST",
		os.Getenv("GENAI_ENDPOINT"),
		bytes.NewBuffer(body),
	)

	req.Header.Set("Authorization", "Bearer "+os.Getenv("GENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return InfraCode{Content: "", Format: "terraform"}
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	return InfraCode{
		Content: result["code"],
		Format:  "terraform",
	}
}