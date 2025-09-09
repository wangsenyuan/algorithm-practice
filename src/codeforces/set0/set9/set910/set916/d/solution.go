package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	drive(reader, writer)
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

func drive(reader *bufio.Reader, writer *bufio.Writer) {
	n := readNum(reader)
	solve(n, reader, writer)
}

const inf = 1e9

func solve(n int, reader *bufio.Reader, writer *bufio.Writer) {

	scores := make([]int, n+1)
	tr2 := NewTree2()
	ids := make(map[string]int)

	add := func(name string) int {
		if _, ok := ids[name]; !ok {
			ids[name] = len(ids)
		}
		return ids[name]
	}

	trs := make([]*Tree, n+1)
	trs[0] = new(Tree)

	for i := range n {
		op := readString(reader)
		s := strings.Split(op, " ")
		if strings.HasPrefix(op, "set") {
			j := add(s[1])
			w := tr2.Get(scores[i], 0, n-1, j)
			tmp_tr := trs[i]
			tmp_score := scores[i]

			if w > 0 {
				tmp_tr = tmp_tr.Set(0, inf, w, -1)
				tmp_score = tr2.Set(tmp_score, 0, n-1, j, 0)
			}
			v, _ := strconv.Atoi(s[2])

			trs[i+1] = tmp_tr.Set(0, inf, v, 1)
			scores[i+1] = tr2.Set(tmp_score, 0, n-1, j, v)
		} else if strings.HasPrefix(op, "remove") {
			if j, ok := ids[s[1]]; ok {
				v := tr2.Get(scores[i], 0, n-1, j)
				if v != 0 {
					// 之前设置了分数
					trs[i+1] = trs[i].Set(0, inf, v, -1)
					scores[i+1] = tr2.Set(scores[i], 0, n-1, j, 0)
					continue
				}
			}
			scores[i+1] = scores[i]
			trs[i+1] = trs[i]
		} else if strings.HasPrefix(op, "query") {
			if j, ok := ids[s[1]]; ok {
				v := tr2.Get(scores[i], 0, n-1, j)
				if v == 0 {
					writer.WriteString("-1\n")
				} else {
					ans := trs[i].FindPosition(0, inf, v)
					writer.WriteString(fmt.Sprintf("%d\n", ans))
				}
			} else {
				writer.WriteString("-1\n")
			}
			writer.Flush()
			scores[i+1] = scores[i]
			trs[i+1] = trs[i]
		} else {
			// undo
			days, _ := strconv.Atoi(s[1])
			trs[i+1] = trs[i-days]
			scores[i+1] = scores[i-days]
		}
	}
}

type Tree struct {
	left, right *Tree
	cnt         int
}

func (tr *Tree) clone() *Tree {
	if tr == nil {
		return new(Tree)
	}
	return &Tree{
		left:  tr.left,
		right: tr.right,
		cnt:   tr.cnt,
	}
}

func (tr *Tree) Count() int {
	if tr == nil {
		return 0
	}
	return tr.cnt
}

func (tr *Tree) Set(l int, r int, p int, v int) *Tree {
	res := tr.clone()
	if l == r {
		res.cnt += v
	} else {
		mid := (l + r) >> 1
		if p <= mid {
			res.left = res.left.Set(l, mid, p, v)
		} else {
			res.right = res.right.Set(mid+1, r, p, v)
		}
		res.cnt = res.left.Count() + res.right.Count()
	}
	return res
}

func (tr *Tree) FindPosition(l int, r int, p int) int {
	if tr == nil || l == r {
		return 0
	}
	mid := (l + r) >> 1
	if p <= mid {
		return tr.left.FindPosition(l, mid, p)
	}
	return tr.left.Count() + tr.right.FindPosition(mid+1, r, p)
}

type Tree2 struct {
	left  []int
	right []int
	val   []int
}

func NewTree2() *Tree2 {
	return &Tree2{
		left:  make([]int, 1),
		right: make([]int, 1),
		val:   make([]int, 1),
	}
}

func (t *Tree2) expand() int {
	t.left = append(t.left, 0)
	t.right = append(t.right, 0)
	t.val = append(t.val, 0)
	return len(t.left) - 1
}

func (t *Tree2) clone(i int) int {
	ni := t.expand()
	t.left[ni] = t.left[i]
	t.right[ni] = t.right[i]
	t.val[ni] = t.val[i]
	return ni
}

func (t *Tree2) Set(i int, l int, r int, p int, v int) int {
	ni := t.clone(i)
	if l == r {
		t.val[ni] = v
	} else {
		mid := (l + r) >> 1
		if p <= mid {
			t.left[ni] = t.Set(t.left[ni], l, mid, p, v)
		} else {
			t.right[ni] = t.Set(t.right[ni], mid+1, r, p, v)
		}
	}
	return ni
}

func (t *Tree2) Get(i int, l int, r int, p int) int {
	if i == 0 {
		return 0
	}
	if l == r {
		return t.val[i]
	}
	mid := (l + r) >> 1
	if p <= mid {
		return t.Get(t.left[i], l, mid, p)
	}
	return t.Get(t.right[i], mid+1, r, p)
}
