package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type JError uint

const (
	READ_FILE_JERR JError = iota
	UNMARSHAL_JERR
	INVALID_PATH_JERR
	NOT_FOUND_JERR
	REMARSHAL_JERR
	TG_UNMARSHAL_JERR
	WRITE_FILE_JERR
)

type JsonError struct {
	Code JError
	Msg  error
}

func (err JsonError) Error() string {
	return err.Msg.Error()
}

func NewJsonError(c JError, m error) JsonError {
	return JsonError{
		c,
		m,
	}
}

// LoadJSONPathToStruct finds a JSON file, navigates to a specific element
// using a dot-notation path, and unmarshals that element into a new instance
// of the provided generic type T.
//
// Parameters:
//   - filePath: The path to the JSON file.
//   - jsonPath: A dot-separated path (e.g., "data.user.profile").
//
// Returns:
//   - A pointer to the populated struct of type T, or an error if any step fails.
func LoadJSONPathToStruct[T any](filePath, jsonPath string) (*T, error) {
	// 1. Read the entire JSON file
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, NewJsonError(READ_FILE_JERR, fmt.Errorf("failed to read file '%s': %w", filePath, err))
	}

	// 2. Unmarshal the whole file into a generic interface
	var fullData interface{}
	if err := json.Unmarshal(fileBytes, &fullData); err != nil {
		return nil, NewJsonError(UNMARSHAL_JERR, fmt.Errorf("failed to unmarshal json from file '%s': %w", filePath, err))
	}

	// 3. Navigate through the data using the jsonPath
	pathParts := strings.Split(jsonPath, ".")
	currentElement := fullData

	for _, part := range pathParts {
		// Assert that the current element is a map to navigate deeper
		currentMap, ok := currentElement.(map[string]interface{})
		if !ok {
			return nil, NewJsonError(INVALID_PATH_JERR, fmt.Errorf("invalid json path: segment '%s' is not in a JSON object", part))
		}

		// Find the next element in the map
		element, found := currentMap[part]
		if !found {
			return nil, NewJsonError(NOT_FOUND_JERR, fmt.Errorf("json path not found: key '%s' does not exist", part))
		}
		currentElement = element
	}

	// 4. Marshal the found element back to JSON bytes
	// This is the simplest way to prepare it for unmarshaling into a specific struct
	targetBytes, err := json.Marshal(currentElement)
	if err != nil {
		return nil, NewJsonError(REMARSHAL_JERR, fmt.Errorf("failed to re-marshal target element: %w", err))
	}

	// 5. Unmarshal the target bytes into a new instance of the desired struct type
	var result T
	if err := json.Unmarshal(targetBytes, &result); err != nil {
		return nil, NewJsonError(TG_UNMARSHAL_JERR, fmt.Errorf("failed to unmarshal target element into struct: %w", err))
	}

	return &result, nil
}

// AddElementToJSON finds a JSON object by its key in a file and adds a new key-value pair to it.
//
// Parameters:
//
//	filePath: The path to the JSON file.
//	parentKey: The key of the JSON object to which the new element should be added.
//	           If the parentKey is empty (""), the element is added to the root object.
//	newKey: The key of the new element to add.
//	newValue: The value of the new element to add.
//
// Returns:
//
//	A *JsonError if any step of the process fails.
func AddElementToJSON(filePath, parentKey, newKey string, newValue interface{}) error {
	// 1. Read the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return NewJsonError(READ_FILE_JERR, fmt.Errorf("failed to open file: %w", err))
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return NewJsonError(READ_FILE_JERR, fmt.Errorf("failed to read file: %w", err))
	}

	// 2. Unmarshal the JSON into a map
	var data map[string]interface{}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return NewJsonError(UNMARSHAL_JERR, fmt.Errorf("failed to unmarshal JSON: %w", err))
	}

	// 3. Find the target element and add the new key-value pair
	if parentKey == "" {
		// Add to the root object
		data[newKey] = newValue
	} else {
		// Find the nested object
		targetElement, ok := data[parentKey]
		if !ok {
			return NewJsonError(NOT_FOUND_JERR, fmt.Errorf("parent key '%s' not found in JSON", parentKey))
		}

		targetMap, ok := targetElement.(map[string]interface{})
		if !ok {
			return NewJsonError(TG_UNMARSHAL_JERR, fmt.Errorf("element with key '%s' is not a JSON object (map)", parentKey))
		}
		targetMap[newKey] = newValue
	}

	// 4. Marshal the modified data back to JSON with indentation
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return NewJsonError(REMARSHAL_JERR, fmt.Errorf("failed to marshal updated JSON: %w", err))
	}

	// 5. Write the updated JSON back to the file
	if err := ioutil.WriteFile(filePath, updatedJSON, 0644); err != nil {
		return NewJsonError(WRITE_FILE_JERR, fmt.Errorf("failed to write updated JSON to file: %w", err))
	}

	return nil
}
