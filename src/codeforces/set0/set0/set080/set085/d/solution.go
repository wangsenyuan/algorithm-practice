package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	ops := make([]string, n)
	for i := range n {
		ops[i] = readString(reader)
	}
	return solve(ops)
}

func solve(ops []string) []int {
	var arr []int
	for _, op := range ops {
		if op[0] == 'a' {
			var x int
			readInt([]byte(op), 4, &x)
			arr = append(arr, x)
		}
	}
	sort.Ints(arr)
	// 还要去重
	var n int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] != arr[i-1] {
			arr[n] = arr[i-1]
			n++
		}
	}
	arr = arr[:n]
	set := NewTree(n)

	var ans []int

	for _, op := range ops {
		if op[0] == 'a' {
			var x int
			readInt([]byte(op), 4, &x)
			i := sort.SearchInts(arr, x)
			set.Update(i, x, true)
		} else if op[0] == 'd' {
			var x int
			readInt([]byte(op), 4, &x)
			i := sort.SearchInts(arr, x)
			set.Update(i, x, false)
		} else {
			if n == 0 || set.cnt[0] < 3 {
				ans = append(ans, 0)
			} else {
				ans = append(ans, set.sum[0][2])
			}
		}
	}
	return ans
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

const inf = 1 << 60

type Tree struct {
	sum [][]int
	cnt []int
	sz  int
}

func NewTree(n int) *Tree {
	sum := make([][]int, 4*n)
	cnt := make([]int, 4*n)
	for i := range 4 * n {
		sum[i] = make([]int, 5)
	}
	return &Tree{sum, cnt, n}
}

func (t *Tree) pull(i int) {
	l := 2*i + 1
	r := l + 1
	u := t.cnt[l] % 5
	copy(t.sum[i], t.sum[l])
	for v := range 5 {
		t.sum[i][(u+v)%5] += t.sum[r][v]
	}
	t.cnt[i] = t.cnt[l] + t.cnt[r]
}

func (t *Tree) Update(p int, v int, add bool) {

	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			if add {
				t.sum[i][0] = v
				t.cnt[i] = 1
			} else {
				t.sum[i][0] = 0
				t.cnt[i] = 0
			}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		t.pull(i)
	}

	loop(0, 0, t.sz-1)
}
