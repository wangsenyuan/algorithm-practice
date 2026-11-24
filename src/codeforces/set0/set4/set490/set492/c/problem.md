## Problem: Vanya and Exams

Vanya is a student who wants to get an average grade of at least `avg` over `n` exams. Initially, he has grade `a_i` for exam `i`, and he may improve that grade to `r` by spending `b_i` rubles per essay.

Find the minimum amount of money Vanya needs to spend to reach the required average grade.

### Input

The first line contains three integers `n`, `r`, and `avg` (`1 <= n <= 10^5`, `1 <= r <= 10^9`, `1 <= avg <= r`) — the number of exams, the maximum grade, and the required average grade.

Each of the next `n` lines contains two integers `a_i` and `b_i` (`0 <= a_i <= r`, `1 <= b_i <= 10^9`) — Vanya's initial grade and the cost of raising the grade by one point for that exam.

### Output

Print the minimum amount of money needed.

### Example

```text
Input
3 5 4
4 3
4 2
4 1

Output
3
```
