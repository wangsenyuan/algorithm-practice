package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var r, g int
	fmt.Fscanf(reader, "%d %d", &r, &g)
	res := solve(r, g)
	fmt.Println(res)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

type pair struct {
	first  int
	second int
}

func solve(red int, green int) int {
	if red > green {
		red, green = green, red
	}
	if green == 0 {
		return 0
	}
	if red == 0 {
		// 全部都是green
		return 1
	}

	H := sort.Search(green+1, func(h int) bool {
		return h*(h+1)/2 > (red + green)
	})

	H--

	dp := make([]int, green+1)
	dp[0] = 1
	for i := 1; i <= H; i++ {
		for j := green; j >= i; j-- {
			dp[j] = add(dp[j], dp[j-i])
		}
	}

	var res int
	for i := range red + 1 {
		if green+i >= H*(H+1)/2 {
			res = add(res, dp[i])
		}
	}
	return res
}
