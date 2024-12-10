package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	levels := getLevels("input.txt")

	safeCounter := 0
	for _, level := range levels {
		if isLevelSafe(level) {
			fmt.Println(level)
			safeCounter++
		}
	}

	fmt.Println(safeCounter)
}

func getLevels(fileName string) [][]int {
	file, err := os.Open(fileName)
	errCheck(err)

	txt, err := io.ReadAll(file)
	errCheck(err)

	levelsTxt := strings.Split(string(txt), "\r\n")
	var levelsArray [][]int

	for _, v := range levelsTxt {
		level := strings.Split(v, " ")
		levelsArray = append(levelsArray, stringToIntArray(level))
	}

	return levelsArray
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func stringToIntArray(arr []string) []int {
	var intArr = []int{}
	for _, v := range arr {
		i, err := strconv.Atoi(v)
		errCheck(err)
		intArr = append(intArr, i)
	}

	return intArr
}

func isLevelSafe(arr []int) bool {
	increasing := true
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] || arr[i] == arr[i+1] || (arr[i+1]-arr[i] > 3) {
			increasing = false
			break
		}
	}

	if increasing {
		return increasing
	}

	decreasing := true
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] || arr[i] == arr[i+1] || (arr[i]-arr[i+1] > 3) {
			decreasing = false
			break
		}
	}

	return decreasing
}
