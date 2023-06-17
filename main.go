package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello")
}

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
