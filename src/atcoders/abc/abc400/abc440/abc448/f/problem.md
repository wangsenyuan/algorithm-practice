### Problem

There are `N` locations numbered from `1` to `N` on a 2D plane.  
Location `i` is at coordinates `(X_i, Y_i)`.

You will:

- start from location `1`,
- visit **every** location exactly once,
- then return to location `1`.

Moving from location `i` to location `j` takes:

\[
d(i, j) = |X_i - X_j| + |Y_i - Y_j|
\]

seconds (Manhattan distance).

You must output **one route** that visits all locations and returns to location `1` such that the **total travel time** is at most \(10^{10}\) seconds.  
It is guaranteed that under the given constraints, at least one such route exists.

This is an **approximate / constructive** problem: any route satisfying the conditions is accepted.

### Constraints

- `1 ≤ N ≤ 6 × 10^4`
- `0 ≤ X_i ≤ 2 × 10^7`
- `0 ≤ Y_i ≤ 2 × 10^7`
- All `(X_i, Y_i)` are distinct.
- All input values are integers.

### Input

Input is given from Standard Input in the following format:

```text
N
X_1 Y_1
X_2 Y_2
⋮
X_N Y_N
```

### Output

Let `p_i` be the `i`-th location visited (`1 ≤ i ≤ N`).  
Output your route in the following format:

```text
p_1 p_2 … p_N
```

Your answer is considered correct if **all** of the following hold:

1. `(p_1, p_2, …, p_N)` is a permutation of `(1, 2, …, N)`.
2. `p_1 = 1`.
3. Letting `d(i, j)` be the travel time from location `i` to location `j`, the total travel time

   \[
   \sum_{i=1}^{N} d\bigl(p_i,\ p_{(i \bmod N) + 1}\bigr) \le 10^{10}.
   \]

   Here, we interpret `p_{N+1} = p_1`, so the last term goes from the last visited location back to location `1`.

If there are multiple valid routes, you may print **any** of them.

### Sample Input 1

```text
3
0 6
3 5
2 4
```

### Sample Output 1

```text
1 3 2
```

Explanation:

\[
\sum_{i=1}^{3} d\bigl(p_i, p_{(i \bmod 3) + 1}\bigr)
  = d(1,3) + d(3,2) + d(2,1)
  = 4 + 2 + 4 = 10 \le 10^{10}.
\]

### Sample Input 2

```text
10
9706344 19786176
19341349 15565412
5711023 19068083
12521132 14054301
14767612 17088029
14961700 18526945
13801766 5740101
6581153 8643675
13176196 16586661
4086263 5172719
```

### Sample Output 2

```text
1 5 2 6 4 7 9 8 3 10
```

### ideas
1. 一个直观的想法是，先移动到最左下的点（然后，移动到距离不超过1000）的范围内的，最近的上方的点，然后往下，
2. 迭代，再从最后一个点返回1
3. 这样可以保证在x轴方向上，移动不会炒股过 1000 * n
4. 貌似可行
