By the age of three Smart Beaver mastered all arithmetic operations and got this summer homework from the amazed teacher:

You are given a sequence of integers $a_1, a_2, \ldots, a_n$. Your task is to perform on it $m$ consecutive operations of the following type:

- For given numbers $x_i$ and $v_i$ assign value $v_i$ to element $a_{x_i}$.
- For given numbers $l_i$ and $r_i$ you've got to calculate sum $\sum_{j=l_i}^{r_i} f_{a_j}$, where $f_0 = f_1 = 1$ and for $i \geq 2$: $f_i = f_{i-1} + f_{i-2}$.
- For a group of three numbers $l_i, r_i, d_i$ you should increase value $a_x$ by $d_i$ for all $x$ ($l_i \leq x \leq r_i$).

Smart Beaver planned a tour around great Canadian lakes, so he asked you to help him solve the given problem.

---

## Input

The first line contains two integers $n$ and $m$ ($1 \leq n, m \leq 2 \cdot 10^5$) — the number of integers in the sequence and the number of operations, correspondingly.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($0 \leq a_i \leq 10^5$).

Then follow $m$ lines, each describes an operation. Each line starts with an integer $t_i$ ($1 \leq t_i \leq 3$) — the operation type:

- if $t_i = 1$, then next follow two integers $x_i$ $v_i$ ($1 \leq x_i \leq n$, $0 \leq v_i \leq 10^5$);
- if $t_i = 2$, then next follow two integers $l_i$ $r_i$ ($1 \leq l_i \leq r_i \leq n$);
- if $t_i = 3$, then next follow three integers $l_i$ $r_i$ $d_i$ ($1 \leq l_i \leq r_i \leq n$, $0 \leq d_i \leq 10^5$).

### Subtasks

- The input limits for scoring **30 points** are (subproblem E1):

  > It is guaranteed that $n$ does not exceed 100, $m$ does not exceed 10000 and there will be no queries of the 3-rd type.

- The input limits for scoring **70 points** are (subproblems E1+E2):

  > It is guaranteed that there will be queries of the 1-st and 2-nd type only.

- The input limits for scoring **100 points** are (subproblems E1+E2+E3):

  > No extra limitations.

---

## Output

For each query print the calculated sum modulo $1000000000$ ($10^9$).


## ideas

To solve this problem, one should use a **segment tree**.

Let's consider a line segment $[l, r]$ on this tree. For this purpose, introduce $S(x)$, where $x$ is an integer:

$$
S(x) = F_{0+x} \cdot A_l + F_{1+x} \cdot A_{l+1} + \ldots + F_{r-l+x} \cdot A_r
$$
where $F_i$ is the $i$-th Fibonacci number, and $A_i$ is the array element at position $i$.

One should note that $S(x)$ with $x = 0, 1, 2, \ldots$ are similar to Fibonacci numbers, i.e.,

$$
S(x) = S(x-1) + S(x-2) \quad \text{for } x \geq 2
$$

This means that for every line segment $[l, r]$, one should memorize $S(0)$ and $S(1)$.

To calculate $S(x)$ for $x \geq 2$, one can use the formula:

$$
S(x) = S(0) \cdot F_{x-2} + S(1) \cdot F_{x-1}
$$

So, solving a type-2 query is reduced to calculating at most a few values of $S(x)$ for different segment lines.

- The type-1 modification can be done by walking down the tree.
- For type-3 queries, one should use additional information in the tree and partial sums of Fibonacci numbers.