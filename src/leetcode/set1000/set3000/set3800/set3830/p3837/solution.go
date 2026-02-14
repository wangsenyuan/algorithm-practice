package p3837

func mapWordWeights(words []string, weights []int) string {
	var buf []byte

	for _, word := range words {
		var sum int
		for i := range word {
			x := int(word[i] - 'a')
			sum += weights[x]
		}
		sum %= 26
		sum = 25 - sum
		buf = append(buf, byte(sum+'a'))
	}
	return string(buf)
}
