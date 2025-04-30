package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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

func process(reader *bufio.Reader) int {
	line := readNNums(reader, 4)
	mice := readNNums(reader, line[0])
	chess := readNNums(reader, line[1])
	return solve(mice, chess)
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 60

func solve(mice []int, chess []int) int {
	n := len(mice)
	m := len(chess)

	var res int

	dist := make([]int, m)
	for i := range dist {
		dist[i] = inf
	}

	for i, l := 0, 0; i < n; i++ {
		for l < m && chess[l] < mice[i] {
			l++
		}
		tmp := inf
		// l == m or chess[l] >= mice[i]
		if l < m {
			tmp = min(tmp, chess[l]-mice[i])
		}
		if l > 0 {
			tmp = min(tmp, mice[i]-chess[l-1])
		}

		j := l
		if l == m || chess[l]-mice[i] > tmp ||
			l > 0 && mice[i]-chess[l-1] == tmp && (dist[l-1] == inf || dist[l-1] == tmp) {
			j--
		}
		if dist[j] == inf || dist[j] == tmp {
			res++
		}
		dist[j] = min(dist[j], tmp)
	}

	return n - res
}

func abs(num int) int {
	return max(num, -num)
}
