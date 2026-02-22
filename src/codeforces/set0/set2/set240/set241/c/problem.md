# Problem C

Mirror Box is a name of a popular game in the Iranian National Amusement Park (INAP). There is a wooden box, 105 cm long and 100 cm high in this game. Some parts of the box's ceiling and floor are covered by mirrors. There are two negligibly small holes in the opposite sides of the box at heights hₗ and hᵣ centimeters above the floor. The picture below shows what the box looks like.

In the game, you will be given a laser gun to shoot once. The laser beam must enter from one hole and exit from the other one. Each mirror has a preset number vᵢ, which shows the number of points players gain if their laser beam hits that mirror. Also — to make things even funnier — the beam must not hit any mirror more than once.

Given the information about the box, your task is to find the maximum score a player may gain. Please note that the reflection obeys the law "the angle of incidence equals the angle of reflection".

## Input

The first line of the input contains three space-separated integers hₗ, hᵣ, n (0 < hₗ, hᵣ < 100, 0 ≤ n ≤ 100) — the heights of the holes and the number of the mirrors.

Next n lines contain the descriptions of the mirrors. The i-th line contains space-separated vᵢ, cᵢ, aᵢ, bᵢ; the integer vᵢ (1 ≤ vᵢ ≤ 1000) is the score for the i-th mirror; the character cᵢ denotes i-th mirror's position — the mirror is on the ceiling if cᵢ equals "T" and on the floor if cᵢ equals "F"; integers aᵢ and bᵢ (0 ≤ aᵢ < bᵢ ≤ 10⁵) represent the x-coordinates of the beginning and the end of the mirror.

No two mirrors will share a common point. Consider that the x coordinate increases in the direction from left to right, so the border with the hole at height hₗ has the x coordinate equal to 0 and the border with the hole at height hᵣ has the x coordinate equal to 105.

## Output

The only line of output should contain a single integer — the maximum possible score a player could gain.

## Examples

**Input:**
```
50 50 7
10 F 1 80000
20 T 1 80000
30 T 81000 82000
40 T 83000 84000
50 T 85000 86000
60 T 87000 88000
70 F 81000 89000
```

**Output:**
```
100
```

**Input:**
```
80 72 9
15 T 8210 15679
10 F 11940 22399
50 T 30600 44789
50 F 32090 36579
5 F 45520 48519
120 F 49250 55229
8 F 59700 80609
35 T 61940 64939
2 T 92540 97769
```

**Output:**
```
120
```

## Note

The second sample is depicted above. The red beam gets 10 + 50 + 5 + 35 + 8 + 2 = 110 points and the blue one gets 120.

The red beam on the picture given in the statement shows how the laser beam can go approximately; this is just an illustration of how the laser beam can gain score. So for the second sample there is no such beam that gains score 110.

## Algorithm — Unfolded Box

### Core observation

When a laser beam reflects off a horizontal surface (ceiling or floor), the angle of incidence equals the angle of reflection. This means the absolute slope |dy/dx| is preserved across every reflection. Equivalently, the x-distance between consecutive reflections on opposite walls is constant: dx = 100/s, where s is the slope magnitude.

### The unfolding trick

Instead of simulating reflections, we "unfold" the box vertically by stacking mirrored copies. In this infinite strip the beam becomes a **single straight line** from (0, hₗ) to (10⁵, Y), where Y is an "unfolded" y-coordinate of the exit hole.

In the unfolded view, surfaces repeat with period 200:

- **Floor mirrors** (y = 0 in original) appear at y-levels y = 100j where j is **even** (j = …, −2, 0, 2, 4, …)
- **Ceiling mirrors** (y = 100 in original) appear at y-levels y = 100j where j is **odd** (j = …, −1, 1, 3, 5, …)

The exit at height hᵣ on the right wall folds correctly when:

- Y = 200k + hᵣ, or
- Y = 200k − hᵣ

for any integer k.

### Beam validity

Each y-level crossing of the straight line corresponds to one reflection in the original box. At crossing y = 100j, the beam's x-coordinate is:

x = (100j − hₗ) × 10⁵ / (Y − hₗ)

For the path to be physically valid:

1. **Every crossing must hit a mirror** — x must fall within some mirror's [aᵢ, bᵢ] range on the correct wall (floor if j even, ceiling if j odd). If any crossing misses, the beam is absorbed by bare wood.
2. **Each mirror hit at most once** — the same original mirror appearing at different y-levels counts as the same mirror. A repeated hit makes the path invalid.

### Bounding the search

Each crossing uses a distinct mirror, so the number of crossings ≤ n. This bounds |Y − hₗ| ≤ (n + 1) × 100. Since Y = 200k ± hᵣ, there are only about 2 × (n + 1) candidate Y values to check.

### Exact integer arithmetic

To avoid floating-point precision issues, the check "x ∈ [aᵢ, bᵢ]" is rewritten as a cross-multiplication:

- Let num = (100j − hₗ) × 10⁵ and den = Y − hₗ
- If den > 0: check aᵢ × den ≤ num ≤ bᵢ × den
- If den < 0: check bᵢ × den ≤ num ≤ aᵢ × den

All values fit in 64-bit integers (max ≈ 10⁹).

### Complexity

- ~200 candidate Y values
- ≤ 100 crossings per Y
- ≤ 100 mirrors to check per crossing

Total: O(n³) ≈ 2 × 10⁶ operations. Very fast.

### Example 1 walkthrough

hₗ = 50, hᵣ = 50. Try Y = −250 (base = −50, k = −1):

- den = −250 − 50 = −300
- j = 0 (floor): x = 50 × 10⁵ / 300 = 16667 → hits floor mirror [1, 80000], score += 10
- j = −1 (ceiling): x = 150 × 10⁵ / 300 = 50000 → hits ceiling mirror [1, 80000], score += 20
- j = −2 (floor): x = 250 × 10⁵ / 300 = 83333 → hits floor mirror [81000, 89000], score += 70

Total = 100. All mirrors distinct. This is the maximum.
