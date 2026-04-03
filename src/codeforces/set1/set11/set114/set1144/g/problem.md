# G. Split into Increasing and Decreasing

Two integer sequences existed initially: one was **strictly increasing**, the other **strictly decreasing**.

- A **strictly increasing** sequence is \([x_1 < x_2 < \cdots < x_k]\).
- A **strictly decreasing** sequence is \([y_1 > y_2 > \cdots > y_l]\).

The **empty** sequence and a sequence of **one** element count as both increasing and decreasing.

Elements of the increasing sequence were **inserted** into the decreasing one (possibly before the first element and/or after the last) **without changing the relative order** within each original sequence.

**Example:** \([1, 3, 4]\) (increasing) and \([10, 4, 2]\) (decreasing) can yield results such as \([10, 1, 3, 4, 2, 4]\) or \([1, 3, 4, 10, 4, 2]\). The sequence \([1, 10, 4, 4, 3, 2]\) is **not** possible, because the order inside the increasing part was broken.

Let the merged sequence be \(a\). You are given \(a\). Find **any** pair of suitable original sequences (one strictly increasing, one strictly decreasing) that could produce \(a\) by this process, or decide it is **impossible**.

## Input

The first line contains one integer \(n\) (\(1 \le n \le 2 \cdot 10^5\)) — the length of \(a\).

The second line contains \(n\) integers \(a_1, a_2, \ldots, a_n\) (\(0 \le a_i \le 2 \cdot 10^5\)).

## Output

If no valid split exists, print `NO` on the first line.

Otherwise print `YES` on the first line. On the second line print \(n\) integers `res_1, res_2, …, res_n`, each `res_i` being **0** or **1**:

- `res_i = 0` — \(a_i\) belongs to the **increasing** subsequence;
- `res_i = 1` — \(a_i\) belongs to the **decreasing** subsequence.

(Again, empty or singleton subsequences are allowed as increasing/decreasing.)

## Examples

### Example 1

**Input**

```
9
5 1 3 6 8 2 9 0 10
```

**Output**

```
YES
1 0 0 0 0 1 0 1 0
```

### Example 2

**Input**

```
5
1 2 4 0 2
```

**Output**

```
NO
```


### ideas
1. 假设最小值是x，如果只有一个x，那么x要么是a的起点，要么是b的终点
2. 如果有两个x，那么其中一个是a的起点，另外一个是b的终点
3. 如果有3个x， => no answer
4. 如果有两个x, 如果两个中间存在数，那么第一个肯定是a的起点，第二个肯定是b的终点
5. 如果中间没有数字，那哪个做起点，哪个做终点没有关系。这时候，而且这时候左边是b，右边是a
6. 假设是2个x，且中间存在数字的情况，把x删掉（剩余的进行处理），它们就是一个子问题
7. 

## solution

Let color `0` mean the element goes to the increasing subsequence, and color `1` mean it goes to the decreasing subsequence.

We scan the array from left to right and do DP on the color of the current element.

### DP state

For each position `i`:

- `dp[i][0]`: `a[i]` is placed into the increasing subsequence, and this value stores the maximum possible last value of the decreasing subsequence after processing `a[0..i]`
- `dp[i][1]`: `a[i]` is placed into the decreasing subsequence, and this value stores the minimum possible last value of the increasing subsequence after processing `a[0..i]`

Why keep these extremal values?

- If `a[i]` is in the increasing subsequence, then the increasing side is fixed at `a[i]`, and for the decreasing side we want the largest possible last value, because that leaves the most freedom for future elements to still go down.
- Symmetrically, if `a[i]` is in the decreasing subsequence, then the decreasing side is fixed at `a[i]`, and for the increasing side we want the smallest possible last value, because that leaves the most freedom for future elements to still go up.

### Transition

Suppose we know a valid state at position `i` and want to place `a[i+1]`.

1. If `a[i]` is in the increasing subsequence:
   - the current last increasing value is `a[i]`
   - the current last decreasing value is `dp[i][0]`

   Then:
   - `a[i+1]` can also go to increasing iff `a[i+1] > a[i]`
   - `a[i+1]` can go to decreasing iff `a[i+1] < dp[i][0]`

2. If `a[i]` is in the decreasing subsequence:
   - the current last decreasing value is `a[i]`
   - the current last increasing value is `dp[i][1]`

   Then:
   - `a[i+1]` can also go to decreasing iff `a[i+1] < a[i]`
   - `a[i+1]` can go to increasing iff `a[i+1] > dp[i][1]`

When several transitions lead to the same state, keep only the best extremal value:

- maximize `dp[*][0]`
- minimize `dp[*][1]`

### Initialization

At position `0`, either choice is possible:

- put `a[0]` into increasing, then the decreasing subsequence is still empty
- put `a[0]` into decreasing, then the increasing subsequence is still empty

### Reconstruction

Store the predecessor color for every reachable state.
At the end:

- if neither `dp[n-1][0]` nor `dp[n-1][1]` is reachable, answer `NO`
- otherwise backtrack the predecessor array to recover one valid `0/1` assignment

### Why this DP is correct

For every prefix and for each choice of the current element's color, the only information that matters for future transitions is:

- the last value used by the increasing subsequence
- the last value used by the decreasing subsequence

One of them is exactly `a[i]`, because we know the color of `a[i]`.
The other one only matters through the most permissive boundary:

- as large as possible for the decreasing side
- as small as possible for the increasing side

So the DP keeps precisely the optimal summary of all valid ways to split each prefix.

### Complexity

- `O(n)` states
- `O(1)` transitions per state
- total `O(n)` time and `O(n)` memory
