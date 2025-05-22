package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, s := range res {
		buf.WriteString(fmt.Sprintf("%d ", s))
	}
	buf.WriteByte('\n')

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
	n, m := readTwoNums(reader)
	instructions := make([][]int, m)
	for i := range m {
		instructions[i] = readNNums(reader, 2)
	}
	return solve(n, instructions)
}

const inf = 1 << 60

func solve(n int, instructions [][]int) []int {
	m := len(instructions)

	at := make([][]pair, n)

	diff := make([]int, n+1)

	for i, cur := range instructions {
		l, r := cur[0]-1, cur[0]+cur[1]-2
		at[r] = append(at[r], pair{l, i})
		diff[l]++
		diff[r+1]--
	}

	dp := NewSegTree(n)
	fp := make([]pair, n)
	for i := range n {
		if i > 0 {
			diff[i] += diff[i-1]
		}
		if diff[i] == 0 {
			if i == 0 {
				dp.Update(0, 0)
			} else {
				dp.Update(i, dp.Get(i-1, i).first)
			}
		} else {
			val := pair{inf, -1}
			for _, cur := range at[i] {
				l := cur.first
				if l == 0 {
					fp[i] = pair{-1, cur.second}
					dp.Update(i, 1)
					break
				} else {
					// 需要找到l-1...i的最小值
					tmp := dp.Get(l-1, i)
					if min_pair(val, tmp) != val {
						val = tmp
						fp[i] = pair{tmp.second, cur.second}
					}
				}
			}

			if val.first < inf {
				val.first++
				dp.Update(i, val.first)
			}
		}
	}

	ok := make([]bool, m)

	for i := n - 1; i >= 0; {
		if diff[i] == 0 {
			i--
			continue
		}

		ok[fp[i].second] = true
		i = fp[i].first
	}

	var res []int
	for i := range m {
		if !ok[i] {
			res = append(res, i+1)
		}
	}

	return res
}

type pair struct {
	first  int
	second int
}

func min_pair(a, b pair) pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

type SegTree []pair

func NewSegTree(n int) SegTree {
	arr := make([]pair, 2*n)
	for i := n; i < len(arr); i++ {
		arr[i] = pair{inf, i - n}
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = min_pair(arr[i*2], arr[i*2+1])
	}
	return arr
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p].first = v

	for p > 1 {
		tr[p>>1] = min_pair(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) pair {
	n := len(tr) / 2
	l += n
	r += n
	rl := pair{inf, -1}
	rr := pair{inf, -1}
	for l < r {
		if l&1 == 1 {
			rl = min_pair(tr[l], rl)
			l++
		}
		if r&1 == 1 {
			r--
			rr = min_pair(rr, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return min_pair(rl, rr)
}
