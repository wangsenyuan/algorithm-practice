package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var l, r int
	fmt.Fscanf(reader, "%d %d", &l, &r)
	res := solve(l, r)
	fmt.Println(res)
}

func solve(l int, r int) int {
	if l == r {
		return 0
	}
	h := bits.Len(uint(r))

	for h > 0 && (l>>h)&1 == (r>>h)&1 {
		h--
	}

	return 1<<(h+1) - 1
}
