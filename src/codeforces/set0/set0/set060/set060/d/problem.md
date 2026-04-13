# D. Mushrooms

## Story

Misha wants to help Pasha and Akim become friends again. His plan is to make all the **laughing mushrooms** burst — they burst easily when they laugh. The mushrooms grow on lawns; there are `a[t]` mushrooms on the `t`-th lawn.

## Mechanics

Each lawn has a special property. Lawn `i` can **transmit laughter** to lawn `j` (`i ≠ j`) if there exists an integer `b` such that **some permutation** of the three numbers `a[i]`, `a[j]`, and `b` forms a **beautiful triple**.

A **beautiful triple** is three **pairwise coprime** positive integers `x`, `y`, `z` that satisfy `x^2 + y^2 = z^2`.

Misha needs the smallest number of lawns on which he must laugh so that **every** mushroom bursts (i.e. minimal number of “starting” lawns under the transmission rule above).

## Input

- The first line contains an integer `n` (`1 ≤ n ≤ 10^6`) — the number of lawns.
- The second line contains `n` distinct integers `a_1, …, a_n` (`1 ≤ a_i ≤ 10^7`) — the number of mushrooms on each lawn.

## Output

Print one integer — the minimum number of lawns on which Misha should laugh so that all mushrooms burst.

## Examples

### Example 1

**Input**

```
1
2
```

**Output**

```
1
```

### Example 2

**Input**

```
2
1 2
```

**Output**

```
2
```

### Example 3

**Input**

```
2
3 5
```

**Output**

```
1
```

---

## 中文题面

### 题目大意

Misha 想让 Pasha 和 Akim 和好，打算通过让所有**爱笑蘑菇**爆掉来实现——蘑菇一笑就容易爆。蘑菇长在草坪上，第 `t` 块草坪上有 `a[t]` 朵蘑菇。

### 规则

草坪有特殊能力：若存在整数 `b`，使得 `a[i]`、`a[j]`、`b` 的**某个排列**构成**优美三元组**，则草坪 `i` 可以把「笑」传给草坪 `j`（`i ≠ j`）。

**优美三元组**：三个**两两互质**的正整数 `x, y, z`，满足 `x² + y² = z²`。

问：Misha 最少需要在多少块草坪上「笑」，才能让**所有**蘑菇都爆掉。

### 输入输出

- **输入**：第一行 `n`（`1 ≤ n ≤ 10^6`），第二行 `n` 个**互不相同**的整数 `a_i`（`1 ≤ a_i ≤ 10^7`）。
- **输出**：一个整数，即所求最少草坪数。

（样例与上文英文部分相同。）

---

## Hints

### Hint 1 — Graph / DSU

Treat each lawn as a vertex. Put an undirected edge between `i` and `j` when the rule allows laughter to pass between them (i.e. when their `a`-values can be two numbers of some beautiful triple together with a suitable third integer `b`).

If laughing once on a lawn eventually reaches every lawn in its connected component, then the answer is the number of **connected components**. A **Disjoint Set Union (DSU)** is the usual way to build that graph implicitly.

### Hint 2 — What “edge” means algebraically

Fix distinct lawn values `u` and `v`. You need an integer `b > 0` such that `{u, v, b}` is a **beautiful triple** in some order. So the three numbers must be two legs and the hypotenuse of a right triangle: one of them equals `sqrt` of the sum of squares of the other two, and all three must be **pairwise coprime**.

For fixed `u, v` there are only three cases: `u² + v² = ?²`, `u² + ?² = v²` (needs `v > u`), `v² + ?² = u²` (needs `u > v`). Each gives at most one candidate for `b`; then check **gcd** conditions. You cannot afford to do this for every pair (`n` is up to `10⁶`).

### Hint 3 — Why not all pairs

Checking every lawn pair is **O(n²)** — too slow. You need a **global** construction: generate all *candidate* beautiful triples up to a value bound implied by the constraints, and only connect lawns whose values appear in the same triple.

### Hint 4 — Enumerate primitive triples (expanded)

**Why “primitive” is enough.** A beautiful triple is three **pairwise coprime** positive integers `x, y, z` with `x² + y² = z²`. Such a triple cannot have a common prime factor (that would break pairwise coprimality), so it is exactly a **primitive Pythagorean triple**: the unordered triple `{x,y,z}` is `{leg, leg, hypotenuse}` of a primitive right triangle.

So every valid edge between two lawns must come from **some** primitive triple `(p, q, r)` with `p² + q² = r²` (with `r` the hypotenuse), where **two** of `{p,q,r}` equal the two lawn values (and the third is the `b` from the statement — it does **not** need to appear among the inputs).

**Bounding the coordinates.** All given `a_i` lie in `[1, 10⁷]`. If the two lawn numbers are both legs, the hypotenuse is `√(u²+v²) ≤ √2·10⁷ < 1.5·10⁷`. If one lawn number is the hypotenuse `c ≤ 10⁷` and the other is a leg `a`, then the missing leg is `√(c²−a²) < c ≤ 10⁷`. So every number that can appear in any beautiful triple **together with two inputs** is **O(10⁷)** (with a modest constant; a safe global upper bound for generation is on the order of `1.5·10⁷` to `2·10⁷`, depending on how you implement the loop).

**Euclid’s formula (generating primitives).** All primitive triples (up to swapping the two legs) come from integers `m > n ≥ 1` with `gcd(m,n)=1` and **not both odd**:

- `leg₁ = m² − n²`
- `leg₂ = 2mn`
- `hyp  = m² + n²`

Then `leg₁² + leg₂² = hyp²`, and `(leg₁, leg₂, hyp)` is primitive. Iterate `(m, n)` so that `m² + n²` (and hence all three) stay below your chosen limit.

**Connecting lawns.** Build a map `value → lawn index` for all inputs (values are **distinct**, so at most one index per value). For each generated primitive triple `(x, y, z)` with `z` the hypotenuse:

- If both `x` and `y` appear in the map, **union** those two indices.
- If both `x` and `z` appear, union those indices.
- If both `y` and `z` appear, union those indices.

Any edge that the problem allows will show up as **some** such triple (because the beautiful triple for that edge is primitive and fits inside the bound above).

**Complexity intuition.** The number of primitive triples with hypotenuse ≤ `M` grows **Θ(M)**. So the loop is linear in your generation limit, not in `n`. With DSU (nearly inverse-Ackermann per union), the whole approach is feasible if the constant factor and limit are tuned for the language / time limit.

### Hint 5 — Implementation detail

Use a fast associative map from mushroom count to lawn index (e.g. hash map). Precompute nothing per pair of lawns; only per generated triple, do O(1) lookups and up to three `union` operations.

---

### 提示（中文概要）

1. 把草坪当点，能传笑就连边；答案即**连通块个数**，常用 **DSU**。  
2. 两点有边 ⟺ 两处的蘑菇数能作为某个**本原勾股数**三元组中的两个数（第三个数为 `b`，不必在输入里）。  
3. 不能枚举点对；应**枚举本原勾股三元组**，用 **欧几里得公式** 生成，再用 `值 → 下标` 对出现在输入中的两个数 `union`。  
4. 由 `a_i ≤ 10⁷` 可推出参与三元组的各数上界约为 **`O(10⁷)`**（常数略大于 1），据此设定生成循环上界。

---

## Solution notes — thoughts

- **Graph model.** Lawns are vertices; connect two lawns if their mushroom counts can appear as **two** of the three numbers in some beautiful triple (the third is `b`, not required to be in the input). Laughter spreads along edges, so the answer is the number of **connected components**; maintain them with **DSU**.
- **Number theory.** “Beautiful” here is exactly a **primitive Pythagorean triple**: two legs and a hypotenuse, pairwise coprime, `leg₁² + leg₂² = hyp²`. So every edge is witnessed by **some** primitive triple generated in a bounded range.
- **Why not all pairs.** `n ≤ 10⁶` forbids `O(n²)` checks. Enumerate primitive triples with **Euclid’s formula** (`m > n`, `gcd(m,n)=1`, not both odd) up to a hypotenuse bound derived from `a_i ≤ 10⁷`.
- **Two equivalent DSU builds (same graph).**
  1. **Adjacency-style (Go-style):** precompute lists such as “larger leg → smaller legs” and “hypotenuse → legs”, then a **sorted** pass over `a` with **`lastPos[value]`** so each leg–leg edge is united once when you see the larger leg.
  2. **Position array (C++ / low-memory):** one array **`pos[value] → lawn index`** (or 0). For **each** generated triple, if two sides both appear in the input, **union** those two indices immediately. No giant per-value slice/vector buckets.

## Solution notes — fixes (pitfalls we hit)

- **Euclid loop must enforce `m > n`.** If `n ≥ m`, the leg `m² − n²` is non-positive; bad triples can pollute structures or make **`lastPos[invalid]`** / wrong unions. Filter with **`gcd(m,n)=1`** and **opposite parity** early to skip non-primitive `(m,n)` pairs.
- **Hypotenuse bound vs input bound.** Two legs can both be `≤ 10⁷` while **`hyp = √(leg₁² + leg₂²)` exceeds `10⁷`** (up to about **`⌈√2 · 10⁷⌉`**). If generation (or discarding triples) caps everything at **`10⁷`**, you **drop valid leg–leg triples** and may output **too many** components. **Fix:** generate triples with a larger **`maxHyp`** (e.g. **`1.5 · 10⁷`** or **`⌈√2 · 10⁷⌉`**). Only treat **`hyp` as a lawn value** when **`hyp ≤ 10⁷`** (inputs never exceed `10⁷`); leg–leg edges still come from **`case1` / larger-leg logic** or from **`unite(leg₁, leg₂)`** in the on-the-fly variant.
- **Memory (Go).** Storing **`case1` / `case2` as ~10⁷ slices** (each with a slice header plus backing arrays for all listed partners) can **MLE**. The **`pos[]` + enumerate-and-union** pattern matches the graph with about **`O(maxA)`** integer storage plus DSU — the **C++** solution in `cpp/solution.cpp` followed this and **got AC**.
- **I/O.** Contest `main` should read stdin and print the component count (repo template: `drive` + `Println` in Go).

### 解题小结（中文）

- **思路：** 传笑关系是无向图，边由「两数同在一个本原勾股三元组中」定义；答案为连通块个数，用 DSU。不能枚举点对，用欧几里得公式枚举本原三元组，再按值把草坪合并。
- **易错：** 内层必须 **`m > n`**；斜边可 **大于 `10⁷`**（两直角边都在输入范围内时），生成上界要用 **`maxHyp`**，不能与输入上界混为一谈；`hyp` 作为草坪上的数时才需 `hyp ≤ 10⁷`。
- **实现取舍：** 预计算邻接 + `lastPos` 扫一遍（省 union 次数） vs **`pos` + 枚举时 union**（省内存）；后者在 CF 上更易 **AC**。

---

### ideas
1. 假设（x, y) 和某个b组成了一个优美三元组
2. (y, z) 和某个b组成了3元组
3. 那么(x, z) 满足吗？
4. x * x + y * y = b1 * b1
5. y * y + z * z = b2 * b2
6. z * z - x * x = (b2 - b1) * (b2 + b1)
7. b3 * b3 = (b2 - b1) * (b2 + b1) ?
8. 似乎对b1, b2是有要求的
9. 比如 2 * 8 = 16 => b2 - b1 = 2, b2 + b1 = 8 => b2 = 5, b1 = 3
10. 