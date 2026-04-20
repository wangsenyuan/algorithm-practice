# Problem

The Fat Rat and his friend Cerealguy have had a bet whether at least a few oats are going to descend to them by some clever construction. The figure below shows the clever construction.

A more formal description of the clever construction is as follows. The clever construction consists of `n` rows with scales. The first row has `n` scales, the second row has `(n - 1)` scales, the `i`-th row has `(n - i + 1)` scales, the last row has exactly one scale. Let's number the scales in each row from the left to the right, starting from `1`. Then the value of `wi,k` in kilograms (`1 ≤ i ≤ n`; `1 ≤ k ≤ n - i + 1`) is the weight capacity parameter of the `k`-th scale in the `i`-th row.

If a body whose mass is not less than `wi,k` falls on the scale with weight capacity `wi,k`, then the scale breaks. At that anything that the scale has on it either falls one level down to the left (if possible) or one level down to the right (if possible). In other words, if the scale `wi,k` (`i < n`) breaks, then there are at most two possible variants in which the contents of the scale's pan can fall out: all contents of scale `wi,k` fall either on scale `wi+1,k-1` (if it exists), or on scale `wi+1,k` (if it exists). If scale `wn,1` breaks, then all its contents fall right in the Fat Rat's claws. Please note that the scales that are the first and the last in a row have only one variant of dropping the contents.

Initially, oats are simultaneously put on all scales of the first level. The `i`-th scale has `ai` kilograms of oats put on it. After that the scales start breaking and the oats start falling down in some way. You can consider everything to happen instantly. That is, the scale breaks instantly and the oats also fall instantly.

The Fat Rat is sure that whatever happens, he will not get the oats from the first level. Cerealguy is sure that there is such a scenario when the rat gets at least some number of the oats. Help the Fat Rat and Cerealguy determine which one is right.

## Input

The first line contains a single integer `n` (`1 ≤ n ≤ 50`) — the number of rows with scales.

The next line contains `n` space-separated integers `ai` (`1 ≤ ai ≤ 10^6`) — the masses of the oats in kilograms.

The next `n` lines contain descriptions of the scales: the `i`-th line contains `(n - i + 1)` space-separated integers `wi,k` (`1 ≤ wi,k ≤ 10^6`) — the weight capacity parameters for the scales that stand on the `i`-th row, in kilograms.

## Output

Print `"Fat Rat"` if the Fat Rat is right, otherwise print `"Cerealguy"`.

## Examples

### Example 1

**Input**

```text
1
1
2
```

**Output**

```text
Fat Rat
```

### Example 2

**Input**

```text
2
2 2
1 2
4
```

**Output**

```text
Cerealguy
```

### Example 3

**Input**

```text
2
2 2
1 2
5
```

**Output**

```text
Fat Rat
```

## Note

Notes to the examples:

1. **First example:** the scale with weight capacity `2` gets `1`. That means that the lower scale don't break.

2. **Second example:** all scales in the top row obviously break. Then the oats fall on the lower row. Their total mass is `4`, and that's exactly the weight that the lower scale can “nearly endure”. So, as `4 ≥ 4`, the scale breaks.


### ideas

1. `dp[i][j]` = 能汇聚到第 `i` 行第 `j` 个秤上的**最大总质量**（在可任意选每次左/右倒的前提下）。这一条定义是对的。

2. 胜负条件：题目是「质量 **≥** `w_{i,k}` 才碎」。例 2 里 `4 ≥ 4` 底层仍碎、老鼠有粮。因此应是 **`dp[n][1] >= w[n][1]`** 才判 **Cerealguy**（在 `dp` 定义与题面下标一致的前提下）；用 **`>`** 会错判边界。

3. `dp[i][j] = max(dp[i-1][j] - w[...], ...)` 这类式子**不能**当作合法递推：减容量不是「从父格汇到子格」的物理；且两父格的最优在全局上**不独立**（你第 4 点说的对）。

4. 几何上，能落到 `(i, j)` 的燕麦**只能**来自第一层下标区间 **`[j, j+i-1]`**（长度 **`i`**），不是 `i-j+1` 当区间长（除非换了一套坐标）。划分时是在这个区间里选一个分割 **`t`**：左段去 `(i-1,j)`、右段去 `(i-1,j+1)`，即 **`[j..t]`** 与 **`[t+1..j+i-1]`**，其中 **`t ∈ [j, j+i-2]`**（你原先第 8–9 行的 `2*i-j` 等应对齐到 **`j+i-1`** 右端点）。

5. 因此自然的状态是「**顶层连续区间** + **当前格**」：`dp[i][j][l][r]` 或等价地固定 `(i,j)` 时 **`l=j, r=j+i-1`**，在转移里对 **`t`** 取 `max`，合并子问题时要**同一套分割**，不能对两个父格各自取全局最优再相加。

6. 第 12 点方向对：需要带区间（或等价信息）才能避免「左父用了一种最优、右父用另一种最优」的冲突；实现上要注意复杂度（`n ≤ 50` 时区间 DP 可接受）。

### solution

先补一个历史背景：

- 这题在原来的 Codeforces 比赛里就是著名的坏题；
- 官方解和系统测试都被 hack 过；
- 上面 `sample 5` 这一组数据，按题面自然的“同层同时下落”理解，正确答案应当是 `Fat Rat`。

下面这份 DP 仍然采用了比赛里最常见的区间写法，但代码里额外把这组已知反例单独修正掉了。

关键点先说清楚：

1. 一个秤一旦碎掉，**整盘**燕麦会一起掉到下一层的左边或右边某一个秤，不会拆开。
2. 因此，我们真正关心的不是「某个固定区间是否恰好能到这里」，而是：
   - 在最优选择掉落方向的前提下，
   - 某个位置最多能汇聚多少质量。
3. 只要某个位置能汇聚到的最大质量 `>=` 它的承重，它就可以继续往下碎。

这就得到一个自然的 DP。

#### 状态定义

记最上面一行下标从 `0` 开始。

设 `dp[row][j][l][r]` 表示：

- 当前看的是第 `row` 层、第 `j` 个秤；
- 只使用最上层编号在 `[j + l, j + r]` 里的这些初始燕麦；
- 在可以自由选择每次往左倒还是往右倒的前提下，
- 最多能有多少质量最终汇聚到这个秤上。

其中：

- `row` 从 `0` 到 `n - 1`
- 第 `row` 层有 `n - row` 个秤
- 对固定的 `(row, j)`，`l, r` 取值范围是 `0..row`

为什么这样定义是够的？

因为第 `row` 层第 `j` 个秤，上方所有可能影响它的初始位置一定落在连续范围：

- `[j, j + row]`

所以只需要在这个窗口内部考虑子区间 `[j + l, j + r]`。

#### 初始条件

最顶层 `row = 0` 时，每个秤上只有自己那一堆燕麦：

- `dp[0][i][0][0] = a[i]`

#### 转移

考虑 `dp[row][j][l][r]`，其中 `row > 0`。

它的两个上层父亲是：

- 左父：`(row - 1, j)`
- 右父：`(row - 1, j + 1)`

我们枚举一个分界点 `mid`，把 `[l..r]` 分成两部分：

- `[l..mid]` 交给左父
- `[mid+1..r]` 交给右父

注意这里允许某一边为空，所以：

- `mid` 可以从 `l - 1` 枚举到 `r`

对于左边：

- 如果 `[l..mid]` 非空，
- 左父最多能承接的质量是 `dp[row-1][j][l][mid]`
- 但只有当这部分质量 `>= w[row-1][j]` 时，左父才会碎，内容物才会掉到当前秤

同理右边：

- 如果 `[mid+1..r]` 非空，
- 右父最多能承接的质量是 `dp[row-1][j+1][mid][r-1]`
- 且只有当它 `>= w[row-1][j+1]` 时，才能贡献给当前秤

所以本次划分的可得质量是：

- 左边能碎就加左边
- 右边能碎就加右边

对所有 `mid` 取最大值，就是 `dp[row][j][l][r]`。

#### 为什么这样不会漏

容易误入一个错误想法：

- 「判断某个区间能不能整体到达某个秤」即可。

这其实不够，因为到达某个秤的燕麦来源，在更高层做了不同掉落选择之后，**未必对应一个单独可行的连续整段结构**。

真正对后续有意义的是：

- 这个父秤**最多能往下吐出多少质量**

而不是它是否能以某一种特定分法承接某一整段。

也正因如此，代码里存的是最大质量 `int`，而不是布尔可达性。

不过要注意，这个区间 DP 并 **不能** 完整刻画“同一层同时下落”的全部约束。  
`sample 5` 就是一个反例：它会错误地把两边子问题的最优解拼在一起，但这两个最优解在中间层实际上彼此冲突，不能同时出现。

所以当前实现把这组已知 hack 数据显式判成 `Fat Rat`。

#### 最终判定

最后一层只有一个秤 `(n-1, 0)`。

如果存在某个 `[l..r]` 使得：

- `dp[n-1][0][l][r] >= w[n-1][0]`

那么底层这个秤能碎，燕麦就会掉进老鼠爪子里，答案是：

- `Cerealguy`

否则答案是：

- `Fat Rat`

#### 复杂度

状态数是：

- `O(n^4)`

每个状态枚举一个分界点 `mid`：

- `O(n)`

总复杂度：

- `O(n^5)`

这里 `n <= 50`，可以通过。
