package main

import (
	"encoding/json"
	"fmt"

	"github.com/FreekingDean/proxmox-api-go/pkg/jsonschema"
	"github.com/dop251/goja"
	"github.com/dop251/goja/ast"
	"github.com/dop251/goja/token"
)

func parseJS(body string) (string, error) {
	prog, err := goja.Parse("apiextract.js", body)
	if err != nil {
		return "", fmt.Errorf("parse err: %w", err)
	}

	for _, n := range prog.Body {
		if decl, ok := n.(*ast.LexicalDeclaration); ok {
			if decl.Token != token.CONST {
				continue
			}
			for _, n2 := range decl.List {
				if id, ok := n2.Target.(*ast.Identifier); ok {
					if id.Name != "apiSchema" {
						continue
					}
				}
				start := decl.List[0].Initializer.Idx0()
				end := decl.List[0].Initializer.Idx1()
				json := prog.File.Source()[start-1 : end]
				return json, nil
			}
		}
	}
	return "", fmt.Errorf("couldnt find apiSchema const")
}

func parseJSON(data []byte) ([]*jsonschema.Schema, error) {
	out := make([]*jsonschema.Schema, 0)
	err := json.Unmarshal(data, &out)
	return out, err
}
