package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) int {
	n := len(a)

	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}

	slices.SortFunc(arr, func(x, y pair) int {
		return cmp.Or(y.first-x.first, x.second-y.second)
	})

	fa := make([]int, n)
	flag := make([][2]int, n)
	sz := make([]int, n)

	for i := range n {
		fa[i] = i
		sz[i] = 1
		if arr[i].first == arr[0].first {
			flag[arr[i].second][0] = 1
		}
	}
	var now, cnt int

	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	check := func(u int) int {
		if flag[u][0]+flag[u][1] == 0 {
			return 0
		}
		if sz[u]&1 == 1 {
			// flag[u][0] 表示u所在的component是否能取到这个component的最大值？
			return flag[u][0]
		}
		return 1
	}

	union := func(x, y int) {
		x = find(x)
		y = find(y)
		if x == y {
			return
		}
		if sz[x]&1 == 1 && sz[y]&1 == 1 {
			now--
		}

		cnt -= check(x) + check(y)

		fa[x] = y

		// 这个flag我没有理解
		flag[y][0] |= flag[x][sz[y]&1]
		flag[y][1] |= flag[x][sz[y]&1^1]

		sz[y] += sz[x]

		cnt += check(y)
	}

	pos := make([]int, n)
	for i := range n {
		pos[i] = n
	}

	var ans int
	for i := range n {
		x := arr[i].second
		pos[x] = i
		now++
		cnt += flag[x][0]
		if x > 0 && pos[x-1] < i {
			union(x-1, x)
		}
		if x+1 < n && pos[x+1] < i {
			union(x, x+1)
		}
		tmp := arr[0].first + arr[i].first + now
		if cnt == 0 {
			tmp--
		}
		ans = max(ans, tmp)
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}
