package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	jiro := make([]string, n)
	for i := range n {
		jiro[i] = readString(reader)
	}
	ciel := make([]int, m)
	for i := range m {
		ciel[i] = readNum(reader)
	}
	return solve(jiro, ciel)
}

const inf = 1 << 60

func solve(jiro []string, ciel []int) int {
	sort.Ints(ciel)

	var attack []int
	var defense []int
	for _, card := range jiro {
		var score int
		readInt([]byte(card), 4, &score)
		if card[0] == 'A' {
			attack = append(attack, score)
		} else {
			defense = append(defense, score)
		}
	}
	sort.Ints(attack)
	sort.Ints(defense)

	// 先将def消耗掉
	marked := make([]bool, len(ciel))

	get := func(ok bool) int {
		// 那么就不考虑消耗完所有的attack了，应该是选择前j个attack，和后j个ciel
		var arr []int
		for i := 0; i < len(ciel); i++ {
			if !marked[i] {
				arr = append(arr, ciel[i])
			}
		}
		var ans int
		for j := 0; j <= len(attack) && j <= len(arr); j++ {
			yes := true
			var sum int
			for i := 0; i < j; i++ {
				if attack[i] > arr[len(arr)-j+i] {
					yes = false
					break
				}
				sum += arr[len(arr)-j+i] - attack[i]
			}
			if !yes {
				break
			}
			if ok && j == len(attack) {
				for i := 0; i < len(arr)-j; i++ {
					sum += arr[i]
				}
			}

			ans = max(ans, sum)
		}

		return ans
	}

	best := get(false)

	ok := true

	for i, j := 0, 0; i < len(defense); i++ {
		for j < len(ciel) && ciel[j] <= defense[i] {
			j++
		}
		if j == len(ciel) {
			ok = false
			break
		}
		marked[j] = true
		j++
	}

	best = max(best, get(ok))

	return best
}
