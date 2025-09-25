package main

import (
	"cmd/amnezia_xray_to_vless_url/internal"
	"fmt"
	"log"
)

func main() {
	// reading ARGS, validating and returning first one
	fileName, err := internal.ReadArg()

	if err != nil {
		log.Fatalf("Error arguments reading: %s", err.Error())
	}

	cfg, err := internal.ReadConfig(fileName)
	if err != nil {
		log.Fatalf("Error reading config: %s", err.Error())
	}

	// printing out connection url
	fmt.Println(cfg.ComposeConnectionUrl())
}
