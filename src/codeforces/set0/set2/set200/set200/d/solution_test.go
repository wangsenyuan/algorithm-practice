package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	reader = bufio.NewReader(strings.NewReader(expect))
	for _, x := range res {
		y := readNum(reader)
		if x != y {
			t.Errorf("Sample expect %d, but got %d", y, x)
		}
	}
}
func TestSample1(t *testing.T) {
	s := `4
void f(int,T)
void  f(T, T)
 void foo123   ( int,  double,  string,string  ) 
  void  p(T,double)
3
int a
 string    s
double x123 
5
f(a,  a)
  f(s,a   )
foo   (a,s,s)
 f  (  s  ,x123)
proc(a)
`
	expect := `2
1
0
1
0
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
void f(string,double,int)
void f(int)
   void f  ( T  )
void procedure(int,double)
void f  (T, double,int)   
void f(string, T,T)
4
 int a
 int x
string  t
double  val  
5
f(t, a, a)
f(t,val,a)
f(val,a, val)
 solve300(val, val)
f  (x)
`
	expect := `1
3
0
0
2
`
	runSample(t, s, expect)
}
