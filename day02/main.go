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

func RemoveFromReport(slice []int, i int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	return append(result[:i], result[i+1:]...)
}

func CountSafeReports(path string) (int, int) {
	file, err := os.Open("input/day_2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	countSafe := 0
	countSafeTolerant := 0
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
		if ValidateReportTolerant(report) {
			countSafeTolerant++
		} else {
		}

	}
	return countSafe, countSafeTolerant
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

func ValidateReportTolerant(report []int) bool {
	if len(report) <= 2 {
		return InValidRange(report[0] - report[1])
	}
	if (report[0] < report[1] && report[1] > report[2]) || (report[0] > report[1] && report[1] < report[2]) {
		report_a := RemoveFromReport(report, 0)
		if ValidateReport(report_a) {
			return true
		}
	}

	i := 1
	if report[i] > report[i-1] { // increasing
		for i < len(report) {
			if report[i] < report[i-1] || !InValidRange(report[i]-report[i-1]) {
				report_a := RemoveFromReport(report, i)
				report_b := RemoveFromReport(report, i-1)
				return ValidateReport(report_a) || ValidateReport(report_b)
			}
			i++
		}
	} else { // decreassing
		for i < len(report) {
			if report[i] > report[i-1] || !InValidRange(report[i]-report[i-1]) {
				report_a := RemoveFromReport(report, i)
				report_b := RemoveFromReport(report, i-1)
				return ValidateReport(report_a) || ValidateReport(report_b)
			}
			i++
		}
	}
	return true
}

func main() {
	safeReportCount, safeReportTolerantCount := CountSafeReports("input/day_2.txt")
	fmt.Println("# safe reports: ", safeReportCount)
	fmt.Println("# safe reports with tolerance: ", safeReportTolerantCount)
}
