package csv

import (
	"io"
)

func (p *YourCSVParser) ReadLine(r io.Reader) (string, error) {
	var inQuote bool
	var currentField []byte
	var line []byte

	buffer := make([]byte, 1)

	for {
		n, err := r.Read(buffer)
		if n == 0 {
			if len(line) == 0 {
				if err == io.EOF {
					p.isEOF = true
					return "", io.EOF
				}
				return "", err
			}
			break
		}
		if err != nil {
			return "", err
		}

		char := buffer[0]
		switch char {
		case '\n':
			if !inQuote {
				line = append(line, currentField...)
				p.Line = line
				p.parseFields()
				if p.expectedFields == 0 {
					p.expectedFields = len(p.Fields)
				} else if len(p.Fields) != p.expectedFields {
					return "", ErrFieldCount
				}

				p.wasRead = true
				return string(p.Line), nil
			} else {
				currentField = append(currentField, char)
			}
		case '\r':
			_, err := r.Read(buffer)
			if err != nil && err != io.EOF {
				return "", err
			}
			if buffer[0] == '\n' {
			} else {
				line = append(line, currentField...)
				line = append(line, '\r')
				currentField = nil
				if err != nil {
					return "", err
				}
			}
		case ',':
			if !inQuote {
				line = append(line, currentField...)
				line = append(line, ',')
				currentField = nil
			} else {
				currentField = append(currentField, char)
			}
		case '"':
			inQuote = !inQuote
			currentField = append(currentField, char)
		default:
			currentField = append(currentField, char)
		}
	}

	if inQuote {
		return "", ErrQuote
	}

	line = append(line, currentField...)
	p.Line = line
	p.parseFields()
	if p.expectedFields == 0 {
		p.expectedFields = len(p.Fields)
	} else if len(p.Fields) != p.expectedFields {
		return "", ErrFieldCount
	}

	p.wasRead = true
	p.isEOF = false
	return string(p.Line), nil
}

func (p *YourCSVParser) parseFields() {
	var inQuote bool
	var fields []string
	var currentField []byte
	for _, b := range p.Line {
		if b == ',' {
			if inQuote {
				currentField = append(currentField, b)
			} else {
				fields = append(fields, string(currentField))
				currentField = nil
			}
		} else if b == '\n' && !inQuote {
			fields = append(fields, string(currentField))
			break
		} else if b == '"' {
			inQuote = !inQuote
		} else {
			currentField = append(currentField, b)
		}
	}
	if len(currentField) > 0 || len(fields) == 0 {
		fields = append(fields, string(currentField))
	}
	if len(p.Line) > 0 && p.Line[len(p.Line)-1] == ',' {
		fields = append(fields, "")
	}

	p.Fields = fields
}

func (p *YourCSVParser) GetNumberOfFields() int {
	if !p.wasRead && !p.isEOF {
		return 0
	}
	return len(p.Fields)
}

func (p *YourCSVParser) GetField(n int) (string, error) {
	if !p.wasRead || n < 0 || n >= len(p.Fields) {
		return "", ErrFieldCount
	}
	return p.Fields[n], nil
}
