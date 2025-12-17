package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	fmt.Println(len(res))
	w := fmt.Sprintf("%v", res)
	fmt.Println(w[1 : len(w)-1])
}

func drive(reader *bufio.Reader) (s int, k int, res []int) {
	fmt.Fscan(reader, &s, &k)
	res = solve(s, k)
	return
}

func solve(s int, k int) []int {
	var seq []int
	seq = append(seq, 1)
	sum := 1
	cnt := 1
	var head int
	for sum <= s {
		seq = append(seq, sum)
		sum += sum
		cnt++
		if cnt > k {
			sum -= seq[head]
			head++
			cnt--
		}
	}
	// len(seq) < 30
	var res []int

	for i := len(seq) - 1; i >= 0 && s > 0; i-- {
		if s >= seq[i] {
			res = append(res, seq[i])
			s -= seq[i]
		}
	}

	if len(res) == 1 {
		res = append(res, 0)
	}

	return res
}
