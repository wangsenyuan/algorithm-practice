package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
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

func process(reader *bufio.Reader) []int {
	q := readNum(reader)
	queries := make([][]int, q)
	for i := range q {
		s, _ := reader.ReadBytes('\n')
		var x int
		pos := readInt(s, 0, &x)
		if x < 3 {
			queries[i] = []int{x}
		} else {
			var k int
			readInt(s, pos+1, &k)
			queries[i] = []int{x, k}
		}
	}
	return solve(queries)
}

func solve(queries [][]int) []int {
	ans := make([]int, len(queries))
	var n, s, ss, revS int
	var l, r []int
	for i, cur := range queries {
		op := cur[0]
		if op == 1 {
			var v int
			if len(r) > 0 {
				v = r[len(r)-1]
				r = r[:len(r)-1]
			} else {
				v = l[0]
				l = l[1:]
			}
			l = append(l, v)
			ss += s - v*n
			revS += v*n - s
		} else if op == 2 {
			l, r = r, l
			ss, revS = revS, ss
		} else {
			v := cur[1]
			r = append(r, v)
			n++
			ss += v * n
			s += v
			revS += s
		}
		ans[i] = ss
	}
	return ans
}

func solve1(queries [][]int) []int {
	m := len(queries)
	todo := NewTodo(2 * m)

	ans := make([]int, m)
	for i, cur := range queries {
		if cur[0] == 1 {
			todo.RightShift()
		} else if cur[0] == 2 {
			todo.Reverse()
		} else {
			todo.AddLast(cur[1])
		}
		ans[i] = todo.GetValue()
	}
	return ans
}

type Todo struct {
	sum     int
	lscore  int
	rscore  int
	que     []int
	front   int
	end     int
	reverse bool
}

func NewTodo(reserve int) *Todo {
	var sum int
	var lscore int
	var rscore int
	que := make([]int, 2*reserve)
	return &Todo{
		sum:     sum,
		lscore:  lscore,
		rscore:  rscore,
		que:     que,
		front:   reserve,
		end:     reserve - 1,
		reverse: false,
	}
}

func abs(num int) int {
	return max(num, -num)
}

func (t *Todo) AddLast(x int) {
	sz := t.end - t.front + 1
	if !t.reverse {
		t.end++
		t.que[t.end] = x
		t.rscore += t.sum + x
		t.sum += x
		t.lscore += x * (sz + 1)
	} else {
		t.front--
		t.que[t.front] = x
		t.lscore += t.sum + x
		t.sum += x
		t.rscore += x * (sz + 1)
	}
}

func (t *Todo) Reverse() {
	t.reverse = !t.reverse
}

func (t *Todo) GetValue() int {
	if !t.reverse {
		return t.lscore
	}
	return t.rscore
}

func (t *Todo) RightShift() {
	sz := t.end - t.front + 1
	if !t.reverse {
		t.lscore -= t.que[t.end] * sz
		t.lscore += t.sum
		t.rscore -= t.sum
		t.rscore += t.que[t.end] * sz
		x := t.que[t.end]
		t.end--
		t.front--
		t.que[t.front] = x
	} else {
		t.rscore -= t.que[t.front] * sz
		t.rscore += t.sum
		t.lscore -= t.sum
		t.lscore += t.que[t.front] * sz
		x := t.que[t.front]
		t.front++
		t.end++
		t.que[t.end] = x
	}
}
