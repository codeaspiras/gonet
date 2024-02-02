package utils

import (
	"encoding/json"
	"os"
)

func GetJSON(path string, dest interface{}) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, dest)
	if err != nil {
		panic(err)
	}
}
