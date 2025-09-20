package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

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
		return nil, fmt.Errorf("failed to read file '%s': %w", filePath, err)
	}

	// 2. Unmarshal the whole file into a generic interface
	var fullData interface{}
	if err := json.Unmarshal(fileBytes, &fullData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json from file '%s': %w", filePath, err)
	}

	// 3. Navigate through the data using the jsonPath
	pathParts := strings.Split(jsonPath, ".")
	currentElement := fullData

	for _, part := range pathParts {
		// Assert that the current element is a map to navigate deeper
		currentMap, ok := currentElement.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid json path: segment '%s' is not in a JSON object", part)
		}

		// Find the next element in the map
		element, found := currentMap[part]
		if !found {
			return nil, fmt.Errorf("json path not found: key '%s' does not exist", part)
		}
		currentElement = element
	}

	// 4. Marshal the found element back to JSON bytes
	// This is the simplest way to prepare it for unmarshaling into a specific struct
	targetBytes, err := json.Marshal(currentElement)
	if err != nil {
		return nil, fmt.Errorf("failed to re-marshal target element: %w", err)
	}

	// 5. Unmarshal the target bytes into a new instance of the desired struct type
	var result T
	if err := json.Unmarshal(targetBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal target element into struct: %w", err)
	}

	return &result, nil
}
