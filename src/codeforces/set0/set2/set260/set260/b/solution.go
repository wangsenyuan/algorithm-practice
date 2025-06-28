package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := solve(input)
	fmt.Println(result)
}

func solve(s string) string {
	dateCounts := make(map[string]int)
	n := len(s)

	for i := 0; i <= n-10; i++ {
		sub := s[i : i+10]
		if isValidDate(sub) {
			dateCounts[sub]++
		}
	}

	maxCount := 0
	apocalypseDate := ""

	for date, count := range dateCounts {
		if count > maxCount {
			maxCount = count
			apocalypseDate = date
		}
	}

	return apocalypseDate
}

func isValidDate(dateStr string) bool {
	if len(dateStr) != 10 {
		return false
	}
	if dateStr[2] != '-' || dateStr[5] != '-' {
		return false
	}

	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return false
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}
	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}
	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return false
	}

	// Year validation
	if year < 2013 || year > 2015 {
		return false
	}

	// Month validation
	if month < 1 || month > 12 {
		return false
	}

	// Day validation
	daysInMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31} // Index 0 is dummy, Feb has 28 days for 2013-2015
	if day < 1 || day > daysInMonth[month] {
		return false
	}

	return true
}
