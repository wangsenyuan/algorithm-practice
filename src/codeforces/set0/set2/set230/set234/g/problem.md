# Problem

There are `n` players, numbered `1` to `n` (all distinct). Each **practice** is one game: split **all** `n` players into two **non-empty** teams (sizes may differ).

Schedule as few practices as possible so that **every pair of distinct players** appears in **at least one** practice on **different teams**.

Print a schedule achieving that minimum `m`: first `m`, then `m` lines. On line `i` print `f_i` — the size of the first team — then `f_i` distinct player numbers in the first team (in any order; the remaining `n - f_i` players form the second team).

## Input

A single integer `n` (`2 <= n <= 1000`).

## Output

- Line 1: `m` — minimum number of practices.
- Next `m` lines: line `i` contains `f_i` (`1 <= f_i < n`) followed by `f_i` integers — the players in the first team for that practice.

If several optimal schedules exist, print any.

## Examples

### Example 1

Input

```text
2
```

Output

```text
1
1 1
```

### Example 2

Input

```text
3
```

Output

```text
2
2 1 2
1 1
```

## Key Insights

这题的核心其实只有两个结论：

1. 至少需要 `ceil(log2 n)` 场训练。
2. 用二进制分组，恰好可以用 `ceil(log2 n)` 场做到。

一旦这两点都想明白，答案就非常自然了。

### 1. 为什么至少要 `ceil(log2 n)` 场

假设一共安排了 `m` 场训练。

对每个选手，我们记录一个长度为 `m` 的 0/1 串：

- 第 `i` 位是 `1`，表示他在第 `i` 场被分到第一队
- 第 `i` 位是 `0`，表示他在第 `i` 场被分到第二队

这个 0/1 串可以看成这个选手在整个训练计划中的“签名”。

题目的要求是：

- 任意两个不同选手，至少在某一场训练中分到不同的队

翻译成签名语言就是：

- 任意两个不同选手的 0/1 串必须不同

因为如果两个人的签名完全一样，那么每一场他们都在同一队，从来没有被分开过，不满足要求。

现在问题就变成：

- 用长度为 `m` 的 0/1 串，最多能区分多少个人？

答案显然是：

- 最多 `2^m` 个人

所以必须有：

- `2^m >= n`

也就是：

- `m >= ceil(log2 n)`

这就是下界。  
它不是拍脑袋出来的，而是一个非常标准的“编码”视角：

- 每场训练提供 1 bit 信息
- 要区分 `n` 个不同对象，至少要 `ceil(log2 n)` bit

### 2. 为什么二进制分组正好可以做到

既然下界来自“给每个人一个不同的二进制编码”，那最自然的构造就是：

- 直接把每个选手编号写成二进制

比如选手编号从 `1` 到 `n`。

对于第 `i` 场训练：

- 把二进制第 `i` 位为 `1` 的选手放进第一队
- 其余选手放进第二队

这样一共只需要处理：

- `i = 0, 1, ..., ceil(log2 n)-1`

也就是 `ceil(log2 n)` 场。

为什么它一定正确？

因为任意两个不同的整数，它们的二进制表示一定至少有一位不同。  
设选手 `u` 和 `v` 的二进制在第 `i` 位不同，那么在第 `i` 场训练里：

- 一个会被分到第一队
- 另一个会被分到第二队

于是这对选手就被成功分开了。

所以：

- 任意一对选手都能在某一场训练中分处两队

这正好满足题意。

### 3. 为什么不会出现空队

题目要求每场都要把所有人分成两个非空队。

在这个构造里，只要我们只枚举真正需要的那些二进制位，就不会出问题。

具体来说，我们只做满足下面条件的位：

- `2^i < n`

这是因为：

- 如果 `2^i >= n`，那么 `1..n` 里根本没有数的第 `i` 位是 `1`
- 这一场第一队就会是空的，显然无效

所以代码里循环应当是：

```go
for i := 0; 1<<i < n; i++ {
    ...
}
```

而不是 `<= n`。

这也是之前那个 bug 的根源：

- 当 `n` 恰好是 2 的幂时，`1<<i == n` 这一位会多算一场
- 那一场其实没有必要，而且会导致答案不是最优

### 4. 一个小例子

例如 `n = 5`，编号写成三位二进制：

- `1 = 001`
- `2 = 010`
- `3 = 011`
- `4 = 100`
- `5 = 101`

因为：

- `ceil(log2 5) = 3`

所以我们做 3 场训练：

第 0 位：

- 第一队：`1, 3, 5`
- 第二队：`2, 4`

第 1 位：

- 第一队：`2, 3`
- 第二队：`1, 4, 5`

第 2 位：

- 第一队：`4, 5`
- 第二队：`1, 2, 3`

任意两个人的二进制表示总有一位不同，因此总能在某一场训练中分开。

### 5. 总结成一句话

这题本质上不是“分组构造”，而是“给每个选手分配一个唯一的二进制签名”：

- 下界：`m` 场训练最多只能给出 `2^m` 种签名，所以 `m >= ceil(log2 n)`
- 上界：直接用编号的二进制位做分组，就能用 `ceil(log2 n)` 场实现

上下界一合，答案就唯一确定了：

- 最少训练场数就是 `ceil(log2 n)`

