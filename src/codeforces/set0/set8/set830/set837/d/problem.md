# Problem Description

Let's call the roundness of the number the number of zeros to which it ends.

You have an array of n numbers. You need to choose a subset of exactly k numbers so that the roundness of the product of the selected numbers will be maximum possible.

## Input

The first line contains two integer numbers n and k (1 ≤ n ≤ 200, 1 ≤ k ≤ n).

The second line contains n space-separated integer numbers a1, a2, ..., an (1 ≤ ai ≤ 10^18).

## Output

Print maximal roundness of product of the chosen subset of length k.

## Examples

### Example 1

**Input:**

```text
3 2
50 4 20
```

**Output:**

```text
3
```

### Example 2

**Input:**

```text
5 3
15 16 3 25 9
```

**Output:**

```text
3
```

### Example 3

**Input:**

```text
3 3
9 77 13
```

**Output:**

```text
0
```

## Note

In the first example there are 3 subsets of 2 numbers. [50, 4] has product 200 with roundness 2, [4, 20] — product 80, roundness 1, [50, 20] — product 1000, roundness 3.

In the second example subset [15, 16, 25] has product 6000, roundness 3.

In the third example all subsets has product with roundness 0.


### ideas
1. 假设一个集合中，一共有x个0，除去这些0后，还剩余w个2，v个5, 那么又可以造出min(w, v)个0
2. 所以，所有的数，都可以变成这样一个结构{0, 2, 5}
3. dp[i][0][5] = 最多5的个数
4. 前m个数子，选择i个数，其中0的数量是 cnt0, 2的数量 = cnt[5]时，最大的cnt5的数量
5. 0的个数，可以一开始算出来， 不会超过n * 18  = 2000
6. 但是这样子， 5的个数，差不多也是2000， 这样子就不行了
7. 5在遇到2的时候，就会被消耗掉，假设目前的这个数500000
8. 那么就不断的去消耗之前的2，并且转化成0