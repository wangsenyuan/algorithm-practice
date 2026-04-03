# F. Tree Edge XOR Constraints

You are given a **tree** on \(n\) vertices numbered \(1\) to \(n\).

Assign an integer weight to each edge. Let the weight of the \(i\)-th edge be \(a_i\) (\(1 \le i \le n - 1\)). Every edge weight must satisfy \(0 \le a_i \le 2^{30} - 1\).

You are given **\(q\)** conditions. Each condition is three integers \(u\), \(v\), and \(x\): the **bitwise XOR** of all edge weights on the (unique) simple path from \(u\) to \(v\) must equal \(x\).

Determine whether there exist weights \(a_1, \ldots, a_{n-1}\) satisfying all conditions. If **yes**, output a valid assignment such that

\[
a_1 \oplus a_2 \oplus \cdots \oplus a_{n-1}
\]

is **minimized** (where \(\oplus\) is bitwise XOR). If several assignments achieve this minimum XOR, **any** of them may be printed.

## Input

The first line contains two integers \(n\) and \(q\) (\(2 \le n \le 2.5 \cdot 10^5\), \(0 \le q \le 2.5 \cdot 10^5\)).

The next \(n - 1\) lines describe the edges: the \(i\)-th line contains two integers \(x_i\) and \(y_i\) (\(1 \le x_i, y_i \le n\), \(x_i \ne y_i\)) — the endpoints of the \(i\)-th edge. These edges form a tree.

The next \(q\) lines describe the conditions. Each line has three integers \(u\), \(v\), and \(x\) (\(1 \le u, v \le n\), \(u \ne v\), \(0 \le x \le 2^{30} - 1\)): the XOR of edge weights on the path from \(u\) to \(v\) must be \(x\).

## Output

If **no** assignment exists, print `No`.

Otherwise print `Yes` on the first line, then one line with \(n - 1\) integers — the weight of the \(1\)-st edge through the \((n-1)\)-st edge in the **same order** as in the input, minimizing \(a_1 \oplus a_2 \oplus \cdots \oplus a_{n-1}\) among all feasible assignments. If several assignments tie on that minimum, print any.

**Case:** You may print `Yes` / `No` in any letter case (e.g. `yes`, `YES`, `yEs`).

## Examples

### Example 1

**Input**

```
4 4
1 2
2 3
3 4
1 4 3
2 4 2
1 3 1
2 3 1
```

**Output**

```
No
```

### Example 2

**Input**

```
6 2
1 2
2 3
3 4
2 5
5 6
1 4 2
2 6 7
```

**Output**

```
Yes
4 2 4 1 6
```

### Example 3

**Input**

```
6 2
1 2
2 3
3 4
2 5
5 6
1 4 3
1 6 5
```

**Output**

```
Yes
6 1 4 3 0
```

## Note

**Example 1:** No edge weights satisfy all constraints.

**Example 2:** The conditions are \(a_1 \oplus a_2 \oplus a_3 = 2\) and \(a_4 \oplus a_5 = 7\) (using edge indices as in the statement). Many solutions exist; one is \((a_1, a_2, a_3, a_4, a_5) = (1, 2, 1, 4, 3)\).

**Example 3:** The conditions are \(a_1 \oplus a_2 \oplus a_3 = 3\) and \(a_1 \oplus a_4 \oplus a_5 = 5\). The assignment \((1, 1, 3, 4, 0)\) satisfies the constraints, but the XOR of **all** edge weights is \(7\), which is **not** minimal among valid solutions, so it is not a required answer.


### ideas
1. 什么时候没有答案呢？考虑 1 -> 2 -> 3, 如果 (1, 2, 1), (2, 3, 1), (1, 3, 1) => -1
2. 