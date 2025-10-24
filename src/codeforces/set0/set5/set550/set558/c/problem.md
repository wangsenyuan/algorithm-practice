## Problem Statement

Amr loves Chemistry, and specially doing experiments. He is preparing for a new interesting experiment.

Amr has $n$ different types of chemicals. Each chemical $i$ has an initial volume of $a_i$ liters. For this experiment, Amr has to mix all the chemicals together, but all the chemicals volumes must be equal first. So his task is to make all the chemicals volumes equal.

To do this, Amr can do two different kind of operations:

1. Choose some chemical $i$ and double its current volume so the new volume will be $2a_i$
2. Choose some chemical $i$ and divide its volume by two (integer division) so the new volume will be $\lfloor a_i/2 \rfloor$

Suppose that each chemical is contained in a vessel of infinite volume. Now Amr wonders what is the minimum number of operations required to make all the chemicals volumes equal?

## Input

The first line contains one number $n$ ($1 \leq n \leq 10^5$), the number of chemicals.

The second line contains $n$ space separated integers $a_i$ ($1 \leq a_i \leq 10^5$), representing the initial volume of the $i$-th chemical in liters.

## Output

Output one integer the minimum number of operations required to make all the chemicals volumes equal.

## Examples

### Example 1
**Input:**
```
3
4 8 2
```

**Output:**
```
2
```

### Example 2
**Input:**
```
3
3 5 6
```

**Output:**
```
5
```

## Note

In the first sample test, the optimal solution is to divide the second chemical volume by two, and multiply the third chemical volume by two to make all the volumes equal 4.

In the second sample test, the optimal solution is to divide the first chemical volume by two, and divide the second and the third chemical volumes by two twice to make all the volumes equal 1.


## ideas
1. 如果已经构造好了图，那么就是所有选中一个节点，其他所有节点离它的最小距离之和
2. 问题是先要构造图；
3. x和x/2之间有一条边，但是2-4, 2 - 5之间有一条长度为1的边
4. dp[x] = dp[x/2] + cnt[x/2] cnt[x/2]表示x/2的子树中的节点的数量
5. cnt[x] = freq[x] + cnt[x/2] (最多到 2 * mx - 1)
6. fp[x] = fp[2 * x] + suf[2 * x]
7. 但是不大对，比如4和5的距离，在上面体现不出来
8. x/2相当于右移，x * 2相当于左移
9. 也就是操作，指导大家的相同
10. 大概有点想法了，那么就固定两个位置，表示，要得到相同的数（然后看大家需要操作多少次，可以得到同样的数）
11. 