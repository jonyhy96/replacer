package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// JSON handles transform json file to map.
type JSON struct {
}

// Transform implements Transformer.Transform
func (j JSON) Transform(filename string) (map[string]interface{}, error) {
	var result map[string]interface{}
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile %w", err)
	}
	err = json.Unmarshal(fileBytes, &result)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarsha %w", err)
	}
	return result, nil
}
