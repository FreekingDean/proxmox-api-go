package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
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
		Enums:   make(map[string][]string),
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
							if !strings.HasSuffix(methodName, textName) && !strings.HasPrefix(methodName, textName) {
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
						if !strings.HasSuffix(methodName, textName) && !strings.HasPrefix(methodName, textName) {
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
	t, err := template.New("package.go.tmpl").Funcs(template.FuncMap{
		"Enumify":     Enumify,
		"TrimN":       TrimN,
		"TrimPrefArr": TrimPrefArr,
	}).ParseFiles("cmd/download-schema/templates/package.go.tmpl")
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
	Enums     map[string][]string
	ImportURL bool
	Strconv   bool
}

type OperationTempl struct {
	Operation   string
	Description string
	Method      string
	Path        string
	Request     *Type
	Response    *Type
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
	name = strings.Replace(name, ".", "-", -1)
	name = strings.Replace(name, "_", "-", -1)
	name = caser.String(name)
	return strings.Replace(name, "-", "", -1)
}

func (p *Package) StrType(schemaType *jsonschema.SchemaOrString) string {
	switch schemaType.String {
	case "integer":
		return "int"
	case "object":
		return "struct"
	case "boolean":
		p.Util = true
		return "util.PVEBool"
	case "array":
		return "[]"
	case "string":
		return "string"
	case "null":
		return "null"
	case "number":
		return "float64"
	case "any":
		return "interface{}"
	default:
		panic(schemaType.String)
	}
}

func Format(w io.Writer, source string) error {
	formatted, err := format.Source([]byte(source))
	if err != nil {
		return err
	}

	_, err = w.Write(formatted)
	return err
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

func (p *Package) AddEnum(name string, enum []string) {
	if enums, ok := p.Enums[name]; ok {
		for _, val := range enum {
			if !inSlice(enums, val) {
				p.Enums[name] = append(p.Enums[name], val)
			}
		}
		return
	}
	p.Enums[name] = enum
}

func inSlice(sl []string, val string) bool {
	for _, v := range sl {
		if v == val {
			return true
		}
	}

	return false
}

func Enumify(s string) string {
	s = strings.ToUpper(s)
	s = strings.Replace(s, "-", "_", -1)
	s = strings.Replace(s, ".", "_", -1)
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, "+", "_AND_", -1)
	return s
}

func TrimN(s string) string {
	return strings.TrimSuffix(s, "[n]")
}

func TrimPrefArr(s string) string {
	return strings.TrimPrefix(s, "map[string]*")
}
