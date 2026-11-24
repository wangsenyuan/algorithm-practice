package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintln(writer, drive(reader))
}

func drive(reader *bufio.Reader) string {
	var m, s int
	fmt.Fscan(reader, &m, &s)
	return solve(m, s)
}

func solve(m, s int) string {
	return fmt.Sprintf("%s %s", smallest(m, s), largest(m, s))
}

func smallest(m int, s int) string {
	if s == 0 {
		if m == 1 {
			return "0"
		}
		return "-1"
	}
	// 1...99
	buf := make([]int, m)
	for d := 1; d <= 9; d++ {
		if d+(m-1)*9 >= s {
			buf[0] = d
			break
		}
	}
	s -= buf[0]
	// d0...99
	for i := m - 1; i > 0 && s > 0; i-- {
		x := min(s, 9)
		buf[i] = x
		s -= x
	}
	if s != 0 {
		return "-1"
	}
	var ans []byte
	for i := range m {
		ans = append(ans, byte(buf[i]+'0'))
	}
	return string(ans)
}

func largest(m int, s int) string {
	if s == 0 && m > 1 {
		return "-1"
	}
	// 999900
	buf := make([]int, m)
	for i := range m {
		buf[i] = min(9, s)
		s -= buf[i]
		if s == 0 {
			break
		}
	}
	if s != 0 {
		return "-1"
	}
	var ans []byte
	for i := range m {
		ans = append(ans, byte(buf[i]+'0'))
	}
	return string(ans)
}
