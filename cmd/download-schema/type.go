package main

import (
	"log"
	"sort"
	"strings"

	"github.com/FreekingDean/proxmox-api-go/pkg/jsonschema"
)

func (p *Package) defineType(name string, jsonName string, schema *jsonschema.JSONSchema) *Type {
	if schema.Type == nil {
		if schema.Properties != nil {
			schema.Type = &jsonschema.SchemaOrString{String: "object"}
		} else {
			return nil
		}
	}

	if schema.Type.String == "null" && schema.Properties == nil {
		return nil
	}

	t := &Type{
		Type:               p.StrType(schema.Type),
		Optional:           bool(schema.Optional),
		Name:               name,
		JSONName:           jsonName,
		Properties:         make([]*Type, 0),
		OptionalProperties: make([]*Type, 0),
		Description:        removeNewline(schema.Description),
	}

	if t.Type == "struct" {
		if schema.Properties == nil || len(schema.Properties) == 0 {
			t.Type = "map[string]interface{}"
		} else {
			t.Type = name

			nt := p.defineStruct(name, schema.Description, schema.Properties)
			p.AddType(nt)
		}
	}

	if schema.Format != nil && len(schema.Format.Map) > 0 {
		if t.Type != "string" {
			log.Fatalf("Unkown type %s, expected (*)string", t.Type)
		}

		t.Type = name
		p.ImportURL = true
		p.Util = true
		nt := p.defineStruct(name, schema.Description, schema.Format.Map)
		nt.Format = schema.TypeText
		p.AddType(nt)
	}

	if t.Type == "[]" {
		if schema.Items == nil {
			t.Type = "[]map[string]interface{}"
		} else {
			nt := p.defineType(name, jsonName, schema.Items)
			t.Type = "[]" + nt.Type
		}
	}

	if strings.HasSuffix(t.Name, "[N]") {
		t.Name = strings.TrimSuffix(t.Name, "[N]")
		arrType := &Type{
			Type:               "[]" + strings.TrimSuffix(t.Type, "[N]"),
			Name:               t.Name + "s",
			Properties:         make([]*Type, 0),
			OptionalProperties: make([]*Type, 0),
			Description:        "Array of " + t.Name,
			Format:             "array",
		}
		t.Type = t.Name + "s"
		t.Name = t.Name + "s"
		p.AddType(arrType)
	}

	if strings.HasSuffix(t.Type, "[N]") {
		t.Type = strings.TrimSuffix(t.Type, "[N]") + "s"
	}

	return t
}

func removeNewline(in string) string {
	return strings.Replace(strings.Replace(in, "\n", " ", -1), "  ", " ", -1)
}

func (p *Package) defineStruct(name, description string, properties map[string]*jsonschema.JSONSchema) *Type {
	t := &Type{
		Name:               name,
		Type:               "struct",
		Properties:         make([]*Type, 0),
		OptionalProperties: make([]*Type, 0),
		Description:        removeNewline(description),
	}
	for name, param := range properties {
		pt := p.defineType(Nameify(name), name, param)
		if pt == nil {
			continue
		}
		if param.Optional {
			t.OptionalProperties = append(t.OptionalProperties, pt)
		} else {
			t.Properties = append(t.Properties, pt)
		}
	}
	sort.Sort(&TypeSorter{t.Properties})
	sort.Sort(&TypeSorter{t.OptionalProperties})
	t.Name = strings.TrimSuffix(t.Name, "[N]")
	return t
}
