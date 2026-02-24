# Problem C

You are given an array $a$ of $n$ integers.

In one operation, you perform the following two steps:

1. Choose an integer $x$ ($0 \le x \le 10^9$).
2. Replace each $a_i$ with $|a_i - x|$, where $|v|$ denotes the absolute value of $v$.

For example, choosing $x = 8$, the array $[5, 7, 10]$ becomes $[|5-8|, |7-8|, |10-8|] = [3, 1, 2]$.

Construct a sequence of operations to make all elements of $a$ equal to $0$ in at most $40$ operations, or determine that it is impossible. You do not need to minimize the number of operations.

---

## Input

- First line: a single integer $t$ ($1 \le t \le 10^4$) — the number of test cases.
- For each test case:
  - First line: a single integer $n$ ($1 \le n \le 2 \cdot 10^5$) — the length of the array $a$.
  - Second line: $n$ integers $a_1, a_2, \ldots, a_n$ ($0 \le a_i \le 10^9$) — the elements of $a$.

The sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case:

- If it is impossible to make all elements $0$ in at most $40$ operations, output a single integer $-1$.
- Otherwise, output two lines:
  - First line: a single integer $k$ ($0 \le k \le 40$) — the number of operations.
  - Second line: $k$ integers $x_1, x_2, \ldots, x_k$ ($0 \le x_i \le 10^9$) — the chosen $x$ for each operation in order.

If there are multiple solutions, output any of them. You do not need to minimize the number of operations.

---

## Example

**Input:**

```
5
1
5
2
0 0
3
4 6 8
4
80 40 20 10
5
1 2 3 4 5
```

**Output:**

```
1
5
0

3
6 1 1
7
60 40 20 10 30 25 5
-1
```

**Note:**

- **Test 1:** One operation with $x = 5$ changes $[5]$ to $[0]$.
- **Test 2:** All elements are already $0$; no operations needed.
- **Test 3:** $x = 6$ gives $[4,6,8] \to [2,0,2]$, then $x = 1$ gives $[1,1,1]$, then $x = 1$ again gives $[0,0,0]$.
- **Test 4:** The sequence $(60, 40, 20, 10, 30, 25, 5)$ makes all elements $0$.
- **Test 5:** It is impossible within $40$ operations; output $-1$.
