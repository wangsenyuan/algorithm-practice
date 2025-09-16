# Problem Description

After you have read all the problems, probably, you think Alex is genius person. That's true! One day he came up with the following task.

Given a sequence of integer numbers a₁, a₂, ..., aₙ. You are to find a longest sequence b₁, b₂, ..., b₄ₘ, that satisfies the following conditions:

- **b₄ₖ₊₁ = b₄ₖ₊₃** for all valid integer k
- **b₄ₖ₊₂ = b₄ₖ₊₄** for all valid integer k  
- sequence b is subsequence of a (not necessarily contiguous subsequence)

And finally... Alex had given this complicated task to George, and George gave it to you. Help George to cope with the task.

## Input

The first line contains a single integer n (1 ≤ n ≤ 5·10⁵). The next line contains n integers a₁, a₂, ..., aₙ (1 ≤ aᵢ ≤ 10⁹).

## Output

In the first line print a single integer 4m — the maximal possible length of required sequence b. In the second line print 4m integers b₁, b₂, ..., b₄ₘ, that is required sequence.

If there are multiple optimal answers you may print any of them.

## Examples

### Example 1
**Input:**
```
4
3 5 3 5
```

**Output:**
```
4
3 5 3 5
```

### Example 2
**Input:**
```
10
35 1 2 1 2 35 100 200 100 200
```

**Output:**
```
8
1 2 1 2 100 200 100 200
```


### ideas
1. 每4个一组，每组里面 (1, 3)相同，（2， 4)相同
2. dp[i][0/1] 表示到i为止最多的b的长度(dp[i][0]表示正好4的倍数，dp[i][1]表示 % 4 = 2)
3. dp[i][0] = dp[i-1][0] (不使用i), 或者 dp[j][1] + 2 (j是last(a[i]) - 1)
4. dp[i][1] = dp[i-1][1], dp[j][0] + 2 (这个不对，因为是要1, 3 相同， 2, 4 相同) 
5. fp[i] = j 表示最大的j, 在j的后面存在一个 a[k] = a[i], 且在k的后面存在一个a[j]
6. 所以问题的关键在于fp的计算
7. 在给定i的情况下，j其实也是不确定的，就是它可以是之前a[i]出现过的任何数
8. 如果j是a[i]出现的前一个位置，那么 fp[i] >= fp[j] 肯定是成立的（将j替换成i，4元组还是成立的）
9. 新的fp[i]只可能是那些，只好跨过了a[j]出现在它两边的数
10. 现在(j, i)确定了，如何找到fp[i]?
11. 3 4 1 2 3 1 2 4
12. 虽然从fp的角度看, 最后一个4应该去匹配第一个3，但是从整体的角度看， 这个fp不可能给出最优的结果，反而是使用最后一个2是更优的
13. fp[i] = max(fp[i-1], fp[j] j是上一个a[i]) 那这样子确实是可以滚动的，就是可以用前面处理过的结果去更新当前的位置
14. y a x .z. x y a
15. 在两个x中间出现的z, 把x前面的z都舍弃掉，似乎是没问题的。它的贡献要么在x中，如果后面某一个数能指向到这个z，那么它肯定不会指向前面去
16. 但是在考虑z对x的贡献时，似乎又是必须要考虑进去的