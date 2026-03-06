# Problem

Let $p$ and $q$ be non-empty strings called the **container** and the **key**, respectively. String $q$ consists only of characters `0` and `1`. Consider the following algorithm that extracts a message $s$ from the given container $p$.

## Algorithm

```text
i = 0;
j = 0;
s = <>;
while i is less than the length of the string p
{
    if q[j] == 1, then add to the right of string s character p[i];
    increase variables i, j by one;
    if the value of the variable j equals the length of the string q, then j = 0;
}
```

In the pseudocode above, $i$ and $j$ are integer variables, $s$ is a string, `=` is assignment, `==` is comparison, `[]` accesses a character by index, `<>` is the empty string. Characters are indexed from **0** in all strings.

Your task is to construct the **lexicographically minimum** key $q$ of length $k$ such that, when used in the algorithm above, it extracts exactly the given message $s$ from container $p$. If no such key exists, report that it is impossible.

## Input

The first two lines contain two non-empty strings $p$ and $s$ ($1 \\le |p| \\le 10^6$, $1 \\le |s| \\le 200$), describing the container and the message. The strings may contain any ASCII characters with codes from 32 to 126 inclusive.

The third line contains a single integer $k$ ($1 \\le k \\le 2000$) — the key length.

## Output

Print the required key (a string of length $k$ consisting only of characters `0` and `1`). If the key does not exist, print a single character `0`.

## Examples

### Example 1

**Input:**

```text
abacaba
aba
6
```

**Output:**

```text
100001
```

### Example 2

**Input:**

```text
abacaba
aba
3
```

**Output:**

```text
0
```

## Note

String $x = x_1 x_2 \\dots x_p$ is lexicographically smaller than string $y = y_1 y_2 \\dots y_q$ if either:

- $p < q$ and $x_1 = y_1, x_2 = y_2, \\dots, x_p = y_p$, or
- there exists an integer $r$ ($0 \\le r < \\min(p, q)$) such that $x_1 = y_1, \\dots, x_r = y_r$ and $x_{r+1} < y_{r+1}$.

Characters are compared by their ASCII codes.


### ideas
1. key[i] = 1, 那么第i个, 第i+k个，i+2*k, ... 个字符都被选中了
2. 但是没法知道在s中具体的位置（除了它前面的选择会影响到，后面的选择也会被影响到）
3. 假设key中有x个1, 那么一次迭代选中x个，一共迭代 m = len(p) / k
4. 那么 m * x <= len(s) => x <= len(s) / m and x >= len(s) % m, 且 m * (x + 1) > len(s)
5. 先假设 len(p) % k = 0, 那么 m = len(p) / k
6. 那么 x = len(s) / m, len(s) % m <= x
7. key[i] = 1, 且假设它之前有w个1，
8. 那么s[w] = p[i], s[w + x] = p[i+k], s[w+2*x] = p[i+2*k], ...
9. 那么就可以有状态 dp[i][w] = 1
10. fp[i][w] 表示到i为止，是否可以有w个1， fp[i][w] = fp[i-1][w-1] and dp[i][w] or fp[i-1][w]
11. 第一项表示，在i-1处，可以有w-1个1，且在i处可以有w个1
12. 第二项表示，在i-1处，就已经有w个1的情况
13. 这里的前提是x是确定的（我感觉x应该是比较少的几种情况）

### review

设：

- `m = len(p)`
- `n = len(s)`
- `c = m / k`
- `r = m % k`
- `x = key` 中 `1` 的总个数

那么每个完整轮次会贡献 `x` 个字符，一共会有 `c` 个完整轮次；最后还有一个长度为 `r` 的残缺轮次，只会再贡献前 `r` 个位置里为 `1` 的那些字符。

因此固定 `x` 以后，最终消息长度一定满足：

```text
c * x <= n <= c * x + min(r, x)
```

当前代码里枚举 `x` 用的是更宽松但安全的条件：

```text
c * x <= n <= (c + 1) * x
```

它不会漏掉答案，只是会多试一些不可能的 `x`。

#### 1. 时间复杂度

对固定的一个 `x`：

- `reset()` 要清空 `dp[k][n]` 和 `fp[k][n]`，复杂度是 `O(k * n)`
- 计算 `dp[i][w]` 时，虽然有一层关于 `w` 的循环，但对每个固定的 `i`，所有 `w` 的内层跳跃检查总和仍然是 `O(n)` 量级，因此整段是 `O(k * n)`
- 计算 `fp[i][w]` 是 `O(k * x)`，而 `x <= n`，所以也是 `O(k * n)`
- 回溯答案是 `O(k)`

所以一次 `play(x)` 的复杂度是：

```text
O(k * n)
```

外层会枚举若干个 `x`。最坏情况下这个数量是 `O(n)`（例如 `c = 1` 时，会从大约 `n` 枚举到 `ceil(n / 2)`）。

因此总时间复杂度是：

```text
O(k * n^2)
```

空间复杂度是：

```text
O(k * n)
```

在本题里 `n = len(s) <= 200`，所以这个复杂度是完全可接受的。

#### 2. 有没有更好的解

如果“不改主思路”，最直接的优化是把 `x` 的枚举范围收紧到真正可能的值：

```text
c * x <= n <= c * x + min(r, x)
```

这样能减少不少无效尝试，尤其是在 `m % k` 很小的时候。这个优化主要是减少常数，最坏复杂度仍然可能写成 `O(k * n^2)`，但实际运行会更干净。

如果问“是否存在明显更优、并且同样容易实现的解”，我认为没有特别值得替换当前方案的版本。因为本题真正小的维度是 `len(s) <= 200`，你的做法就是抓住这个维度做 DP，方向是对的，复杂度也已经压在可用范围内了。
