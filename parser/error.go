package parser

type ParseError struct {
	Type    ParseErrorType
	Message string
	Raw     string
}

type ParseErrorType string

const (
	ErrInvalidPrefix ParseErrorType = "invalid url prefix"
	ErrInvalidStruct ParseErrorType = "invalid struct"
	ErrInvalidPort   ParseErrorType = "invalid port number"
)

func (e *ParseError) Error() string {
	if e.Message != "" {
		return string(e.Type) + ": " + e.Message + " \"" + e.Raw + "\""
	}
	return string(e.Type)
}
