# D - Jumping Takahashi 2

[Problem link](https://atcoder.jp/contests/abc257/tasks/abc257_d)

**Contest:** [AtCoder Beginner Contest 257](https://atcoder.jp/contests/abc257)

time limit: 3 sec

memory limit: 1024 MiB

score: 400 points

There are N trampolines on a 2D plane. The i-th trampoline is at (x_i, y_i) and has power P_i.

Takahashi's jumping ability is S. Initially S = 0; each training increases S by 1.

Takahashi can jump from trampoline i to trampoline j if and only if:

P_i * S >= |x_i - x_j| + |y_i - y_j|

Choose a starting trampoline so that every other trampoline is reachable by a sequence of jumps.
Find the minimum number of trainings needed to make this possible.

## Constraints

- 2 <= N <= 200
- -10^9 <= x_i, y_i <= 10^9
- 1 <= P_i <= 10^9
- (x_i, y_i) != (x_j, y_j) for i != j
- All input values are integers

## Input

```text
N
x_1 y_1 P_1
...
x_N y_N P_N
```

## Output

Print the minimum number of trainings required.

## Sample Input 1

```text
4
-10 0 1
0 0 5
10 0 1
11 0 1
```

## Sample Output 1

```text
2
```

After 2 trainings, S = 2. Starting from trampoline 2, every trampoline is reachable.
For example, to reach trampoline 4: jump 2 -> 3 (P_2 * S = 10, distance 10), then 3 -> 4
(P_3 * S = 2, distance 1).

## Sample Input 2

```text
7
20 31 1
13 4 3
-10 -15 2
34 26 5
-2 39 4
0 -50 1
5 -20 2
```

## Sample Output 2

```text
18
```
