# Problem Description

Iahub is a big fan of tourists. He wants to become a tourist himself, so he planned a trip. There are n destinations on a straight road that Iahub wants to visit. Iahub starts the excursion from kilometer 0. The n destinations are described by a non-negative integers sequence a₁, a₂, ..., aₙ. The number aₖ represents that the kth destination is at distance aₖ kilometers from the starting point. No two destinations are located in the same place.

Iahub wants to visit each destination only once. Note that, crossing through a destination is not considered visiting, unless Iahub explicitly wants to visit it at that point. Also, after Iahub visits his last destination, he doesn't come back to kilometer 0, as he stops his trip at the last destination.

The distance between destination located at kilometer x and next destination, located at kilometer y, is |x - y| kilometers. We call a "route" an order of visiting the destinations. Iahub can visit destinations in any order he wants, as long as he visits all n destinations and he doesn't visit a destination more than once.

Iahub starts writing out on a paper all possible routes and for each of them, he notes the total distance he would walk. He's interested in the average number of kilometers he would walk by choosing a route. As he got bored of writing out all the routes, he asks you to help him.

## Input

The first line contains integer n (2 ≤ n ≤ 10⁵). Next line contains n distinct integers a₁, a₂, ..., aₙ (1 ≤ aᵢ ≤ 10⁷).

## Output

Output two integers — the numerator and denominator of a fraction which is equal to the wanted average number. The fraction must be irreducible.

## Examples

### Input
```
3
2 3 5
```

### Output
```
22 3
```

## Note

Consider 6 possible routes:

- **[2, 3, 5]**: total distance traveled: |2 – 0| + |3 – 2| + |5 – 3| = 5
- **[2, 5, 3]**: |2 – 0| + |5 – 2| + |3 – 5| = 7
- **[3, 2, 5]**: |3 – 0| + |2 – 3| + |5 – 2| = 7
- **[3, 5, 2]**: |3 – 0| + |5 – 3| + |2 – 5| = 8
- **[5, 2, 3]**: |5 – 0| + |2 – 5| + |3 – 2| = 9
- **[5, 3, 2]**: |5 – 0| + |3 – 5| + |2 – 3| = 8

The average travel distance is $\frac{5 + 7 + 7 + 8 + 9 + 8}{6} = \frac{44}{6} = \frac{22}{3}$.


## ideas
1. 这个题目有点怪啊。正常情况下，除数是F[n], 因为有F[n]中排列，每种排列代表一条路径。 但显然不大可能
2. 考虑一个点的贡献， a[i], 在前后各有l, n - 1 - l 个点
3. 它的整体的贡献 = (l * a[i] - pref(i-1) + suf[i+1] - (n - 1 - l) * a[i]) * F(n-1)
4. 所以，不考虑F(n), F(n-1), 最后除以n即可
5. 2, 3, 5   = (8 - 4) + (1 + 2) + (3 + 2) = 4 + 3 + 5 = 12
6. 12 / 3 = 4 (不大对) 除以n没有问题
7. 考虑2的贡献 = (2, 3)一共出现了4次， 4 + 2 * 4 + 3 * 4 = 24 + （2 + 3 + 5） * 2 = 44
8.  F(n-1) / F(n) = ((1 + 2 + 3) * 2) / 3 + (2 + 3 + 5) / 3