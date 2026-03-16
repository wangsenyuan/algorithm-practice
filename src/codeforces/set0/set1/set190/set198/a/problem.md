### Problem

Qwerty the Ranger is running experiments on bacteria in a test tube.

In the **first experiment**:

- Initially, there is **1** bacterium in the tube.
- Every second:
  - Each bacterium divides into `k` bacteria.
  - Then an additional `b` bacteria appear due to abnormal effects.
- So if the tube has `x` bacteria at the **beginning** of a second, it has `k * x + b` bacteria at the **end** of that second.

It is known that after exactly `n` seconds there were **`z` bacteria** in the tube (and the first experiment stopped then).

In the **second experiment**:

- The tube is sterilized.
- Qwerty places `t` bacteria into it.
- The bacteria will grow according to the **same rule** as above: each second, `x -> k * x + b`.

Qwerty wants to know: **how many seconds** are needed in the second experiment to have **at least `z` bacteria** in the tube?

### Input

- One line with four integers: `k`, `b`, `n`, `t`  
  - `1 ≤ k, b, n, t ≤ 10^6`  
  - Growth rule parameters: `k`, `b`  
  - In the first experiment, after `n` seconds, the tube contained `z` bacteria  
  - In the second experiment, the initial number of bacteria is `t`

(The value `z` is implicit, determined from the first experiment’s rule and `n`.)

### Output

Print a single integer — the **minimum number of seconds** needed in the second experiment to obtain **at least `z` bacteria**.

### Examples

**Input**
```text
3 1 3 5
```

**Output**
```text
2
```

**Input**
```text
1 4 4 7
```

**Output**
```text
3
```

**Input**
```text
2 2 4 100
```

**Output**
```text
0
```

### Short summary

- First compute `z` by simulating the recurrence `x_{i+1} = k * x_i + b` starting from `x_0 = 1` for `n` steps.
- Then, starting from `t`, repeatedly apply the same recurrence until the value is at least `z`.
- The answer is the minimal number of steps needed; handle up to `10^6` steps safely without overflow.***


### ideas
1. 如果先计算出z（理论上可以），这个z会非常大
2. 假设目前有w个，那么下一时刻 = k * w + b
3. 所以，应该计算多久可以有t个，假设花了n1时间, 那么答案 = n - n1