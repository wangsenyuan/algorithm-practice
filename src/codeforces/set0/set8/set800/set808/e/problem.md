# Problem E

## Description

After several latest reforms many tourists are planning to visit Berland, and Berland people understood that it's an opportunity to earn money and changed their jobs to attract tourists. Petya, for example, left the IT corporation he had been working for and started to sell souvenirs at the market.

This morning, as usual, Petya will come to the market. Petya has n different souvenirs to sell; ith souvenir is characterised by its weight wi and cost ci. Petya knows that he might not be able to carry all the souvenirs to the market. So Petya wants to choose a subset of souvenirs such that its total weight is not greater than m, and total cost is maximum possible.

Help Petya to determine maximum possible total cost.

## Input

The first line contains two integers n and m (1 ≤ n ≤ 100000, 1 ≤ m ≤ 300000) — the number of Petya's souvenirs and total weight that he can carry to the market.

Then n lines follow. ith line contains two integers wi and ci (1 ≤ wi ≤ 3, 1 ≤ ci ≤ 10^9) — the weight and the cost of ith souvenir.

## Output

Print one number — maximum possible total cost of souvenirs that Petya can carry to the market.

## Examples

### Example 1

**Input:**

```text
1 1
2 1
```

**Output:**

```text
0
```

### Example 2

**Input:**

```text
2 2
1 3
2 2
```

**Output:**

```text
3
```

### Example 3

**Input:**

```text
4 3
3 10
2 7
2 8
1 1
```

**Output:**

```text
10
```


### ideas
1. n * m 肯定不行
2. w = {1, 2, 3}
3. 假设选择了a个3，b个2，剩下的全部选择1， 那么要求 a * 3 + 2 * b + c <= m
4. a + b + c <= n
5. 迭代a的数量，从0.。。n, 需要知道 m - a * 3 时的最优解
6. 那么目前假定只能选择2和1的时候， 能不能快速的计算的？
7. 好像是同样的套路 