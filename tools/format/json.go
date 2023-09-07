package format

import "encoding/json"

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
