package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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
	nums := readNNums(reader, 4)
	return solve(nums[0], nums[1], nums[2], nums[3])
}

func solve(l1 int, r1 int, l2 int, r2 int) int {

	var dfs func(L int, R int, l int, r int) []int
	dfs = func(L int, R int, l int, r int) []int {
		if L <= l && r <= R {
			return []int{r - l}
		}
		mid := (l + r) / 2
		if R <= mid {
			return dfs(L, R, l, mid)
		}
		if mid <= L {
			return dfs(L, R, mid, r)
		}
		res := dfs(L, mid, l, mid)
		res2 := dfs(mid, R, mid, r)
		return append(res, res2...)
	}

	x := dfs(l1, r1, 0, 1<<25)
	y := dfs(l2, r2, 0, 1<<25)

	sort.Ints(x)
	sort.Ints(y)

	var sum int
	for _, u := range x {
		sum += u
	}

	var ans int

	for i, j := 0, 0; i < len(y); i++ {
		for j < len(x) && x[j] < y[i] {
			sum -= x[j]
			j++
		}
		ans += sum / y[i]
	}
	sum = 0
	for _, v := range y {
		sum += v
	}

	for i, j := 0, 0; i < len(x); i++ {
		for j < len(y) && y[j] <= x[i] {
			sum -= y[j]
			j++
		}
		ans += sum / x[i]
	}

	return ans
}
