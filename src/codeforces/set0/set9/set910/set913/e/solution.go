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

	var n int
	fmt.Fscan(reader, &n)
	best := precompute()
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(reader, &s)
		mask := 0
		for j := 0; j < 8; j++ {
			if s[j] == '1' {
				mask |= 1 << j
			}
		}
		fmt.Fprintln(writer, best[mask])
	}
}

func precompute() [256]string {
	var bestE [256]string
	var bestT [256]string
	var bestF [256]string

	update := func(dst *[256]string, mask int, cand string) bool {
		cur := dst[mask]
		if cur == "" || len(cand) < len(cur) || len(cand) == len(cur) && cand < cur {
			dst[mask] = cand
			return true
		}
		return false
	}

	const (
		maskX = 240 // 00001111
		maskY = 204 // 00110011
		maskZ = 170 // 01010101
	)
	bestF[maskX] = "x"
	bestF[maskY] = "y"
	bestF[maskZ] = "z"

	changed := true
	for changed {
		changed = false

		for mask := 0; mask < 256; mask++ {
			if bestF[mask] != "" && update(&bestT, mask, bestF[mask]) {
				changed = true
			}
			if bestT[mask] != "" && update(&bestE, mask, bestT[mask]) {
				changed = true
			}
			if bestE[mask] != "" {
				if update(&bestF, mask, "("+bestE[mask]+")") {
					changed = true
				}
			}
			if bestF[mask] != "" {
				if update(&bestF, (^mask)&255, "!"+bestF[mask]) {
					changed = true
				}
			}
		}

		for a := 0; a < 256; a++ {
			if bestT[a] != "" {
				for b := 0; b < 256; b++ {
					if bestF[b] != "" {
						if update(&bestT, a&b, bestT[a]+"&"+bestF[b]) {
							changed = true
						}
					}
				}
			}
			if bestE[a] != "" {
				for b := 0; b < 256; b++ {
					if bestT[b] != "" {
						if update(&bestE, a|b, bestE[a]+"|"+bestT[b]) {
							changed = true
						}
					}
				}
			}
		}
	}

	return bestE
}
