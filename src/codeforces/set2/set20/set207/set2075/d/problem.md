You are given two non-negative integers 𝑥
 and 𝑦
.

You can perform the following operation any number of times (possibly zero): choose a positive integer 𝑘
 and divide either 𝑥
 or 𝑦
 by 2𝑘
 rounding down. The cost of this operation is 2𝑘
. However, there is an additional constraint: you cannot select the same value of 𝑘
 more than once.

Your task is to calculate the minimum possible cost in order to make 𝑥
 equal to 𝑦
.

### ideas
1. 7 / 4 = 1   0111
2. 9 / 4 = 2   1001
3. 除 pow(2, k), 右移k次
4. 假设要右移k次，最优的方案，应该是进行k次右移1，这个时候cost = 2 * k（而不是 pow(2, k))
5. 直到高位部分一致为止
6. 完蛋了，每个k只能被用一次
7. 假设x需要被右移kx次，y需要被右移ky次
8. (1 + 2 + .. + i) = kx
9. (1 + 2 + .. + j) = ky
10. 有个感觉，就是一边把最大的数k操作掉，另外一边，尽可能的操作小的部分 