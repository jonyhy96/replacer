package replacer

import "io"

// Replacer replace something of the file.
type Replacer interface {
	Replace(in io.Reader, out io.Writer) error
}

// New creates a new Replacer based on replacer file's type.
func New(fileType string, filecontent interface{}) Replacer {
	switch fileType {
	case "json":
		return &JSON{
			replace: filecontent.(map[string]interface{}),
		}
	default:
		return &JSON{}
	}
}
