package internal

import (
	"encoding/json"
	"io"
	"os"
)

func ReadConfig(fileName string) (SourceConfig, error) {
	var cfg SourceConfig

	// opening the file
	file, err := os.Open(fileName)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	// reading the file contents
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return cfg, err
	}

	//unmarshalling file contents to struct
	err = json.Unmarshal(fileContent, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, err
}
