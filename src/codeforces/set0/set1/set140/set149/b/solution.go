package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	cnt, radixes := solve(s)
	if cnt > 0 {
		s = fmt.Sprintf("%v", radixes)
		fmt.Println(s[1 : len(s)-1])
	} else {
		fmt.Println(cnt)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) (cnt int, radixes []int) {
	// 如果24可以 => -1
	ss := strings.Split(s, ":")
	hour := ss[0]
	min := ss[1]

	for len(hour) > 1 && hour[0] == '0' {
		hour = hour[1:]
	}

	for len(min) > 1 && min[0] == '0' {
		min = min[1:]
	}

	if len(hour) == 1 && len(min) == 1 {
		if get(hour[0]) >= 24 {
			return 0, nil
		}

		return -1, nil
	}

	check := func(radix int, w string, limit int) bool {
		var time int
		for i := range len(w) {
			time = time*radix + get(w[i])
		}
		return time < limit
	}

	var lo int
	for i := range len(hour) {
		lo = max(lo, get(hour[i]))
	}
	for i := range len(min) {
		lo = max(lo, get(min[i]))
	}
	lo++

	var res []int
	for i := lo; ; i++ {
		if !check(i, hour, 24) || !check(i, min, 60) {
			break
		}
		res = append(res, i)
	}

	return len(res), res
}

func get(x byte) int {
	if x >= '0' && x <= '9' {
		return int(x - '0')
	}
	return 10 + int(x-'A')
}
