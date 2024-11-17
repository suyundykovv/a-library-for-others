# CSV Parser in Go

## Overview

This project is about building a simple **CSV Parser** library in Go. The goal is to implement the `CSVParser` interface with the following methods:

- **ReadLine**: Reads a line from a CSV file.
- **GetField**: Returns the n-th field of the last line read.
- **GetNumberOfFields**: Returns the number of fields in the last line read.

The parser should handle standard CSV formatting, including handling quoted fields, and return appropriate errors for malformed lines.

## Requirements

Your program must follow these guidelines:

1. **gofumpt**: Ensure your code adheres to Go's formatting standards enforced by `gofumpt`.
2. **No panics**: Your code should never exit unexpectedly due to a panic (e.g., `nil` pointer dereference, out-of-range index).
3. **Use `io.Reader`**: The parser must read from an `io.Reader`, and you cannot use any other method for reading the file.
4. **Functionality**: Implement and test the following methods within the `CSVParser` interface.

---

## CSVParser Interface

```go
type CSVParser interface {
    ReadLine(r io.Reader) (string, error)
    GetField(n int) (string, error)
    GetNumberOfFields() int
}
```

You need to implement the above methods in your CSVParser type.

---

## Error Handling

There are two predefined error variables:

```go
var (
    ErrQuote      = errors.New("excess or missing \" in quoted-field")
    ErrFieldCount = errors.New("wrong number of fields")
)
```

You will use these to handle specific errors related to malformed CSV fields:

- **ErrQuote**: This is returned when a quoted field has excess or missing quotes.
- **ErrFieldCount**: This is returned when the number of fields in a line does not match the expected count.

---

## Method Implementations

### 1. ReadLine

The `ReadLine` method reads one line from the provided `io.Reader` and removes the line terminator (i.e., `\n`, `\r`, or `\r\n`). If a quote is malformed (extra or missing quote), it should return `ErrQuote`.

#### Function Signature:
```go
func (p *CSVParser) ReadLine(r io.Reader) (string, error)
```

- **Input**: An `io.Reader` (usually a file or string).
- **Output**: A string containing the line, with the terminator removed, or an empty string and `ErrQuote` if a malformed quote is detected.

### 2. GetField

The `GetField` method returns the n-th field from the last line read by `ReadLine`. The fields in CSV are separated by commas, and fields may be enclosed in quotes. If the requested field is out of range (either negative or beyond the available fields), `ErrFieldCount` should be returned.

#### Function Signature:
```go
func (p *CSVParser) GetField(n int) (string, error)
```

- **Input**: The index `n` (1-based) of the field to retrieve.
- **Output**: The field's value as a string, or `ErrFieldCount` if `n` is out of range.

### 3. GetNumberOfFields

The `GetNumberOfFields` method returns the number of fields in the last line read by `ReadLine`. This should only return a valid value after `ReadLine` has been called.

#### Function Signature:
```go
func (p *CSVParser) GetNumberOfFields() int
```

- **Input**: None.
- **Output**: The number of fields in the last line read by `ReadLine`.

---

## Example of Usage

Here is a simple example of how your CSV parser can be used to read from a file and access its contents:

```go
package main

import (
    "errors"
    "fmt"
    "io"
    "os"
)

type CSVParser interface {
    ReadLine(r io.Reader) (string, error)
    GetField(n int) (string, error)
    GetNumberOfFields() int
}

var (
    ErrQuote      = errors.New("excess or missing \" in quoted-field")
    ErrFieldCount = errors.New("wrong number of fields")
)

func main() {
    file, err := os.Open("example.csv")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var csvparser CSVParser = YourCSVParser{} // Use your CSVParser implementation

    for {
        line, err := csvparser.ReadLine(file)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("Error reading line:", err)
            return
        }
        
        // Example of getting fields
        numFields := csvparser.GetNumberOfFields()
        fmt.Println("Number of fields:", numFields)

        for i := 0; i < numFields; i++ {
            field, err := csvparser.GetField(i)
            if err != nil {
                fmt.Println("Error getting field:", err)
                return
            }
            fmt.Println("Field", i+1, ":", field)
        }
    }
}
```

### Example of Input CSV (example.csv):

```csv
Name, Age, Occupation
John Doe, 25, Engineer
"Jane, Doe", 30, "Software Developer"
```

---

## Notes on Implementation

- **Handling Quotes**: Ensure that quoted fields are handled correctly. For example, `"John, Doe"` should be parsed as a single field, not two.
- **Field Separation**: Fields are separated by commas, but if a field is enclosed in quotes, it may contain commas inside, so you need to handle quoted fields carefully.
- **Error Handling**: Ensure that errors like missing quotes or extra commas are handled gracefully and return the correct error (e.g., `ErrQuote`, `ErrFieldCount`).
- **EOF**: When the end of the file is reached, `ReadLine` should return `io.EOF`.

---

## Testing Your Code

Make sure to test your library using different cases:

1. **Normal CSV files**: Files with basic data and no errors.
2. **Files with malformed quotes**: Fields with missing or extra quotes.
3. **Files with a variable number of fields**: Some rows may have more or fewer fields than others.
4. **Empty files**: Files with no data or empty lines.
5. **Binary files**: Test what happens if a binary file is passed.

Use the provided `main` function template for basic testing, but also write unit tests for your methods (`ReadLine`, `GetField`, `GetNumberOfFields`).

---

## Submission

Submit your implementation of the `CSVParser` interface along with any helper functions you used to accomplish the task. Ensure that your code is properly formatted with `gofumpt` and is well-tested.