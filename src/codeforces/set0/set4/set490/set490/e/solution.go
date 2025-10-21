package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Fprintln(writer, "NO")
		return
	}
	fmt.Fprintln(writer, "YES")
	for _, cur := range res {
		fmt.Fprintln(writer, cur)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func drive(reader *bufio.Reader) (a []string, res []string) {
	n := readNum(reader)
	a = make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	res = solve(n, a)
	return
}

func solve(n int, a []string) []string {

	play := func(x string, y string) string {
		if len(x) > len(y) || len(y) > 1 && y[0] == '0' {
			return ""
		}
		buf := []byte(y)
		// len(x) <= len(y)
		if len(x) < len(y) {
			if buf[0] == '?' {
				buf[0] = '1'
			}
			// 后面的？都可以使用0
			for i := 1; i < len(buf); i++ {
				if buf[i] == '?' {
					buf[i] = '0'
				}
			}
			return string(buf)
		}
		// len(x) == len(y)
		eq := true
		for i := 0; i < len(x); i++ {
			if y[i] != '?' {
				if eq && x[i] > y[i] {
					return ""
				}
				if x[i] < y[i] {
					eq = false
				}
				continue
			}
			if !eq {
				buf[i] = '0'
				continue
			}
			// eq
			if x[i] == '9' {
				buf[i] = x[i]
				continue
			}
			// 如果buf[i] == x[i]， 那么需要后面能够比x大的数
			ok := false
			for j := i + 1; j < len(x); j++ {
				if x[j] == y[j] || x[j] == '9' && y[j] == '?' {
					// y[j] = ? must be 9
					continue
				}
				// x[j] != y[j]
				if y[j] == '?' || x[j] < y[j] {
					ok = true
				}
				break
			}
			if ok {
				buf[i] = x[i]
				continue
			}
			buf[i] = x[i] + 1
			eq = false
		}

		if eq {
			return ""
		}
		return string(buf)
	}

	ans := make([]string, len(a))
	for i, cur := range a {
		if i == 0 {
			ans[i] = play("", cur)
		} else {
			ans[i] = play(ans[i-1], cur)
		}
		if len(ans[i]) == 0 {
			return nil
		}
	}
	return ans
}
