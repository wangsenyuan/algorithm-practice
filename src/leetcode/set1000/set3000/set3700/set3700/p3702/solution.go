package p3702

type pair struct {
	first  int
	second byte
}

func removeSubstring(s string, k int) string {
	// stack[l]...stack[top]是连续的
	n := len(s)

	var arr []pair

	for i := range n {
		m := len(arr)

		if m == 0 || arr[m-1].second != s[i] {
			arr = append(arr, pair{1, s[i]})
		} else {
			arr = append(arr, pair{arr[m-1].first + 1, s[i]})
		}

		m++

		if s[i] == ')' && m >= 2*k && arr[m-1].first == k && arr[m-k-1].first >= k {
			for range 2 * k {
				arr = arr[:m-1]
				m--
			}
		}
	}
	var buf []byte
	for i := range arr {
		buf = append(buf, arr[i].second)
	}

	return string(buf)
}
