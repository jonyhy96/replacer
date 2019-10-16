package file

// Transformer transforms file to builtin.
type Transformer interface {
	Transform(filename string) (map[string]interface{}, error)
}

// New creates a new Transformer.
func New(fileType string) Transformer {
	switch fileType {
	case "json":
		return &JSON{}
	default:
		return &JSON{}
	}
}
