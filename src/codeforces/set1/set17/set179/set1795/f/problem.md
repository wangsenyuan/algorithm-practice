# F. Chip Game on a Tree

You are given a **tree** with \(n\) vertices. There are **\(k\)** chips on distinct vertices \(a_1, a_2, \ldots, a_k\). Those \(k\) vertices are **black** initially; all other vertices are **white**.

You play a game with **moves** (possibly zero). On move \(i\) (1-indexed), you move the **\(((i - 1) \bmod k) + 1\)-th chip** from its current vertex to an **adjacent white** vertex, then color that vertex **black**.

So for \(k = 3\), the order of chips moved is \(1, 2, 3, 1, 2, 3, \ldots\). If a chip has **no** adjacent white vertex on its turn, the **game ends** immediately.

What is the **maximum** number of moves you can make?

## Input

The first line contains a single integer \(t\) (\(1 \le t \le 10^4\)) — the number of test cases.

The first line of each test case contains a single integer \(n\) (\(1 \le n \le 2 \cdot 10^5\)) — the number of vertices.

Each of the next \(n - 1\) lines contains two integers \(v\) and \(u\) (\(1 \le v, u \le n\)) — an edge of the tree.

The next line contains a single integer \(k\) (\(1 \le k \le n\)) — the number of chips.

The next line contains \(k\) integers \(a_1, a_2, \ldots, a_k\) (\(1 \le a_i \le n\)) — the vertices that initially hold chips. All \(a_i\) are **distinct**.

It is guaranteed that the sum of \(n\) over all test cases does not exceed \(2 \cdot 10^5\).

## Output

For each test case, print a single integer — the maximum number of moves you can perform.

## Example

**Input**

```
5
5
1 2
2 3
3 4
4 5
1
3
5
1 2
2 3
3 4
4 5
2
1 2
5
1 2
2 3
3 4
4 5
2
2 1
6
1 2
1 3
2 4
2 5
3 6
3
1 4 6
1
1
1
```

**Output**

```
2
0
1
2
0
```


### ideas
1. 考虑最后的状态，假设在第i步时，走不了了。那么0...i%k移动的路径长度时 (i + k-1)/k
2. 后面移动的距离 = i / k; 而且他们所有的只能往一个方向（超叶子节点移动）
3. 那么这个时候，可以选择最外围的节点；如果能往外走，还好（但是可能存在它必须往里面走的情况）
4. 

## solution summary

1. Binary search the answer.
   - Suppose we want to know whether it is possible to make at least `m` moves.
   - Chip `i` will be moved either `floor(m / k)` or `ceil(m / k)` times, depending on whether its turn appears in the first `m` moves.
   - So for fixed `m`, every chip has a fixed required path length.

2. Root the tree at `a[0]`.
   - In any feasible schedule, the used edges form `k` vertex-disjoint paths starting from the initial chip positions.
   - Since chip `1` is the only chip whose path may go through the root side, every other chip must be completely contained in its own branch when viewed from `a[0]`.
   - Therefore a subtree DP is enough to test feasibility.

3. DFS state meaning.
   - For a node `u`, DFS returns one of three kinds of values for the subtree of `u`:
   - positive value: there is exactly one unfinished path in this subtree, and it still needs that many edges upward
   - `0`: the subtree can be completed internally
   - negative value: there is no unfinished path, and the subtree provides `-value` free white vertices that can still be used as room
   - a special impossible value means this subtree configuration cannot be realized.

4. Merge rule.
   - If two different children both return positive values, then two unfinished chip-paths would need to pass through `u`, which is impossible.
   - Also, if `u` itself already contains a chip and some child still wants to pass upward, that is impossible too.
   - So at most one positive child contribution may survive.

5. What does `room` mean?
   - A negative return value counts how many extra vertices are available in that direction.
   - If a chip still needs `need` more steps and the current subtree already has at least `need` free vertices below, then the path can be finished completely inside this subtree and the return becomes `0`.
   - Otherwise the unfinished requirement is propagated upward.

6. Transition at a chip node.
   - For fixed `m`, compute how many times this chip must move:
     `ceil(m / k)` for earlier chips in the cycle and `floor(m / k)` for the others.
   - If the subtree below this chip has enough room, it is finished locally.
   - Otherwise it must send an unfinished requirement upward.

7. Why binary search works.
   - If at least `m` moves are possible, then any smaller number of moves is also possible.
   - So feasibility is monotone, and we can binary search the maximum `m`.

8. Complexity.
   - Each feasibility check is one DFS, so `O(n)`.
   - Binary search adds a `log n` factor.
   - Total complexity is `O(n log n)` per test case.
