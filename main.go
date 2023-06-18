package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
}

func readCSVFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return rows, nil
}
