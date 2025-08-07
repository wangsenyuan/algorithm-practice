package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

const inf = 1 << 30

func solve(a []int) []int {
	n := len(a)
	left := make([]int, n+1)
	for _, v := range a {
		left[v]++
	}
	stack := make([]int, n)
	var top int
	marked := make([]bool, n+1)
	for _, v := range a {
		left[v]--
		if marked[v] {
			continue
		}
		for top > 0 && left[stack[top-1]] > 0 {
			x := stack[top-1]
			if top%2 == 1 {
				// x是在奇数位的，如果把x去掉后，v就要放到奇数位，如果v < x, 那么就需要把v放在x的前面的位置
				// 但是如果无法放置，比如只有一个，或者倒数第二位更小，或者它后续无法再被选中了
				if v < x && (top == 1 || v > stack[top-2] || left[stack[top-2]] == 0) {
					break
				}
			} else {
				if v > x && (top == 1 || v < stack[top-2] || left[stack[top-2]] == 0) {
					break
				}
			}
			top--
			marked[x] = false
		}
		stack[top] = v
		top++
		marked[v] = true
	}
	return stack[:top]
}

func solve1(a []int) []int {
	n := len(a)
	pos := NewSegTree(n+1, inf, min)
	x := NewSegTree(n, -inf, max)
	y := NewSegTree(n, inf, min)

	at := make([][]int, n+1)

	for i, num := range a {
		pos.Update(num, i)
		x.Update(i, num)
		y.Update(i, num)
		at[num] = append(at[num], i)
	}

	var arr []int

	for i := 0; i < n; i++ {
		ni := pos.Get(1, n+1)
		if ni == inf {
			break
		}
		// 可以在获得区间[i...ni]的任何数，且保证能勾获取到最长的序列
		var v int
		if len(arr)%2 == 0 {
			// 获取最大值
			v = x.Get(i, ni+1)
		} else {
			v = y.Get(i, ni+1)
		}

		arr = append(arr, v)
		pos.Update(v, inf)
		found := false
		for _, j := range at[v] {
			x.Update(j, -inf)
			y.Update(j, inf)
			if !found && j >= i {
				i = j
				found = true
			}
		}
	}

	return arr
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
