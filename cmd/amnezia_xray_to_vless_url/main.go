package main

import (
	"cmd/amnezia_xray_to_vless_url/cmd/amnezia_xray_to_vless_url/internal"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// reading ARGS, validating and returning first one
	fileName, err := internal.ReadArg()

	if err != nil {
		log.Fatalf("Error arguments reading: %s", err)
	}

	// opening the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// reading the file contents
	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var cfg internal.SourceConfig

	//unmarshalling file contents to struct
	err = json.Unmarshal(fileContent, &cfg)
	if err != nil {
		log.Fatalf("Error unmarshalling file contents: %v", err)
	}

	// printing out connection url
	fmt.Println(cfg.ComposeConnectionUrl())
}
