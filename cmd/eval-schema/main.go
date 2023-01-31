package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/FreekingDean/proxmox-api-go/pkg/jsonschema"
)

func main() {
	data, err := os.ReadFile("scripts/schema.json")
	if err != nil {
		log.Fatal(err)
	}

	out := make([]*jsonschema.Schema, 0)
	err = json.Unmarshal(data, &out)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range out {
		checkSchemaType(s)
	}
}

func checkSchemaType(s *jsonschema.Schema) {
	for _, c := range s.Children {
		checkSchemaType(c)
	}

	for _, i := range s.Info {
		if i.Parameters == nil {
			log.Printf("Path: %s Method %s No Parameters\n", i.Name, i.Method)
		} else {
			checkJSONSchemaType(i.Name+"Req", i.Parameters)
		}
		if i.Returns == nil {
			log.Printf("Path: %s Method %s No Returns\n", i.Name, i.Method)
		} else {
			checkJSONSchemaType(i.Name+"Resp", i.Returns)
		}
	}
}

func checkJSONSchemaType(name string, s *jsonschema.JSONSchema) {
	for k, v := range s.Properties {
		checkJSONSchemaType(k, v)
	}
	if s.Type != nil && s.Properties != nil && s.Type.String != "object" {
		log.Printf("Nil type: %s: \n%+v\n%+v\n\n", name, s, s.AdditionalProperties)
	}
}
