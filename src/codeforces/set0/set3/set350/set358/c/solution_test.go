package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)
	score := make([]int, 2)

	st := make([]*Stack, 2)
	que := make([]*Queue, 2)
	d := make([]*Deck, 2)
	for i := range 2 {
		st[i] = &Stack{}
		que[i] = &Queue{}
		d[i] = &Deck{}
	}

	play := func(x int, op string, i int) {
		switch op {
		case "pushStack":
			st[i].Push(x)
		case "pushQueue":
			que[i].Push(x)
		case "pushFront":
			d[i].PushFront(x)
		case "pushBack":
			d[i].PushBack(x)
		}
	}

	kiss := func(op string, j int) {
		xs := strings.Split(op, " ")
		for i := 1; i < len(xs); i++ {
			switch xs[i] {
			case "popStack":
				score[j] += st[j].Pop()
			case "popQueue":
				score[j] += que[j].Pop()
			case "popFront":
				score[j] += d[j].PopFront()
			case "popBack":
				score[j] += d[j].PopBack()
			}
		}
	}

	for i := range len(a) {
		if a[i] > 0 {
			play(a[i], res[i], 0)
			play(a[i], expect[i], 1)
		} else {
			kiss(res[i], 0)
			kiss(expect[i], 1)

			if score[0] != score[1] {
				t.Fatalf("Sample expect %v, but got %v", expect, res)
			}
		}
	}
}

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	n := len(*s)
	if n == 0 {
		panic("stack is empty")
	}
	res := (*s)[n-1]
	*s = (*s)[:n-1]
	return res
}

type Queue []int

func (q *Queue) Push(x int) {
	*q = append(*q, x)
}

func (q *Queue) Pop() int {
	n := len(*q)
	if n == 0 {
		panic("queue is empty")
	}
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

type Deck []int

func (d *Deck) PushFront(x int) {
	*d = append([]int{x}, *d...)
}

func (d *Deck) PopFront() int {
	n := len(*d)
	if n == 0 {
		panic("deck is empty")
	}
	res := (*d)[0]
	*d = (*d)[1:]
	return res
}

func (d *Deck) PushBack(x int) {
	*d = append(*d, x)
}

func (d *Deck) PopBack() int {
	n := len(*d)
	if n == 0 {
		panic("deck is empty")
	}
	res := (*d)[n-1]
	*d = (*d)[:n-1]
	return res
}

func TestSample1(t *testing.T) {
	s := `10
0
1
0
1
2
0
1
2
3
0
`
	expect := []string{
		"0",
		"pushStack",
		"1 popStack",
		"pushStack",
		"pushQueue",
		"2 popStack popQueue",
		"pushStack",
		"pushQueue",
		"pushFront",
		"3 popStack popQueue popFront",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1
2
3
0
`
	expect := []string{
		"pushStack",
		"pushQueue",
		"pushFront",
		"3 popStack popQueue popFront",
	}
	runSample(t, s, expect)
}
