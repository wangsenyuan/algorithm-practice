package p3527

import "sort"

func findCommonResponse(responses [][]string) string {
	freq := make(map[string]int)

	for _, cur := range responses {
		sort.Strings(cur)
		for i := 0; i < len(cur); i++ {
			if i == 0 || cur[i] != cur[i-1] {
				freq[cur[i]]++
			}
		}
	}
	var ans string

	for k, v := range freq {
		if len(ans) == 0 || freq[ans] < v || freq[ans] == v && k < ans {
			ans = k
		}
	}

	return ans
}
