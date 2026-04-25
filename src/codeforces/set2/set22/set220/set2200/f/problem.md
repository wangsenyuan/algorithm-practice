# Problem

Bessie needs to produce as much energy as possible in her nuclear reactor. She has `n` different particles.

Each particle is defined by two integers `x` and `y`. The particle generates `x` units of energy but has a reactivity `y`, meaning that it can only exist together with at most `y` other particles in the reactor. Formally, if this particle is chosen to generate energy, then at most `y` particles (other than itself) can also be chosen to generate energy.

Bessie must choose a subset of particles that satisfy this constraint to generate energy. The amount of energy that she generates is equal to the sum of the energies of the particles in the subset.

There is a shop with `m` particles. Bessie can buy exactly one particle from the shop. For each particle in the shop, determine the maximum total energy that Bessie would be able to produce if she were to buy only that particle from the shop. Bessie is not required to use the particle that is purchased from the shop.

## Input

The first line contains a single integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.

The first line of each test case contains two integers `n` and `m` (`1 ≤ n, m ≤ 2 * 10^5`) — the number of particles that Bessie has and the number of particles in the shop, respectively.

The `i`-th of the next `n` lines contains two integers `x` and `y` (`1 ≤ x ≤ 10^9`, `0 ≤ y ≤ n`) — the energy and reactivity of Bessie's `i`-th particle.

The `j`-th of the next `m` lines contains two integers `x` and `y` (`1 ≤ x ≤ 10^9`, `0 ≤ y ≤ n`) — the energy and reactivity of the shop's `j`-th particle.

It is guaranteed that the sum of `n` over all test cases and the sum of `m` over all test cases do not exceed `2 * 10^5`.

## Output

For each test case, print `m` integers. The `i`-th integer should be the maximum total energy that Bessie can produce if she purchases the `i`-th particle from the shop.

## Example

**Input**

```text
3
3 3
67 0
6 1
7 1
1 0
100 0
62 1
2 1
2 2
4 2
3 1
1 2
6 1
7 0
8 1
```

**Output**

```text
67 100 69
7
7 14
```

## Note

In the first test case:

- If Bessie buys particle `4`, then it is optimal for her to only use particle `1` and generate `67` energy units.
- If she buys particle `5`, then she should only use particle `5`, which generates `100` energy units.
- If she buys particle `6`, then she should use particles `3` and `6`, for a total of `69` energy units.

## key insights

1. For a fixed chosen subset size `k`, every selected particle must satisfy `y >= k - 1`. So among particles with `y >= k - 1`, the best valid subset is simply the top-`k` values of `x`.

2. The code processes original particles offline by decreasing `y`:
   - `ys[y]` stores all energies `x` whose reactivity is exactly `y`.
   - While iterating `y = n..0`, it inserts all `ys[y]` into a structure that represents all particles with reactivity at least current `y`.

3. To get "sum of largest `k` energies" fast, `Tree` is a segment tree over compressed `x` values:
   - each node stores `cnt` (how many particles in this value range) and `val` (sum of energies),
   - `GetBest(k)` walks from large `x` to small `x` and returns the sum of the largest `k` inserted energies.

4. During the same sweep, the solution builds:
   - `best = max over y of GetBest(y + 1)` (best answer using only original particles),
   - `dp[y] = GetBest(y)` (best sum for selecting exactly `y` original particles from those with reactivity at least `y`).

5. For a shop particle `(x, y)`, if it is used, we can add at most `y` other particles, each needing reactivity at least `y`. So candidate value is:
   - `x + dp.Get(0, y + 1)` where `dp.Get(0, y + 1)` is the maximum of `dp[k]` for `0 <= k <= y`.

6. Final answer per shop particle:
   - `max(best, x + max_{0<=k<=y} dp[k])`,
   which also covers the case "buy but do not use it".

7. Complexity per test case:
   - `O((n + m) log n)` time (segment tree operations),
   - `O(n)` extra memory (plus compressed-value tree).
