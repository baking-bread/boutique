package schema

import (
	"encoding/json"
	"errors"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

func ValidateJSON(schema []byte, input []byte) error {
	// check if the schema and the input is a valid json
	schemaIsValid := json.Valid(schema)
	if !schemaIsValid {
		return errors.New("schema is not a valid json")
	}

	inputIsValid := json.Valid(input)
	if !inputIsValid {
		return errors.New("input is not a valid json")
	}

	results, err := gojsonschema.Validate(
		gojsonschema.NewBytesLoader(schema),
		gojsonschema.NewBytesLoader(input))

	if len(results.Errors()) > 0 {
		err = errors.New("validation failed")
		for _, validationError := range results.Errors() {
			errors.Join(err, errors.New(validationError.Description()))
		}
	}

	return err
}

func ValidateYAML(schema []byte, input []byte) error {
	var container map[string]interface{}

	err := yaml.Unmarshal(input, &container)
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(container)
	if err != nil {
		return err
	}

	return ValidateJSON(schema, jsonBytes)
}
