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
		_, score, res := drive(reader)
		fmt.Fprintln(writer, score)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) (a [][]int, score int, res string) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a = make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	score, res = solve(a)
	return
}

func solve(a [][]int) (score int, res string) {
	n := len(a)
	m := len(a[0])
	var sum int
	for _, row := range a {
		for _, v := range row {
			sum += v
		}
	}
	// x + y = sum
	// n * m 还挺大的
	// x = sum / 2， 这个结果是不是一定可以得到？
	// 好像可以得到的。假设一开始通过某种分配，得到了x0 > x
	// 那么这个时候，可以从最底下开始调整，如果这一行右边有 x0 - x 个1，那么就一直往右边移动
	// 如果不够，就继续往上, 所以肯定可以
	x := sum / 2
	var buf []byte

	var cur int
	for i, row := range a {
		var tmp int
		for _, v := range row {
			tmp += v
		}
		if cur+tmp < x {
			cur += tmp
			buf = append(buf, 'D')
			continue
		}
		if cur+tmp == x {
			buf = append(buf, 'D')
			for range m {
				buf = append(buf, 'R')
			}
		} else {
			// cur + tmp > x
			down := false
			for _, v := range row {
				buf = append(buf, 'R')
				tmp -= v
				if cur+tmp == x && !down {
					down = true
					buf = append(buf, 'D')
				}
			}
		}

		for i1 := i + 1; i1 < n; i1++ {
			buf = append(buf, 'D')
		}

		break
	}

	return x * (sum - x), string(buf)
}
