package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	buf, _ := io.ReadAll(reader)
	s := string(buf)
	i := 0
	for i < len(s) && (s[i] < '0' || s[i] > '9') {
		i++
	}
	j := i
	for j < len(s) && s[j] >= '0' && s[j] <= '9' {
		j++
	}
	target, _ := strconv.Atoi(s[i:j])
	return solve(target, s[j:])
}

func solve(target int, code string) int {
	p := NewParser(code)
	ops := p.parse()
	ev := &Evaluator{ops: ops, cache: make([]int, mod)}
	for i := range ev.cache {
		ev.cache[i] = -1
	}
	ans := -1
	for n := 0; n < mod; n++ {
		if ev.call(n) == target {
			ans = n
		}
	}
	return ans
}

const mod = 32768

type Node struct {
	kind        byte
	value       int
	left, right *Node
}

type Cond struct {
	op          string
	left, right *Node
}

type Operator struct {
	cond *Cond
	expr *Node
}

type Parser struct {
	s   string
	pos int
}

func NewParser(code string) *Parser {
	var buf strings.Builder
	for _, ch := range code {
		if ch != ' ' && ch != '\n' && ch != '\r' && ch != '\t' {
			buf.WriteRune(ch)
		}
	}
	return &Parser{s: buf.String()}
}

func (p *Parser) parse() []Operator {
	for p.pos < len(p.s) && p.s[p.pos] != '{' {
		p.pos++
	}
	p.pos++
	var ops []Operator
	for p.pos < len(p.s) && p.s[p.pos] != '}' {
		if strings.HasPrefix(p.s[p.pos:], "if(") {
			p.pos += 3
			cond := p.parseCond()
			p.expect(")")
			p.expect("return")
			expr := p.parseSum()
			p.expect(";")
			ops = append(ops, Operator{cond: cond, expr: expr})
		} else {
			p.expect("return")
			expr := p.parseSum()
			p.expect(";")
			ops = append(ops, Operator{expr: expr})
		}
	}
	return ops
}

func (p *Parser) parseCond() *Cond {
	left := p.parseSum()
	if p.pos+1 < len(p.s) && p.s[p.pos:p.pos+2] == "==" {
		p.pos += 2
		return &Cond{op: "==", left: left, right: p.parseSum()}
	}
	op := p.s[p.pos : p.pos+1]
	p.pos++
	return &Cond{op: op, left: left, right: p.parseSum()}
}

func (p *Parser) parseSum() *Node {
	node := p.parseProduct()
	for p.pos < len(p.s) && (p.s[p.pos] == '+' || p.s[p.pos] == '-') {
		op := p.s[p.pos]
		p.pos++
		node = &Node{kind: op, left: node, right: p.parseProduct()}
	}
	return node
}

func (p *Parser) parseProduct() *Node {
	node := p.parseMultiplier()
	for p.pos < len(p.s) && (p.s[p.pos] == '*' || p.s[p.pos] == '/') {
		op := p.s[p.pos]
		p.pos++
		node = &Node{kind: op, left: node, right: p.parseMultiplier()}
	}
	return node
}

func (p *Parser) parseMultiplier() *Node {
	if p.s[p.pos] == 'n' {
		p.pos++
		return &Node{kind: 'n'}
	}
	if p.s[p.pos] == 'f' {
		p.pos += 2
		arg := p.parseSum()
		p.expect(")")
		return &Node{kind: 'f', left: arg}
	}
	start := p.pos
	for p.pos < len(p.s) && p.s[p.pos] >= '0' && p.s[p.pos] <= '9' {
		p.pos++
	}
	num, _ := strconv.Atoi(p.s[start:p.pos])
	return &Node{kind: 'c', value: num}
}

func (p *Parser) expect(token string) {
	p.pos += len(token)
}

type Evaluator struct {
	ops   []Operator
	cache []int
}

func (ev *Evaluator) call(n int) int {
	if ev.cache[n] >= 0 {
		return ev.cache[n]
	}
	for _, op := range ev.ops {
		if op.cond == nil || ev.evalCond(op.cond, n) {
			ev.cache[n] = ev.eval(op.expr, n)
			return ev.cache[n]
		}
	}
	return 0
}

func (ev *Evaluator) evalCond(cond *Cond, n int) bool {
	a := ev.eval(cond.left, n)
	b := ev.eval(cond.right, n)
	if cond.op == "==" {
		return a == b
	}
	if cond.op == "<" {
		return a < b
	}
	return a > b
}

func (ev *Evaluator) eval(node *Node, n int) int {
	switch node.kind {
	case 'c':
		return node.value
	case 'n':
		return n
	case 'f':
		return ev.call(ev.eval(node.left, n))
	case '+':
		return (ev.eval(node.left, n) + ev.eval(node.right, n)) % mod
	case '-':
		return (ev.eval(node.left, n) - ev.eval(node.right, n) + mod) % mod
	case '*':
		return ev.eval(node.left, n) * ev.eval(node.right, n) % mod
	default:
		return ev.eval(node.left, n) / ev.eval(node.right, n)
	}
}
