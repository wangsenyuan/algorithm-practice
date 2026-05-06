package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

const N = 200010

var divisors [N][]int

func init() {
	for d := 2; d < N; d++ {
		for x := d; x < N; x += d {
			divisors[x] = append(divisors[x], d)
		}
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	n := len(a)
	res := make([]int, n)

	freq := make([]int, n+1)
	buckets := NewBuckets(n + 1)
	incr := func(v int) {
		if freq[v] > 0 {
			buckets.Add(freq[v], -1)
		}
		freq[v]++
		buckets.Add(freq[v], 1)
	}

	for i, num := range a {
		for _, d := range divisors[num] {
			incr(d)
		}

		if i > 0 {
			res[i] = buckets.Get(i)
		}
	}

	return res
}

const B = 450

type Buckets struct {
	cnt []int
	sum []int
}

func NewBuckets(n int) *Buckets {
	return &Buckets{
		cnt: make([]int, n),
		sum: make([]int, (n+B-1)/B),
	}
}

func (b *Buckets) Add(p int, v int) {
	b.cnt[p] += v
	b.sum[p/B] += v
}

func (b *Buckets) Get(r int) int {
	for block := r / B; block >= 0; block-- {
		if b.sum[block] == 0 {
			continue
		}
		for i := min(r, (block+1)*B-1); i >= max(1, block*B); i-- {
			if b.cnt[i] > 0 {
				return i
			}
		}
	}
	return 0
}
