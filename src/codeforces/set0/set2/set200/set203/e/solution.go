package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res[0], res[1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	var d, s int
	fmt.Fscan(reader, &n, &d, &s)

	allDrive := make([]int, 0, n)
	goodCarry := make([]int, 0, n)
	simpleDrive := make([]int, 0, n)

	var carriers int
	var baseSlots int
	var others int

	for i := 0; i < n; i++ {
		var c, f, l int
		fmt.Fscan(reader, &c, &f, &l)

		if l >= d {
			allDrive = append(allDrive, f)
		}

		if c > 0 {
			carriers++
			baseSlots += c - 1
			if l >= d {
				goodCarry = append(goodCarry, f)
			}
		} else {
			others++
			if l >= d {
				simpleDrive = append(simpleDrive, f)
			}
		}
	}

	sort.Slice(allDrive, func(i, j int) bool {
		return allDrive[i] < allDrive[j]
	})
	sort.Slice(goodCarry, func(i, j int) bool {
		return goodCarry[i] < goodCarry[j]
	})
	sort.Slice(simpleDrive, func(i, j int) bool {
		return simpleDrive[i] < simpleDrive[j]
	})

	sumAll := prefixSums(allDrive)
	sumCarry := prefixSums(goodCarry)
	sumSimple := prefixSums(simpleDrive)

	bestCnt, bestFuel := maximizeWithBudget(sumAll, s)

	for t := 1; t <= len(goodCarry); t++ {
		carryFuel := sumCarry[t]
		if carryFuel > s {
			break
		}

		remain := s - carryFuel
		afford, _ := maximizeWithBudget(sumSimple, remain)
		freeSlots := baseSlots + t
		additional := min(others, freeSlots+afford)
		needDrive := max(0, additional-freeSlots)
		totalFuel := carryFuel + sumSimple[needDrive]
		totalCnt := carriers + additional

		if totalCnt > bestCnt || totalCnt == bestCnt && totalFuel < bestFuel {
			bestCnt = totalCnt
			bestFuel = totalFuel
		}
	}

	return []int{bestCnt, bestFuel}
}

func prefixSums(arr []int) []int {
	res := make([]int, len(arr)+1)
	for i, x := range arr {
		res[i+1] = res[i] + x
	}
	return res
}

func maximizeWithBudget(pref []int, budget int) (int, int) {
	k := sort.Search(len(pref), func(i int) bool {
		return pref[i] > budget
	}) - 1
	if k < 0 {
		return 0, 0
	}
	return k, pref[k]
}
