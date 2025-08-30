# Red-Green Tower

There are r red and g green blocks for construction of the red-green tower. Red-green tower can be built following next rules:

## Rules

1. **Red-green tower consists of some number of levels**

2. **Level structure**: Let the red-green tower consist of n levels, then:
   - The first level of this tower should consist of n blocks
   - Second level — of n - 1 blocks
   - The third one — of n - 2 blocks
   - And so on — the last level of such tower should consist of one block
   
   In other words, each successive level should contain one block less than the previous one.

3. **Color consistency**: Each level of the red-green tower should contain blocks of the same color.

## Problem Statement

Let h be the maximum possible number of levels of red-green tower, that can be built out of r red and g green blocks meeting the rules above. The task is to determine how many different red-green towers having h levels can be built out of the available blocks.

Two red-green towers are considered different if there exists some level, that consists of red blocks in the one tower and consists of green blocks in the other tower.

You are to write a program that will find the number of different red-green towers of height h modulo 10^9 + 7.

## Input

The only line of input contains two integers r and g, separated by a single space — the number of available red and green blocks respectively (0 ≤ r, g ≤ 2·10^5, r + g ≥ 1).

## Output

Output the only integer — the number of different possible red-green towers of height h modulo 10^9 + 7.

## Examples

### Example 1
**Input:**
```
4 6
```
**Output:**
```
2
```

### Example 2
**Input:**
```
9 7
```
**Output:**
```
6
```

### Example 3
**Input:**
```
1 1
```
**Output:**
```
2
```

## Note

The image in the problem statement shows all possible red-green towers for the first sample.


### ideas
1. h要怎么计算出来。h <= 1000的样子
2. h * (h + 1) / 2 <= r + g
3. dp[h][r] = 在第h列填满后，剩余r个时的计数
4. dp[h+1][r - (h + 1)] +=dp[h][r] (如果新的一列使用red)
5. dp[h+1][r] += dp[h][r] 新的一列使用green
6. 但是必须知道green剩余多少 g0 =  h * (h + 1) / 2 - r0 (使用掉的red)
7. 剩余的g = green - g0
8. 