Petya sometimes has to water his field. To water the field, Petya needs a tank with exactly `V` ml of water.

Petya has `N` tanks. The `i`-th tank initially contains `ai` ml of water. The tanks are really large: any of them can contain any amount of water.

Also Petya has a scoop that can contain up to `K` ml of water (initially the scoop is empty). This scoop can be used to get some water from one tank and then pour it all into one tank (it is impossible to get water from multiple tanks without pouring it, or leave some water in the scoop when pouring it). When Petya tries to get some water from a tank, he gets `min(v, K)` water, where `v` is the current volume of water in the tank.

Is it possible to obtain a tank with exactly `V` ml of water using these operations? If it is possible, print a sequence of operations that allows this. If there are multiple ways, print any of them.

## Input

The first line contains `3` integers: `N` (`2 <= N <= 5000`), `K` (`1 <= K <= 5000`), and `V` (`0 <= V <= 10^9`) — the number of tanks, the maximum volume of water the scoop can contain, and the required amount of water in some tank, respectively.

The second line contains `N` integers `ai` (`0 <= ai <= 10^5`), where `ai` is the initial volume of water in the `i`-th tank.

## Output

If it is impossible to obtain a tank with exactly `V` ml of water, print `NO`.

Otherwise print `YES` in the first line, and beginning from the second line, print the sequence of operations in the following format:

Each line has to contain `3` numbers denoting a compressed operation: `cnt x y` (`1 <= cnt <= 10^9`, `1 <= x, y <= N`), where `x` is the index of the tank where we get water, `y` is the index of the tank where we pour water, and `cnt` is the number of times we transfer water from tank `x` to tank `y`.

The number of these lines must not exceed `N + 5`.

## Examples

### Example 1
**Input**
```
2 3 5
2 3
```
**Output**
```
YES
1 2 1
```

### Example 2
**Input**
```
2 3 4
2 3
```
**Output**
```
NO
```

### Example 3
**Input**
```
5 2 0
1 3 5 7 9
```
**Output**
```
YES
2 2 1
3 3 1
4 4 1
5 5 1
```

## Editorial

Eliminate the obvious corner case when we don't have enough water (the total amount of water in all tanks is less than `V`). We do not consider it below.

Fix some set of tanks `S`, and let `d` be the total amount of water in `S`. If `d ≡ V (mod K)` (i.e. `d` and `V` have the same remainder modulo `K`), then we can transfer all water from `S` to one tank `x`, transfer all water from the remaining tanks to another tank `y`, and then, using some number of operations, transfer the required amount of water from `x` to `y` (or from `y` to `x`). So we have a solution when there exists a set `S` such that `d ≡ V (mod K)`.

If we don't have such a set, the problem is impossible: we cannot obtain a tank with amount `d` such that `d ≡ V (mod K)` (and therefore we cannot obtain a tank with exactly `V`).

To find such a set `S`, we may use knapsack-style dynamic programming.
