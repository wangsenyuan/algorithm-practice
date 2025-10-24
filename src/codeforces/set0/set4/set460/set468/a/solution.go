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
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	if len(res) == 0 {
		fmt.Fprintln(writer, "NO")
		return
	}
	fmt.Fprintln(writer, "YES")
	for _, line := range res {
		fmt.Fprintln(writer, line)
	}
}

func solve(n int) []string {
	if n <= 3 {
		return nil
	}
	// n >= 4
	r := n % 4
	if r == 0 {
		return solve0(n)
	}
	if r == 1 {
		return solve1(n)
	}

	if r == 2 {
		return solve2(n)
	}
	return solve3(n)
}

func play(n int, x int) []string {
	var res []string
	for i := x; i < n; i += 4 {
		if i > x {
			res = append(res, fmt.Sprintf("0 + %d = %d", i, i))
		}
		res = append(res, fmt.Sprintf("%d + %d = %d", i, i+3, i+i+3))
		res = append(res, fmt.Sprintf("%d + %d = %d", i+1, i+2, i+i+3))
		res = append(res, fmt.Sprintf("%d - %d = 0", i+i+3, i+i+3))
	}
	if len(res) > 0 {
		res = append(res, "0 + 1 = 1")
	}
	return res
}

func solve0(n int) []string {
	res := play(n, 5)

	res = append(res, "1 * 2 = 2")
	res = append(res, "2 * 3 = 6")
	res = append(res, "6 * 4 = 24")
	return res
}

func solve1(n int) []string {
	res := play(n, 6)

	res = append(res, "4 * 5 = 20")
	res = append(res, "20 + 3 = 23")
	res = append(res, "23 + 2 = 25")
	res = append(res, "25 - 1 = 24")
	return res
}

func solve2(n int) []string {
	// 1...6
	res := play(n, 7)

	res = append(res, "4 * 6 = 24")
	res = append(res, "1 * 2 = 2")
	res = append(res, "2 + 3 = 5")
	res = append(res, "5 - 5 = 0")
	res = append(res, "24 + 0 = 24")
	return res
}

func solve3(n int) []string {
	res := play(n, 8)
	res = append(res, "3 * 7 = 21")
	res = append(res, "21 + 6 = 27")
	res = append(res, "27 - 5 = 22")
	res = append(res, "22 + 4 = 26")
	res = append(res, "26 - 2 = 24")
	res = append(res, "24 * 1 = 24")
	return res
}
