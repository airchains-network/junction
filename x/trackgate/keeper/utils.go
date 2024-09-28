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

func DynamicUnmarshal(schemaDef types.SchemaDef, jsonData string) (string, error) {
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
		return "", err
	}
	structType := reflect.StructOf(fields)
	structInstance := reflect.New(structType).Elem()

	// Unmarshal JSON into the dynamically created struct
	if err := json.Unmarshal([]byte(jsonData), structInstance.Addr().Interface()); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Marshal the struct back to JSON with indentation
	formattedJson, err := json.MarshalIndent(structInstance.Interface(), "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal struct to JSON: %w", err)
	}

	return string(formattedJson), nil
}
