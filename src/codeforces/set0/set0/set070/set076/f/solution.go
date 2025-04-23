package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0], res[1])
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
	n := readNum(reader)
	events := make([][]int, n)
	for i := range n {
		events[i] = readNNums(reader, 2)
	}
	v := readNum(reader)
	return solve(v, events)
}

const inf = 1 << 60

type state struct {
	first  int
	second int
	x      int
}

func solve(v int, events [][]int) []int {
	var arr, arr2 []state

	var found bool
	for _, e := range events {
		x, t := e[0], e[1]
		arr = append(arr, state{-x + v*t, x + v*t, x})
		arr2 = append(arr2, state{-x + v*t, x + v*t, x})
		if x == 0 {
			found = true
		}
	}
	if !found {
		arr = append(arr, state{0, 0, 0})
	}

	slices.SortFunc(arr, func(a, b state) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	for i := 0; i < len(arr); i++ {
		if arr[i].x == 0 {
			arr = arr[i:]
			break
		}
	}

	slices.SortFunc(arr2, func(a, b state) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	que := make([]state, len(arr))
	que[0] = arr[0]
	pos := 1
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		j := sort.Search(pos, func(j int) bool {
			return cur.second < que[j].second
		})
		if j == 0 {
			// 不能替换头
			continue
		}
		que[j] = arr[i]
		if j == pos {
			pos++
		}
	}
	ans := []int{0, 0}
	ans[0] = pos
	if que[0].first == 0 && que[0].second == 0 {
		// 增加的那个点
		ans[0]--
	}
	que = make([]state, len(arr2))
	pos = 0
	for _, cur := range arr2 {
		j := sort.Search(pos, func(j int) bool {
			return cur.second < que[j].second
		})
		que[j] = cur
		if j == pos {
			pos++
		}
	}

	ans[1] = pos

	return ans
}
