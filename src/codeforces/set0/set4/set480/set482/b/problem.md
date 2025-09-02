# Problem B: Interesting Array

## Problem Description

We'll call an array of $n$ non-negative integers $a[1], a[2], \ldots, a[n]$ **interesting**, if it meets $m$
constraints. The $i$-th of the $m$ constraints consists of three
integers $l_i$, $r_i$, $q_i$ ($1 \leq l_i \leq r_i \leq n$) meaning that the
value $a[l_i] \& a[l_i+1] \& \ldots \& a[r_i]$ should be equal to $q_i$.

Your task is to find any interesting array of $n$ elements or state that such array doesn't exist.

**Note:** Expression $x \& y$ means the bitwise AND of numbers $x$ and $y$. In programming languages C++, Java and
Python this operation is represented as "&", in Pascal — as "and".

## Input

- The first line contains two integers $n$, $m$ ($1 \leq n \leq 10^5$, $1 \leq m \leq 10^5$) — the number of elements in
  the array and the number of limits.

- Each of the next $m$ lines contains three
  integers $l_i$, $r_i$, $q_i$ ($1 \leq l_i \leq r_i \leq n$, $0 \leq q_i < 2^{30}$) describing the $i$-th limit.

## Output

If the interesting array exists:

- In the first line print "YES" (without the quotes)
- In the second line print $n$ integers $a[1], a[2], \ldots, a[n]$ ($0 \leq a[i] < 2^{30}$) describing the interesting
  array. If there are multiple answers, print any of them.

If the interesting array doesn't exist:

- Print "NO" (without the quotes) in a single line.

## Examples

### Example 1

**Input:**

```
3 1
1 3 3
```

**Output:**

```
YES
3 3 3
```

**Explanation:** The array $[3, 3, 3]$ satisfies the constraint because $3 \& 3 \& 3 = 3$.

### Example 2

**Input:**

```
3 2
1 3 3
1 3 2
```

**Output:**

```
NO
```

**Explanation:** No array can satisfy both constraints simultaneously. The first constraint
requires $a[1] \& a[2] \& a[3] = 3$, while the second requires $a[1] \& a[2] \& a[3] = 2$. These are contradictory.

## Notes

- The bitwise AND operation $\&$ is applied to all elements in the range $[l_i, r_i]$
- Each constraint specifies that the bitwise AND of elements in a range must equal a specific value
- If multiple constraints are contradictory, the answer is "NO"
- The array elements must be non-negative integers less than $2^{30}$
- If a valid array exists, any valid array is acceptable as output

### ideas

1. 按位考虑，对于任何位i，默认全部是1，只有当它必须为0时，考虑把它设置成0
2. 考虑按照左端点升序。
3. 对于每一段，如果q[d]= 1， 那么这一段必须都是1
4. 如果q[d] = 0, 那么就使用尽量靠前段的为0