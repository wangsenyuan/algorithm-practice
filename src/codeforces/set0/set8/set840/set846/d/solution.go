package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, k, q int
	fmt.Fscan(reader, &n, &m, &k, &q)
	events := make([][]int, q)
	for i := range q {
		events[i] = make([]int, 3)
		fmt.Fscan(reader, &events[i][0], &events[i][1], &events[i][2])
	}
	return solve(n, m, k, events)
}

func solve(n int, m int, k int, events [][]int) int {
	tv := make([][]int, n)
	for i := range n {
		tv[i] = make([]int, m)
		for j := range m {
			tv[i][j] = -1
		}
	}

	for _, cur := range events {
		r, c, t := cur[0], cur[1], cur[2]
		tv[r-1][c-1] = t
	}

	sum := make([][]int, n+1)
	for i := range n + 1 {
		sum[i] = make([]int, m+1)
	}

	for i := range n {
		for j := range m {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j]
			if tv[i][j] != -1 {
				sum[i+1][j+1]++
			}
		}
	}

	cols := make([]*Queue, m)
	for i := range m {
		cols[i] = NewQueue(n)
	}

	for i := range k - 1 {
		for j := range m {
			if tv[i][j] >= 0 {
				cols[j].Add(i, tv[i][j])
			}
		}
	}
	row := NewQueue(m)
	ans := 1 << 60

	getSum := func(i int, j int) int {
		return sum[i+1][j+1] - sum[i+1][j-k+1] - sum[i-k+1][j+1] + sum[i-k+1][j-k+1]
	}

	for i := k - 1; i < n; i++ {
		row.Reset()
		for j := range m {
			if tv[i][j] >= 0 {
				cols[j].Add(i, tv[i][j])
			}
			if i >= k {
				cols[j].RemoveTailIfAt(i - k)
			}
			if cols[j].tail < cols[j].head {
				row.Add(j, cols[j].GetTail().first)
			}
			if j >= k {
				row.RemoveTailIfAt(j - k)
			}
			if j >= k-1 && getSum(i, j) == k*k {
				ans = min(ans, row.GetTail().first)
			}
		}
	}

	if ans < 1<<60 {
		return ans
	}
	return -1
}

type pair struct {
	first  int
	second int
}
type Queue struct {
	arr        []pair
	head, tail int
}

func NewQueue(n int) *Queue {
	return &Queue{
		arr:  make([]pair, n),
		head: 0,
		tail: 0,
	}
}

func (q *Queue) Add(p int, v int) {
	for q.head > q.tail && q.arr[q.head-1].first <= v {
		q.head--
	}
	q.arr[q.head] = pair{v, p}
	q.head++
}

func (q *Queue) RemoveTailIfAt(p int) {
	if q.tail < q.head && q.arr[q.tail].second == p {
		q.tail++
	}
}

func (q *Queue) GetTail() pair {
	return q.arr[q.tail]
}

func (q *Queue) Reset() {
	q.head = 0
	q.tail = 0
}
