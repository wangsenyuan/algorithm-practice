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

	var t int
	fmt.Fscan(reader, &t)
	for range t {
		ans, _ := drive(reader)
		if len(ans) == 0 {
			fmt.Fprintln(writer, "NO")
		} else {
			fmt.Fprintln(writer, "YES")
			fmt.Fprintln(writer, len(ans[0]))
			for _, v := range ans {
				fmt.Fprintln(writer, v)
			}
		}
	}
}

func drive(reader *bufio.Reader) (res []string, k int) {
	fmt.Fscan(reader, &k)
	res = solve(k)
	return
}

func solve(k int) []string {
	if k%5 == 2 || k%5 == 4 {
		return nil
	}
	buf := make([][]byte, 2)
	if k%5 > 0 {
		// 1 or 3
		// ...
		// *..
		m := k / 5
		for range m {
			buf[0] = append(buf[0], '.', '.', '.')
			buf[1] = append(buf[1], '*', '.', '.')
		}
		buf[0] = append(buf[0], '.')
		buf[1] = append(buf[1], '*')
		if k%5 == 3 {
			buf[0] = append(buf[0], '.')
			buf[1] = append(buf[1], '.')
		}
	} else {
		// 5
		m := k / 5
		for range m {
			buf[0] = append(buf[0], '.', '.', '.')
			buf[1] = append(buf[1], '.', '*', '.')
		}
	}
	ans := make([]string, 2)

	for i := range 2 {
		ans[i] = string(buf[i])
	}
	return ans
}
