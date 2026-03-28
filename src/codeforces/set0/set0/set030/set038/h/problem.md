On the Berland Dependence Day it was decided to organize a great marathon. Berland consists of n cities, some of which are linked by two-way roads. Each road has a certain length. The cities are numbered from 1 to n. It is known that one can get from any city to any other one by the roads.

n runners take part in the competition, one from each city. But Berland runners are talkative by nature and that's why the juries took measures to avoid large crowds of marathon participants. The jury decided that every runner should start the marathon from their hometown. Before the start every sportsman will get a piece of paper containing the name of the city where the sportsman's finishing line is. The finish is chosen randomly for every sportsman but it can't coincide with the sportsman's starting point. Several sportsmen are allowed to finish in one and the same city. All the sportsmen start simultaneously and everyone runs the shortest route from the starting point to the finishing one. All the sportsmen run at one speed which equals to 1.

After the competition a follow-up table of the results will be composed where the sportsmen will be sorted according to the nondecrease of time they spent to cover the distance. The first g sportsmen in the table will get golden medals, the next s sportsmen will get silver medals and the rest will get bronze medals. Besides, if two or more sportsmen spend the same amount of time to cover the distance, they are sorted according to the number of the city where a sportsman started to run in the ascending order. That means no two sportsmen share one and the same place.

According to the rules of the competition the number of gold medals g must satisfy the inequation g1 ≤ g ≤ g2, where g1 and g2 are values formed historically. In a similar way, the number of silver medals s must satisfy the inequation s1 ≤ s ≤ s2, where s1 and s2 are also values formed historically.

At present, before the start of the competition, the destination points of every sportsman are unknown. However, the press demands details and that's why you are given the task of counting the number of the ways to distribute the medals. Two ways to distribute the medals are considered different if at least one sportsman could have received during those distributions different kinds of medals.

## Input

The first input line contains given integers n and m (3 ≤ n ≤ 50, n - 1 ≤ m ≤ 1000), where n is the number of Berland towns and m is the number of roads.

Next in m lines road descriptions are given as groups of three integers v, u, c, which are the numbers of linked towns and its length (1 ≤ v, u ≤ n, v ≠ u, 1 ≤ c ≤ 1000). Every pair of cities have no more than one road between them.

The last line contains integers g1, g2, s1, s2 (1 ≤ g1 ≤ g2, 1 ≤ s1 ≤ s2, g2 + s2 < n). The input data numbers, located on one line, are space-separated.

## Output

Print the single number — the number of ways to distribute the medals. It is guaranteed that the number fits in the standard 64-bit signed data type.

## Examples

### Example 1

**Input**

```
3 2
1 2 1
2 3 1
1 1 1 1
```

**Output**

```
3
```

### Example 2

**Input**

```
4 5
1 2 2
2 3 1
3 4 2
4 1 2
1 3 3
1 2 1 1
```

**Output**

```
19
```

### Example 3

**Input**

```
3 3
1 2 2
2 3 1
3 1 2
1 1 1 1
```

**Output**

```
4
```

## ideas
1. no ideas at all

## algorithm

1. Run Floyd-Warshall to get all-pairs shortest distances.
2. For every runner `i`, collect all distinct values `dist(i, j)` for `j != i`, sort them, and call this set `times[i]`.
   Different destination cities with the same shortest distance are equivalent for the final ranking, so duplicates can be removed.
3. The final order is determined by the pair `(chosen_time[i], i)`, because ties are broken by the start city index.
4. A medal distribution is completely determined by two boundaries in this order:
   `x` separates gold from non-gold, and `y` separates gold-or-silver from bronze, with `x < y`.
5. To avoid double-counting the same medal distribution many times, use canonical boundaries:
   `x` must be equal to the rank-pair of at least one gold runner's earliest achievable position, and `y` must be equal to the first achievable position after `x` of at least one silver runner.
6. For a fixed `x`, define `next[i]` as the smallest pair `(t, i)` such that `(t, i) > x` and `t in times[i]`.
   Then for a fixed pair `(x, y)`:
   - runner `i` can be gold iff `min(times[i]) <= x`
   - runner `i` can be silver iff `next[i]` exists and `next[i] <= y`
   - runner `i` can be bronze iff `max(times[i]) > y`
7. Enumerate all candidate `x` values from runners' minimum pairs, and for each `x` enumerate all candidate `y` values from existing `next[i]`.
8. For each `(x, y)`, run DP over runners:
   `dp[g][s][fx][fy]` = number of ways after processing some prefix, where:
   - `g` runners are gold
   - `s` runners are silver
   - `fx` says whether some chosen gold runner has canonical boundary exactly `x`
   - `fy` says whether some chosen silver runner has canonical boundary exactly `y`
9. Transition each runner to any allowed medal among gold / silver / bronze.
10. Sum all states with:
    - `g1 <= g <= g2`
    - `s1 <= s <= s2`
    - `fx = 1`
    - `fy = 1`

## intuition

The shortest-path part is only preprocessing. After that, runner `i` is completely described by the set of finish times he can achieve:

- `times[i] = {dist(i, j) | j != i}`

He does not care which destination city produced that time. For ranking, only the chosen time matters, and ties are broken by the starting city number `i`.

So the real problem is:

- each runner independently chooses one value from `times[i]`
- all runners are sorted by `(chosen_time, city_id)`
- we want to count how many different `Gold / Silver / Bronze` labelings can appear

That turns the graph problem into a counting problem on ordered pairs.

## why boundaries are enough

Once the final ranking is fixed, medals are assigned by two cut positions:

- everyone before the first cut gets gold
- everyone between the first and second cuts gets silver
- everyone after the second cut gets bronze

Because ranking uses the pair `(time, city)`, it is natural to represent each cut by a pair too.

Let:

- `x` = the last gold position
- `y` = the last silver position

Then `x < y`, and every runner must end up in one of these regions:

- `<= x` for gold
- `(x, y]` for silver
- `> y` for bronze

This is why the whole counting can be organized around enumerating boundary pairs `(x, y)`.

## why duplicates appear if we are not careful

Suppose a runner can be gold with time `3`, and another runner can also be gold with time `4`.  
If we only say "gold runners must be before some boundary", then the same medal assignment might be counted once with boundary `x = (4, ...)`, and again with some larger boundary that still keeps exactly the same runners in gold.

So we need one canonical choice of boundary for every medal assignment.

We choose:

- `x` to be the earliest pair that still contains all gold runners
- `y` to be the earliest pair after `x` that still contains all silver runners

In code this becomes:

- some gold runner must have his earliest possible pair exactly equal to `x`
- some silver runner must have his first possible pair after `x` exactly equal to `y`

Those are exactly the DP flags `fx` and `fy`.

## meaning of `lo`, `hi`, and `next`

For each runner `i`:

- `lo[i]` = the earliest pair he can ever occupy, namely `(min(times[i]), i)`
- `hi[i]` = the latest pair he can ever occupy, namely `(max(times[i]), i)`

For a fixed first boundary `x`, define:

- `next[i]` = the first pair of runner `i` that is strictly larger than `x`

Then:

- runner `i` can be gold iff `lo[i] <= x`
  because if his earliest achievable rank is already after `x`, he can never enter the gold segment
- runner `i` can be silver iff `next[i]` exists and `next[i] <= y`
  because silver means the runner must be placed after `x`, but still not after `y`
- runner `i` can be bronze iff `hi[i] > y`
  because if even his latest possible rank is not after `y`, he can never be bronze

This is the whole logic behind the three transition tests in the DP.

## DP state

For fixed `(x, y)`, process runners one by one.

The DP stores:

- how many gold runners were chosen
- how many silver runners were chosen
- whether we already used a gold runner that certifies the canonical boundary `x`
- whether we already used a silver runner that certifies the canonical boundary `y`

So:

- `dp[g][s][fx][fy]`

means:

- among processed runners, exactly `g` are gold and `s` are silver
- `fx = 1` means at least one chosen gold runner has `lo[i] = x`
- `fy = 1` means at least one chosen silver runner has `next[i] = y`

Each new runner has up to three choices:

- put him in gold if gold is allowed
- put him in silver if silver is allowed
- otherwise or additionally put him in bronze if bronze is allowed

At the end, we only accept states with:

- `g1 <= g <= g2`
- `s1 <= s <= s2`
- `fx = 1`
- `fy = 1`

The two flags guarantee that every medal assignment is counted exactly once.

## why this is fast enough

- Floyd-Warshall is `O(n^3)`, which is fine for `n <= 50`
- there are at most `n` different candidates for `x`
- for each `x`, there are at most `n` different candidates for `y`
- for each pair `(x, y)`, the DP is roughly `O(n * g2 * s2)`

Since all limits are at most `50`, this comfortably fits.

## complexity

- Floyd-Warshall: `O(n^3)`
- Boundary enumeration and DP: `O(n^5)` in the worst case, which is fine for `n <= 50`
- Memory: `O(g2 * s2)`
