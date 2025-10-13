package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) (a []int, res []string) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(n, a)
	return
}

type pair struct {
	first  int
	second int
}

func solve(n int, a []int) []string {
	var res []string
	for i := 0; i < n; i++ {
		j := i
		best := make([]pair, 3)
		for i < n && a[i] != 0 {
			if a[i] >= best[0].first {
				best[2] = best[1]
				best[1] = best[0]
				best[0] = pair{a[i], i + 1}
			} else if a[i] >= best[1].first {
				best[2] = best[1]
				best[1] = pair{a[i], i + 1}
			} else if a[i] >= best[2].first {
				best[2] = pair{a[i], i + 1}
			}
			i++
		}
		if i == j {
			// 没有数据
			res = append(res, "0")
			continue
		}

		var flag int
		// 第一个放入stack， 第二个放入 queue, 第三个放入deck， 其他的都放入deck
		for j < i {
			if best[0].second == j+1 || best[1].second == j+1 || best[2].second == j+1 {
				if flag&1 == 0 {
					flag |= 1
					res = append(res, "pushStack")
				} else if flag&2 == 0 {
					flag |= 2
					res = append(res, "pushQueue")
				} else if flag&4 == 0 {
					flag |= 4
					res = append(res, "pushFront")
				}
			} else {
				res = append(res, "pushBack")
			}
			j++
		}

		if i == n {
			break
		}
		cnt := bits.OnesCount(uint(flag))

		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%d", cnt))

		if flag&1 == 1 {
			buf.WriteString(" popStack")
		}
		if flag&2 == 2 {
			buf.WriteString(" popQueue")
		}

		if flag&4 == 4 {
			buf.WriteString(" popFront")
		}
		res = append(res, buf.String())
	}

	return res
}
