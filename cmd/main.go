package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/ostapneko/syllabus"
)

func main() {
	var after string
	if len(os.Args) > 1 {
		after = os.Args[1]
	} else {
		after = "a"
	}

	println(len(syllabus.StartConsonnants) * len(syllabus.MiddleVowels) * len(syllabus.EndConsonnants))

	for _, start := range syllabus.StartConsonnants {
		for _, middle := range syllabus.MiddleVowels {
			for _, end := range syllabus.EndConsonnants {
				name := start + middle + end

				if name > after {

					cells := make([]string, 3)
					cells[0] = name
					dotComExists := syllabus.DomainExists(name + ".com")
					dotIOExists := syllabus.DomainExists(name + ".io")
					cells[1] = strconv.FormatBool(dotComExists)
					cells[2] = strconv.FormatBool(dotIOExists)
					row := strings.Join(cells, ",")
					os.Stdout.Write([]byte(row + "\n"))
				}
			}
		}
	}
}
