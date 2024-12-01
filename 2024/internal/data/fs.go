package data

import (
	"embed"
	"fmt"
	"log"
	"strings"
)

//go:embed *.txt
var InputFiles embed.FS

func ReadInput(day int) ([]string, error) {
	filename := fmt.Sprintf("day%d.txt", day)

	log.Printf("Loading in input from file '%s'\n", filename)
	data, err := InputFiles.ReadFile(filename)

	lines := strings.Split(string(data), "\n")

	if err != nil {
		log.Printf("Succesfully loaded %d lines!", len(lines))
	}

	return lines, err
}
