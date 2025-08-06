package calculator

func eval(express string) int {
	express = normalize(express)
	n := len(express)

	stack := make([]int, n)
	op := make([]byte, n)
	var p, q int

	for i := 0; i < n; i++ {
		if isOp(express[i]) {
			if express[i] == '+' || express[i] == '-' {
				for p > 0 && compare(express[i], op[p-1]) <= 0 {
					a, b := stack[q-2], stack[q-1]
					c := op[p-1]
					x := calc(a, b, c)
					stack[q-2] = x
					p--
					q--
				}
				op[p] = express[i]
				p++
			} else {
				op[p] = express[i]
				p++
			}
		} else {
			var num int
			j := i
			for j < n && !isOp(express[j]) {
				num = num*10 + int(express[j]-'0')
				j++
			}
			i = j
			stack[q] = num
			q++
		}
	}

	for p > 0 {
		a, b := stack[q-2], stack[q-1]
		c := op[p-1]
		x := calc(a, b, c)
		stack[q-2] = x
		p--
		q--
	}
	return stack[0]
}

func calc(a, b int, c byte) int {
	if c == '+' {
		return a + b
	}
	if c == '-' {
		return a - b
	}
	if c == '*' {
		return a * b
	}
	if c == '/' {
		if b == 0 {
			panic("division by zero")
		}
		return a / b
	}
	panic("unknown operator")
}

func isOp(x byte) bool {
	return x == '+' || x == '-' || x == '*' || x == '/'
}

func normalize(s string) string {
	bs := []byte(s)

	var j int

	for i := 0; i < len(bs); i++ {
		if s[i] == ' ' {
			continue
		}
		bs[j] = bs[i]
		j++
	}

	return string(bs[:j])
}

func compare(a, b byte) int {
	if a == '+' || a == '-' {
		return -1
	}
	if b == '*' || b == '/' {
		return 0
	}
	return 1
}
