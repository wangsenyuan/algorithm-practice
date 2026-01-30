For an arbitrary array $a$ of length $m$, call an index $i$ ($1 \le i \le m - 2$) bad if $a_i$, $a_{i+1}$, and $a_{i+2}$ are all pairwise coprime. More formally, $i$ is a bad index if and only if:

$$\gcd(a_i, a_{i+1}) = \gcd(a_i, a_{i+2}) = \gcd(a_{i+1}, a_{i+2}) = 1$$

Furthermore, call $a$ good if it has at most $6$ bad indices.

You are given an integer $n$. Construct a good permutation† $p$ of length $n$. It can be shown that such a permutation always exists.

Note that you do not have to minimize the number of bad indices.

* $\gcd(x, y)$ denotes the greatest common divisor of $x$ and $y$.
* † A permutation of length $n$ is an array that contains every integer from $1$ to $n$ exactly once, in any order.

## Input

The first line contains a single integer $t$ ($1 \le t \le 10^4$) — the number of test cases.

The only line of each test case contains a single integer $n$ ($3 \le n \le 2 \cdot 10^5$).

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output on a single line $n$ integers $p_1, p_2, \ldots, p_n$, an example of a good permutation of length $n$. If there are multiple good permutations, you may output any of them.

## Example

### Input
```
4
3
6
8
9
```

### Output
```
2 1 3
4 1 6 3 5 2
4 1 6 3 5 2 8 7
5 4 8 1 9 3 6 2 7
```