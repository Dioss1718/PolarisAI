package feedback

import (
	"encoding/json"
	"os"
)

const file = "backend/feedback/data.json"

func Load() []Record {
	data, err := os.ReadFile(file)
	if err != nil {
		return []Record{}
	}

	var records []Record
	json.Unmarshal(data, &records)

	return records
}

func Save(records []Record) {
	data, _ := json.MarshalIndent(records, "", "  ")
	os.WriteFile(file, data, 0644)
}
