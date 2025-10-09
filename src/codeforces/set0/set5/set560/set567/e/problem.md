Berland has n cities, the capital is located in city s, and the historic home town of the President is in city t (s ≠ t). The cities are connected by one-way roads, the travel time for each of the road is a positive integer.

Once a year the President visited his historic home town t, for which his motorcade passes along some path from s to t (he always returns on a personal plane). Since the president is a very busy man, he always chooses the path from s to t, along which he will travel the fastest.

The ministry of Roads and Railways wants to learn for each of the road: whether the President will definitely pass through it during his travels, and if not, whether it is possible to repair it so that it would definitely be included in the shortest path from the capital to the historic home town of the President. Obviously, the road can not be repaired so that the travel time on it was less than one. The ministry of Berland, like any other, is interested in maintaining the budget, so it wants to know the minimum cost of repairing the road. Also, it is very fond of accuracy, so it repairs the roads so that the travel time on them is always a positive integer.

## Input

The first lines contain four integers n, m, s and t (2 ≤ n ≤ 105; 1 ≤ m ≤ 105; 1 ≤ s, t ≤ n) — the number of cities and roads in Berland, the numbers of the capital and of the Presidents' home town (s ≠ t).

Next m lines contain the roads. Each road is given as a group of three integers ai, bi, li (1 ≤ ai, bi ≤ n; ai ≠ bi; 1 ≤ li ≤ 106) — the cities that are connected by the i-th road and the time needed to ride along it. The road is directed from city ai to city bi.

The cities are numbered from 1 to n. Each pair of cities can have multiple roads between them. It is guaranteed that there is a path from s to t along the roads.

## Output

Print m lines. The i-th line should contain information about the i-th road (the roads are numbered in the order of appearance in the input).

If the president will definitely ride along it during his travels, the line must contain a single word "YES" (without the quotes).

Otherwise, if the i-th road can be repaired so that the travel time on it remains positive and then president will definitely ride along it, print space-separated word "CAN" (without the quotes), and the minimum cost of repairing.

If we can't make the road be such that president will definitely ride along it, print "NO" (without the quotes).

## Examples

### Input

```
6 7 1 6
1 2 2
1 3 10
2 3 7
2 4 8
3 5 3
4 5 2
5 6 1
```

### Output

```
YES
CAN 2
CAN 1
CAN 1
CAN 1
CAN 1
YES
```

### Input

```
3 3 1 3
1 2 10
2 3 10
1 3 100
```

### Output

```
YES
YES
CAN 81
```

### Input

```
2 2 1 2
1 2 1
1 2 2
```

### Output

```
YES
NO
```

## Note

The cost of repairing the road is the difference between the time needed to ride along it before and after the repairing.

In the first sample president initially may choose one of the two following ways for a ride: 1 → 2 → 4 → 5 → 6 or 1 → 2 → 3 → 5 → 6.


### ideas
1. 有点绕～
2. 先找到s->t的一条最短路径。但是怎么知道那些肯定会被选中的路径呢？
3. 因为会出现类似 a -> b -> c -> d -> e
4.  a -> u -> c -> v -> e 这样的情况
5.  先找出从s -> t 的所有的路径节点，组成一个新的无向图。然后找出这个图中的critical bridge 
6.  那么这些bridge肯定会被选中 => YES
7.  然后对于那些CAN的，就是它们必须提供更短的路径长度( > 0)
8.  如果它到s -> u + f(u, v) + v -> t  < min_dist => CAN