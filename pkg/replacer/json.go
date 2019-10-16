package replacer

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	// ErrorUnknownType handles unknown types error.
	ErrorUnknownType = errors.New("unknown type")
)

// JSON handles json format replace file.
type JSON struct {
	replace map[string]interface{}
}

// Replace replace all matched string of in by replace.key and write to out.
func (j *JSON) Replace(in io.Reader, out io.Writer) error {
	inBytes, err := ioutil.ReadAll(in)
	result := string(inBytes)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll %w", err)
	}
	for key, value := range j.replace {
		result = strings.ReplaceAll(result, key, interface2string(value))
	}
	_, err = out.Write([]byte(result))
	if err != nil {
		return fmt.Errorf("write to out %w", err)
	}
	return nil
}

func interface2string(in interface{}) string {
	if val, ok := in.(string); ok {
		return val
	}
	if val, ok := in.(int); ok {
		return strconv.Itoa(val)
	}
	if val, ok := in.(float64); ok {
		return strconv.FormatFloat(val, 'E', -1, 64)
	}
	return ErrorUnknownType.Error()
}
