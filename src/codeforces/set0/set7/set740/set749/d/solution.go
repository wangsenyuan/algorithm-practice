package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", ans[0], ans[1]))
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

func process(reader *bufio.Reader) [][]int {
	n := readNum(reader)
	bids := make([][]int, n)
	for i := range n {
		bids[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	queries := make([][]int, m)
	for i := range m {
		var k, pos int
		s, _ := reader.ReadBytes('\n')
		pos = readInt(s, 0, &k) + 1
		queries[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}
	return solve(bids, queries)
}

func solve(bids [][]int, queries [][]int) [][]int {
	n := len(bids)
	pos := make([][]int, n)
	last := make([]int, n)
	id := make([]int, n)

	for i := range n {
		id[i] = i
		last[i] = -1
	}
	for i, bid := range bids {
		x := bid[0] - 1
		pos[x] = append(pos[x], i)
		last[x] = i
	}

	// 越晚结束的，越早
	slices.SortFunc(id, func(a int, b int) int {
		return cmp.Or(last[b]-last[a], a-b)
	})

	ans := make([][]int, len(queries))

	for i, cur := range queries {
		for j := range cur {
			cur[j]--
		}
		slices.SortFunc(cur, func(a int, b int) int {
			// make result stable
			return cmp.Or(last[b]-last[a], a-b)
		})
		var first int
		for first < len(cur) && first < n && cur[first] == id[first] {
			first++
		}
		if first == n || last[id[first]] == -1 {
			ans[i] = []int{0, 0}
			continue
		}
		// u是没有弃权的，最后bid的人
		u := id[first]
		second := first + 1
		for second-1 < len(cur) && second < n && cur[second-1] == id[second] {
			second++
		}
		if second == n || last[id[second]] == -1 {
			ans[i] = bids[pos[u][0]]
		} else {
			j := sort.Search(len(pos[u]), func(j int) bool {
				return pos[u][j] > last[id[second]]
			})
			ans[i] = bids[pos[u][j]]
		}
	}

	return ans
}

func solve1(bids [][]int, queries [][]int) [][]int {
	n := len(bids)
	pos := make([][]int, n+1)
	var arr []int
	for i, bid := range bids {
		x := bid[0]
		pos[x] = append(pos[x], i)
		arr = append(arr, x)
	}
	sort.Ints(arr)
	arr = slices.Compact(arr)

	check := func(mid int, absent []int) bool {
		var cnt int
		w := bids[mid][0]
		for _, x := range absent {
			if w == x {
				w = -1
			}
			j := sort.SearchInts(pos[x], mid+1)
			// pos[x][j] >= mid
			cnt += len(pos[x]) - j
		}
		// 还要把mid位置所在的出价者的，也计算在内
		if w >= 0 {
			j := sort.SearchInts(pos[w], mid+1)
			cnt += len(pos[w]) - j
		}
		return cnt == n-mid-1
	}

	ans := make([][]int, len(queries))

	for i, absent := range queries {
		if len(absent) >= len(arr) && contains(absent, arr) {
			ans[i] = []int{0, 0}
			continue
		}
		j := sort.Search(n, func(mid int) bool {
			return check(mid, absent)
		})
		x := bids[j][0]
		// 还要对pos[x]进行二分
		u := sort.Search(len(pos[x]), func(u int) bool {
			return check(pos[x][u], absent)
		})
		ans[i] = bids[pos[x][u]]
	}
	return ans
}

func contains(a []int, b []int) bool {
	for i, j := 0, 0; i < len(b); i++ {
		for j < len(a) && a[j] != b[i] {
			j++
		}
		if j == len(a) {
			return false
		}
		j++
	}
	return true
}
