# Problem F: Rain and Umbrellas

## Problem Description

Polycarp lives on a coordinate line at the point $x = 0$. He goes to his friend that lives at the point $x = a$. Polycarp can move only from left to right, he can pass one unit of length each second.

Now it's raining, so some segments of his way are in the rain. Formally, it's raining on $n$ non-intersecting segments, the $i$-th segment which is in the rain is represented as $[l_i, r_i]$ ($0 \leq l_i < r_i \leq a$).

There are $m$ umbrellas lying on the line, the $i$-th umbrella is located at point $x_i$ ($0 \leq x_i \leq a$) and has weight $p_i$. When Polycarp begins his journey, he doesn't have any umbrellas.

During his journey from $x = 0$ to $x = a$ Polycarp can pick up and throw away umbrellas. Polycarp picks up and throws down any umbrella instantly. He can carry any number of umbrellas at any moment of time. Because Polycarp doesn't want to get wet, he must carry at least one umbrella while he moves from $x$ to $x + 1$ if a segment $[x, x + 1]$ is in the rain (i.e., if there exists some $i$ such that $l_i \leq x$ and $x + 1 \leq r_i$).

The condition above is the only requirement. For example, it is possible to go without any umbrellas to a point where some rain segment starts, pick up an umbrella at this point and move along with an umbrella. Polycarp can swap umbrellas while he is in the rain.

Each unit of length passed increases Polycarp's fatigue by the sum of the weights of umbrellas he carries while moving.

**Question:** Can Polycarp make his way from point $x = 0$ to point $x = a$? If yes, find the minimum total fatigue after reaching $x = a$, if Polycarp picks up and throws away umbrellas optimally.

## Input

The first line contains three integers $a$, $n$ and $m$ ($1 \leq a, m \leq 2000$, $1 \leq n \leq \lceil \frac{a}{2} \rceil$) — the point at which Polycarp's friend lives, the number of the segments in the rain and the number of umbrellas.

Each of the next $n$ lines contains two integers $l_i$ and $r_i$ ($0 \leq l_i < r_i \leq a$) — the borders of the $i$-th segment under rain. It is guaranteed that there is no pair of intersecting segments. In other words, for each pair of segments $i$ and $j$ either $r_i < l_j$ or $r_j < l_i$.

Each of the next $m$ lines contains two integers $x_i$ and $p_i$ ($0 \leq x_i \leq a$, $1 \leq p_i \leq 10^5$) — the location and the weight of the $i$-th umbrella.

## Output

Print `-1` (without quotes) if Polycarp can't make his way from point $x = 0$ to point $x = a$. Otherwise print one integer — the minimum total fatigue after reaching $x = a$, if Polycarp picks up and throws away umbrellas optimally.

## Examples

### Example 1

**Input:**
```
10 2 4
3 7
8 10
0 10
3 4
8 1
1 2
```

**Output:**
```
14
```

**Explanation:** The only possible strategy is to take the fourth umbrella at the point $x = 1$, keep it till the point $x = 7$ (the total fatigue at $x = 7$ will be equal to $12$), throw it away, move on from $x = 7$ to $x = 8$ without an umbrella, take the third umbrella at $x = 8$ and keep it till the end (the total fatigue at $x = 10$ will be equal to $14$).

### Example 2

**Input:**
```
10 1 1
0 9
0 5
```

**Output:**
```
45
```

**Explanation:** The only possible strategy is to take the first umbrella, move with it till the point $x = 9$, throw it away and proceed without an umbrella till the end.

### Example 3

**Input:**
```
10 1 1
0 9
1 5
```

**Output:**
```
-1
```

**Explanation:** Polycarp cannot reach his destination because there's no umbrella available at the starting point $x = 0$, and the rain segment starts immediately at $x = 0$.


### ideas
1. 在每个下雨的时候，都需要持有一把伞；如果现在持有的伞的重量是w1,遇到了一个更轻的，肯定要马上换掉
2. 如果现在不下雨，把伞舍弃掉的前提是，在下次下雨时，能够获得一把伞。好像问题出在这里
3. 假设当前的伞的重量比较轻，后面虽然能获得一把伞，但是重的多；那么持有这把伞有可能是更优的选择
4. dp[i]表示使用第i把伞后能得到的最优解
5. dp[i] = min(dp[j] + (u[j] - pos[i]) * w[i]) 
6. u[j] 表示在j的左边结束的下雨的位置r[?]
7. 