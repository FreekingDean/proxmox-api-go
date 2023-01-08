package main

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"sort"
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
	p := Package{
		Name:    packageName,
		Methods: make([]*OperationTempl, 0),
	}
	if s.Info["GET"] != nil {
		p.Methods = append(
			p.Methods,
			genMethod("Index", s.Info["GET"], s.Path),
		)
	}
	if s.Info["POST"] != nil {
		p.Methods = append(
			p.Methods,
			genMethod("Create", s.Info["POST"], s.Path),
		)
	}
	if s.Info["PUT"] != nil {
		p.Methods = append(
			p.Methods,
			genMethod("MassUpdate", s.Info["PUT"], s.Path),
		)
	}
	if s.Info["DELETE"] != nil {
		p.Methods = append(
			p.Methods,
			genMethod("MassDelete", s.Info["DELETE"], s.Path),
		)
	}
	for _, c := range s.Children {
		if c.Text[0] == '{' && c.Text[len(c.Text)-1] == '}' {
			if c.Info["GET"] != nil {
				p.Methods = append(
					p.Methods,
					genMethod("Find", c.Info["GET"], c.Path),
				)
			}
			if c.Info["POST"] != nil {
				p.Methods = append(
					p.Methods,
					genMethod("ChildCreate", c.Info["POST"], c.Path),
				)
			}
			if c.Info["PUT"] != nil {
				p.Methods = append(
					p.Methods,
					genMethod("Update", c.Info["PUT"], c.Path),
				)
			}
			if c.Info["DELETE"] != nil {
				p.Methods = append(
					p.Methods,
					genMethod("Delete", c.Info["DELETE"], c.Path),
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
								genMethod(methodName, info, cc.Path),
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
							genMethod(methodName, info, c.Path),
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
	return Format(f, b.String())
}

func genMethod(operation string, info *jsonschema.InfoSchema, path string) *OperationTempl {
	ot := &OperationTempl{
		Operation:   operation,
		Description: info.Description,
		Method:      info.Method,
		Path:        path,
	}
	if info.Parameters != nil {
		ot.Request = defineType(info.Parameters)
	}
	if info.Returns != nil {
		ot.Response = defineType(info.Returns)
	}
	return ot
}

type Package struct {
	Name    string
	Methods []*OperationTempl
}

type OperationTempl struct {
	Operation   string
	Description string
	Request     *Type
	Response    *Type
	Method      string
	Path        string
}

type Type struct {
	Properties         []*Type
	OptionalProperties []*Type
	Name               string
	JSONName           string
	Type               string
	SubType            string
	Description        string
}

func defineType(schema *jsonschema.JSONSchema) *Type {
	if schema.Type == nil {
		if schema.Properties != nil {
			schema.Type = &jsonschema.SchemaOrString{String: "object"}
		} else {
			return nil
		}
	}
	t := &Type{
		Type:               StrType(schema.Type, bool(schema.Optional)),
		Properties:         make([]*Type, 0),
		OptionalProperties: make([]*Type, 0),
		Description:        strings.Replace(schema.Description, "\n", "", -1),
	}
	for name, param := range schema.Properties {
		p := defineType(param)
		if p == nil {
			continue
		}
		p.Name = Nameify(name)
		p.JSONName = name
		if param.Optional {
			t.OptionalProperties = append(t.OptionalProperties, p)
		} else {
			t.Properties = append(t.Properties, p)
		}
	}
	if schema.Items == nil && schema.Type.String == "array" {
		t.Properties = []*Type{
			&Type{Type: "struct"},
		}
	}
	if schema.Items != nil {
		t.Properties = []*Type{
			defineType(schema.Items),
		}
	}
	sort.Sort(&TypeSorter{t.Properties})
	sort.Sort(&TypeSorter{t.OptionalProperties})
	return t
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

func StrType(schemaType *jsonschema.SchemaOrString, optional bool) string {
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
		if optional {
			return "*bool"
		} else {
			return "bool"
		}
	case "array":
		return "slice"
	case "string":
		if optional {
			return "*string"
		} else {
			return "string"
		}
	case "null":
		return "struct"
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
