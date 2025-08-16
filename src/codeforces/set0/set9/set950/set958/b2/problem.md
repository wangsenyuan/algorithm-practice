# Problem B2: The Resistance

## Problem Description

The Resistance is trying to take control over as many planets of a particular solar system as possible. Princess Heidi is in charge of the fleet, and she must send ships to some planets in order to maximize the number of controlled planets.

The Galaxy contains **N** planets, connected by bidirectional hyperspace tunnels in such a way that there is a unique path between every pair of the planets.

A planet is controlled by the Resistance if:
- There is a Resistance ship in its orbit, **OR**
- The planet lies on the shortest path between some two planets that have Resistance ships in their orbits.

Heidi has not yet made up her mind as to how many ships to use. Therefore, she is asking you to compute, for every **K = 1, 2, 3, ..., N**, the maximum number of planets that can be controlled with a fleet consisting of **K** ships.

## Input

- The first line contains an integer **N** (1 ≤ N ≤ 10^5) – the number of planets in the galaxy.
- The next **N-1** lines describe the hyperspace tunnels between the planets.
- Each line contains two space-separated integers **u** and **v** (1 ≤ u, v ≤ N) indicating a bidirectional hyperspace tunnel between planets **u** and **v**.
- It is guaranteed that every two planets are connected by a path of tunnels, and each tunnel connects a different pair of planets.

## Output

Print **N** space-separated integers on a single line. The **K**-th number should correspond to the maximum number of planets that can be controlled by the Resistance using a fleet of **K** ships.

## Examples

### Example 1

**Input:**
```
3
1 2
2 3
```

**Output:**
```
1 3 3
```

### Example 2

**Input:**
```
4
1 2
3 2
4 2
```

**Output:**
```
1 3 4 4
```

## Note

Consider the first example:
- If **K = 1**, then Heidi can only send one ship to some planet and control it.
- For **K ≥ 2**, sending ships to planets 1 and 3 will allow the Resistance to control all planets.


### ideas
1. 只需要在叶子节点上，部署飞船；
2. 每次部署，是不是应该在前面的基础上进行呢？
3. 