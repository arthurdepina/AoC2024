package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AddMults(path string) int {
	sum := 0
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}

	return sum
}

func main() {
	sum := AddMults("input/day_3.txt")
	fmt.Println(sum)
}
