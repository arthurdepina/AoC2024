package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFirstInput(path string) ([]int, []int) {
	input_1, err := os.Open(path)
	check(err)
	defer input_1.Close()

	var locations_1 []int
	var locations_2 []int

	scanner := bufio.NewScanner(input_1)

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "   ")
		first, _ := strconv.Atoi(numbers[0])
		second, _ := strconv.Atoi(numbers[1])
		locations_1 = append(locations_1, first)
		locations_2 = append(locations_2, second)
	}

	return locations_1, locations_2
}

func AddedDistances(locations_1 []int, locations_2 []int) int {
	totalDistance := 0
	slices.Sort(locations_1)
	slices.Sort(locations_2)

	for index := range locations_1 {
		distance := locations_1[index] - locations_2[index]
		if distance < 0 {
			totalDistance -= distance
		} else {
			totalDistance += distance
		}
	}

	return totalDistance
}

func main() {
	locations_1, locations_2 := ReadFirstInput("input/input_1.txt")
	totalDistance := AddedDistances(locations_1, locations_2)
	fmt.Println("total distance: ", totalDistance)
}
