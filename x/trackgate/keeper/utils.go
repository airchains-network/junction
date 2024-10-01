package keeper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/airchains-network/junction/x/trackgate/types"
)

// ValidateVersion checks if the provided version string conforms to Semantic Versioning 2.0.0
func ValidateVersion(version string) bool {
	// Regex for Semantic Versioning (https://semver.org/)
	var semverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-((0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(\.(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(\+([0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*))?$`)
	return semverRegex.MatchString(version)
}

// BuildExtTrackSchemaPath constructs a path for storing ext-track-schema objects in a database.
func BuildExtTrackSchemaPath(extTrackStationID string) string {
	dbPath := fmt.Sprintf("%s/%s", types.ExtTrackSchemaStoreKey, extTrackStationID)

	return dbPath
}

// BuildExtTrackVersionFinderStoreKey constructs the database key for storing external track version data using a station ID.
func BuildExtTrackVersionFinderStoreKey(extTrackStationID string) string {
	dbPath := fmt.Sprintf("%s/%s", types.ExtTrackVersionFinderStoreKey, extTrackStationID)
	return dbPath
}

// BuildExtTrackSchemaEngagementsStoreKey constructs the database key for storing engagement data by combining the base key and station ID.
func BuildExtTrackSchemaEngagementsStoreKey(extTrackStationID string) string {
	dbPath := fmt.Sprintf("%s/%s", types.ExtTrackSchemaEngagementsStoreKey, extTrackStationID)
	return dbPath
}

// ContainsOperator checks if the slice 'operators' contains the 'operator'.
func ContainsOperator(operators []string, operator string) bool {
	for _, op := range operators {
		if op == operator {
			return true
		}
	}
	return false
}

func validateJsonStructure(schemaFields map[string]interface{}, jsonData map[string]interface{}) bool {
	// Ensure every field in schema is present in the JSON
	for key, fieldType := range schemaFields {
		if jsonValue, exists := jsonData[key]; exists {
			switch ft := fieldType.(type) {
			case string:
				if !validateFieldType(ft, jsonValue) {
					fmt.Printf("Field '%s' has mismatched type or invalid value.\n", key)
					return false
				}
			case map[string]interface{}:
				if nestedJsonData, ok := jsonValue.(map[string]interface{}); ok {
					if !validateJsonStructure(ft, nestedJsonData) {
						fmt.Printf("Field '%s' has mismatched structure.\n", key)
						return false
					}
				} else {
					fmt.Printf("Expected nested structure for '%s', got something else.\n", key)
					return false
				}
			default:
				fmt.Printf("Unsupported type for '%s'.\n", key)
				return false
			}
		} else {
			// Error when a required field is missing from JSON data
			fmt.Printf("Field '%s' is missing from the JSON data.\n", key)
			return false
		}
	}

	// Check for extra fields in the JSON that aren't in the schema
	for key := range jsonData {
		if _, exists := schemaFields[key]; !exists {
			fmt.Printf("Extra field '%s' found in the JSON data but not in the schema.\n", key)
			return false
		}
	}

	return true
}

// Validate individual field types
func validateFieldType(fieldType string, value interface{}) bool {
	switch fieldType {
	case "string":
		_, ok := value.(string)
		return ok
	case "int":
		_, ok := value.(float64) // JSON unmarshalling in Go converts numbers to float64
		return ok
	case "uint":
		val, ok := value.(float64)
		return ok && val >= 0
	case "bytes":
		_, ok := value.(string) // Assuming bytes are encoded as base64 strings
		return ok
	default:
		return false
	}
}

func DynamicUnmarshal(schemaDef types.SchemaDef, jsonData string) (string, error) {
	// Temporarily unmarshal JSON into a map to validate structure
	var tempData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &tempData); err != nil {
		return "error", fmt.Errorf("failed to unmarshal JSON for validation: %w", err)
	}

	// Validate JSON against struct definition
	if !validateJsonStructure(schemaDef.Fields, tempData) {
		return "error", fmt.Errorf("JSON data does not match the expected schema")
	}

	// Function to recursively create fields
	var createFields func(fields map[string]interface{}) ([]reflect.StructField, error)
	createFields = func(fields map[string]interface{}) ([]reflect.StructField, error) {
		var result []reflect.StructField
		for name, field := range fields {
			switch f := field.(type) {
			case string: // Base types
				var typ reflect.Type
				switch f {
				case "string":
					typ = reflect.TypeOf("")
				case "int":
					typ = reflect.TypeOf(0)
				case "uint":
					typ = reflect.TypeOf(uint(0))
				case "bytes":
					typ = reflect.TypeOf([]byte{})
				default:
					return nil, fmt.Errorf("unsupported type: %s", f)
				}
				result = append(result, reflect.StructField{
					Name: strings.Title(name),
					Type: typ,
					Tag:  reflect.StructTag(`json:"` + name + `"`),
				})
			case map[string]interface{}: // Nested structs
				nestedFields, err := createFields(f)
				if err != nil {
					return nil, err
				}
				structType := reflect.StructOf(nestedFields)
				result = append(result, reflect.StructField{
					Name: strings.Title(name),
					Type: structType,
					Tag:  reflect.StructTag(`json:"` + name + `"`),
				})
			default:
				return nil, fmt.Errorf("invalid field definition: %v", field)
			}
		}
		return result, nil
	}

	// Create struct type from fields
	fields, err := createFields(schemaDef.Fields)
	if err != nil {
		return "error", err
	}

	structType := reflect.StructOf(fields)
	structInstance := reflect.New(structType).Elem()

	// Unmarshal JSON into the dynamically created struct
	if err := json.Unmarshal([]byte(jsonData), structInstance.Addr().Interface()); err != nil {
		return "error", fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Marshal the struct back to JSON with indentation
	formattedJson, err := json.MarshalIndent(structInstance.Interface(), "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal struct to JSON: %w", err)
	}

	return string(formattedJson), nil
}
