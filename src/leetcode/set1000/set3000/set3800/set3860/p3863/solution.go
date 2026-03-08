package p3863

import "slices"

func minOperations(s string) int {
	n := len(s)
	convert := func() []int {
		res := make([]int, n)
		for i := range n {
			res[i] = int(s[i] - 'a')
		}
		return res
	}

	arr := convert()
	if slices.IsSorted(arr) {
		return 0
	}
	if n == 2 {
		return -1
	}
	x := slices.Min(arr)
	y := slices.Max(arr)
	if int(s[0]-'a') == x || int(s[n-1]-'a') == y {
		return 1
	}
	slices.Sort(arr[:n-1])
	slices.Sort(arr[1:])
	if slices.IsSorted(arr) {
		return 2
	}
	arr = convert()
	slices.Sort(arr[1:])
	slices.Sort(arr[:n-1])
	if slices.IsSorted(arr) {
		return 2
	}
	return 3
}
