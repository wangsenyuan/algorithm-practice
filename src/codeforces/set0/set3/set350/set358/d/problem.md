# Problem D: Dima and Hares

Dima liked the present he got from Inna very much. He liked the present he got from Seryozha even more.

Dima felt so grateful to Inna about the present that he decided to buy her $n$ hares. Inna was very happy. She lined up the hares in a row, numbered them from $1$ to $n$ from left to right and started feeding them with carrots. Inna was determined to feed each hare exactly once. But in what order should she feed them?

Inna noticed that each hare radiates joy when she feeds it. And the joy of the specific hare depends on whether Inna fed its adjacent hares before feeding it. Inna knows how much joy a hare radiates if it eats when either both of his adjacent hares are hungry, or one of the adjacent hares is full (that is, has been fed), or both of the adjacent hares are full. Please note that hares number $1$ and $n$ don't have a left and a right-adjacent hare correspondingly, so they can never have two full adjacent hares.

Help Inna maximize the total joy the hares radiate. :)

## Input

- The first line of the input contains integer $n$ ($1 \leq n \leq 3000$) — the number of hares.
- Then three lines follow, each line has $n$ integers:
  - The first line contains integers $a_1, a_2, \ldots, a_n$.
  - The second line contains $b_1, b_2, \ldots, b_n$.
  - The third line contains $c_1, c_2, \ldots, c_n$.
- The following limits are fulfilled: $0 \leq a_i, b_i, c_i \leq 10^5$.

Number $a_i$ in the first line shows the joy that hare number $i$ gets if his adjacent hares are both hungry. Number $b_i$ in the second line shows the joy that hare number $i$ radiates if he has exactly one full adjacent hare. Number $c_i$ in the third line shows the joy that hare number $i$ radiates if both his adjacent hares are full.

## Output

In a single line, print the maximum possible total joy of the hares Inna can get by feeding them.

## ideas
1. 假设现在处理到了i, i可能处在两种状态（可能更多，先假设两种）
2. 一种是它已经被喂过了；一种是没有
3. 如果已经喂过了，那么对于i+1来说，只能考虑有一个，还是有两个邻居被喂过的情况
4. 如果没有被喂过，那么要考虑3种情况
5. 所以，这里要考虑连续的3个兔子的情况
6. dp[i][state]表示到i为止，state表示的前两个兔子的情况
7. 0表示都没有喂过，1表示，只喂过了i-1, 2表示只喂过了i-2, 3表示两个都喂过了
8. 现在处理i
9. 如果0，那么在处理到i+1前，必须要吧i-2给喂掉（这里必须保证i-3已经被喂过了，或者在边界外）
10. 但是这里的选择似乎有点麻烦。但应该可以枚举，只要保证i-2被喂掉就好了
11. 但是这样的算法似乎是 o(n)的。似乎不大可能呐
12. 估计这个是错的，算不出来最优解
13. dp[l][r][s] 表示建l...r处理掉后，且两头的状态是s, 0/1/2/3
14. dp[l][r][s] = 先处理l, 再处理l+1...r, 或者先处理l+1...r, 再处理l
15.               先处理r, 再处理l...r-1, 或者先处理l...r-1, 再处理r
16. 