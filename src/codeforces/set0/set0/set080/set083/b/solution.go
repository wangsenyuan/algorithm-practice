package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	less, res := process(reader)
	if less {
		fmt.Println("-1")
	} else if len(res) == 0 {
		fmt.Println()
	} else {
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
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

func process(reader *bufio.Reader) (bool, []int) {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(a, k)
}

func solve(a []int, k int) (bool, []int) {
	var sum int
	for _, v := range a {
		sum += v
	}
	if sum <= k {
		return sum < k, nil
	}
	type pair struct {
		first  int
		second int
	}

	n := len(a)
	arr := make([]pair, n)
	for i, v := range a {
		arr[i] = pair{v, i}
	}

	slices.SortFunc(arr, func(x, y pair) int {
		return cmp.Or(x.first-y.first, x.second-y.second)
	})

	marked := make([]bool, n)

	var cnt int
	var prev int
	var w int

	for i := 0; i < n; i++ {
		w = arr[i].first
		if (w-prev)*(n-i)+cnt > k {
			// u * (n - i) + cnt < k
			u := (k - cnt) / (n - i)
			v := (k - cnt + n - i - 1) / (n - i)
			cnt += u * (n - i)
			k -= cnt
			w = v + prev
			break
		}
		marked[arr[i].second] = true
		cnt += (w - prev) * (n - i)
		prev = arr[i].first
	}

	var lastPos int
	for k > 0 {
		if !marked[lastPos] && a[lastPos] >= w {
			k--
			if a[lastPos] == w {
				marked[lastPos] = true
			}
		}
		lastPos++
	}

	var res []int

	for i := range n {
		j := (lastPos + i) % n
		if !marked[j] {
			res = append(res, j+1)
		}
	}

	return false, res
}
