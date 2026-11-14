There are `n` people, and person `i` is willing to attend `a_i` productive meetings.

A productive meeting is a pair of two different people `(u, v)`. If they meet, both of their counts decrease by `1`, and the pair `(u, v)` is recorded. We want to maximize the total number of productive meetings.

For each test case:
- the first line contains `n`,
- the second line contains `a_1, a_2, ..., a_n`.

Output the maximum number of meetings, followed by the chosen meetings as `u v` pairs (1-indexed).

### idea
1. meetings count <= sum(a) / 2
2. 貌似就是找最大的两个人, 去talk