package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

type pair struct {
	first  int
	second int
}

func solve(s string) int {
	var arr []pair
	n := len(s)
	for i := 0; i < n; {
		j := i
		for i < n && s[i] == s[j] {
			i++
		}
		arr = append(arr, pair{first: i - j, second: int(s[j] - 'a')})
	}
	var res int
	for len(arr) > 1 {
		x := arr[0].first
		for i := 0; i < len(arr); i++ {
			if i == 0 || i == len(arr)-1 {
				x = min(x, arr[i].first)
			} else {
				x = min(x, (arr[i].first+1)/2)
			}
		}
		res += x
		// 经过x步骤后，至少有一个被删除了
		var j int
		for i := 0; i < len(arr); i++ {
			var y int
			if i == 0 || i == len(arr)-1 {
				y = x
			} else {
				y = 2 * x
			}
			if arr[i].first > y {
				z := arr[i].first - y
				if j > 0 && arr[j-1].second == arr[i].second {
					arr[j-1].first += z
				} else {
					arr[j] = pair{first: z, second: arr[i].second}
					j++
				}
			}
		}
		arr = arr[:j]
	}
	return res
}
