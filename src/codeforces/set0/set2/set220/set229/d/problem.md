# Problem D - Beautiful City

The city of D consists of $n$ towers, built consecutively on a straight line. The height of the tower that goes $i$-th (from left to right) in the sequence equals $h_i$. The city mayor decided to rebuild the city to make it beautiful. In a beautiful city, all towers are arranged in non-descending order of their height from left to right.

The rebuilding consists of performing several (perhaps zero) operations. An operation constitutes using a crane to take any tower and put it altogether on the top of some other neighboring tower. In other words, we can take the tower that stands $i$-th and put it on the top of either the $(i-1)$-th tower (if it exists), or the $(i+1)$-th tower (if it exists). The height of the resulting tower equals the sum of heights of the two towers that were put together. After that, the two towers can't be split by any means, but more similar operations can be performed on the resulting tower. Note that after each operation the total number of towers on the straight line decreases by 1.

Help the mayor determine the minimum number of operations required to make the city beautiful.

---

## Input

The first line contains a single integer $n$ ($1 \leq n \leq 5000$) — the number of towers in the city. The next line contains $n$ space-separated integers: the $i$-th number $h_i$ ($1 \leq h_i \leq 10^5$) determines the height of the tower that is $i$-th (from left to right) in the initial tower sequence.

## Output

Print a single integer — the minimum number of operations needed to make the city beautiful.

### ideas
1. dp[i][x] 表示前i个楼，且保证第i个不高于x时的最优解
2. dp[i][x] = dp[j][y] + i - j if sum[j+1..i] >= x
3. 这个复杂度太大。不可行
4. dp[i] 表示在i处，满足good条件的队列， arr[j] = pair{i的最大值, 最优的cnt} 
5. 按照最大值升序，同时按照cnt降序
6. 这是因为，在i处的最大值越小，越有利于后面的值，只有当最大值，能够带来更有解时，才应该被保留
7. 那么这样子，似乎就可以用二分的方式去查了