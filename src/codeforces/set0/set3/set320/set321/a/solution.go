package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) bool {
	a, b := readTwoNums(reader)
	s := readString(reader)
	return solve(a, b, s)
}

const inf = 1 << 60

func solve(a int, b int, s string) bool {
	if a < 0 {
		s = replace(s, 'L', 'R')
		a *= -1
	}
	if b < 0 {
		s = replace(s, 'D', 'U')
		b *= -1
	}

	// a >= 0, b >= 0
	var x, y int
	n := len(s)

	for range n {
		if a == 0 && b == 0 {
			return true
		}

		for i := range n {
			switch s[i] {
			case 'R':
				x++
			case 'L':
				x--
			case 'U':
				y++
			default:
				y--
			}
			if x == a && y == b {
				return true
			}
		}
	}

	if x < 0 || y < 0 || x+y == 0 {
		return false
	}
	x /= n
	y /= n

	c := 1 << 60
	if x > 0 {
		c = min(c, a/x)
	}
	if y > 0 {
		c = min(c, b/y)
	}
	if c < n {
		return false
	}

	x *= (c - n)
	y *= (c - n)

	for range 2 * n {
		if x == a && y == b {
			return true
		}

		for i := range n {
			switch s[i] {
			case 'R':
				x++
			case 'L':
				x--
			case 'U':
				y++
			default:
				y--
			}
			if x == a && y == b {
				return true
			}
		}
	}

	return false
}

func replace(s string, a byte, b byte) string {
	buf := []byte(s)
	for i := range buf {
		if buf[i] == a {
			buf[i] = b
		} else if buf[i] == b {
			buf[i] = a
		}
	}
	return string(buf)
}
