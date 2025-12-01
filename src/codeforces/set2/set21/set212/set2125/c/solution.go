package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var l, r int
	fmt.Fscan(reader, &l, &r)
	return solve(l, r)
}

func solve(l int, r int) int {
	calc := func(num int) int {
		arr := []int{2, 3, 5, 7}

		T := 1 << 4
		res := num
		for mask := 1; mask < T; mask++ {
			prod := 1
			for i := range 4 {
				if (mask>>i)&1 == 1 {
					prod *= arr[i]
				}
			}
			cnt := num / prod
			if bits.OnesCount(uint(mask))&1 == 1 {
				res -= cnt
			} else {
				res += cnt
			}
		}
		return res
	}

	return calc(r) - calc(l-1)
}
