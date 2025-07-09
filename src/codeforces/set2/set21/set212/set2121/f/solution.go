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
	n, s, x := readThreeNums(reader)
	a := readNNums(reader, n)
	return solve(s, x, a)
}

func solve(s int, x int, a []int) int {
	n := len(a)

	dq1 := NewDque(n)
	dq2 := NewDque(n)

	freq1 := make(map[int]int)
	freq2 := make(map[int]int)
	freq1[0]++
	freq2[0]++
	pref := make([]int, n+1)
	var l1, l2 int
	var res int
	for i := range n {
		pref[i+1] = pref[i] + a[i]

		cur := pair{a[i], i}

		for !dq1.Empty() && dq1.PeekRight().first < cur.first {
			dq1.PopRight()
		}
		dq1.Push(cur)

		for !dq1.Empty() && dq1.PeekLeft().first > x {
			if dq1.PeekLeft().second == l1 {
				dq1.PopLeft()
			}
			freq1[pref[l1]]--
			l1++
		}

		for !dq2.Empty() && dq2.PeekRight().first < cur.first {
			dq2.PopRight()
		}

		dq2.Push(cur)

		for !dq2.Empty() && dq2.PeekLeft().first >= x {
			if dq2.PeekLeft().second == l2 {
				dq2.PopLeft()
			}
			freq2[pref[l2]]--
			l2++
		}

		res += freq1[pref[i+1]-s] - freq2[pref[i+1]-s]

		freq1[pref[i+1]]++
		freq2[pref[i+1]]++
	}

	return res
}

type pair struct {
	first  int
	second int
}

type Dque struct {
	arr  []pair
	head int
	tail int
}

func NewDque(n int) *Dque {
	arr := make([]pair, n)
	return &Dque{
		arr:  arr,
		head: 0,
		tail: 0,
	}
}

func (dq *Dque) Push(v pair) {
	dq.arr[dq.head] = v
	dq.head++
}

func (dq Dque) PeekRight() pair {
	return dq.arr[dq.head-1]
}

func (dq *Dque) PopRight() pair {
	v := dq.arr[dq.head-1]
	dq.head--
	return v
}

func (dq *Dque) PopLeft() pair {
	v := dq.arr[dq.tail]
	dq.tail++
	return v
}

func (dq *Dque) PeekLeft() pair {
	return dq.arr[dq.tail]
}

func (dq *Dque) Empty() bool {
	return dq.head == dq.tail
}
