package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []string {
	n := readNum(reader)
	bombs := make([][]int, n)
	for i := 0; i < n; i++ {
		bombs[i] = readNNums(reader, 2)
	}
	return solve(bombs)
}

func solve(bombs [][]int) []string {
	slices.SortFunc(bombs, func(a, b []int) int {
		return a[0]*a[0] + a[1]*a[1] - b[0]*b[0] - b[1]*b[1]
	})
	var ans []string

	play := func(x int, y int) {
		if x != 0 {
			ans = append(ans, fmt.Sprintf("1 %d %c", abs(x), horizontal(x)))
		}
		if y != 0 {
			ans = append(ans, fmt.Sprintf("1 %d %c", abs(y), vertical(y)))
		}
		ans = append(ans, "2")
		if y != 0 {
			ans = append(ans, fmt.Sprintf("1 %d %c", abs(y), vertical(-y)))
		}
		if x != 0 {
			ans = append(ans, fmt.Sprintf("1 %d %c", abs(x), horizontal(-x)))
		}
		ans = append(ans, "3")
	}

	for _, cur := range bombs {
		play(cur[0], cur[1])
	}

	return ans
}

func vertical(y int) byte {
	if y < 0 {
		return 'D'
	}
	return 'U'
}

func horizontal(x int) byte {
	if x < 0 {
		return 'L'
	}
	return 'R'
}

func abs(num int) int {
	return max(num, -num)
}
