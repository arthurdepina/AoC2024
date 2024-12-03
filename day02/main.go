package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func InValidRange(n int) bool {
	return -3 <= n && n <= 3 && n != 0
}

func CountSafeReports(path string) int {
	file, err := os.Open("input/day_2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	countSafe := 0
	for scanner.Scan() {
		var report []int
		line := strings.Split(scanner.Text(), " ")
		for _, numStr := range line {
			num, _ := strconv.Atoi(numStr)
			report = append(report, num)
		}
		if ValidateReport(report) {
			countSafe++
		}
	}
	return countSafe
}

func ValidateReport(report []int) bool {
	if len(report) <= 2 {
		return InValidRange(report[0] - report[1])
	}

	i := 1
	if report[i] > report[i-1] { // increasing
		for i < len(report) {
			if report[i] < report[i-1] || !InValidRange(report[i]-report[i-1]) {
				return false
			}
			i++
		}
	} else { // decreassing
		for i < len(report) {
			if report[i] > report[i-1] || !InValidRange(report[i]-report[i-1]) {
				return false
			}
			i++
		}
	}

	return true
}

func main() {
	fmt.Println("# safe reports: ", CountSafeReports("input/day_2.txt"))
}
