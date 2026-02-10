Today is Devu's birthday. For celebrating the occasion, he bought n sweets from the nearby market. He has invited his f friends. He would like to distribute the sweets among them. As he is a nice guy and the occasion is great, he doesn't want any friend to be sad, so he would ensure to give at least one sweet to each friend.

He wants to celebrate it in a unique style, so he would like to ensure following condition for the distribution of sweets. Assume that he has distributed n sweets to his friends such that ith friend is given ai sweets. He wants to make sure that there should not be any positive integer x > 1, which divides every ai.

Please find the number of ways he can distribute sweets to his friends in the required way. Note that the order of distribution is important, for example [1, 2] and [2, 1] are distinct distributions. As the answer could be very large, output answer modulo 1000000007 (109 + 7).

To make the problem more interesting, you are given q queries. Each query contains an n, f pair. For each query please output the required number of ways modulo 1000000007 (109 + 7).

## Input

The first line contains an integer q representing the number of queries (1 ≤ q ≤ 105). Each of the next q lines contains two space space-separated integers n, f (1 ≤ f ≤ n ≤ 105).

## Output

For each query, output a single integer in a line corresponding to the answer of each query.

## Examples

**Input**

```
5
6 2
7 2
6 3
6 4
7 4
```

**Output**

```
2
6
9
10
20
```

## Note

For first query: n = 6, f = 2. Possible partitions are [1, 5] and [5, 1].

For second query: n = 7, f = 2. Possible partitions are [1, 6] and [2, 5] and [3, 4] and [4, 3] and [5, 3] and [6, 1]. So in total there are 6 possible ways of partitioning.


### ideas
1. gcd(a[?]) = 1
2. sum(a[?]) = n
3. 如果不考虑gcd的限制 nCr(n + f - 1, f - 1)
4. nCr(n - f + f - 1, f - 1) (先给每个人分配一个, 剩下的n-f个，再通过隔板进行分配)
5. 怎么保证gcd的限制呢？上面的是没有限制的然后限制 gcd(?) = 2, gcd(?) = 3, gcd(?) = ... 
6. 这个gcd(?) * f <= n, 那么似乎不会超过n
7. 那么在gcd(?) = x 的时候，怎么计算呢？
8. 且这个 n % x = 0 (肯定要能整除n)
9. 那么就是分成 n / x 份，把这些份分给f个人，保证每人分到一份 