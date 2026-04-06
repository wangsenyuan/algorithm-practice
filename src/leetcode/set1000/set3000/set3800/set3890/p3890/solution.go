package p3890

import "slices"

func findGoodIntegers(n int) []int {
	var cubes []int
	for x := 1; x*x <= n/x; x++ {
		cubes = append(cubes, x*x*x)
	}
	// len(x) <= 1000

	freq := make(map[int]int)

	var res []int
	for i, x := range cubes {
		for j := i; j < len(cubes) && x+cubes[j] <= n; j++ {
			y := cubes[j]
			freq[x+y]++
			if freq[x+y] == 2 {
				res = append(res, x+y)
			}
		}
	}
	slices.Sort(res)
	return res
}
