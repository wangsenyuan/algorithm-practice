# Xenia's Weight Scales Problem

## Problem Description

Xenia has a set of weights and pan scales. Each weight has an integer weight from 1 to 10 kilos. Xenia is going to play with scales and weights a little. For this, she puts weights on the scalepans, one by one. The first weight goes on the left scalepan, the second weight goes on the right scalepan, the third one goes on the left scalepan, the fourth one goes on the right scalepan and so on. Xenia wants to put the total of m weights on the scalepans.

Simply putting weights on the scales is not interesting, so Xenia has set some rules:

1. **No consecutive same weights**: She does not put on the scales two consecutive weights of the same weight. That is, the weight that goes i-th should be different from the (i + 1)-th weight for any i (1 ≤ i < m).

2. **Always outweigh the other pan**: Every time Xenia puts a weight on some scalepan, she wants this scalepan to outweigh the other one. That is, the sum of the weights on the corresponding scalepan must be strictly greater than the sum on the other pan.

You are given all types of weights available for Xenia. You can assume that the girl has an infinite number of weights of each specified type. Your task is to help Xenia lay m weights on the scales or to say that it can't be done.

## Input

The first line contains a string consisting of exactly ten zeroes and ones: the i-th (i ≥ 1) character in the line equals "1" if Xenia has i kilo weights, otherwise the character equals "0". The second line contains integer m (1 ≤ m ≤ 1000).

## Output

In the first line print "YES", if there is a way to put m weights on the scales by all rules. Otherwise, print in the first line "NO". If you can put m weights on the scales, then print in the next line m integers — the weights' weights in the order you put them on the scales.

If there are multiple solutions, you can print any of them.

## Examples

### Example 1

**Input:**
```
0000000101
3
```

**Output:**
```
YES
8 10 8
```

### Example 2

**Input:**
```
1000000000
2
```

**Output:**
```
NO
```

## ideas
1. 假设有两个重量 a, b, 假设 a < b
   ```
   i.  n * a < (n + 1) * b
   ii. (n + 1) * a > n * b
   ``` 
2. ii => a > n * (b - a)， a最大是9（b = 10), 9 > n => n = 8也就是说，如果只有(9, 10), m不能超过8
3. 假设现在两边分别是sl, sr，且sl <= sr成立，目前要在左边放置一个w
4. sl + w1 > sr成立， 
5. sl + w1 < sr + w2
6. w1 - w2 < sr - sl
7. w1 > sr - sl 的 
8. dp[w][d][0/1][i] 表示最后一次放置w在（0表示左边,1表示右边), 且得到的两边的差值是d（也和0/1有关），已经进行了i步后的结果