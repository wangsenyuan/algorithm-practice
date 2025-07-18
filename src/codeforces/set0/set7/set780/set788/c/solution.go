package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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
	n, k := readTwoNums(reader)
	a := readNNums(reader, k)
	return solve(n, a)
}

const X = 1000

type pair struct {
	first  int
	second int
}

func solve(n int, a []int) int {
	sort.Ints(a)
	a = slices.Compact(a)
	k := len(a)
	var que []pair
	vis := make(map[int]bool)

	for i := 0; i < k; i++ {
		que = append(que, pair{a[i] - n, 1})
		vis[a[i]-n] = true
	}

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]

		if cur.first == 0 {
			return cur.second
		}

		for _, v := range a {
			if cur.first+v-n >= -1000000 && cur.first+v-n <= 1000000 && !vis[cur.first+v-n] {
				vis[cur.first+v-n] = true
				que = append(que, pair{cur.first + v - n, cur.second + 1})
			}
		}
	}

	return -1
}
