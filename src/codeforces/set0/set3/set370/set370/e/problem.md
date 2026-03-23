# Problem

At school Vasya got an impressive list of summer reading books. Unlike other modern schoolchildren, Vasya loves reading, so he read some book each day of the summer.

As Vasya was reading books, he was making notes in the Reader's Diary. Each day he wrote the **ordinal** number of the book he was reading. The books in the list are numbered starting from `1` and Vasya was reading them in the order they go in the list. Vasya never reads a new book until he finishes reading the previous one. Unfortunately, Vasya wasn't accurate and some days he forgot to note the number of the book and the notes for those days remained empty.

As Vasya knows that the literature teacher will want to check the Reader's Diary, he needs to restore the lost records. Help him do it and fill all the blanks. Vasya is sure that he spends at least two and at most five days for each book. Vasya finished reading all the books he had started. Assume that the reading list contained many books. So many, in fact, that it is impossible to read all of them in a summer. If there are multiple valid ways to restore the diary records, Vasya prefers the one that shows the maximum number of read books.

## Input

- First line: integer `n` — the number of summer days (`2 <= n <= 2 * 10^5`).
- Second line: `n` integers `a1, a2, ..., an` — the records in the diary in the order they were written (`0 <= ai <= 10^5`). If Vasya forgot to write the number of the book on the `i`-th day, then `ai` equals `0`.

## Output

If it is impossible to correctly fill the blanks in the diary (the diary may contain mistakes initially), print `-1`.

Otherwise, print in the first line the maximum number of books Vasya could have read in the summer if we stick to the diary. In the second line print `n` integers — the diary with correctly inserted records. If there are multiple optimal solutions, you can print any of them.

## Examples

### Example 1

Input

```text
7
0 1 0 0 0 3 0
```

Output

```text
3
1 1 2 2 3 3 3
```

### Example 2

Input

```text
8
0 0 0 0 0 0 0 0
```

Output

```text
4
1 1 2 2 3 3 4 4
```

### Example 3

Input

```text
4
0 0 1 0
```

Output

```text
1
1 1 1 1
```

### Example 4

Input

```text
4
0 0 0 3
```

Output

```text
-1
```

### ideas
1. 有两个限制， a, 按顺序读, b, 每本书最多5天，最少2天
2. 如果有一段连续的空格, 那么考虑怎么处理它们
3. 如果中间都是新的一段，那么最多可以放置 k / 2 个连续的数字
4. 最少可以放置 (k + 4) / 5 个连续的数字
5. 如果最后一个地方有数字，那还要检查是否能覆盖最后一个数字
6. 可以从后开始处理

## Notes

这题最容易卡住的地方是：

- 我们不仅要判断能不能填
- 还要在所有合法填法里，让读过的书本数尽量多

一个比较自然的思路是把“每天写了哪本书”转成“每本书对应一个连续段”。

### 1. 每本书一定对应一个连续段

因为 Vasya 读书顺序固定，而且只有读完前一本才会读下一本，所以：

- 第 `1` 本书对应一段连续区间
- 第 `2` 本书对应下一段连续区间
- ...

并且每段长度都满足：

- 至少 `2`
- 至多 `5`

所以最终答案一定长成：

- `1 1 ... 1`
- `2 2 ... 2`
- `3 3 ... 3`
- ...

只是每段具体有多长不确定。

### 2. 已知数字能告诉我们什么

如果日记里某一天写了数字 `x`，说明：

- 这一天一定落在第 `x` 本书对应的那一段里

因此对于每个出现过的数字 `x`，我们可以先看它在数组中出现的最左位置 `L[x]` 和最右位置 `R[x]`。

那么第 `x` 本书的整段必须满足：

- 覆盖区间 `[L[x], R[x]]`
- 长度在 `[2, 5]`

换句话说，第 `x` 本书的那段只可能是若干个候选区间之一。

例如如果：

- `L[x] = 10`
- `R[x] = 12`

那么它的长度至少要 `3`，最多 `5`，并且整段必须把 `10..12` 包进去。

这就把“数字 x 怎么放”变成了：

- 枚举 `x` 的若干个合法连续段候选

### 3. 为什么只需要关心出现过的书号

没有出现过的书号，整段都落在原数组的 `0` 中。

这些书不会被已有记录强行约束，所以它们只起到一个作用：

- 填补两个已知书号之间的空隙
- 或者填补开头、结尾的空隙

所以真正需要精确枚举位置的，是那些：

- 在原数组中至少出现过一次的书号

对于没出现过的书号，我们只需要知道：

- 某个空档长度 `days`
- 要塞进 `books` 本书

是否可行。

### 4. 一个空档能塞多少本书

如果一个空档长度是 `days`，要放 `books` 本书，那么每本书长度在 `[2, 5]`，因此必须满足：

- `2 * books <= days <= 5 * books`

这就是最核心的可行性条件。

反过来，如果想在一个全空白的后缀里尽量多放书，显然每本书尽量短最好，所以：

- 最多能放 `days / 2` 本书

但有一个细节：

- `days = 1` 时是完全不可能的，因为没有书能只读一天

所以对于纯空白区间：

- `0` 天可以放 `0` 本
- `1` 天无解
- `days >= 2` 时最多放 `days / 2` 本

### 5. 相邻两个已知书号之间怎么连

设两个相邻出现过的书号是：

- `x`
- `y`

并且 `y > x`

假设我们已经选好了：

- 书 `x` 的连续段
- 书 `y` 的连续段

那么它们之间的空档天数是：

- `gapDays`

而中间缺失的书本数是：

- `gapBooks = y - x - 1`

要能连起来，充要条件就是：

- `2 * gapBooks <= gapDays <= 5 * gapBooks`

如果不满足，就说明这两段选法不能同时成立。

这一步特别关键，因为它把“中间一堆 0 怎么填”压缩成了一个非常简单的区间判定。

### 6. DP 怎么做

把所有出现过的不同书号按顺序记为：

- `v1 < v2 < ... < vm`

对每个 `vi`，先枚举它所有合法的连续段候选。

然后做一个顺序 DP：

- `dp[i][j]` 表示第 `i` 个出现过的书号 `vi` 选择第 `j` 个候选段时，前面是否可行

转移时只需要检查：

1. 前一个出现过的书号 `v(i-1)` 选的段是否可行
2. 两段之间的空档是否能用缺失书号补上

开头也类似：

- 第一个出现过的书号是 `v1`
- 它前面必须正好放下 `v1 - 1` 本书
- 所以前缀长度 `prefixDays` 必须满足  
  `2 * (v1-1) <= prefixDays <= 5 * (v1-1)`

结尾则是优化目标：

- 在最后一个已知书号后面的空白里，尽量再多塞几本书
- 最优就是尽量放 `suffixDays / 2` 本

因此我们在 DP 终点只要挑：

- 可行
- 并且结尾还能延伸出最多书本数

的那个状态即可。

### 7. 为什么这样一定最优

因为所有已知书号的位置限制，都已经被候选段精确表达了。

在这些限制之外，剩下的自由度只有：

- 前缀能否放下 `1..v1-1`
- 中间空档能否放下缺失书号
- 后缀还能再接多少本书

而前两者只决定“可不可行”，最后一项才决定“总书本数最大”。

所以：

1. 先枚举所有可能的已知段摆法
2. 只保留可行摆法
3. 在可行终点里选让后缀书本数最大的

就一定得到全局最优。

### 8. 最后怎么生成答案数组

一旦确定了每个出现过的书号对应的那段区间，剩下就简单了：

1. 前缀空档按顺序放 `1, 2, ..., v1-1`
2. 两个已知书号之间的空档按顺序放缺失书号
3. 后缀继续往后放更多的新书

每个空档的具体长度分配方法很简单：

- 先给每本书 `2` 天
- 剩余的天数再从前往后，每本最多多加 `3` 天

因为只要满足总天数在 `[2*books, 5*books]` 之间，这种分法一定能成功。

### 9. 这一题真正的思维转换

最关键的转换是：

- 不要直接按“每天填什么”去想
- 而要按“每本书是一段连续区间”去想

这样以后：

- 已知数字 => 某些书的区间必须覆盖这些位置
- 未知数字 => 只是若干空档需要用长度 `[2,5]` 的段去填满

于是原题就变成了：

- 枚举已知书号的合法区间
- 检查相邻区间之间的空档能否用缺失书号填满
- 最后在后缀尽量多放新书

这就是实现里“候选段 + DP + 回溯构造”的本质。


### editorial

For each book number that is in the sequence, find the leftmost and the rightmost position of this number. In other words, for each such book number we find a segment of positions that should consist of this number. If for some pair of numbers there segments intersect, it is impossible to construct the answer. The same thing happens if some segment has length more than 5. It is reasonable to separately handle the case when all given numbers are zeroes. In this case, fill in the numbers greedily, spending 2 days on each book (probably, except the last one).

So, we have some blocks of numbers and gaps between them. Lets do the following DP: each state of DP is described by two values (i, j): i means the number of block (lets enumerate them consecutively), j means how far to the right will this block eventually extend (if there is a gap after this block, it is possible that we fill some prefix of this gap with the same book number that is in the block). It is clear that j - i will not exceed 5, so we actually can describe the state by values (i, j - i), which may sound more convenient. So, the number of states is linear. Lets say that D(i, j) is true if it it possible to correctly fill all the gaps that come before the i-th block, under condition that the i-th block extends to the position j, and D(i, j) is false otherwise. To calculate the value of D(i, j), lets try to extend the i-th block to the left in all (not so many) possible ways (to replace some number of consecutive zeroes that are in the gap just before the i-th block). Then, try to fix where the previous block can actually end (fix the state D(i - 1, k), where D(i - 1, k) is true, of course). To make a transition in DP, we should check whether it possible or not to fill the rest of the gap between the (i - 1)-th block and the i-th block. Lets say that (i - 1)-th block consists of number x, the i-th block consists of number y, and there are f still unfilled positions in the gap. Than the gap can be correctly filled if and only if 2·(y - x - 1) ≤ f ≤ 5·(y - x - 1).

If you understand this DP, it won’t be difficult for you to find out how to construct the answer from it.