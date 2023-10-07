package json

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type JsonFormat struct {
}

func (obj *JsonFormat) Format(input string) (string, error) {
	jsonObj := make(map[string]any)
	err := json.Unmarshal([]byte(input), &jsonObj)
	if err != nil {
		return "", err
	}

	b, err := json.MarshalIndent(jsonObj, "", "	")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (obj *JsonFormat) FormatFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	str, err := obj.Format(string(content))
	if err != nil {
		return "", err
	}
	err = os.WriteFile(filename, []byte(str), 0644)
	if err != nil {
		return "", err
	}

	output, _ := filepath.Abs(filename)
	return output, nil
}
