Twilight Sparkle wanted to track the path of Nightmare Moon. Unfortunately, she didn't know the exact path. What she knew is the parity of the number of times that each place Nightmare Moon visited. Can you help Twilight Sparkle to restore any path that is consistent with this information?

Ponyville can be represented as an undirected graph (vertices are places, edges are roads between places) without self-loops and multi-edges. The path can start and end at any place (also it can be empty). Each place can be visited multiple times. The path must not visit more than 4n places.

## Input

The first line contains two integers n and m (2 ≤ n ≤ 105; 0 ≤ m ≤ 105) — the number of places and the number of roads in Ponyville. Each of the following m lines contains two integers ui, vi (1 ≤ ui, vi ≤ n; ui ≠ vi), these integers describe a road between places ui and vi.

The next line contains n integers: x1, x2, ..., xn (0 ≤ xi ≤ 1) — the parity of the number of times that each place must be visited. If xi = 0, then the i-th place must be visited even number of times, else it must be visited odd number of times.

## Output

Output the number of visited places k in the first line (0 ≤ k ≤ 4n). Then output k integers — the numbers of places in the order of path. If xi = 0, then the i-th place must appear in the path even number of times, else i-th place must appear in the path odd number of times. Note, that given road system has no self-loops, therefore any two neighbouring places in the path must be distinct.

If there is no required path, output -1. If there multiple possible paths, you can output any of them.

## Examples

**Input**

```
3 2
1 2
2 3
1 1 1
```

**Output**

```
3
1 2 3
```

**Input**

```
5 7
1 2
1 3
1 4
1 5
3 4
3 5
4 5
0 1 0 1 0
```

**Output**

```
10
2 1 3 4 5 4 5 4 3 1
```

**Input**

```
2 0
0 0
```

**Output**

```
0
```


### ideas
1. 选择一条路径的时候，会修改这条路径连接的目的地的奇偶性
2. 假设在一个欧拉回路上运动（所有的边被选择了一遍）
3. 那么所有的点，都会被访问deg/2次
4. 如果 (deg / 2) % 2 = a[i] 那么就可以用欧拉回路的算法进行访问
5. 假设 a[v]它是奇数次，那么需要deg = 6, 10, 14， （6 / 2)%2 = 1
6. 如果a[v] = 0， 那么要么它完全没有被访问到, 要么 deg = 4 (被访问到2次) 那么对于每个v，根据a[v]，算出它的最小的deg（0, 4，或者6）
7. 然后构造一个图，满足 sum(deg) % 2 = 0, 因为一条边贡献2
8. 否则没有答案。
9. 然后用deg构造一个新的图（如果构造出来，就可以用欧拉回路访问这个图）
10. 但是 (0, 4, 6)是否足够呢？
11. 或者反过来，删除边（假设u需要的边是6，如果删除一条边(u, v))
12. 可以同时满足(u, v)的约束，那么就删除它
13. 否则就要传递（dfs）