package main

import (
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileCloseAndErrorCheck(file *os.File) {
	check(file.Close())
}

func convertStringToInt(inputAsString string) int {
	inputAsInt, err := strconv.Atoi(inputAsString)
	check(err)
	return inputAsInt
}
