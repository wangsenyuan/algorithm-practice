package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) string {
	// sum 会不会太大？
	var sum int
	var sum2 int
	n := len(s)
	for i := 0; i < n; {
		for i < n && s[i] >= 'a' && s[i] <= 'z' {
			i++
		}

		j := i
		for i < n && (s[i] >= '0' && s[i] <= '9' || s[i] == '.') {
			i++
		}

		m := i - j
		x := s[j:i]
		if m <= 2 {
			num, _ := strconv.Atoi(x)
			sum += num
			continue
		}
		// m > 2
		if x[m-3] == '.' {
			// 小数部分
			num, _ := strconv.Atoi(x[m-2 : m])
			sum2 += num
			x = x[:m-3]
		}

		var val int

		for i := 0; i < len(x); i++ {
			if x[i] >= '0' && x[i] <= '9' {
				val = val*10 + int(x[i]-'0')
			}
		}

		sum += val
	}

	sum += sum2 / 100
	sum2 %= 100

	var buf []byte
	var cnt int
	for sum > 0 {
		c := sum % 10
		buf = append(buf, byte(c+'0'))
		sum /= 10
		cnt++
		if cnt%3 == 0 && sum > 0 {
			buf = append(buf, '.')
		}
	}
	if len(buf) == 0 {
		buf = append(buf, '0')
	}

	slices.Reverse(buf)

	if sum2 > 0 {
		buf = append(buf, '.')
		buf = append(buf, byte(sum2/10+'0'))
		buf = append(buf, byte(sum2%10+'0'))
	}

	return string(buf)
}
