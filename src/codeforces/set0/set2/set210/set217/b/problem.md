Fibonacci numbers are defined as the sequence of integers: `f0 = 0`, `f1 = 1`, `f2 = 1`, `f3 = 2`, `f4 = 3`, `f5 = 5`, `...`, `fn = f(n-2) + f(n-1)`. So every next number is the sum of the previous two.

Bajtek has developed a nice way to compute Fibonacci numbers on a blackboard. First, he writes a `0`. Then, below it, he writes a `1`. Then he performs the following two operations:

- operation `T`: replace the top number with the sum of both numbers;
- operation `B`: replace the bottom number with the sum of both numbers.

If he performs `n` operations, starting with `T` and then choosing operations alternately (so that the sequence of operations looks like `TBTBTB...`), the last number written will be equal to `f(n+1)`.

Unfortunately, Bajtek sometimes makes mistakes and repeats an operation two or more times in a row. For example, if Bajtek wanted to compute `f7`, then he would want to do `n = 6` operations: `TBTBTB`. If he instead performs the sequence of operations `TTTBBT`, then he will have made `3` mistakes, and he will incorrectly compute that the seventh Fibonacci number is `10`.

The number of mistakes in the sequence of operations is the number of neighbouring equal operations (`TT` or `BB`).

You are given the number `n` of operations that Bajtek has made in an attempt to compute `f(n+1)` and the number `r` that is the result of his computations (i.e. the last written number). Find the minimum possible number of mistakes that Bajtek must have made and any possible sequence of `n` operations resulting in `r` with that number of mistakes.

Assume that Bajtek always correctly starts with operation `T`.

## Input
The first line contains the integers `n` and `r` (`1 <= n, r <= 10^6`).

## Output
The first line of the output should contain one number - the minimum possible number of mistakes made by Bajtek. The second line should contain `n` characters, starting with `T`, describing one possible sequence of operations with that number of mistakes. Each character must be either `T` or `B`.

If the required sequence doesn't exist, output `IMPOSSIBLE` (without quotes).

## Examples

### Example 1
**Input**
```
6 10
```
**Output**
```
2
TBBTTB
```

### Example 2
**Input**
```
4 5
```
**Output**
```
0
TBTB
```

### Example 3
**Input**
```
2 1
```
**Output**
```
IMPOSSIBLE
```

### ideas
1. 在正确的情况下， 得到的序列应该是 0, 1, 1, 2, 3, 5, 8, ....
2. 1, 1, 2, 3, 4, 5, 6, 7, ...  
3. 假设到了某一步，目前的board = (a, b), 且剩余m步，那么如果从(a, b)开始能够正确的得到r
4. 这个验证是很快的（不超过20步）
5. 如果不行，那么就必须是错误的步骤 (这个好像不对)
6. 能不能找到两个序列(a, b)在 <= n的步骤内得到r?
7. 

## Notes

这题最关键的转化是：

- 不要正着想“怎么做 `T/B` 才能得到 `r`”
- 而要倒着想“最后黑板上的两个数可能是什么，然后怎么一步步还原回起点 `(0, 1)`”

一旦这样看，问题会变得非常像欧几里得算法。

### 1. 操作本质上在维护一对正整数

开始时黑板上是：

- `(top, bottom) = (0, 1)`

操作：

- `T`：`top = top + bottom`
- `B`：`bottom = top + bottom`

注意每做一步，最后写下的数字就是：

- 如果做的是 `T`，答案是新的 `top`
- 如果做的是 `B`，答案是新的 `bottom`

所以做完 `n` 步后，最后写下来的数 `r`，意味着最后状态一定是：

- `(r, x)`，或者
- `(x, r)`

其中 `x` 是某个正整数。

也就是说：

- 我们其实只是不知道“另一个数是多少”

### 2. 为什么倒推像欧几里得算法

假设最后状态是 `(x, y)`，并且 `x > y`。

那最后一步一定是：

- `T`

因为只有 `T` 才会修改上面的数，让它变成 `x = oldTop + y`。

所以倒推时我们只能做：

- `x = x - y`

同理，如果 `y > x`，最后一步一定是：

- `B`

倒推时只能做：

- `y = y - x`

这就和欧几里得算法的减法版本完全一样。

### 3. 为什么可以一次减很多次

如果当前 `x > y`，并且在正向过程中连续做了很多次 `T`，那么倒推时就是连续做很多次：

- `x -= y`

设连续做了 `t` 次，那么：

- `x = x - t * y`

这里为了保证中间每一步都仍然是正数，最大的合法 `t` 是：

- `t = (x - 1) / y`

同理，在 `y > x` 时：

- `t = (y - 1) / x`

这样每次可以整段地把一串连续相同操作压缩成一个 run。

这也是实现里 `analyze` 的核心。

### 4. 为什么终点必须回到 `(1, 1)`

从初始 `(0, 1)` 开始，第一步固定是 `T`，所以一步之后变成：

- `(1, 1)`

之后所有状态里的两个数都严格为正。

因此如果一个终态是合法的，倒推到最后一定会落到：

- `(1, 1)`

反过来，只要能通过这样的减法过程恰好回到 `(1, 1)`，那么这条倒推链就对应着一条合法正向操作序列。

### 5. 操作数为什么是所有 run 长度之和再加 1

倒推时我们得到的是一串 run，例如：

- 先连续减 `T` 类型 3 次
- 再连续减 `B` 类型 2 次
- ...

这些 run 反过来就是正向操作里的一段段连续相同字符。

但是别忘了：

- 正向第一步一定是额外的一个 `T`

它把 `(0, 1)` 变成 `(1, 1)`，而这一步在倒推到 `(1, 1)` 的过程中是看不见的。

所以：

- 总操作数 = `1 + 所有 run 长度之和`

这也是代码里 `total` 从 `1` 开始算的原因。

### 6. 为什么“最少 mistakes”等价于“run 数最多”

如果一个长度为 `n` 的操作串被分成了 `runs` 段连续相同字符，那么其中有：

- 每一段内部的相邻字符都相同

所以错误数就是：

- `n - runs`

因为每增加一个新的 run，就少一个“相邻相同”的位置。

所以题目的优化目标：

- 最少 mistakes

等价于：

- 在合法长度为 `n` 的方案里，让 run 数尽量多

而用倒推得到的 run 分解以后，mistakes 可以直接算成：

- `mistakes = n - runs`

### 7. 为什么枚举另一个数 `x` 就够了

因为最后写下来的结果是 `r`，所以终态只能是：

- `(r, x)` 或 `(x, r)`

其中 `x` 一定满足：

- `1 <= x < r`

`x = r` 不可能，因为倒推时两数相等且大于 `1` 无法继续，还到不了 `(1,1)`。

因此只要枚举所有可能的 `x`，分别尝试：

- `analyze(r, x, n)`
- `analyze(x, r, n)`

就不会漏解。

约束里 `r <= 10^6`，所以这样枚举是足够快的。

### 8. 一个例子

比如样例：

- `n = 6`
- `r = 10`

考虑终态 `(10, 7)`：

倒推：

- `(10, 7)`，因为 `10 > 7`，减去一次 `7`，得到 `(3, 7)`，对应一段 `T`
- `(3, 7)`，因为 `7 > 3`，最多减两次 `3`，得到 `(3, 1)`，对应一段 `BB`
- `(3, 1)`，因为 `3 > 1`，最多减两次 `1`，得到 `(1, 1)`，对应一段 `TT`

所以倒推 run 是：

- `T, BB, TT`

正向翻转以后，再补上开头固定的那个 `T`，可得：

- `TBBTTB`

它长度是 `6`，结果确实是 `10`，并且 mistakes 为：

- `1` 次 `BB`
- `1` 次 `TT`

总共 `2`

### 9. 这题真正的核心 insight

最重要的不是 Fibonacci，而是这个过程实际上在生成一对互质数，并且：

- 正向是加法
- 反向是减法
- 连续相同操作可以按 run 压缩

所以整题本质上是：

1. 枚举可能的终态 `(r, x)` / `(x, r)`
2. 用欧几里得式倒推恢复 run
3. 检查总长度是否正好为 `n`
4. 在所有合法方案里选 mistakes 最少的

这样问题就从“搜索指数级操作串”变成了“对终态做线性枚举，对每个终态做一次很快的倒推”。
