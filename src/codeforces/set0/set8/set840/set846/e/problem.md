# Problem Description

Igor is a post-graduate student of chemistry faculty in Berland State University (BerSU). He needs to conduct a complicated experiment to write his thesis, but laboratory of BerSU doesn't contain all the materials required for this experiment.

Fortunately, chemical laws allow material transformations (yes, chemistry in Berland differs from ours). But the rules of transformation are a bit strange.

Berland chemists are aware of n materials, numbered in the order they were discovered. Each material can be transformed into some other material (or vice versa). Formally, for each i (2 ≤ i ≤ n) there exist two numbers xi and ki that denote a possible transformation: ki kilograms of material xi can be transformed into 1 kilogram of material i, and 1 kilogram of material i can be transformed into 1 kilogram of material xi. Chemical processing equipment in BerSU allows only such transformation that the amount of resulting material is always an integer number of kilograms.

For each i (1 ≤ i ≤ n) Igor knows that the experiment requires ai kilograms of material i, and the laboratory contains bi kilograms of this material. Is it possible to conduct an experiment after transforming some materials (or none)?

## Input

The first line contains one integer number n (1 ≤ n ≤ 10^5) — the number of materials discovered by Berland chemists.

The second line contains n integer numbers b1, b2... bn (1 ≤ bi ≤ 10^12) — supplies of BerSU laboratory.

The third line contains n integer numbers a1, a2... an (1 ≤ ai ≤ 10^12) — the amounts required for the experiment.

Then n - 1 lines follow. j-th of them contains two numbers xj + 1 and kj + 1 that denote transformation of (j + 1)-th material (1 ≤ xj + 1 ≤ j, 1 ≤ kj + 1 ≤ 10^9).

## Output

Print YES if it is possible to conduct an experiment. Otherwise print NO.

## Examples

### Example 1

**Input:**

```text
3
1 2 3
3 2 1
1 1
1 1
```

**Output:**

```text
YES
```

### Example 2

**Input:**

```text
3
3 2 1
1 2 3
1 1
1 2
```

**Output:**

```text
NO
```


### ideas
1. 对于物质i, 当前有b[i]，需要a[i]， 问是否能使的所有i，最后的容量w[i] >= a[i]
2. 如果 b[i] > a[i] => 那么就将其转化，但是目标是什么呢？
3. i = x[j] < j, 表示 k[j]的i物质，转化为1单位的j位置，1单位的j物质，可以变成1单位的i物质
4. 假设 j = 2, i = 1, 切k[j] = 2, 那么开始时有[8, 1]
5. [6, 2] => [4, 3], [2, 4], [0, 5]
6. 从前往后，如果b[1] > a[1], 那么多出来的，肯定要往后面推（但可以先记录下来）
7. 如果少了，那么必须从后面的x[?] = 1的那里借过来。但是从谁那里借，也还是有问题（因为有的借走了，可能不大好补回去）
8. 先不管，假设到了i, 目前它不够了，那么可能从前面多的拿来用？（也还是面临选择问题）
9. 感觉它们是组成了一颗以1为root的树
10. 对于某个子树u，先处理它的子树v,
11. 然后就可以得到这样一个关系，有些子树会有盈余，这些都需要聚集的u上面，有些是有亏损，如果有亏损，那么就必须使用u的盈余去填补
12. 好搞了
13. 