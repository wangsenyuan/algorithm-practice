# Problem Statement

Ivan has $n$ different boxes. The first of them contains some balls of $n$ different colors.

Ivan wants to play a strange game. He wants to distribute the balls into boxes in such a way that for every $i$ $(1 \leq i \leq n)$ the $i$-th box will contain all balls with color $i$.

In order to do this, Ivan will make some turns. Each turn he does the following:

1. Ivan chooses any non-empty box and takes all balls from this box.
2. Then Ivan chooses any $k$ empty boxes (the box from the first step becomes empty, and Ivan is allowed to choose it), separates the balls he took on the previous step into $k$ non-empty groups and puts each group into one of the boxes. He should put each group into a separate box. He can choose either $k = 2$ or $k = 3$.

The penalty of the turn is the number of balls Ivan takes from the box during the first step of the turn. And the penalty of the game is the total penalty of turns made by Ivan until he distributes all balls to corresponding boxes.

Help Ivan to determine the minimum possible penalty of the game!

## Input

The first line contains one integer $n$ $(1 \leq n \leq 200000)$ — the number of boxes and colors.

The second line contains $n$ integer numbers $a_1, a_2, \ldots, a_n$ $(1 \leq a_i \leq 10^9)$, where $a_i$ is the number of balls with color $i$.

## Output

Print one number — the minimum possible penalty of the game.

## Examples

### Input 1

```text
3
1 2 3
```

### Output 1

```text
6
```

### Input 2

```text
4
2 3 4 5
```

### Output 2

```text
19
```

## Note

In the first example, you take all the balls from the first box, choose $k = 3$ and sort all colors to corresponding boxes. Penalty is $6$.

In the second example, you make two turns:

1. Take all the balls from the first box, choose $k = 3$, put balls of color $3$ to the third box, of color $4$ — to the fourth box and the rest put back into the first box. Penalty is $14$.
2. Take all the balls from the first box, choose $k = 2$, put balls of color $1$ to the first box, of color $2$ — to the second box. Penalty is $5$.

Total penalty is $19$.


### ideas
1. 目标是把球按照颜色放入不同的box里面；
2. 每次操作，可以选择一个box，从中拿出所有的球（代价 = 球的数量）
3. 然后把这些球分到k = 2/3个空的box中；
4. 考虑一个不是那么好的策略，就是从box1中拿出所有的球，然后把颜色1的放回box1, 把颜色2的到box2, 把其他的放置到box3, 
5. 然后从box3,开始处理；
6. 这样子是不是一个好的策略？比如不是放入box2，而是把除去1，最多的的i，放入box[i], 然后把剩余的放入box[2]中
7. 例子
  ```
    // 21 + 13 + 5 = 39
    // 这个要怎么做到 38呢？
    // 21 分成 1 + 4 + 4， 4 + 4， 4 （3堆），其中有一堆是ok的
    // 然后将 1 + 4 + 4 分好, 9
    // 然后将 4 + 4 分好， 8
    // 刚好 38
    // 这个策略还有点复杂了～
    // 构造性的贪心问题，好难搞～
    // 一共要进行多少次操作？ 一共n个桶，每个桶
  ```
8. 上面的例子表明之前的策略是不对的。 考虑操作的序列3, 3, 3, 3, 2, 2, 2是不是最优的？
9. 假设存在一个序列3, 3, ..2, 3, .., 2, 2, 
10. 证明交换2, 3不会更差？假设第一个操作2时，该堆的大小为x, 把它分到了两个更小的堆, 且这堆中，至少包含3个数
11. 那么至少有一个堆，分到了至少两个数，（可以继续分）那么在这里把它替换成3，显然是更好的选择
12. 如果每一步都有一个颜色被分离出来，那么这个颜色最多，肯定是更优的（这样它参与的计算就更好）
13. 