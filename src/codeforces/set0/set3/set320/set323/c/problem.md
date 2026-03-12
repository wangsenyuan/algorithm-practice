# Problem

- **Given**: two permutations `p` and `q` of size `n`, and `m` adaptively-defined range queries.
- **Query**: count how many values `v` have their index in `p` within `[l1, r1]` and their index in `q` within `[l2, r2]`.
- **Twist**: each query’s ranges are obfuscated by `f(z)` using `x` that depends on the **previous query’s answer**, so queries must be processed online in order.

---

## Input

- The first line contains an integer `n` (`1 ≤ n ≤ 10^6`) — the number of elements in both permutations.
- The second line contains `n` integers `p1, p2, ..., pn` (`1 ≤ pi ≤ n`) — the first permutation.
- The third line contains `n` integers `q1, q2, ..., qn` — the second permutation, in the same format.
- The fourth line contains an integer `m` (`1 ≤ m ≤ 2·10^5`) — the number of queries.
- Each of the following `m` lines contains four integers `a, b, c, d` (`1 ≤ a, b, c, d ≤ n`) describing a query.

The actual query parameters \((l1, r1, l2, r2)\) are derived from `a, b, c, d` as follows:

1. Introduce a variable `x`:
   - For the **first** query, `x = 0`.
   - For each **subsequent** query, `x = (answer to previous query) + 1`.
2. Define a function:
   \[
   f(z) = ((z - 1 + x) \bmod n) + 1
   \]
3. Then compute:
   - `l1 = min(f(a), f(b))`, `r1 = max(f(a), f(b))`
   - `l2 = min(f(c), f(d))`, `r2 = max(f(c), f(d))`

---

## Output

- For each query, print the answer (the number of such integers `v`) on a separate line.

---

## Examples

### Example 1

Input:

```text
3
3 1 2
3 2 1
1
1 2 3 3
```

Output:

```text
1
```

### Example 2

Input:

```text
4
4 3 2 1
2 3 4 1
3
1 2 3 4
1 3 2 1
1 4 2 3
```

Output:

```text
1
1
2
```

---

## ideas

1. 对于给定的查询 [l1, r1, l2, r2], 询问有多少v满足，同时在区间p[l1...r1], q[l2...r2]
2. 每个v在p\q中只会出现一次
3. 假设有一棵树，记录了到r为止，v在q序列的位置
4. 那么result = get(r1, l2, r2) - get(l1-1, l2, r2) ?
5. 好像是的

---

## Key insights

1. Treat each value `v` as a point `(pos_p[v], pos_q[v])`.
2. Then each query asks for the number of points inside the rectangle `[l1, r1] x [l2, r2]`.
3. If we enumerate positions in permutation `p`, we can build an array `a` where `a[i] = pos_q[p[i]]`.
4. The query becomes: count how many `a[i]` lie in `[l2, r2]` for `i` in `[l1, r1]`.
5. So the problem reduces to a static range counting problem on one array.
6. A persistent segment tree can solve this, but with `n <= 10^6` it is very likely to exceed memory because it creates too many nodes.
7. A wavelet matrix (or other compact static range counting structure) is a better fit here because it answers rectangle/range-frequency queries in `O(log n)` while using much less memory.
8. The adaptive decoding with `x = previous_answer + 1` means queries must be processed online, but this does not change the underlying data structure.

---

## Wavelet matrix steps

1. Build the transformed array `a`, where `a[i] = pos_q[p[i]]`.
2. All values in `a` are in the range `[0, n - 1]`, so we can process them bit by bit from the highest bit to the lowest bit.
3. At each level, split the current sequence into two groups:
   - values whose current bit is `0`
   - values whose current bit is `1`
4. Record a bitvector for that level, where the `i`-th bit tells whether the `i`-th value went to the `1` group.
5. Also record `mid[level]`, the number of values that went to the `0` group. After stable partitioning, the next level's sequence is:
   - all `0`-bit values first
   - all `1`-bit values after them
6. Repeat this for all bits. This produces a compact layered structure instead of storing many segment tree versions.
7. To count how many values in subarray `a[l...r]` are `< upper`, walk down the levels:
   - if the current bit of `upper` is `0`, move only inside the `0` group
   - if the current bit of `upper` is `1`, add all values from the `0` group, then move into the `1` group
8. Use the bitvector rank operation to map the range `[l, r)` from one level to the next in `O(1)` per level.
9. Then `countLess(l, r, upper)` is computed in `O(log n)`.
10. Finally, the answer to a rectangle query is:
    - `countLess(l1 - 1, r1, r2) - countLess(l1 - 1, r1, l2 - 1)` if using inclusive value bounds, or equivalently
    - `countLess(L, R, upper) - countLess(L, R, lower)` with half-open conventions
11. So each query is answered in `O(log n)`, and the whole structure uses about `O(n log n)` bits plus prefix counts, which is much smaller than the pointer-heavy persistent segment tree.

---

## Implementation note

1. The persistent segment tree approach is still algorithmically correct here.
2. In C++, an array-backed implementation with fast input usually fits the time and memory limits comfortably.
3. In Go, the same idea is much more sensitive to constant factors, recursion overhead, allocation behavior, and package-level tooling constraints.
4. So for this problem:
   - C++ persistent segment tree is a practical AC solution.
   - Go wavelet matrix is often the safer choice when memory is tight.
