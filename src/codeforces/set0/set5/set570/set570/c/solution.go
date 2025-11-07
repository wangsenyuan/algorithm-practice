package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	_, m := readTwoNums(reader)
	s := readString(reader)
	replaces := make([]string, m)
	for i := range m {
		replaces[i] = readString(reader)
	}
	return solve(s, replaces)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func solve(s string, replaces []string) []int {
	n := len(s)
	t1 := NewSegTree(n, -1, func(a int, b int) int {
		return max(a, b)
	})

	t2 := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})

	var sum, cnt int
	for i := range n {
		if s[i] != '.' {
			t1.Update(i, i)
			t2.Update(i, i)
			sum += max(0, cnt-1)
			cnt = 0
		} else {
			cnt++
		}
	}
	sum += max(0, cnt-1)

	buf := []byte(s)

	ans := make([]int, len(replaces))

	for i, cur := range replaces {
		var x int
		pos := readInt([]byte(cur), 0, &x)
		x--
		if buf[x] == cur[pos+1] || buf[x] != '.' && cur[pos+1] != '.' {
			buf[x] = cur[pos+1]
			ans[i] = sum
			continue
		}
		if buf[x] == '.' {
			l := t1.Find(0, x)
			r := t2.Find(x, n)
			sum -= max(0, r-l-2)
			sum += max(0, x-l-2)
			sum += max(0, r-x-2)
			t1.Update(x, x)
			t2.Update(x, x)
		} else {
			// 从c变成.
			t1.Update(x, -1)
			t2.Update(x, n)
			l := t1.Find(0, x)
			r := t2.Find(x, n)
			sum -= max(0, x-l-2)
			sum -= max(0, r-x-2)
			sum += max(0, r-l-2)
		}
		buf[x] = cur[pos+1]
		ans[i] = sum
	}

	return ans
}

type SegTree struct {
	arr       []int
	sz        int
	initValue int
	fn        func(int, int) int
}

func NewSegTree(n int, initValue int, fn func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = initValue
	}
	return &SegTree{
		arr:       arr,
		sz:        n,
		initValue: initValue,
		fn:        fn,
	}
}

func (tree *SegTree) Update(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] = v
	for pos > 0 {
		tree.arr[pos>>1] = tree.fn(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Find(l, r int) int {
	l += tree.sz
	r += tree.sz

	ans := tree.initValue

	for l < r {
		if l&1 == 1 {
			ans = tree.fn(ans, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = tree.fn(ans, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
