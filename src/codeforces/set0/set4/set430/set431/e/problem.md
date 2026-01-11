# Problem E

One day two students, Grisha and Diana, found themselves in the university chemistry lab. In the lab the students found $n$ test tubes with mercury numbered from 1 to $n$ and decided to conduct an experiment.

The experiment consists of $q$ steps. On each step, one of the following actions occurs:

1. Diana pours all the contents from tube number $p_i$ and then pours there exactly $x_i$ liters of mercury.
2. Let's consider all the ways to add $v_i$ liters of water into the tubes; for each way let's count the volume of liquid (water and mercury) in the tube with water with maximum amount of liquid; finally let's find the minimum among counted maximums. That is the number the students want to count. At that, the students don't actually pour the mercury. They perform calculations without changing the contents of the tubes.

Unfortunately, the calculations proved to be too complex and the students asked you to help them. Help them conduct the described experiment.

## Input

The first line contains two integers $n$ and $q$ ($1 \leq n, q \leq 10^5$) — the number of tubes and the number of experiment steps. The next line contains $n$ space-separated integers: $h_1, h_2, \ldots, h_n$ ($0 \leq h_i \leq 10^9$), where $h_i$ is the volume of mercury in the $i$-th tube at the beginning of the experiment.

The next $q$ lines contain the game actions in the following format:

- A line of form "1 $p_i$ $x_i$" means an action of the first type ($1 \leq p_i \leq n$; $0 \leq x_i \leq 10^9$).
- A line of form "2 $v_i$" means an action of the second type ($1 \leq v_i \leq 10^{15}$).

It is guaranteed that there is at least one action of the second type. It is guaranteed that all numbers that describe the experiment are integers.

## Output

For each action of the second type print the calculated value. The answer will be considered correct if its relative or absolute error doesn't exceed $10^{-4}$.

## Examples

**Input:**

```text
3 3
1 2 0
2 2
1 2 1
2 3
```

**Output:**

```text
1.50000
1.66667
```

**Input:**

```text
4 5
1 3 0 1
2 3
2 1
1 3 2
2 3
2 4
```

**Output:**

```text
1.66667
1.00000
2.33333
2.66667
```

## 题目解释

### 问题概述

这个题目涉及两个学生进行化学实验，有 $n$ 个试管，每个试管里有一定量的汞（mercury）。实验过程中有两种操作：

### 操作类型1：更新汞的量

- 清空第 $p_i$ 个试管的所有内容
- 然后倒入 $x_i$ 升汞
- 这个操作会实际改变试管的状态

### 操作类型2：计算最优水分配方案（关键理解）

这是题目的核心难点。给定 $v_i$ 升水，我们需要：

1. **分配水**：将 $v_i$ 升水分配到 $n$ 个试管中（可以任意分配，只要总和等于 $v_i$）
   - 可以只分配到某些试管
   - 可以分配到所有试管
   - 每个试管分配的水量可以是任意非负实数

2. **计算最大值**：对于每种分配方案，找出**所有含有水的试管**中，液体总量（汞 + 水）最大的那个值

3. **找最小值**：在所有可能的分配方案中，找出上述最大值的最小值

**注意**：类型2的操作只是计算，不会实际改变试管中的汞量。

### 示例分析

以第一个例子为例：

```text
初始状态：试管1有1升汞，试管2有2升汞，试管3有0升汞
操作：2 2  (分配2升水)
```

我们需要将2升水分配到3个试管中。最优策略是：

- 将水分配到汞量较少的试管，以平衡各试管的液体总量
- 如果分配到试管1和试管3：试管1得到1升水（总量=1+1=2），试管3得到1升水（总量=0+1=1），最大值是2
- 如果只分配到试管3：试管3得到2升水（总量=0+2=2），最大值是2
- 更优方案：将水分配到试管1和试管3，使得它们的总量尽可能接近

实际上，最优方案是让所有含有水的试管的液体总量尽可能相等，这样最大值最小。

### 解题思路

这是一个**最小化最大值**的问题，可以用以下方法：

1. **二分搜索**：二分搜索答案（最大值），判断是否可以用 $v_i$ 升水使得所有含有水的试管的液体总量不超过这个值

2. **贪心策略**：优先将水分配给汞量较少的试管，使得各试管的液体总量尽可能平衡

3. **关键观察**：最优方案中，所有含有水的试管的液体总量应该尽可能相等（或接近）

### 算法复杂度

- 对于每个类型2的操作，需要高效地计算最优分配方案
- 时间复杂度：$O(n \log n)$ 或 $O(n)$ 取决于具体实现
- 总复杂度：$O(q \cdot n \log n)$ 或更优


### ideas
1. 操作1，改变了试管的量，也就是顺序
2. 操作2，就是把水尽量分配到少的试管里面
3. 但是，那些只有水银的试管，是不用算的。
4. 假设放入的试管是第i个，且它的量是x[i]
5. 那么 (v - (x[i] * i - s[i])) / i <= x[i+1] 否则的话，就使用i+1
6. 所以，这里可以二分(如果没有操作1，那么可以这么搞)
7. 但是有操作1，所以，必须使用更高效的结构，使用segment tree?
8. 4秒钟的话，可以试试 lgn * lgn 的算法
