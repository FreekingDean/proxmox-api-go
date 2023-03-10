package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/FreekingDean/proxmox-api-go/pkg/jsonschema"
)

type Type struct {
	Properties         []*Type
	OptionalProperties []*Type
	FormatArrayTypes   []*Type
	ElemType           *Type
	Optional           bool
	Name               string
	JSONName           string
	Type               string
	Description        string
	Format             string
	Enum               []string
}

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
			Type:               "[]*" + strings.TrimSuffix(t.Type, "[N]"),
			Name:               t.Name + "s",
			Properties:         make([]*Type, 0),
			OptionalProperties: make([]*Type, 0),
			Description:        "Array of " + t.Name,
			Format:             "array",
		}
		t.Type = t.Name + "s"
		t.ElemType = arrType
		t.Name = t.Name + "s"
		p.AddType(arrType)
	}

	if strings.HasSuffix(t.Type, "[N]") {
		t.Type = strings.TrimSuffix(t.Type, "[N]") + "s"
	}

	if len(schema.Enum) > 0 {
		t.Enum = schema.Enum
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
		FormatArrayTypes:   make([]*Type, 0),
	}
	for name, param := range properties {
		pt := p.defineType(Nameify(name), name, param)
		if pt == nil {
			continue
		}
		if strings.HasSuffix(name, "[n]") {
			p.Strconv = true
			t.FormatArrayTypes = append(
				t.FormatArrayTypes,
				pt,
			)
		}
		if param.Optional {
			t.OptionalProperties = append(t.OptionalProperties, pt)
		} else {
			t.Properties = append(t.Properties, pt)
		}
		if len(pt.Enum) > 0 {
			prefix := ""
			if !strings.HasSuffix(t.Name, "Request") && !strings.HasSuffix(t.Name, "Response") {
				prefix = t.Name
				prefix = strings.TrimSuffix(prefix, "[N]")
			}
			pt.Type = fmt.Sprintf("%s%s", prefix, Nameify(name))
			p.AddEnum(pt.Type, pt.Enum)
		}
	}
	sort.Sort(&TypeSorter{t.Properties})
	sort.Sort(&TypeSorter{t.OptionalProperties})
	sort.Sort(&TypeSorter{t.FormatArrayTypes})
	t.Name = strings.TrimSuffix(t.Name, "[N]")
	return t
}
