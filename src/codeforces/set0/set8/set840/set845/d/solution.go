package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	events := make([][]int, n)
	for i := range n {
		var tp int
		fmt.Fscan(reader, &tp)
		if tp == 1 || tp == 3 {
			var v int
			fmt.Fscan(reader, &v)
			events[i] = []int{tp, v}
		} else {
			events[i] = []int{tp}
		}
	}
	return solve(events)
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func solve(events [][]int) int {
	n := len(events)
	var res int

	stack := make([]pair, n+1)
	var top int
	stack[top] = pair{inf, 1}
	top++
	var speed int
	canOverTake := true
	var notOverTakeCnt int

	workSpeed := func() {
		for top > 0 && speed > stack[top-1].first {
			res += stack[top-1].second
			top--
		}
	}

	for i := range n {
		switch events[i][0] {
		case 1:
			// change speed
			speed = events[i][1]
			workSpeed()
			// speed <= stack[top-1]
		case 2:
			// overtake
			if !canOverTake {
				res += notOverTakeCnt
			}
			canOverTake = true
			notOverTakeCnt = 0
		case 3:
			v := events[i][1]
			cnt := 1
			for v >= stack[top-1].first {
				cnt += stack[top-1].second
				top--
			}
			stack[top] = pair{v, cnt}
			top++
			workSpeed()
		case 4:
			canOverTake = true
			notOverTakeCnt = 0
		case 5:
			top = 0
			stack[top] = pair{inf, 1}
			top++
		default:
			canOverTake = false
			notOverTakeCnt++
		}
	}

	return res
}
