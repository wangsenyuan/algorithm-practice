package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func abs(num int) int {
	return max(num, -num)
}

func sign(num int) int32 {
	if num > 0 {
		return 1
	}
	if num < 0 {
		return -1
	}
	return 0
}

func checkSquare(num int) bool {
	if num < 0 {
		return false
	}
	tmp := int(math.Sqrt(float64(num)))
	return tmp*tmp == num
}

func solve(a []int) []int {

	var nums [][]int
	var zeros []int

	n := len(a)
	belong := make([]int, n)

	for i, v := range a {
		if v == 0 {
			belong[i] = -1
			zeros = append(zeros, i)
		} else {
			found := false
			for j, row := range nums {
				tmp := a[row[0]]
				if checkSquare(tmp * v) {
					belong[i] = j
					nums[j] = append(nums[j], i)
					found = true
					break
				}
			}
			if !found {
				belong[i] = len(nums)
				nums = append(nums, []int{i})
			}
		}
	}

	ans := make([]int, n)

	seen := make([]bool, len(nums))
	for i := range n {
		var k int
		for j := i; j < n; j++ {
			if belong[j] != -1 {
				if !seen[belong[j]] {
					k++
				}
				seen[belong[j]] = true
			}
			ans[max(1, k)-1]++
		}
		for j := i; j < n; j++ {
			if belong[j] != -1 {
				seen[belong[j]] = false
			}
		}
	}

	return ans
}
