# Monster Eating Problem

## Problem Description

There was an epidemic in Monstropolis and all monsters became sick. To recover, all monsters lined up in queue for an appointment to the only doctor in the city.

Soon, monsters became hungry and began to eat each other.

### Eating Rules

One monster can eat another monster if:
- Its weight is **strictly greater** than the weight of the monster being eaten
- They stand in the queue **next to each other** (neighbors)

**Important Notes:**
- Monsters eat each other instantly
- There are no monsters which are being eaten at the same moment
- After monster A eats monster B, the weight of monster A increases by the weight of the eaten monster B
- The length of the queue decreases by one
- All monsters after the eaten one step forward so that there are no empty places in the queue
- A monster can eat several monsters one after another

Initially there were $n$ monsters in the queue, the $i$-th of which had weight $a_i$.

### Example Scenario

If weights are `[1, 2, 2, 2, 1, 2]` (in order of queue, monsters are numbered from 1 to 6 from left to right) then:

- The first monster **can't** eat the second monster because $a_1 = 1$ is not greater than $a_2 = 2$
- The second monster **can't** eat the third monster because $a_2 = 2$ is not greater than $a_3 = 2$
- The second monster **can't** eat the fifth monster because they are not neighbors
- The second monster **can** eat the first monster, the queue will be transformed to `[3, 2, 2, 1, 2]`

### Problem Statement

After some time, someone said a good joke and all monsters recovered. At that moment there were $k$ ($k \leq n$) monsters in the queue, the $j$-th of which had weight $b_j$. Both sequences ($a$ and $b$) contain the weights of the monsters in the order from the first to the last.

You are required to provide one of the possible orders of eating monsters which led to the current queue, or to determine that this could not happen. Assume that the doctor didn't make any appointments while monsters were eating each other.

## Input

- **Line 1:** Single integer $n$ ($1 \leq n \leq 500$) — the number of monsters in the initial queue
- **Line 2:** $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^6$) — the initial weights of the monsters
- **Line 3:** Single integer $k$ ($1 \leq k \leq n$) — the number of monsters in the queue after the joke
- **Line 4:** $k$ integers $b_1, b_2, \ldots, b_k$ ($1 \leq b_j \leq 5 \cdot 10^8$) — the weights of the monsters after the joke

Monsters are listed in the order from the beginning of the queue to the end.

## Output

- **If no actions could lead to the final queue:** Print "NO" (without quotes) in the only line
- **Otherwise:** 
  - Print "YES" (without quotes) in the first line
  - In the next $n - k$ lines print actions in chronological order
  - Each line should contain: `x` — the index number of the monster in the current queue which eats, followed by a space, then:
    - `'L'` if the monster which stays the $x$-th in the queue eats the monster in front of him
    - `'R'` if the monster which stays the $x$-th in the queue eats the monster behind him
  - After each eating, the queue is enumerated again
  - If there are several answers, print any of them

## Examples

### Example 1

**Input:**
```
6
1 2 2 2 1 2
2
5 5
```

**Output:**
```
YES
2 L
1 R
4 L
3 L
```

**Explanation:**
Initially: `[1, 2, 2, 2, 1, 2]` → Final: `[5, 5]`

Sequence of eating:
1. Second monster eats the monster to the left (first monster) → `[3, 2, 2, 1, 2]`
2. First monster (was second before) eats the monster to the right (second monster) → `[5, 2, 1, 2]`
3. Fourth monster eats the monster to the left (third monster) → `[5, 2, 3]`
4. Third monster eats the monster to the left (second monster) → `[5, 5]`

### Example 2

**Input:**
```
5
1 2 3 4 5
1
15
```

**Output:**
```
YES
5 L
4 L
3 L
2 L
```

### Example 3

**Input:**
```
5
1 1 1 3 3
3
2 1 6
```

**Output:**
```
NO
```

**Note:** For each step, the output contains numbers of the monsters in their current order in the queue.


### ideas
1. b[i] >= a[i] 必须成立
2. 可以从头开始
3. b[0] = sum(a[:i])
4. b[1] = sum(a[i:])
5. 且在区间内，不能相等，否则无法开启