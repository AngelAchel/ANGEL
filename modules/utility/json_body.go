package utility

import (
	"encoding/json"
	"fmt"
)

// JSONBodyEngine provides JSON body injection utilities.
type JSONBodyEngine struct{}

// NewJSONBodyEngine creates a new JSONBodyEngine.
func NewJSONBodyEngine() *JSONBodyEngine {
	return &JSONBodyEngine{}
}

// InjectField injects a field into a JSON string.
func (e *JSONBodyEngine) InjectField(jsonStr string, key string, value string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	data[key] = value

	result, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(result), nil
}

// InjectNestedField injects a nested field into a JSON string.
func (e *JSONBodyEngine) InjectNestedField(jsonStr string, path string, value string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	keys := splitPath(path)
	current := data
	for i := 0; i < len(keys)-1; i++ {
		if _, ok := current[keys[i]]; !ok {
			current[keys[i]] = make(map[string]interface{})
		}
		if next, ok := current[keys[i]].(map[string]interface{}); ok {
			current = next
		} else {
			return "", fmt.Errorf("path conflicts with existing value at %s", keys[i])
		}
	}

	current[keys[len(keys)-1]] = value

	result, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(result), nil
}

// OverrideField overrides an existing field in a JSON string.
func (e *JSONBodyEngine) OverrideField(jsonStr string, key string, value string) (string, error) {
	return e.InjectField(jsonStr, key, value)
}

// RemoveField removes a field from a JSON string.
func (e *JSONBodyEngine) RemoveField(jsonStr string, key string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	delete(data, key)

	result, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(result), nil
}

// BuildJSON builds a JSON object from key-value pairs.
func (e *JSONBodyEngine) BuildJSON(pairs map[string]string) (string, error) {
	result, err := json.Marshal(pairs)
	if err != nil {
		return "", fmt.Errorf("failed to build JSON: %w", err)
	}
	return string(result), nil
}

// ArrayElementInject injects an element into a JSON array.
func (e *JSONBodyEngine) ArrayElementInject(jsonStr string, key string, value string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	if arr, ok := data[key].([]interface{}); ok {
		data[key] = append(arr, value)
	} else {
		data[key] = []interface{}{value}
	}

	result, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(result), nil
}

func splitPath(path string) []string {
	var keys []string
	current := ""
	for _, ch := range path {
		if ch == '.' {
			if current != "" {
				keys = append(keys, current)
				current = ""
			}
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		keys = append(keys, current)
	}
	return keys
}