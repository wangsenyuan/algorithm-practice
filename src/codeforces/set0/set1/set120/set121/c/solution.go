package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	fmt.Println(solve(n, k))
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, k int) int {
	var ans int

	F := make([]int, 15)
	F[0] = 1
	for i := 1; i < 15; i++ {
		F[i] = F[i-1] * i
	}

	if n < 15 && F[n] < k {
		return -1
	}

	var x int
	if n > 15 {
		// 只需要剩下30个数
		// 那么在 n - 30 中间，有多少个4, 7, 47, 这些位置的数都是good
		ans = get(n - 15)
		x = n - 15 + 1
	} else {
		x = 1
	}

	var flag int
	// x ... n
	for p := x; p <= n; p++ {
		var cnt int
		y := x
		for cnt < k {
			if flag&(1<<(y-x)) == 0 {
				if cnt+F[n-p] >= k {
					k -= cnt
					break
				}
				cnt += F[n-p]
			}
			y++
		}
		if checkLucky(p) && checkLucky(y) {
			ans++
		}
		flag |= 1 << (y - x)
	}

	return ans
}

func checkLucky(num int) bool {
	for num > 0 {
		r := num % 10
		if r != 4 && r != 7 {
			return false
		}
		num /= 10
	}
	return true
}

func get(num int) int {
	if num < 4 {
		return 0
	}
	if num < 7 {
		return 1
	}
	if num < 10 {
		return 2
	}
	// (4, 7)组成的数 <= num
	var h int
	x := num
	var arr []int
	for x >= 10 {
		h++
		arr = append(arr, x%10)
		x /= 10
	}
	// 4, 7, 47, 74, 447, ... 777 ...
	// 1 << 1 + 1 << 2 + 1 << 3
	res := (1 << (h + 1)) - 2

	arr = append(arr, x)

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < 4 {
			break
		}
		if arr[i] > 7 {
			res += 2 * (1 << i)
			break
		}
		if arr[i] == 7 {
			// 4***
			res += 1 << i
			continue
		}
		if arr[i] == 4 {
			continue
		}
		res += 1 << i
		break
	}

	return res
}

func bruteForce(n int, k int) int {
	// k == 1
	var res int
	for x := 1; x <= n; x++ {
		if checkLucky(x) {
			res++
		}
	}
	return res
}
