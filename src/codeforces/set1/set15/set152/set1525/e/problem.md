Monocarp is playing a game "Assimilation IV". In this game he manages a great empire: builds cities and conquers new lands.

Monocarp's empire has 𝑛 cities. In order to conquer new lands he plans to build one Monument in each city. The game is turn-based and, since Monocarp is still amateur, he builds exactly one Monument per turn.

Monocarp has 𝑚 points on the map he'd like to control using the constructed Monuments. For each point he knows the distance between it and each city. Monuments work as follows: when built in a city, a Monument controls all points at distance at most 1 from that city. The next turn it controls all points at distance at most 2, then at most 3, and so on. Monocarp will build 𝑛 Monuments in 𝑛 turns, and his empire will conquer all points that are controlled by at least one Monument.

Monocarp can't figure out any strategy, so each turn he will choose a city for a Monument uniformly at random among all remaining cities (those without Monuments). He wants to know how many points (among the 𝑚) he will conquer at the end of turn 𝑛. Help him calculate the expected number of conquered points.

## Input

The first line contains two integers 𝑛 and 𝑚 (1≤𝑛≤20, 1≤𝑚≤5⋅10⁴) — the number of cities and the number of points.

The next 𝑛 lines contain 𝑚 integers each: the 𝑗-th integer on the 𝑖-th line, 𝑑ᵢ,ⱼ (1≤𝑑ᵢ,ⱼ≤𝑛+1), is the distance from the 𝑖-th city to the 𝑗-th point.

## Output

It can be shown that the expected number of points Monocarp conquers at the end of turn 𝑛 can be written as an irreducible fraction 𝑥/𝑦. Print this fraction modulo 998244353, i.e. the value 𝑥⋅𝑦⁻¹ mod 998244353, where 𝑦⁻¹ is the modular inverse of 𝑦 (𝑦⋅𝑦⁻¹ ≡ 1 (mod 998244353)).

## Example

**Input**

```
3 5
1 4 4 3 4
1 4 1 4 2
1 4 4 4 3
```

**Output**

```
166374062
```

## Note

All 6 possible orders of building Monuments and the number of points controlled:

- **[1,2,3]:** First city controls points at distance ≤3 (points 1 and 4); second city controls at distance ≤2 (points 1, 3, 5); third at distance ≤1 (point 1). Total: 4 points.
- **[1,3,2]:** First city — points 1 and 4; second — 1 and 3; third — 1. Total: 3 points.
- **[2,1,3]:** First — point 1; second — 1, 3, 5; third — 1. Total: 3 points.
- **[2,3,1]:** First — point 1; second — 1, 3, 5; third — 1. Total: 3 points.
- **[3,1,2]:** First — point 1; second — 1 and 3; third — 1 and 5. Total: 3 points.
- **[3,2,1]:** First — point 1; second — 1, 3, 5; third — 1 and 5. Total: 3 points.

Expected value = (4+3+3+3+3+3)/6 = 19/6, so 19⋅6⁻¹ ≡ 19⋅166374059 ≡ 166374062 (mod 998244353).


### ideas
1. 考虑一个点x, 它离城市i的距离为d[i, x], 假设在第2天才在城市i建设了Monument
2. 那么当 d[i, x] <= n - 1时，这个点x会被占领
3. 反过来考虑它不被占领的机会？
4. 对于point x来说，把城市按照从远到近的顺序排列，看看它能否不被控制？
5. 距离越远的，越可以早点建设
6. 规定了顺序，就没法算概率了。 
7. 对于每个城市来说，它有一个范围 f[i] + dist[i, x] >= n + 1 才行
8. 所以 f[i] >= n + 1 - dist[i, x] (f[i]是这个城市最早可以开始的turn)
9. 按照f[i]降序处理，第一个可以开始的turn = 1 - (n - f[1] + 1) * (n - 1 - f[2] + 1) * .... s / F[n] 就是x被占领的概率