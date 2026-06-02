package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if len(res) == 0 {
		fmt.Println()
	} else {
		var buf bytes.Buffer
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
		buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n, q := readTwoNums(reader)
	queries := make([][]int, q)
	for i := range queries {
		s, _ := reader.ReadBytes('\n')
		var pos, tp int
		pos = readInt(s, 0, &tp) + 1
		if tp == 1 {
			queries[i] = make([]int, 3)
		} else {
			queries[i] = make([]int, 2)
		}
		queries[i][0] = tp
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []int {

	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	var head int

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			p, x := cur[1]-1, cur[2]
			j := (head + p) % n
			arr[j] = x
		} else if cur[0] == 2 {
			p := cur[1] - 1
			j := (head + p) % n
			ans = append(ans, arr[j])
		} else {
			k := cur[1]
			head = (head + k) % n
		}
	}
	return ans
}
