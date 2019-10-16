package replacer

import "io"

// Replacer replace something of the file.
type Replacer interface {
	Replace(in io.Reader) error
}

// New creates a new Replacer based on replacer file's type.
func New(fileType string, filecontent interface{}, filename string) Replacer {
	switch fileType {
	case "json":
		return &JSON{
			FileName: filename,
			replace:  filecontent.(map[string]interface{}),
		}
	default:
		return &JSON{
			FileName: filename,
			replace:  filecontent.(map[string]interface{}),
		}
	}
}
