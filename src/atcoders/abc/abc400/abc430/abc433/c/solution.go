package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	res := solve(s)
	fmt.Println(res)
}

func solve(s string) int {
	var arr []int
	n := len(s)
	var res int
	for i := 0; i < n; {
		j := i
		for i < n && s[i] == s[j] {
			i++
		}
		arr = append(arr, i-j)
		if j > 0 && s[j-1]+1 == s[j] {
			res += min(i-j, arr[len(arr)-2])
		}
	}
	return res
}
