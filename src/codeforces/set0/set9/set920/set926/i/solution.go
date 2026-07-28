package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	alarms := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &alarms[i])
	}
	return solve(alarms)
}

func solve(alarms []string) string {
	if len(alarms) == 1 {
		return "23:59"
	}
	arr := make([]time, len(alarms))
	for i, cur := range alarms {
		arr[i] = Parse(cur)
	}

	slices.SortFunc(arr, func(a, b time) int {
		return cmp.Or(a.hour-b.hour, a.minute-b.minute)
	})

	var best int
	for i := range arr {
		j := (i + 1) % len(arr)
		diff := arr[j].getInMins() - arr[i].getInMins()
		if diff < 0 {
			diff += 24 * 60
		}
		diff--
		best = max(best, diff)
	}
	return fmt.Sprintf("%02d:%02d", best/60, best%60)
}

type time struct {
	hour   int
	minute int
}

func (t time) getInMins() int {
	return t.hour*60 + t.minute
}

func Parse(s string) time {
	ss := strings.Split(s, ":")
	return time{
		hour:   parseInt(ss[0]),
		minute: parseInt(ss[1]),
	}
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
