package main

import (
	"log"
)

const (
	apijs = "https://github.com/proxmox/pve-docs/raw/ab8ba11aa85c2a189a882338033ee288230cca93/api-viewer/apidata.js"
)

func main() {
	data, err := download()
	if err != nil {
		log.Fatal(err)
	}

	json, err := parseJS(data)
	//json, err := os.ReadFile("scripts/schema.json")
	if err != nil {
		log.Fatal(err)
	}
	str, err := parseJSON([]byte(json))
	if err != nil {
		log.Fatal(err)
	}

	err = gen(str)
	if err != nil {
		log.Fatal(err)
	}
}
