Of course our child likes walking in a zoo. The zoo has `n` areas, that are numbered from `1` to `n`. The `i`-th area contains `ai` animals in it. Also there are `m` roads in the zoo, and each road connects two distinct areas. Naturally the zoo is connected, so you can reach any area of the zoo from any other area using the roads.

Our child is very smart. Imagine the child want to go from area `p` to area `q`. Firstly he considers all the simple routes from `p` to `q`. For each route the child writes down the number, that is equal to the minimum number of animals among the route areas. Let's denote the largest of the written numbers as `f(p, q)`. Finally, the child chooses one of the routes for which he writes down the value `f(p, q)`.

After the child has visited the zoo, he thinks about the question: what is the average value of `f(p, q)` for all pairs `p, q` (`p ≠ q`)? Can you answer his question?

## Input

The first line contains two integers `n` and `m` (`2 ≤ n ≤ 10^5`; `0 ≤ m ≤ 10^5`). 

The second line contains `n` integers: `a1, a2, ..., an` (`0 ≤ ai ≤ 10^5`). 

Then follow `m` lines, each line contains two integers `xi` and `yi` (`1 ≤ xi, yi ≤ n`; `xi ≠ yi`), denoting the road between areas `xi` and `yi`.

All roads are bidirectional, each pair of areas is connected by at most one road.

## Output

Output a real number — the average value of `f(p, q)` for all pairs `p, q` (`p ≠ q`).

The answer will be considered correct if its relative or absolute error doesn't exceed `10^-4`.

## Examples

### Example 1

**Input:**
```
4 3
10 20 30 40
1 3
2 3
4 3
```

**Output:**
```
16.666667
```

### Example 2

**Input:**
```
3 3
10 20 30
1 2
2 3
3 1
```

**Output:**
```
13.333333
```

### Example 3

**Input:**
```
7 8
40 20 10 30 20 50 40
1 2
2 3
3 4
4 5
5 6
6 7
1 4
5 7
```

**Output:**
```
18.571429
```

## Note

Consider the first sample. There are 12 possible situations:

- `p = 1, q = 3, f(p, q) = 10`
- `p = 2, q = 3, f(p, q) = 20`
- `p = 4, q = 3, f(p, q) = 30`
- `p = 1, q = 2, f(p, q) = 10`
- `p = 2, q = 4, f(p, q) = 20`
- `p = 4, q = 1, f(p, q) = 10`

Another 6 cases are symmetrical to the above. The average is `(10 + 20 + 30 + 10 + 20 + 10) × 2 / 12 = 200 / 12 = 16.666667`.

Consider the second sample. There are 6 possible situations:

- `p = 1, q = 2, f(p, q) = 10`
- `p = 2, q = 3, f(p, q) = 20`
- `p = 1, q = 3, f(p, q) = 10`

Another 3 cases are symmetrical to the above. The average is `(10 + 20 + 10) × 2 / 6 = 80 / 6 = 13.333333`.

### ideas
1. 考虑一条路径，f(u, v)的值 = 这条路径上的最小值
2. 如果按照x[i]降序考虑, 加入x[i]时，它的贡献是多少呢？
3. 把它和比它大的邻居连接后，假设整体的sz = x, 那么可以计算出它的贡献
4. 最后再除以 n * (n - 1) / 2