# Problem E: Bird Summoning

## Problem Description

To summon birds, Imp needs strong magic. There are $n$ trees in a row on an alley in a park, and there is a nest on each of the trees. In the $i$-th nest there are $c_i$ birds; to summon one bird from this nest, Imp needs to stay under this tree and it costs him $cost_i$ points of mana. However, for each bird summoned, Imp increases his mana capacity by $B$ points. Imp summons birds one by one, and he can summon any number from 0 to $c_i$ birds from the $i$-th nest.

Initially, Imp stands under the first tree and has $W$ points of mana, and his mana capacity equals $W$ as well. He can only go forward, and each time he moves from a tree to the next one, he restores $X$ points of mana (but it can't exceed his current mana capacity). Moving only forward, what is the maximum number of birds Imp can summon?

## Input

The first line contains four integers $n$, $W$, $B$, $X$ ($1 \leq n \leq 10^3$, $0 \leq W, B, X \leq 10^9$) — the number of trees, the initial points of mana, the number of points the mana capacity increases after a bird is summoned, and the number of points restored when Imp moves from a tree to the next one.

The second line contains $n$ integers $c_1, c_2, \ldots, c_n$ ($0 \leq c_i \leq 10^4$) — where $c_i$ is the number of birds living in the $i$-th nest. It is guaranteed that $\sum_{i=1}^{n} c_i \leq 10^4$.

The third line contains $n$ integers $cost_1, cost_2, \ldots, cost_n$ ($0 \leq cost_i \leq 10^9$), where $cost_i$ is the mana cost to summon a bird from the $i$-th nest.

## Output

Print a single integer — the maximum number of birds Imp can summon.

## Examples

### Example 1
**Input:**
```
2 12 0 4
3 4
4 2
```

**Output:**
```
6
```

**Explanation:** In the first sample, the base amount of Imp's mana is equal to 12 (with maximum capacity also equal to 12). After he summons two birds from the first nest, he loses 8 mana points, although his maximum capacity will not increase (since $B = 0$). After this step, his mana will be 4 of 12; during the move, he will replenish 4 mana points, and hence own 8 mana out of 12 possible. Now it's optimal to take 4 birds from the second nest and spend 8 mana. The final answer will be 6.

### Example 2
**Input:**
```
4 1000 10 35
1 2 4 5
1000 500 250 200
```

**Output:**
```
5
```

**Explanation:** In the second sample, the base amount of mana is equal to 1000. The right choice will be to simply pick all birds from the last nest. Note that Imp's mana doesn't restore while moving because it's initially full.

### Example 3
**Input:**
```
2 10 7 11
2 10
6 1
```

**Output:**
```
11
```


### ideas
1. 假设当前有w的能力,召唤每只鸟需要花费cost[i]，如果移动到后一个位置，可以增加X
2. 那么至多召唤a = min(w, X) / cost[i]只鸟（当然不能超过c[i])；
3. 