package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	getLevels("input.txt")
}

func getLevels(fileName string) {
	file, err := os.Open(fileName)
	errCheck(err)

	txt, err := io.ReadAll(file)
	errCheck(err)

	levelsTxt := strings.Split(string(txt), "\n")
	var levelsArray [][]string

	for _, v := range levelsTxt {
		level := strings.Split(v, " ")
		levelsArray = append(levelsArray, level)
	}
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
