package p3597

func partitionString(s string) []string {
	var res []string
	vis := make(map[string]bool)

	for i := 0; i < len(s); {
		j := i
		i++
		for i < len(s) && vis[s[j:i]] {
			i++
		}
		if !vis[s[j:i]] {
			res = append(res, s[j:i])
			vis[s[j:i]] = true
		}
	}
	return res
}
