package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, best, res := drive(reader)
	fmt.Println(best)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) (m int, a []int, best int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &m)
	a = readNNums(reader, n)
	best, res = solve(m, a)
	return
}

func solve(m int, a []int) (best int, res []int) {
	n := len(a)
	avg := n / m
	arr := make([][]int, m)
	cnt := make([]int, m)
	for i, v := range a {
		cnt[v%m]++
		arr[v%m] = append(arr[v%m], i)
	}

	// 从那个位置开始，一直可以增加呢？
	// 假设这个位置是i, 那么必须满足条件 pref[j] - pref[i-1] >= avg * (j - i + 1)
	// pref[j] - avg * j >= pref[i-1] - avg * (i-1)
	// 也就是说要找到这样的i，它后面所有的j都比它要大
	var que []pair
	var pref int
	for i := range m {
		pref += cnt[i]
		for len(que) > 0 && que[len(que)-1].first >= pref-avg*(i+1) {
			que = que[:len(que)-1]
		}
		que = append(que, pair{pref - avg*(i+1), i})
	}
	// first 肯定可以找到
	first := -1
	var sum int
	for i := range m {
		// 如果起点在i
		if sum-avg*i <= que[0].first {
			first = i
			break
		}
		if len(que) > 0 && que[0].second == i {
			que = que[1:]
		}
		pref += cnt[i]
		for len(que) > 0 && que[len(que)-1].first >= pref-avg*(i+m) {
			que = que[:len(que)-1]
		}
		que = append(que, pair{pref - avg*(i+m), i + m})
		sum += cnt[i]
	}

	res = slices.Clone(a)

	for l, r := first, first; l < first+m; l++ {
		for len(arr[l%m]) > avg {
			// 只需要往前加就可以了
			for r < first+2*m && cnt[r%m] >= avg {
				r++
			}
			best += r - l
			i := arr[l%m][0]
			arr[l%m] = arr[l%m][1:]
			res[i] += r - l
			cnt[r%m]++
		}
	}

	return
}

type pair struct {
	first  int
	second int
}
