package main

import (
	"encoding/json"
	"strings"

	"github.com/FreekingDean/proxmox-api-go/pkg/jsonschema"
)

func parseJS(body string) (string, error) {
	body = strings.TrimPrefix(body, "const apiSchema = ")
	body = strings.TrimSuffix(body, ";\n\n")
	return body, nil
}

func parseJSON(data []byte) ([]*jsonschema.Schema, error) {
	out := make([]*jsonschema.Schema, 0)
	err := json.Unmarshal(data, &out)
	return out, err
}
