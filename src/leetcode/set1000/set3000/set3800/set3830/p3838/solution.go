package p3838

import "slices"

func prefixConnected(words []string, k int) int {
	var arr []string
	for _, word := range words {
		if len(word) >= k {
			arr = append(arr, word[:k])
		}
	}

	slices.Sort(arr)

	var res int

	for i := 0; i < len(arr); {
		j := i
		for i < len(arr) && arr[i] == arr[j] {
			i++
		}
		if i-j >= 2 {
			res++
		}
	}
	return res
}
