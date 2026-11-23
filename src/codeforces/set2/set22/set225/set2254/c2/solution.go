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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	var a, b string
	fmt.Fscan(reader, &n, &a, &b)
	return solve(a, b)
}

func solve(a string, b string) int {
	var s1, s2 []byte
	var s3, s4 []byte
	for i := range a {
		if i&1 == 0 {
			s1 = append(s1, a[i])
			s2 = append(s2, b[i])
		} else {
			s3 = append(s3, a[i])
			s4 = append(s4, b[i])
		}
	}
	ans := play(string(s1), string(s2))
	if ans < 0 {
		return -1
	}
	ans2 := play(string(s3), string(s4))
	if ans2 < 0 {
		return -1
	}
	return ans + ans2
}

func play(a string, b string) int {
	// 子串内部只能相邻交换。1 的个数不同就无解。
	// 否则按从左到右把 a 的每个 1 送到 b 对应的 1，距离之和就是最少次数。
	var p, q []int
	for i := range a {
		if a[i] == '1' {
			p = append(p, i)
		}
		if b[i] == '1' {
			q = append(q, i)
		}
	}
	if len(p) != len(q) {
		return -1
	}
	var ans int
	for i := range p {
		d := p[i] - q[i]
		if d < 0 {
			d = -d
		}
		ans += d
	}
	return ans
}
