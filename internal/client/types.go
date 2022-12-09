package client

import (
	"encoding/json"
	"github.com/go-faster/errors"
	"os"
)

type JobConfigJson struct {
	Cmd string `json:"cmd,omitempty"`
}

// readJobConfigJson will read the JobConfigJson from a file
func readJobConfigJson(path string) (JobConfigJson, error) {
	file, err := os.Open(path)
	if err != nil {
		return JobConfigJson{}, errors.Wrap(err, "cannot open file")
	}
	defer file.Close()
	var result JobConfigJson
	err = json.NewDecoder(file).Decode(&result)
	if err != nil {
		return JobConfigJson{}, errors.Wrap(err, "cannot decode json")
	}
	return result, nil
}
