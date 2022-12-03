package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func caloriesPerElf(filename string) []int {
	var caloriesPerElf []int
	file, err := os.Open(filename)
	check(err)
	defer fileCloseAndErrorCheck(file)

	scanner := bufio.NewScanner(file)
	var currentElf = 0
	for scanner.Scan() {
		var inputLine = scanner.Text()
		if len(inputLine) == 0 {
			currentElf++
		} else {
			calories := convertStringToInt(inputLine)
			if len(caloriesPerElf) == currentElf {
				// New elf found. Add a record of their calories starting at 0
				caloriesPerElf = append(caloriesPerElf, 0)
			}
			caloriesPerElf[currentElf] += calories
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return caloriesPerElf
}

func highestCaloriesCarriedByElf(filename string) int {
	var caloriesPerElf = caloriesPerElf(filename)
	sort.Slice(caloriesPerElf, func(i, j int) bool {
		return caloriesPerElf[i] > caloriesPerElf[j]
	})
	return caloriesPerElf[0]
}

func caloriesCarriedByTopThreeElvesWithMostSnacks(filename string) int {
	var caloriesPerElf = caloriesPerElf(filename)
	sort.Slice(caloriesPerElf, func(i, j int) bool {
		return caloriesPerElf[i] > caloriesPerElf[j]
	})
	return caloriesPerElf[0] + caloriesPerElf[1] + caloriesPerElf[2]
}
