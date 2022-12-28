package main

import (
	"log"
	"os"
)

const (
	apijs = "https://github.com/proxmox/pve-docs/raw/master/api-viewer/apidata.js"
)

func main() {
	//data, err := download()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//json, err := parseJS(data)
	json, err := os.ReadFile("scripts/schema.json")
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
