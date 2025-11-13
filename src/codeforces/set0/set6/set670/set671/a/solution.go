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
	fmt.Printf("%.12f\n", res)
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) float64 {
	alice := readNNums(reader, 2)
	bera := readNNums(reader, 2)
	bin := readNNums(reader, 2)
	n := readNNums(reader, 1)[0]
	bottles := make([][]int, n)
	for i := range n {
		bottles[i] = readNNums(reader, 2)
	}
	return solve(alice, bera, bin, bottles)
}

func solve(alice []int, bera []int, bin []int, bottles [][]int) float64 {

	calc := func(a []int, c []int) float64 {
		dx := c[0] - a[0]
		dy := c[1] - a[1]

		return math.Sqrt(float64(dx*dx + dy*dy))
	}
	// alice 和bera 处理过一次后， 就都到达了bin，那么后续的所有的操作的代价都一样
	var sum float64
	for _, cur := range bottles {
		sum += 2 * calc(bin, cur)
	}

	type data struct {
		id  int
		val float64
	}
	// 区别在于，可以选在其中的两个，分别分配给alice和bera，从而更新这个结果
	best_for_alice := []data{{-1, inf}, {-1, inf}}
	best_for_bera := []data{{-1, inf}, {-1, inf}}
	for i, cur := range bottles {
		tmp := calc(alice, cur) - calc(cur, bin)
		if tmp <= best_for_alice[0].val {
			best_for_alice[1] = best_for_alice[0]
			best_for_alice[0] = data{i, tmp}
		} else if tmp <= best_for_alice[1].val {
			best_for_alice[1] = data{i, tmp}
		}
		tmp = calc(bera, cur) - calc(cur, bin)
		if tmp <= best_for_bera[0].val {
			best_for_bera[1] = best_for_bera[0]
			best_for_bera[0] = data{i, tmp}
		} else if tmp <= best_for_bera[1].val {
			best_for_bera[1] = data{i, tmp}
		}
	}

	// 可以只有一个人的最佳
	res := min(best_for_alice[0].val, best_for_bera[0].val) + sum

	if best_for_alice[0].id != best_for_bera[0].id {
		res = min(res, sum+best_for_alice[0].val+best_for_bera[0].val)
	} else {
		res = min(res, sum+best_for_alice[0].val+best_for_bera[1].val)
		res = min(res, sum+best_for_alice[1].val+best_for_bera[0].val)
	}
	return res
}

const inf = 1 << 60
