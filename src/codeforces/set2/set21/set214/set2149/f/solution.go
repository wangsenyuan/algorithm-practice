package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	for range t {
		var hp, d int
		fmt.Fscan(reader, &hp, &d)
		res := solve(hp, d)
		fmt.Fprintln(writer, res)
	}
}

func solve(hp int, d int) int {
	check := func(w int) bool {
		s := w + 1
		a := d / s
		r := d % s
		sum := r*(a+1)*(a+2)/2 + (s-r)*a*(a+1)/2
		return sum <= hp+w-1
	}

	res := sort.Search(d, check)
	return res + d
}
