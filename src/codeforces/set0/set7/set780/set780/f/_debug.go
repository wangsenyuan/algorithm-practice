package main

import "fmt"

func main() {
	n, m := 3, 6
	edges := [][]int{{3,1,1},{1,2,1},{2,1,1},{1,3,0},{3,2,1},{2,2,0}}
	fmt.Println(solve(n, edges))
}
