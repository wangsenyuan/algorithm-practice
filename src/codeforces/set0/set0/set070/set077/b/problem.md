# Problem

**题意：** 参数 $p$ 在 $[0, a]$ 上均匀、独立地选取，$q$ 在 $[-b, b]$ 上均匀、独立地选取（均为实数）。若关于 $x$ 的方程
$$x^2 + \sqrt{p}\cdot x + q = 0$$
至少有一个实根，则视为一次成功命中。求命中概率。多组测试，每组给定 $a$, $b$，输出该概率，绝对或相对误差不超过 $10^{-6}$。

---

（以下为题目原文要点。）

Height parameter $p$ varies in the range $[0, a]$ and wind parameter $q$ in $[-b, b]$, each chosen uniformly and independently (real numbers). The anvil hits the target if and only if the equation

$$x^2 + \sqrt{p}\cdot x + q = 0$$

has **at least one real root** $x$. Determine the probability of a successful hit.

## Input

- First line: integer $t$ ($1 \le t \le 10\,000$) — number of test cases.
- Each of the next $t$ lines: two space-separated integers $a$ and $b$ ($0 \le a, b \le 10^6$).

**Note:** $a$ and $b$ may be zero. Pretests satisfy $0 < a < 10$, $0 \le b < 10$.

## Output

Print $t$ lines — the probability of a successful hit for each test case. The **absolute or relative** error of the answer must not exceed $10^{-6}$.

## Examples

### Example 1

**Input:**

```
2
4 2
1 2
```

**Output:**

```
0.6250000000
0.5312500000
```


## ideas. 
1. 根据一元二次方程有解的前提是 b * b - 4 * a * c >= 0 
2. p - 4 * q >= 0 (p >= 0) => p >= 4 * q
3. 就是在区间[-b...b], 在x轴和线y = 4 * x这条线的上方
4. 