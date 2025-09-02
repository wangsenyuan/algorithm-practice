# Problem D: Ruler Marks

## Problem Description

Valery is a PE teacher at a school in Berland. Soon the students are going to take a test in long jumps, and Valery has lost his favorite ruler!

However, there is no reason for disappointment, as Valery has found another ruler, its length is $l$ centimeters. The ruler already has $n$ marks, with which he can make measurements. We assume that the marks are numbered from 1 to $n$ in the order they appear from the beginning of the ruler to its end. The first point coincides with the beginning of the ruler and represents the origin. The last mark coincides with the end of the ruler, at distance $l$ from the origin. This ruler can be represented by an increasing sequence $a_1, a_2, \ldots, a_n$, where $a_i$ denotes the distance of the $i$-th mark from the origin ($a_1 = 0$, $a_n = l$).

Valery believes that with a ruler he can measure the distance of $d$ centimeters, if there is a pair of integers $i$ and $j$ ($1 \leq i \leq j \leq n$), such that the distance between the $i$-th and the $j$-th mark is exactly equal to $d$ (in other words, $a_j - a_i = d$).

Under the rules:
- The girls should be able to jump at least $x$ centimeters
- The boys should be able to jump at least $y$ centimeters (where $x < y$)

To test the children's abilities, Valery needs a ruler to measure each of the distances $x$ and $y$.

Your task is to determine what is the minimum number of additional marks you need to add on the ruler so that they can be used to measure the distances $x$ and $y$. Valery can add the marks at any integer non-negative distance from the origin not exceeding the length of the ruler.

## Input

- The first line contains four positive space-separated integers $n$, $l$, $x$, $y$:
  - $2 \leq n \leq 10^5$
  - $2 \leq l \leq 10^9$
  - $1 \leq x < y \leq l$
  
  These represent the number of marks, the length of the ruler and the jump norms for girls and boys, correspondingly.

- The second line contains a sequence of $n$ integers $a_1, a_2, \ldots, a_n$ ($0 = a_1 < a_2 < \ldots < a_n = l$), where $a_i$ shows the distance from the $i$-th mark to the origin.

## Output

- In the first line print a single non-negative integer $v$ — the minimum number of marks that you need to add on the ruler.

- In the second line print $v$ space-separated integers $p_1, p_2, \ldots, p_v$ ($0 \leq p_i \leq l$). Number $p_i$ means that the $i$-th mark should be at the distance of $p_i$ centimeters from the origin. Print the marks in any order. If there are multiple solutions, print any of them.

## Examples

### Example 1

**Input:**
```
3 250 185 230
0 185 250
```

**Output:**
```
1
230
```

**Explanation:** It is impossible to initially measure the distance of 230 centimeters. For that it is enough to add a 20 centimeter mark or a 230 centimeter mark.

### Example 2

**Input:**
```
4 250 185 230
0 20 185 250
```

**Output:**
```
0
```

**Explanation:** You already can use the ruler to measure the distances of 185 and 230 centimeters, so you don't have to add new marks.

### Example 3

**Input:**
```
2 300 185 230
0 300
```

**Output:**
```
2
185 230
```

**Explanation:** The ruler only contains the initial and the final marks. We will need to add two marks to be able to test the children's skills.

## Notes

- A ruler can measure distance $d$ if there exists a pair of marks at distance $d$ apart
- You need to add the minimum number of marks to measure both $x$ and $y$ distances
- Marks can be added at any integer position from 0 to $l$
- The solution should be optimal (minimum number of additional marks)