package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) string {
	n := readNum(reader)
	a := make([]int, n)
	for i := range n {
		a[i] = readNum(reader)
	}
	return solve(a)
}

func solve(a []int) string {
	x := slices.Max(a)
	grunt := make([]int, x+1)
	for i := 1; i*(i+1)/2 <= x; i++ {
		grunt[i*(i+1)/2] = i
	}

	for i := 2; i <= x; i++ {
		if grunt[i] == 0 {
			grunt[i] = grunt[i-1]
		}
	}

	// 只要john有办法复制sam的操作，那么john就可以获胜
	var sum int
	for _, v := range a {
		sum ^= grunt[v]
	}

	if sum > 0 {
		return "NO"
	}
	return "YES"
}
