package schema

import "testing"

var schemaJSON = `{
	"$id": "https://example.com/address.schema.json",
	"$schema": "https://json-schema.orfg/draft/2020-12/schema",
	"description": "A schema for input validation tests",
	"type": "object",
	"properties": {
	  "value": {
		"type": "string"
	  },
	  "optional": {
		"type": "string"
	  }
	},
	"required": [ "value" ]
  }`

func TestValidateJSON(t *testing.T) {
	jsonInput := `{
		"value": "test"
	}`

	err := ValidateJSON([]byte(schemaJSON), []byte(jsonInput))
	if err != nil {
		t.Fail()
	}
}

func TestValidateJSONWithAdditionalValue(t *testing.T) {
	jsonInput := `{
		"value": "test",
		"asdf": "something"
	}`

	err := ValidateJSON([]byte(schemaJSON), []byte(jsonInput))
	if err != nil {
		t.Fail()
	}
}

func TestValidateJSONWithInvalidValue(t *testing.T) {
	jsonInput := `{
		"value": "1",
	}`

	err := ValidateJSON([]byte(schemaJSON), []byte(jsonInput))
	if err == nil {
		t.Fail()
	}
}

func TestValidateYAML(t *testing.T) {
	yamlInput := `
value: test 
`

	err := ValidateYAML([]byte(schemaJSON), []byte(yamlInput))
	if err != nil {
		t.Fail()
	}
}
