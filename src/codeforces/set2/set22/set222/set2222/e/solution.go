package main

import (
	"bufio"
	"fmt"
	"os"
)

type Interactor struct {
	Insert func(x int) int
	Query  func(y int) int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	readInt := func() int {
		var v int
		if _, err := fmt.Fscan(reader, &v); err != nil || v == -1 {
			os.Exit(0)
		}
		return v
	}

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n int
		fmt.Fscan(reader, &n)

		chooseInitial := func(a int) {
			fmt.Fprintln(writer, a)
			writer.Flush()
		}

		it := Interactor{
			Insert: func(x int) int {
				fmt.Fprintln(writer, "I", x)
				writer.Flush()
				return readInt()
			},
			Query: func(y int) int {
				fmt.Fprintln(writer, "Q", y)
				writer.Flush()
				return readInt()
			},
		}

		k, c := solve(n, chooseInitial, it)
		fmt.Fprintln(writer, "A", k, c)
		writer.Flush()
	}
}

func solve(n int, chooseInitial func(a int), it Interactor) (k int, c int) {

	set := make(map[int]bool)

	query := func(x int) int {
		res := it.Query(x)
		for w := range set {
			if w >= x {
				res--
			}
		}
		return res
	}

	guess := func() int {
		var pre int
		for i := n - 1; i >= 0; i-- {
			if query(pre|(1<<i)) > 0 {
				pre |= 1 << i
			}
		}
		return pre
	}

	chooseInitial(0)
	// S = {0}
	res1 := it.Insert(0)
	set[0] = true

	if res1 == 1 {
		// k = 1,
		it.Insert(1<<n - 1)
		// S = {0, c}
		c = guess()
		return 1, c
	}
	c = guess()
	set[c] = true

	if c == (1<<n - 1) {
		res2 := it.Insert(1)
		if res2 != res1 {
			k = 3
		} else {
			k = 2
		}
		return
	}
	it.Insert(1<<n - 1)
	if query(1<<n-1) > 0 {
		k = 2
	} else {
		k = 3
	}
	return
}

func applyF(k int, x int, c int) int {
	switch k {
	case 1:
		return x & c
	case 2:
		return x | c
	default:
		return x ^ c
	}
}
