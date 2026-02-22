package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, drive(reader))
}

func drive(reader *bufio.Reader) int {
	var hl, hr, n int
	fmt.Fscan(reader, &hl, &hr, &n)
	mirrors := make([][]int, n)
	for i := range n {
		mirrors[i] = make([]int, 4)
		var pos string
		fmt.Fscan(reader, &mirrors[i][0], &pos, &mirrors[i][2], &mirrors[i][3])
		if pos == "T" {
			mirrors[i][1] = 0
		} else {
			mirrors[i][1] = 1
		}
	}
	return solve(hl, hr, mirrors)
}

// solve finds the maximum laser score using the "unfolded box" model.
// Reflections are replaced by a straight line from (0, hl) to (100000, Y).
// Floor mirrors appear at y = 100j (j even), ceiling at y = 100j (j odd).
// Valid exits: Y = 200k ± hr. Each level crossing must hit a distinct mirror.
func solve(hl, hr int, mirrors [][]int) int {
	n := len(mirrors)
	bound := (n + 1) * 100
	best := 0

	findMirror := func(j, den int, hit []bool) (int, bool) {
		num := int64(100*j-hl) * 100000
		wantFloor := j&1 == 0
		for i, m := range mirrors {
			if (m[1] == 1) != wantFloor {
				continue
			}
			lo, hi := int64(m[2])*int64(den), int64(m[3])*int64(den)
			if den < 0 {
				lo, hi = hi, lo
			}
			if lo <= num && num <= hi {
				if hit[i] {
					return 0, false
				}
				hit[i] = true
				return m[0], true
			}
		}
		return 0, false
	}

	evaluate := func(Y int) {
		if Y == hl {
			return
		}
		den := Y - hl
		score := 0
		hit := make([]bool, n)

		if den > 0 {
			for j := 1; 100*j < Y; j++ {
				if v, ok := findMirror(j, den, hit); ok {
					score += v
				} else {
					return
				}
			}
		} else {
			for j := 0; 100*j > Y; j-- {
				if v, ok := findMirror(j, den, hit); ok {
					score += v
				} else {
					return
				}
			}
		}

		if score > best {
			best = score
		}
	}

	for _, base := range []int{hr, -hr} {
		kLo := (hl-bound-base)/200 - 1
		kHi := (hl+bound-base)/200 + 1
		for k := kLo; k <= kHi; k++ {
			evaluate(200*k + base)
		}
	}

	return best
}
