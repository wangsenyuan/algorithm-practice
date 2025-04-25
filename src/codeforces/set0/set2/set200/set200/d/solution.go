package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

type Signature struct {
	args      int
	args_size int
}

type Procedure struct {
	name string
	Signature
}

var types map[string]int

func init() {
	types = map[string]int{
		"T":      1,
		"int":    2,
		"string": 3,
		"double": 4,
	}
}

// T, int, string, double
func parseProcedure(s string) Procedure {
	s = strings.Replace(s, "void", "", 1)
	s = strings.ReplaceAll(s, " ", "")
	i := strings.Index(s, "(")
	name := s[:i]
	j := strings.Index(s, ")")
	s = s[i+1 : j]
	ss := strings.Split(s, ",")
	var args int
	for _, x := range ss {
		args = args*10 + types[x]
	}
	return Procedure{name, Signature{args, len(ss)}}
}

func convertToSignature(arr []int) Signature {
	var res int
	for _, x := range arr {
		res = res*10 + x
	}
	return Signature{res, len(arr)}
}

func convertToArgs(s Signature) []int {
	res := make([]int, s.args_size)
	v := s.args
	for i := 0; i < s.args_size; i++ {
		res[i] = v % 10
		v /= 10
	}
	return res
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	procedures := make([]string, n)
	for i := 0; i < n; i++ {
		procedures[i] = readString(reader)
	}
	m := readNum(reader)
	variables := make([]string, m)
	for i := 0; i < m; i++ {
		variables[i] = readString(reader)
	}
	q := readNum(reader)
	calls := make([]string, q)
	for i := 0; i < q; i++ {
		calls[i] = readString(reader)
	}
	return solve(procedures, variables, calls)
}

func solve(procedures []string, variables []string, calls []string) []int {
	signatures := make(map[string]map[Signature]int)
	for _, p := range procedures {
		cur := parseProcedure(p)
		if _, ok := signatures[cur.name]; !ok {
			signatures[cur.name] = make(map[Signature]int)
		}
		signatures[cur.name][cur.Signature]++
	}

	variables_map := make(map[string]int)
	for _, v := range variables {
		v = strings.TrimPrefix(v, " ")
		v = strings.TrimPrefix(v, " ")
		i := strings.Index(v, " ")
		var_type := v[:i]
		v = strings.TrimSpace(v[i+1:])
		variables_map[v] = types[var_type]
	}

	parseCall := func(s string) Procedure {
		s = strings.ReplaceAll(s, " ", "")
		i := strings.Index(s, "(")
		f_name := s[:i]
		j := strings.Index(s, ")")
		s = s[i+1 : j]
		ss := strings.Split(s, ",")
		var args int
		for _, x := range ss {
			y := variables_map[x]
			args = args*10 + y
		}
		return Procedure{f_name, Signature{args, len(ss)}}
	}

	check := func(s Signature, arr []int) bool {
		if s.args_size != len(arr) {
			return false
		}
		v := s.args
		for i := 0; i < s.args_size; i++ {
			r := v % 10
			if r != 1 && r != arr[i] {
				return false
			}
			v /= 10
		}
		return true
	}

	find := func(vs map[Signature]int, arr []int) int {
		var cnt int
		for s, v := range vs {
			if check(s, arr) {
				cnt += v
			}
		}
		return cnt
	}

	ans := make([]int, len(calls))

	for i, c := range calls {
		cur := parseCall(c)
		if vs, ok := signatures[cur.name]; ok {
			ans[i] = find(vs, convertToArgs(cur.Signature))
		}
	}

	return ans
}
