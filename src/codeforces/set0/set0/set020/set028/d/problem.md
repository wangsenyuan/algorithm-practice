# D. Don't fear, DravDe is near!

[Problem link](https://codeforces.com/problemset/problem/28/D)

**Contest:** [Codeforces Beta Round #28 (Technocup 2012 — Final Round, 4, Div. 2)](https://codeforces.com/contest/28)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

A motorcade of `n` trucks, driving from city «Z» to city «З», has approached a tunnel known as the
Tunnel of Horror. Each truck is described by four numbers:

- `v` — value of the truck, its passengers, and cargo;
- `c` — number of passengers on the truck, the driver included;
- `l` — total number of people that must enter the tunnel before this truck so the driver can
  overcome his fear;
- `r` — total number of people that must follow this truck so the driver can overcome his fear.

The order of trucks cannot be changed, but any truck may be removed and left near the tunnel
indefinitely. Remove some trucks so that the remaining motorcade can enter the tunnel and the total
value of the removed trucks is maximized.

Equivalently, keep a subsequence of trucks (in original order) such that for every kept truck `i`,
the number of passengers among kept trucks before it is at least `l_i`, and the number of passengers
among kept trucks after it is at least `r_i`.

## Input

The first line contains an integer `n` (`1 <= n <= 10^5`) — the number of trucks.

Each of the next `n` lines contains four integers `v_i`, `c_i`, `l_i`, and `r_i`
(`1 <= v_i <= 10^4`, `1 <= c_i <= 10^5`, `0 <= l_i, r_i <= 10^5`) — the description of the
`i`-th truck from the front of the motorcade.

## Output

In the first line, print `k` — the number of trucks that will drive into the tunnel.

In the second line, print `k` integers — the indices of these trucks in ascending order. The order of
trucks must not be changed. If there are multiple optimal answers, print any.

## Example

### Input

```text
5
1 1 0 3
1 1 1 2
1 1 2 1
1 1 3 0
2 1 3 0
```

### Output

```text
4
1 2 3 5
```

### Input

```text
5
1 1 0 3
10 1 2 1
2 2 1 1
10 1 1 2
3 1 3 0
```

### Output

```text
3
1 3 5
```

### ideas
1. 让一些trunk可以*安心*的通过隧道, 使得value最大
2. 通过的这些需要满足条件, F[i] >= l[i] (在i前面通过的总人数)
3. E[i] >= r[i] (在i后面通过的总人数)
4. 从i的角度看, 它前面的车最好都通过, 因为这样,F[i]才会更大, 同样的它后面的车也应该尽量的通过
5. 但是,后面的车的情况, 是会影响到前面车的情况的(因为假设i, 因为F[i] < l[i], 它没有进入, 那么反过来会影响它前面的车的选择)
6. 假设没有l[i] = 0的车 => 0, 
7. 先按照l[i]的选择, 不断的添加车进来(当某个车满足l[i]的情况时,就添加进来)
8. 这个是理论上的上限(因为,没有被添加进来的, 都是不满足l[i]条件的)
9. 但是,它们不一定满足r[i]的要求.
10. 没法~

## Solution

The useful observation is that a selected truck fixes the total number of people in the final
motorcade.

For a selected truck `i`:

- there must be `l_i` selected people before it;
- the truck itself contributes `c_i` people;
- there must be `r_i` selected people after it.

So the final total number of selected people must be:

```text
T = l_i + c_i + r_i
```

Therefore, trucks with different values of `l_i + c_i + r_i` cannot appear in the same answer. The
solution groups trucks by this value `T`, solves each group independently, and keeps the best result
among all groups.

Inside one fixed group, process trucks in the original order, because trucks cannot be swapped. Let:

```text
dp[x] = maximum total value after choosing some already processed trucks,
        with exactly x selected people already entering the tunnel
```

Initially:

```text
dp[0] = 0
```

For one truck `(v, c, l, r)` in this group, if we choose it, then exactly `l` selected people must be
before it. After this truck enters, the selected people count becomes `l + c`. So the truck is a
transition:

```text
l -> l + c
```

and the DP update is:

```text
dp[l + c] = max(dp[l + c], dp[l] + v)
```

This is the important direction. We do not use `dp[l+c]` as the source. The source is `dp[l]`,
because `l` is the number of selected people before the current truck.

For example, in the first sample all trucks have `T = 4`:

```text
truck 1: 0 -> 1, value 1
truck 2: 1 -> 2, value 1
truck 3: 2 -> 3, value 1
truck 4: 3 -> 4, value 1
truck 5: 3 -> 4, value 2
```

The best path from `0` to `4` is:

```text
0 --1--> 1 --2--> 2 --3--> 3 --5--> 4
```

so the answer is `1 2 3 5`.

To reconstruct the selected trucks, when a transition improves `dp[to]`, we record the path that
created it. A simple pair like `prevPos[to]` and `prevTruck[to]` is not enough, because a later truck
can overwrite the parent of an earlier state. That may reconstruct a sequence such as `3 2`, which
looks as if trucks were swapped.

The implementation avoids that by storing immutable path nodes:

```text
Node{truck, previousNode}
```

When `dp[to]` improves through the current truck, it creates a new node:

```text
new node = current truck + the old chain of dp[from]
```

Then `fromNode[to]` points to this new node. Since the previous chain is stored by node id, later DP
updates cannot rewrite history. Reconstructing from `fromNode[T]` gives the chosen trucks in reverse,
so the final step reverses the list.

The arrays `dp` and `fromNode` are reused for all groups. A `touched` list records which passenger
counts were modified in the current group, so those positions can be reset efficiently before moving
to the next group.

Complexity:

```text
O(n + max(l_i + c_i + r_i))
```

Every truck is processed once, and the DP arrays have size equal to the maximum possible total
passenger count.
