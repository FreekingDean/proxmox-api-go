package jsonschema

import (
	"encoding/json"
	"fmt"
)

type Schema struct {
	Children []*Schema              `json:"children"`
	Info     map[string]*InfoSchema `json:"info"`
	Leaf     int                    `json:"leaf"`
	Path     string                 `json:"path"`
	Text     string                 `json:"text"`
}

type InfoSchema struct {
	AllowToken  int                    `json:"allowToken"`
	Description string                 `json:"description"`
	Method      string                 `json:"method"`
	Name        string                 `json:"name"`
	Parameters  *JSONSchema            `json:"parameters"`
	Returns     *JSONSchema            `json:"returns"`
	Permissions map[string]interface{} `json:"permissions"`
	Protected   int                    `json:"protected"`
}

type JSONSchema struct {
	AdditionalProperties *SchemaOrBool          `json:"additionalProperties"`
	Properties           map[string]*JSONSchema `json:"properties"`
	Default              interface{}            `json:"default"`
	Type                 *SchemaOrString        `json:"type"`
	Items                *JSONSchema            `json:"items"`
	Links                []*Link                `json:"links"`
	Optional             SpecialBool            `json:"optional"`
	TypeText             string                 `json:"typeText"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type SchemaOrBool struct {
	Schema *JSONSchema
	Bool   SpecialBool
}

type SchemaOrString struct {
	Schema *JSONSchema
	String string
}

type SpecialBool bool

func (s *SpecialBool) UnmarshalJSON(data []byte) error {
	if string(data[0:1]) == "1" || string(data[0:3]) == `"1"` {
		*s = SpecialBool(true)
		return nil
	}
	if string(data[0:1]) == "0" || string(data[0:3]) == `"0"` {
		*s = SpecialBool(false)
		return nil
	}
	return fmt.Errorf("Invalid bool")
}

func (s *SchemaOrString) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		s.String = "object"
		var v *JSONSchema
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		s.Schema = v
		return nil
	}
	s.String = string(data[1 : len(data)-1])
	return nil
}

func (s *SchemaOrBool) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		s.Bool = true
		return json.Unmarshal(data, &s.Schema)
	}
	return json.Unmarshal(data, &s.Bool)
}
