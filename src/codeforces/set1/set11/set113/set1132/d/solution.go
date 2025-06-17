package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(a, b, k)
}

func solve(a []int, b []int, k int) int {
	n := len(a)

	ques := make([]*Queue, k)
	for i := range k {
		ques[i] = new(Queue)
	}

	arr := make([]int, n)

	check := func(x int) bool {
		for i := range k {
			ques[i].Reset()
		}

		for i := range n {
			arr[i] = a[i] % b[i]
			t := a[i]/b[i] + 1
			if t < k {
				ques[t].Push(i)
			}
		}
		var last int
		for t := range k {
			for last < k && ques[last].Len() == 0 {
				last++
			}
			if last <= t {
				return false
			}
			if last == k {
				break
			}
			i := ques[last].Back()
			if arr[i]+x < b[i] {
				arr[i] += x
				continue
			}
			ques[last].Pop()
			nt := (arr[i] + x) / b[i]
			arr[i] = (arr[i] + x) % b[i]
			if last+nt < k {
				ques[last+nt].Push(i)
			}
		}

		return true
	}

	inf := slices.Max(b) * k

	if !check(inf) {
		return -1
	}

	return sort.Search(inf, check)
}

type Queue []int

func (q Queue) Len() int {
	return len(q)
}

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q Queue) Back() int {
	return q[len(q)-1]
}

func (q *Queue) Pop() int {
	n := len(*q)
	res := (*q)[n-1]
	*q = (*q)[:n-1]
	return res
}

func (q *Queue) Reset() {
	*q = (*q)[:0]
}
