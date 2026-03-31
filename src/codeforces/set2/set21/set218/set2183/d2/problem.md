You are given a rooted tree consisting of `n` vertices numbered from `1` to `n`, where the root has index `1`, and each vertex is initially white. Let `d_i` be the distance from the root to the `i`-th vertex. You can perform the following operations any number of times:

1. Select a subset `S` of white vertices such that no two distinct vertices in `S` are connected by an edge, and no two vertices in `S` have the same distance to vertex `1`. Formally, for all distinct `x, y ∈ S`, we have `d_x ≠ d_y`, and there is no edge between `x` and `y`.
2. Color every vertex in `S` black.

Your job is to find the minimum number of operations required to color the whole tree black, and output one valid sequence of operations.

\* A tree is a connected graph without cycles. A rooted tree is a tree where one vertex is special and called the root.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 ≤ t ≤ 10^4`). The description of the test cases follows.

The first line of each test case contains a single integer `n` (`2 ≤ n ≤ 2·10^5`) — the number of vertices in the tree.

The next `n − 1` lines describe the edges: the `i`-th of these lines contains two integers `u_i` and `v_i` (`1 ≤ u_i, v_i ≤ n`, `u_i ≠ v_i`) — the ends of the `i`-th edge.

It is guaranteed that the given edges form a tree.

It is guaranteed that the sum of `n` over all test cases does not exceed `2·10^5`.

## Output

For each test case:

Print `k` — your number of operations on the first line (`1 ≤ k ≤ n`).

Then print `k` lines. Each line begins with an integer `m` (`0 ≤ m ≤ n`) — the size of the set for this operation. After that, print `m` distinct integers `u_1, u_2, …, u_m` (`1 ≤ u_i ≤ n`) — the vertices you color black in this operation.

You must guarantee that:

- Each operation is valid.
- You never color the same vertex twice (within one operation or across operations).
- After all operations, every vertex is black.
- The number of operations equals the minimum possible over all valid solutions. If there are multiple optimal solutions, print any.

## Example

**Input**

```
10
5
3 1
1 2
5 1
4 1
5
3 2
2 4
2 5
1 2
5
3 4
4 1
5 1
1 2
5
2 5
3 1
2 1
3 4
5
1 3
1 5
4 3
2 4
13
2 1
3 2
4 2
5 4
6 3
7 1
8 5
9 6
10 4
11 7
12 8
13 10
10
5 7
8 1
1 10
2 8
8 4
9 4
6 1
5 3
7 8
10
7 6
3 7
6 9
7 1
9 8
5 1
3 10
9 2
1 4
10
10 6
2 8
4 10
7 5
1 2
7 10
10 9
9 1
7 3
10
6 8
9 7
4 10
5 9
4 2
3 8
6 5
1 5
1 10
```

**Output**

```
5
1 3
1 2
1 5
1 4
1 1
4
2 3 1
1 4
1 5
1 2
4
1 4
2 5 3
1 2
1 1
3
2 4 2
2 5 3
1 1
3
2 3 2
2 5 4
1 1
3
5 9 12 10 11 2
4 8 6 4 1
4 13 5 3 7
4
4 2 9 3 10
3 4 5 6
2 7 1
1 8
4
2 7 9
3 5 3 2
4 4 6 10 8
1 1
4
4 6 3 8 9
3 4 5 1
1 7
2 10 2
3
4 7 3 4 5
3 8 9 10
3 2 6 1
```

## Note

In the first test case, `d_1 = 1` and `d_2 = d_3 = d_4 = d_5 = 2`. We can show that at least `5` operations are necessary because no two vertices can be colored in the same operation.

In the second test case, the minimum number of operations needed to color the full tree is `4`. The sample output shows one way to achieve that.


## ideas
1. 同一层的，不能同时选两个以上（它们dist相同)
2. 不同层，但是parent-child, 也不能同时选
3. 如果只考虑parent-child，那么需要两次，一层选择偶数层，一次选择奇数层接口
4. 如果只考虑同一层，不能同时选择，那么就是最宽的那层的次数
5. 显然不可能小于 min(2, max-width-layer)
6. 然后考察每一层，假设它有x个，按照parent分成c1, c2, ... ck个
7. 那么对于前c1个，在依次选择它们的时候，不能选择它们的parent，但是可以选择第二个parent
8. 如果只有一个parent， 那么 c1 + 1, (因为parent不能和它们同时选)
9. 如果有两个parent，在选择第一个的时候，可以选择第二个parent，在选择第二组的时候，可以选择第一个（这样子就可以省下一些操作）
10. 如果这一层只有一个分组，（只有一个parent有子节点）， 那么这层的贡献 = sum + 1
11. 

## key insights

1. The problem is equivalent to assigning each vertex an operation id (color).
   - Vertices chosen in the same operation must have different depths.
   - Adjacent vertices cannot share the same operation id.
   - So for every depth, all vertices on that depth must get distinct ids, and every child must get an id different from its parent.

2. The minimum answer is
   - `k = max( max width of a depth, 1 + max number of children of one vertex )`.

3. Why is this a lower bound?
   - If one depth contains `w` vertices, they must all be colored in different operations, so `k >= w`.
   - If one vertex has `c` children, then that vertex and its `c` children need `c + 1` different operations, so `k >= c + 1`.

4. Why is this also sufficient?
   - Process the tree level by level.
   - Assume the previous level has already been assigned distinct ids.
   - For the current level, each vertex forbids exactly one id: its parent's id.
   - Therefore the current task becomes:
     assign distinct ids to all vertices of this level, where each vertex has only one forbidden id.

5. Group vertices of the current level by their forbidden id.
   - Since the previous level uses distinct ids, each forbidden id corresponds to exactly one parent.
   - The size of each group is the number of children of one parent, so every group size is at most `k - 1`.
   - The whole level size is at most `k`.

6. Constructing one level.
   - Let `m` be the size of this level.
   - Collect all forbidden ids that appear on this level.
   - If `m < k`, add one extra free id; otherwise use exactly the `m` ids that appear.
   - Now greedily assign ids:
     always take any still-unused id different from the current forbidden id.
   - The only bad case is when `m = k` and the last remaining id equals the current forbidden id.
   - In that case, swap with any previously assigned vertex coming from a different forbidden group.
   - Such a vertex always exists because the current forbidden group size is at most `k - 1`, so not all previous vertices belong to the same group.

7. This gives an `O(n)` construction after BFS.

## simpler bfs construction

Another clean way to view the construction is to process one BFS level at a time and start from the most natural choice:

1. Give a vertex the same color as its parent whenever possible.
   - While scanning one level, if the parent's color has not been used on this level yet, assign it directly.
   - Put such vertices into `bad`, because they currently conflict with their own parent.

2. All other vertices go into `rest`.
   - For them, greedily assign the smallest unused color on this level.
   - These vertices are already valid, because they do not use the parent's color.

3. Now only the vertices in `bad` need to be fixed.
   - If `|bad| >= 2`, all of them currently hold pairwise distinct colors, and each color comes from another parent on the same level.
   - So we can cyclically shift the colors inside `bad`.
   - After the shift, every vertex in `bad` still has a unique color on this level, but no one keeps its own parent's color anymore.

4. If `|bad| = 1`, call that vertex `u`.
   - Then every other color already used on this level is different from `color(parent(u))`.
   - We simply give `u` the smallest unused color different from its current one.
   - This is exactly the situation where we need one extra color for this level, matching the bound `children(parent(u)) + 1`.

5. Why is this optimal?
   - A level of width `w` needs at least `w` colors.
   - If some vertex has `c` children, then that parent together with its children needs at least `c + 1` colors.
   - The BFS construction above never uses more than `max(max level width, 1 + max children count)`, so it is optimal.

6. The implementation is very short:
   - BFS to get levels.
   - For each level, split nodes into `bad` and `rest`.
   - Fill `rest` greedily with unused colors.
   - Fix `bad` by one cyclic rotation, or by one new color if `|bad| = 1`.
