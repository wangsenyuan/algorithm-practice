package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var k int
	fmt.Fscan(reader, &k)
	res := solve(k)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, layer := range res {
		for _, row := range layer {
			fmt.Fprintln(writer, row)
		}

		fmt.Fprintln(writer)
	}
}

func solve(k int) [][]string {
	if k%2 == 1 {
		return nil
	}

	construct := func(x int) []string {
		res := make([]string, k)
		for i := 0; i < k; i++ {
			buf := make([]byte, k)
			x1 := x ^ (i & 1)
			for j := 0; j < k; j += 2 {
				if x1 == 0 {
					buf[j] = 'w'
					buf[j+1] = 'w'
				} else {
					buf[j] = 'b'
					buf[j+1] = 'b'
				}
				x1 ^= 1
			}
			res[i] = string(buf)
		}
		return res
	}
	res := make([][]string, k)
	var x int
	for i := 0; i < k; i += 2 {
		res[i] = construct(x)
		res[i+1] = construct(x)
		x ^= 1
	}

	return res
}
