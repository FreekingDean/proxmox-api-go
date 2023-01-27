package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/FreekingDean/proxmox-api-go/pkg/jsonschema"
)

var caser = cases.Title(language.English)

func gen(schema []*jsonschema.Schema) error {
	for _, domain := range schema {
		err := LoadPackage("proxmox", domain)
		if err != nil {
			return err
		}
	}
	return nil
}

var (
	keys []string = []string{"GET", "POST", "PUT", "DELETE"}
)

func LoadPackage(curdir string, s *jsonschema.Schema) error {
	packageName := PackageNameify(Nameify(s.Text))
	dir := curdir + "/" + packageName
	err := os.MkdirAll(dir, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	f, err := os.OpenFile(dir+"/"+packageName+".go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	if err != nil {
		return err
	}
	b := &bytes.Buffer{}
	p := &Package{
		Name:    packageName,
		Methods: make([]*OperationTempl, 0),
		Types:   make([]*Type, 0),
	}
	if s.Info["GET"] != nil {
		p.Methods = append(
			p.Methods,
			p.genMethod("Index", s.Info["GET"], s.Path),
		)
	}
	if s.Info["POST"] != nil {
		p.Methods = append(
			p.Methods,
			p.genMethod("Create", s.Info["POST"], s.Path),
		)
	}
	if s.Info["PUT"] != nil {
		p.Methods = append(
			p.Methods,
			p.genMethod("MassUpdate", s.Info["PUT"], s.Path),
		)
	}
	if s.Info["DELETE"] != nil {
		p.Methods = append(
			p.Methods,
			p.genMethod("MassDelete", s.Info["DELETE"], s.Path),
		)
	}
	for _, c := range s.Children {
		if c.Text[0] == '{' && c.Text[len(c.Text)-1] == '}' {
			if c.Info["GET"] != nil {
				p.Methods = append(
					p.Methods,
					p.genMethod("Find", c.Info["GET"], c.Path),
				)
			}
			if c.Info["POST"] != nil {
				p.Methods = append(
					p.Methods,
					p.genMethod("ChildCreate", c.Info["POST"], c.Path),
				)
			}
			if c.Info["PUT"] != nil {
				p.Methods = append(
					p.Methods,
					p.genMethod("Update", c.Info["PUT"], c.Path),
				)
			}
			if c.Info["DELETE"] != nil {
				p.Methods = append(
					p.Methods,
					p.genMethod("Delete", c.Info["DELETE"], c.Path),
				)
			}
			for _, cc := range c.Children {
				if len(cc.Children) > 0 {
					err := LoadPackage(dir, cc)
					if err != nil {
						return err
					}
				} else {
					for _, key := range keys {
						if info, ok := cc.Info[key]; ok {
							methodName := Nameify(info.Name)
							textName := Nameify(cc.Text)
							if !strings.HasSuffix(methodName, textName) {
								methodName += textName
							}
							p.Methods = append(
								p.Methods,
								p.genMethod(methodName, info, cc.Path),
							)
						}
					}
				}
			}
		} else {
			if len(c.Children) > 0 {
				err := LoadPackage(dir, c)
				if err != nil {
					return err
				}
			} else {
				for _, key := range keys {
					if info, ok := c.Info[key]; ok {
						methodName := Nameify(info.Name)
						textName := Nameify(c.Text)
						if !strings.HasSuffix(methodName, textName) {
							methodName += textName
						}
						p.Methods = append(
							p.Methods,
							p.genMethod(methodName, info, c.Path),
						)
					}
				}
			}
		}
	}
	t, err := template.ParseFiles("cmd/download-schema/templates/package.go.tmpl")
	if err != nil {
		return err
	}
	err = t.Execute(b, p)
	if err != nil {
		return err
	}
	data := b.String()
	err = Format(f, data)
	var line1 int
	if err != nil {
		lineParts := strings.Split(err.Error(), ":")
		if len(lineParts) >= 1 {
			line1, _ = strconv.Atoi(lineParts[0])
		}
		return fmt.Errorf("Error in: %s %w\nLines: \n%s\n%s\n%s", p.Name, err,
			strings.Split(data, "\n")[line1-1],
			strings.Split(data, "\n")[line1],
			strings.Split(data, "\n")[line1+1],
		)
	}
	return nil
}

func (p *Package) genMethod(operation string, info *jsonschema.InfoSchema, path string) *OperationTempl {
	ot := &OperationTempl{
		Operation:   operation,
		Description: info.Description,
		Method:      info.Method,
		Path:        path,
	}
	if info.Parameters != nil {
		ot.Request = p.defineType(
			fmt.Sprintf("%sRequest", operation),
			"",
			info.Parameters,
		)
	}
	if info.Returns != nil {
		ot.Response = p.defineType(
			fmt.Sprintf("%sResponse", operation),
			"",
			info.Returns,
		)
	}
	return ot
}

type Package struct {
	Name      string
	Util      bool
	Methods   []*OperationTempl
	Types     []*Type
	ImportURL bool
}

type OperationTempl struct {
	Operation   string
	Description string
	Method      string
	Path        string
	Request     *Type
	Response    *Type
}

type Type struct {
	Properties         []*Type
	OptionalProperties []*Type
	Name               string
	JSONName           string
	Type               string
	Description        string
	Format             string
}

func (p *Package) defineType(name string, jsonName string, schema *jsonschema.JSONSchema) *Type {
	if schema.Type == nil {
		if schema.Properties != nil {
			schema.Type = &jsonschema.SchemaOrString{String: "object"}
		} else {
			return nil
		}
	}
	t := &Type{
		Type:               p.StrType(schema.Type, bool(schema.Optional)),
		Name:               name,
		JSONName:           jsonName,
		Properties:         make([]*Type, 0),
		OptionalProperties: make([]*Type, 0),
		Description: strings.Replace(strings.Replace(schema.Description,
			"\n", " ", -1),
			"  ", " ", -1),
	}
	if t.Type == "struct" {
		if schema.Properties == nil || len(schema.Properties) == 0 {
			t.Type = "map[string]interface{}"
			return t
		}
	}
	if t.Type == "null" && schema.Properties == nil {
		return nil
	}
	if schema.Properties != nil && len(schema.Properties) > 0 {
		t.Type = name
		nt := &Type{
			Name:               name,
			Type:               "struct",
			Properties:         make([]*Type, 0),
			OptionalProperties: make([]*Type, 0),
			Description: strings.Replace(strings.Replace(schema.Description,
				"\n", " ", -1),
				"  ", " ", -1),
		}
		for name, param := range schema.Properties {
			pt := p.defineType(Nameify(name), name, param)
			if pt == nil {
				continue
			}
			if param.Optional {
				nt.OptionalProperties = append(nt.OptionalProperties, pt)
			} else {
				nt.Properties = append(nt.Properties, pt)
			}
		}
		sort.Sort(&TypeSorter{nt.Properties})
		sort.Sort(&TypeSorter{nt.OptionalProperties})
		p.AddType(nt)
	}
	if schema.Format != nil && len(schema.Format.Map) > 0 {
		if t.Type == "string" {
			t.Type = name
		} else if t.Type == "*string" {
			t.Type = "*" + name
		} else {
			log.Fatalf("Unkown type %s, expected (*)string", t.Type)
		}
		p.ImportURL = true
		p.Util = true
		nt := &Type{
			Type:               "[]" + name,
			Name:               name + "Arr",
			Properties:         make([]*Type, 0),
			OptionalProperties: make([]*Type, 0),
			Description:        "Array of " + name,
			Format:             "array",
		}
		p.AddType(nt)
		nt = &Type{
			Type:               "struct",
			Name:               name,
			Properties:         make([]*Type, 0),
			OptionalProperties: make([]*Type, 0),
			Description: strings.Replace(strings.Replace(schema.Description,
				"\n", " ", -1),
				"  ", " ", -1),
		}
		nt.Format = schema.TypeText
		if strings.HasSuffix(jsonName, "[n]") {
			t.Name = t.Name[0:len(t.Name)-1] + "s"
			t.Type = t.Type + "Arr"
		}
		for name, param := range schema.Format.Map {
			pt := p.defineType(Nameify(name), name, param)
			if pt == nil {
				continue
			}
			if param.Optional {
				nt.OptionalProperties = append(nt.OptionalProperties, pt)
			} else {
				nt.Properties = append(nt.Properties, pt)
			}
		}
		sort.Sort(&TypeSorter{nt.Properties})
		sort.Sort(&TypeSorter{nt.OptionalProperties})
		p.AddType(nt)
	}
	if schema.Items == nil && schema.Type.String == "array" {
		t.Type = "[]map[string]interface{}"
	}
	if schema.Items != nil {
		nt := p.defineType(name, jsonName, schema.Items)
		t.Type = "[]" + nt.Type
	}
	return t
}

func (p *Package) AddType(t *Type) {
	if t == nil {
		panic("Invalid nil type")
	}
	for _, ft := range p.Types {
		if ft == nil {
			panic("Invalid nil found_type")
		}
		if t.Name == ft.Name && t.Type == ft.Type {
			if !t.Eq(ft) {
				t.Name = "Sub" + t.Name
			} else {
				return
			}
		}
	}
	p.Types = append(p.Types, t)
}

func (t *Type) Eq(tt *Type) bool {
	if t.Name != tt.Name {
		return false
	}

	if t.Type != tt.Type {
		return false
	}

	if t.JSONName != tt.JSONName {
		return false
	}

	if t.Description != tt.Description {
		return false
	}

	if tt.Format == "" && t.Format != "" {
		tt.Format = t.Format
		return true
	}

	if len(t.Properties) != len(tt.Properties) {
		return false
	}

	for i := range t.Properties {
		if !t.Properties[i].Eq(tt.Properties[i]) {
			return false
		}
	}

	if len(t.OptionalProperties) != len(tt.OptionalProperties) {
		return false
	}

	for i := range t.OptionalProperties {
		if !t.OptionalProperties[i].Eq(tt.OptionalProperties[i]) {
			return false
		}
	}
	return true
}

func PackageNameify(name string) string {
	name = strings.Replace(name, "{", "", -1)
	name = strings.Replace(name, "}", "", -1)
	buf := &bytes.Buffer{}
	for _, c := range name {
		if 'A' <= c && c <= 'Z' {
			// just convert [A-Z] to _[a-z]
			if buf.Len() > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(c - 'A' + 'a')
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

func Nameify(name string) string {
	name = strings.Replace(name, "{", "", -1)
	name = strings.Replace(name, "}", "", -1)
	name = strings.Replace(name, "[", "", -1)
	name = strings.Replace(name, "]", "", -1)
	name = strings.Replace(name, ".", "-", -1)
	name = strings.Replace(name, "_", "-", -1)
	name = caser.String(name)
	return strings.Replace(name, "-", "", -1)
}

func (p *Package) StrType(schemaType *jsonschema.SchemaOrString, optional bool) string {
	switch schemaType.String {
	case "integer":
		if optional {
			return "*int"
		} else {
			return "int"
		}
	case "object":
		return "struct"
	case "boolean":
		p.Util = true
		if optional {
			return "*util.SpecialBool"
		} else {
			return "util.SpecialBool"
		}
	case "array":
		if optional {
			return "*[]"
		} else {
			return "[]"
		}
	case "string":
		if optional {
			return "*string"
		} else {
			return "string"
		}
	case "null":
		return "null"
	case "number":
		if optional {
			return "*float64"
		} else {
			return "float64"
		}
	case "any":
		return "interface{}"
	default:
		panic(schemaType.String)
	}
}

func Format(w io.Writer, source string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", source, parser.ParseComments)
	if err != nil {
		return err
	}

	return format.Node(w, fset, node)
}

type TypeSorter struct {
	types []*Type
}

func (t *TypeSorter) Len() int {
	return len(t.types)
}

func (t *TypeSorter) Less(a, b int) bool {
	return t.types[a].Name < t.types[b].Name
}

func (t *TypeSorter) Swap(i, j int) {
	temp := t.types[i]
	t.types[i] = t.types[j]
	t.types[j] = temp
}
