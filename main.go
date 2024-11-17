package main

import (
	"fmt"
	"io"
	"os"

	. "a-library-for-others/csv"
)

func main() {
	file, err := os.Open("Great.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var csvparser CSVParser = &YourCSVParser{}

	for {
		line, err := csvparser.ReadLine(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}
		fmt.Println("Line:", line)

		numFields := csvparser.GetNumberOfFields()
		fmt.Printf("Number of Fields: %d\n", numFields)
		for i := 0; i < numFields; i++ {
			field, _ := csvparser.GetField(i)
			fmt.Printf("Field %d: %s\n", i+1, field)
		}
	}
}
