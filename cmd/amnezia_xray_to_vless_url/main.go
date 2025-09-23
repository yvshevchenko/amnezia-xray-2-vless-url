package main

import (
	"cmd/amnezia_xray_to_vless_url/cmd/amnezia_xray_to_vless_url/internal"
	"log"
)

func main() {
	_, err := internal.ReadArg()

	if err != nil {
		log.Fatalf("Arguments reading error: %s\n", err)
	}
}
