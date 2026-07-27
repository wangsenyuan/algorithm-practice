package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

func solve(n int) string {
	if n == 1 {
		return "0/1"
	}

	next := func(k int) int { return (2 * k) % n }

	seen := make(map[int]int)
	path := make([]int, 0, 64)
	u := 1
	for u != 0 {
		if idx, ok := seen[u]; ok {
			cyc := path[idx:]
			pref := path[:idx]
			eStart := cycleExpect(cyc) // E[cyc[0]]

			// Back-substitute along the prefix into the cycle.
			E := eStart
			for i := len(pref) - 1; i >= 0; i-- {
				v := pref[i]
				r := next(v)
				// E[v] = 1 + (r/(2v)) * E[r]
				E = new(big.Rat).Add(
					big.NewRat(1, 1),
					new(big.Rat).Mul(big.NewRat(int64(r), int64(2*v)), E),
				)
			}
			return fmt.Sprintf("%s/%s", E.Num().String(), E.Denom().String())
		}
		seen[u] = len(path)
		path = append(path, u)
		u = next(u)
	}

	// Chain ends at 0.
	E := big.NewRat(1, 1) // E for a node with next 0; apply from the end
	for i := len(path) - 1; i >= 0; i-- {
		v := path[i]
		r := next(v)
		if r == 0 {
			E = big.NewRat(1, 1)
		} else {
			E = new(big.Rat).Add(
				big.NewRat(1, 1),
				new(big.Rat).Mul(big.NewRat(int64(r), int64(2*v)), E),
			)
		}
	}
	return fmt.Sprintf("%s/%s", E.Num().String(), E.Denom().String())
}

// cycleExpect returns E[c[0]] for a cycle c under k -> (2k) mod n.
// On such a cycle, ∏ coef = 1/2^L, so
//
//	E[c0] = (Σ_j c[j] * 2^{L-j}) / (c[0] * (2^L - 1)).
func cycleExpect(c []int) *big.Rat {
	L := len(c)
	pow2L := new(big.Int).Lsh(big.NewInt(1), uint(L))
	den := new(big.Int).Sub(pow2L, big.NewInt(1))
	den.Mul(den, big.NewInt(int64(c[0])))

	num := new(big.Int)
	for j, v := range c {
		term := new(big.Int).Lsh(big.NewInt(int64(v)), uint(L-j))
		num.Add(num, term)
	}
	return new(big.Rat).SetFrac(num, den)
}
