# Problem Description

At each moment, you have a set of intervals. You can move from interval \((a, b)\) in your set to interval \((c, d)\) in your set **if and only if**:

- \(c < a < d\) **or** \(c < b < d\)

There is a path from interval \(I_1\) to interval \(I_2\) if there is a sequence of successive moves starting from \(I_1\) so that you can reach \(I_2\).

## Queries

Your program should handle the following two types of queries:

- `1 x y` (where \(x < y\)) — Add the new interval \((x, y)\) to the set of intervals. The length of the new interval is guaranteed to be strictly greater than all previous intervals.
- `2 a b` (where \(a \neq b\)) — Answer the question: **Is there a path from the a-th (one-based) added interval to the b-th (one-based) added interval?**

> Answer all the queries. Note that initially you have an empty set of intervals.