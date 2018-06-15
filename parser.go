package praks

// Parser interface //
type Parser interface {
	TexToStruct(text string) *Struct
}

// NewParser get json or csv parser
func NewParser(t string) Parser {
	switch t {
	case "json":
		return &jsonParser{}
	// TODO
	// Add csv case
	default:
		return &jsonParser{}
	}
}
