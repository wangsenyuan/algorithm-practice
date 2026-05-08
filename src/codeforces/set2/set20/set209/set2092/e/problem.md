# E. She knows...

Source: <https://codeforces.com/problemset/problem/2092/E>

D. Pippy wants to repaint an `n x m` board for a black-and-white party.

Initially, every cell is green except for `k` given cells. Each given cell is already painted either white or black.

All remaining green cells must be repainted either white or black.

After repainting, consider all unordered pairs of edge-adjacent cells with different colors. The requirement is that the number of such pairs must be even.

Count how many ways there are to repaint all green cells so that this parity condition holds.

Output the answer modulo:

```text
1_000_000_007
```

## Input

The first line contains an integer `t`:

```text
1 <= t <= 10^4
```

Each test case starts with three integers:

```text
n m k
```

where:

```text
3 <= n, m <= 10^9
1 <= k <= 2 * 10^5
```

Then follow `k` lines. Each line contains:

```text
x_i y_i c_i
```

where:

```text
1 <= x_i <= n
1 <= y_i <= m
c_i in {0, 1}
```

`c_i = 0` means the cell is white, and `c_i = 1` means the cell is black.

All specified cells are distinct.

It is guaranteed that the sum of `k` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output one integer:

```text
the number of valid repaintings modulo 1_000_000_007
```

## Example

### Input

```text
2
3 3 6
1 1 0
1 2 1
1 3 0
3 1 1
3 2 0
3 3 1
3 4 12
1 1 0
1 2 1
1 3 0
1 4 1
2 1 1
2 2 0
2 3 1
2 4 0
3 1 0
3 2 1
3 3 0
3 4 1
```

### Output

```text
4
0
```

## Notes

In the first sample, there are exactly four valid ways to repaint the three green cells.

In the second sample, all cells are already painted, and the number of adjacent pairs with different colors is odd, so the answer is `0`.

### ideas
1. 把剩余 n * m - k 个格子进行涂色（黑色或者白色), 保证在黑白格子中间的边的数量是偶数
2. 完全想象不出来～
3. 假设目前已经涂完色了，现在要将一个黑色变成白色。分情况考虑；
4. 如果这个格子周围都是黑色格子，那么改变它的颜色， 贡献的边是偶数
5. 如果这个格子周围只有一个白色格子, 贡献的边还是偶数（减少了一条边，但是增加了3条边）
6. 如果这个格子周围有两个白色格子，贡献是0，减少了2条边，但是增加了2条边
7. 如果这个格子周围有3个白色格子，贡献也是偶数，减少了3条边，但是增加了2条边（所以，在内部的格子的颜色，其实是没有影响的）pow(2, (n - 2) * (m - 2) - k1) 
8. 然后考虑在边上的格子（不在corner中的）
9. 如果它的周围都是黑色格子，那么翻转它，边的奇偶性变化了（+3条边）
10. 如果它的周围有两个黑色格子（都在左右）翻转后，变化是+1（+2了两条边，-1了一条边）
11. 如果它的周围是一个黑色格子，翻转后，变化是-1， 少了两条边，+1了一条边
12. 如果它的周围都是白色格子，翻转后，变化是-3
13. 现在考虑corner的情况，如果翻转它，没有变化（不管它靠近的是黑色还是白色，不影响结构）
14. pow(2, (n - 2) * (m - 2) + 4) （不影响结果的部分）
15. 怎么得到偶数的边呢？且全部绿色都变成了黑色的情况下，边是偶数，那么必须翻转偶数个点（剩余部分的偶数个）
16. pow(2, m - 1)不管怎么如何都是这样子的（留下最后一个点，调整就可以）
17. 和现有的颜色没有关系

## Solution explanation

Use `0/1` to represent the two colors.

For an edge `(u, v)`, this edge contributes `1` to the count of different-color adjacent pairs exactly when:

```text
color[u] xor color[v] = 1
```

We only care about whether the total count is even or odd, so work modulo `2`:

```text
sum over edges (color[u] xor color[v])
```

For bits, `xor` is the same as addition modulo `2`:

```text
color[u] xor color[v] = color[u] + color[v] (mod 2)
```

So the parity of all bad edges is:

```text
sum over edges (color[u] + color[v])
```

Every cell color appears once for each incident edge. Therefore:

```text
bad_edge_parity = sum over cells color[cell] * degree[cell] (mod 2)
```

Only cells with odd degree matter.

## Which cells matter

For a normal `n x m` rectangle:

- inner cells have degree `4`, so they do not affect parity;
- corner cells have degree `2`, so they do not affect parity;
- boundary non-corner cells have degree `3`, so they do affect parity.

The code counts these groups as:

```go
inner = (n - 2) * (m - 2)
corner = 4
other = 2*(n-2) + 2*(m-2)
```

Then it subtracts the already fixed cells from the corresponding group.

Here `other` means unpainted boundary non-corner cells, i.e. unpainted odd-degree cells.

## Free cells

Unpainted even-degree cells never affect the parity condition. They can be colored arbitrarily.

So they contribute:

```text
2^(inner + corner)
```

This is:

```go
res := pow(2, inner+corner)
```

## If there is at least one unpainted odd-degree cell

The odd-degree unpainted cells are the only remaining variables in the parity equation.

If there are `other > 0` such variables, exactly half of their assignments satisfy the required parity. We can choose all but one arbitrarily, and the last one is forced.

So their contribution is:

```text
2^(other - 1)
```

The answer is:

```go
res * pow(2, other-1)
```

This case does not need to inspect the fixed colors, because the final odd-degree unpainted cell can always adjust the parity.

## If all odd-degree cells are fixed

When `other == 0`, no variable can adjust the parity. We must directly check whether the currently forced coloring has even bad-edge parity.

The code treats every still-unpainted cell as black. This is safe in this branch because all remaining unpainted cells have even degree, so changing any of them would not affect the parity.

Then it counts bad edges involving fixed cells:

- `cnt` counts edges between a fixed cell and an unpainted black neighbor that have different colors;
- `cnt2` counts edges between two fixed cells of different colors.

The fixed-fixed edges are seen from both endpoints, so:

```go
cnt2 /= 2
```

Finally:

```go
cnt += cnt2
```

If `cnt` is odd, no assignment works:

```go
return 0
```

Otherwise all even-degree unpainted cells are free, so the answer is still:

```go
2^(inner + corner)
```

## Complexity

For each test case, the algorithm only processes the `k` fixed cells and checks their four neighbors.

Time complexity:

```text
O(k log k)
```

because the implementation stores fixed cells in a map.

Memory complexity:

```text
O(k)
```
