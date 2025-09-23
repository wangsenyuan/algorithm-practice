package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var y int
	fmt.Fscan(reader, &y)
	res := solve(y)
	fmt.Println(res)
}

func checkLeapYear(y int) bool {
	if y%400 == 0 {
		return true
	}
	if y%4 == 0 {
		return y%100 != 0
	}
	return false
}

func solve(y int) int {
	// 必须知道1月1日是星期几
	curYear := 2016
	// 2016年1月1日是星期五
	curDay := 5

	if curYear > y {
		for curYear > y {
			curYear--
			sub := 365
			if checkLeapYear(curYear) {
				sub++
			}
			curDay = ((curDay-sub)%7 + 7) % 7
		}
	} else if curYear < y {
		for curYear < y {
			add := 365
			if checkLeapYear(curYear) {
				add++
			}
			curDay = (curDay + add) % 7
			curYear++
		}
	}

	thisDay := curDay

	for {
		add := 365
		if checkLeapYear(curYear) {
			add++
		}
		curDay = (curDay + add) % 7
		curYear++

		if curDay == thisDay && checkLeapYear(y) == checkLeapYear(curYear) {
			return curYear
		}
	}
}
