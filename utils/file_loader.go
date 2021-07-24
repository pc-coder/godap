package utils

import (
	"encoding/json"
	"os"
)

func LoadJSONFile(path string, store interface{}) error {
	fileData, err := os.Open(path)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(fileData)
	err = decoder.Decode(&store)
	if err != nil {
		return err
	}
	return nil
}
