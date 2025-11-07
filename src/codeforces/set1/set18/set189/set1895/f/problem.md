Let's call an array $a$ of $n$ non-negative integers fancy if the following conditions hold:

- at least one from the numbers $x$, $x+1$, ..., $x+k-1$ appears in the array;
- consecutive elements of the array differ by at most $k$ (i.e. $|a_i - a_{i-1}| \leq k$ for each $i \in [2,n]$).

You are given $n$, $x$ and $k$. Your task is to calculate the number of fancy arrays of length $n$. Since the answer can be large, print it modulo $10^9+7$.

### Input

The first line contains a single integer $t$ ($1 \leq t \leq 50$) — the number of test cases.

The only line of each test case contains three integers $n$, $x$ and $k$ ($1 \leq n,k \leq 10^9$; $0 \leq x \leq 40$).

### Output

For each test case, print a single integer — the number of fancy arrays of length $n$, taken modulo $10^9+7$.

### Example

**Input**
```
4
3 0 1
1 4 25
4 7 2
1000000000 40 1000000000
```

**Output**
```
9
25
582
514035484
```

### Note

In the first test case of the example, the following arrays are fancy:

$[0,0,0]$; $[0,0,1]$; $[0,1,0]$; $[0,1,1]$; $[0,1,2]$; $[1,0,0]$; $[1,0,1]$; $[1,1,0]$; $[2,1,0]$.