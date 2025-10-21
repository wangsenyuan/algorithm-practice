package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var r, g, b int
	fmt.Fscan(reader, &r, &g, &b)
	res := solve(r, g, b)
	fmt.Println(res)
}

func solve(r int, g int, b int) int {
	arr := []int{r, g, b}
	slices.Sort(arr)
	x := max(0, 2*(arr[0]+arr[1])-arr[2])
	x /= 3
	res := x
	res += min((arr[2]-x)/2, arr[1]-x+arr[0]-x)

	return res
}
