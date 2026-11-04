package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var a, b, n int
	fmt.Fscan(reader, &a, &b, &n)
	return solve(a, b, n)
}

func solve(a int, b int, n int) string {
	w := a % b
	for d := range 10 {
		v := (w*10 + d) % b
		if v == 0 {
			return play(a, d, n-1)
		}
	}
	return "-1"
}

func play(a int, v int, n int) string {
	s := fmt.Sprintf("%d%d", a, v)
	buf := []byte(s)
	for n > 0 {
		buf = append(buf, '0')
		n--
	}
	return string(buf)
}
