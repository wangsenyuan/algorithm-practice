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
	var yang [3]int
	fmt.Fscan(reader, &yang[0], &yang[1], &yang[2])
	var monster [3]int
	fmt.Fscan(reader, &monster[0], &monster[1], &monster[2])
	var h, a, d int
	fmt.Fscan(reader, &h, &a, &d)
	return solve(yang[:], monster[:], h, a, d)
}

func solve(yang []int, monster []int, h int, a int, d int) int {
	if yang[2] >= monster[1] {
		// monster永远杀不死Yang
		if yang[1] > monster[2] {
			// 只要有足够长的时间，就可以杀死monster
			return 0
		}
		return (monster[2] + 1 - yang[1]) * a
	}

	var best = 1 << 60
	for x := max(monster[2]+1, yang[1]); ; x++ {
		// 当Yang的攻击时x的时候，需要n的时间, monster才能被杀死
		n := (monster[0] + x - monster[2] - 1) / (x - monster[2])

		for y := yang[2]; y <= monster[1]; y++ {
			// h + z > n * (monster[1] - y)
			if y == monster[1] {
				// 不需要增加hp
				best = min(best, (x-yang[1])*a+(y-yang[2])*d)
			} else {
				z := max(0, n*(monster[1]-y)-yang[0]+1)
				best = min(best, (x-yang[1])*a+(y-yang[2])*d+z*h)
			}
		}

		if n == 1 {
			// 不能更快了
			break
		}
	}

	return best
}
