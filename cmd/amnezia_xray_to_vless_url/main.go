package main

import (
	"cmd/amnezia_xray_to_vless_url/internal"
	"fmt"
	"log"
)

func main() {
	// reading ARGS, validating and returning first one
	fileName, err := internal.ReadFileNameArg()
	fileName, err := internal.ReadFileNameArg()

	if err != nil {
		log.Fatalf("Error arguments reading: %s", err.Error())
	}

	// creating new instance of cfg struct
	cfg, err := internal.NewSourceConfig(fileName)
	if err != nil {
		log.Fatalf("Error reading config: %s", err.Error())
	}

	// validating
	if !cfg.ConfigIsValid() {
		log.Fatalf("Error validating config: %+v", cfg)
	}

	// printing out connection url
	fmt.Println(cfg.ComposeConnectionUrl())
}
