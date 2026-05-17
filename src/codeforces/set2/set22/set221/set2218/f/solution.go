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
	for ; tc > 0; tc-- {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		ok, res := solve(x, y)
		if !ok {
			fmt.Fprintln(writer, "NO")
			continue
		}
		fmt.Fprintln(writer, "YES")
		for _, cur := range res {
			fmt.Fprintln(writer, cur[0], cur[1])
		}
	}
}

func solve(x int, y int) (bool, [][]int) {
	if y == 0 || y < x || x == 0 && y%2 == 0 {
		return false, nil
	}

	var res [][]int
	id := 2
	if x == 0 {
		for id <= y {
			res = append(res, []int{1, id})
			id++
		}
		return true, res
	}

	if (x+y)%2 == 0 {
		for i := 0; i < x-1; i++ {
			res = append(res, []int{1, id})
			res = append(res, []int{id, id + 1})
			id += 2
		}
		for id <= x+y {
			res = append(res, []int{1, id})
			id++
		}
		return true, res
	}

	for i := 0; i < x; i++ {
		res = append(res, []int{1, id})
		res = append(res, []int{id, id + 1})
		id += 2
	}
	for id <= x+y {
		res = append(res, []int{1, id})
		id++
	}

	return true, res
}
