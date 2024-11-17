package csv

import (
	"errors"
	"io"
)

var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

type YourCSVParser struct {
	Line           []byte
	Fields         []string
	wasRead        bool
	isEOF          bool
	expectedFields int
}
