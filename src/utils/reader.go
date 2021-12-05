package utils

import (
	"bytes"
	"log"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileBuffer := new(bytes.Buffer)
	fileBuffer.ReadFrom(file)
	fileString := fileBuffer.String()

	return strings.Split(fileString, "\n")
}
