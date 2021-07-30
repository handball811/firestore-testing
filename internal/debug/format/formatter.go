package format

import "errors"

var (
	ErrFin       = errors.New("fin")
	AppendPrefix = "  "
)

type Formatter interface {
	Next() (string, error)
}
