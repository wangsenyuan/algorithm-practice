# Problem E: Iahub and Permutations

## Problem Statement

Iahub is so happy about inventing bubble sort graphs that he's staying all day long at the office and writing permutations. Iahubina is angry that she is no more important for Iahub. When Iahub goes away, Iahubina comes to his office and sabotages his research work.

The girl finds an important permutation for the research. The permutation contains n distinct integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq n$). She replaces some of permutation elements with -1 value as revenge.

When Iahub finds out his important permutation is broken, he tries to recover it. The only thing he remembers about the permutation is it didn't have any fixed point. A **fixed point** for a permutation is an element $a_k$ which has value equal to $k$ ($a_k = k$). 

Your job is to prove to Iahub that trying to recover it is not a good idea. Output the number of permutations which could be originally Iahub's important permutation, modulo $10^9 + 7$.

## Input

The first line contains integer $n$ ($2 \leq n \leq 2000$). On the second line, there are $n$ integers, representing Iahub's important permutation after Iahubina replaces some values with -1.

It's guaranteed that:
- There are no fixed points in the given permutation
- The given sequence contains at least two numbers -1
- Each positive number occurs in the sequence at most once
- There is at least one suitable permutation

## Output

Output a single integer, the number of ways Iahub could recover his permutation, modulo $10^9 + 7$.

## Examples

### Example 1
**Input:**
```
5
-1 -1 4 3 -1
```

**Output:**
```
2
```

**Explanation:**
For the first test example there are two permutations with no fixed points: $[2, 5, 4, 3, 1]$ and $[5, 1, 4, 3, 2]$. Any other permutation would have at least one fixed point.


### ideas
1. 考虑长度为n的序列，组成的good的序列为f(n)
2. f(1) = 0, f(2) = 1
3. f(i) = (i-1) * f(i-1) ? 先排列好(i-1)好的序列，然后把i和任意一个位置交换
4. i = 3,  (3, 1, 2) (2, 3, 1) = 2 * f(1) = 2 
5. 确实是成立的
6. 回到原问题，如果在位置i出现了数字x， 且数字i没有出现时，i是可以随便在任何位置上出现的（它不会造成问题）
7. 那些出现的数字，如果它们组成了一个环，直接忽略它们就可以了
8. 否则，肯定有一个数字变成了free
9. 对于序列[-1 -1 4 3 -1] = f(3) (出现的数字组成了一个环)
10. 但是 [-1 -1 4 -1 -1], 3如果出现在位置4， 那么就是 f(3), 如果3出现在了位置1， 那么就需要考虑1的选择了
11. f(1) + f(2) + f(3) ? 每增加一个数字，且没有组成环的情况下，就多了一个环的可能性；
12. 但是如果是x个free的节点呢？ 
13. 比如序列 【-1， -1， 4， -1， 2]
14. free = [3, 5]
15. dp[n][m] 表示有n个位置时，目前有m个free的数字 dp[n][0] = f(n)
16. dp[n][1] = sum(f[1] + f(2) + ... + f(n))
17. dp[n][2] = dp[n - 1][1] (5放到位置2上， 或者把3放到位置4上面) + dp[n-1][2] * 2 * (n - 1) （5放到了位置1上面，现在新的free = [1, 3])
18. 5如果放到位置3上面 dp[n][m] = dp[n-1][m-1] + dp[n-1][m] * m * (n - 1) 只要不是放到合适的位置m不会减少
19. 当n = m 的时候， = f(n) 