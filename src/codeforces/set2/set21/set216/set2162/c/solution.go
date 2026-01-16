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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		ok, res := solve(a, b)
		if !ok {
			fmt.Fprintln(writer, "-1")
		} else {
			fmt.Fprintln(writer, len(res))
			s := fmt.Sprintf("%v", res)
			fmt.Fprintln(writer, s[1:len(s)-1])
		}
	}
}

func solve(a int, b int) (bool, []int) {
	var res []int
	for i := range 63 {
		if a == b {
			break
		}

		if a&b == a {
			w := a ^ b
			if w > a {
				return false, nil
			}
			res = append(res, w)
			a ^= w
			break
		}

		x := (a >> i) & 1
		y := (b >> i) & 1
		if x != y {
			res = append(res, 1<<i)
			a ^= 1 << i
		}
	}

	return true, res
}
