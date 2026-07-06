# C - Robot Takahashi

[Problem link](https://atcoder.jp/contests/abc257/tasks/abc257_c)

**Contest:** [AtCoder Beginner Contest 257](https://atcoder.jp/contests/abc257)

time limit: 2 sec

memory limit: 1024 MiB

score: 300 points

There are N people, each of whom is either a child or an adult. The i-th person has weight W_i.

Whether each person is a child or an adult is given by a length-N string S of `0` and `1`:

- `S[i] = '0'` — the i-th person is a child
- `S[i] = '1'` — the i-th person is an adult

When Takahashi the robot is given a real number X, it classifies a person with weight **less than** X
as a child and a person with weight **greater than or equal to** X as an adult.

For a real value X, let f(X) be the number of people whom Takahashi correctly classifies.

Find the maximum value of f(X) over all real X.

## Constraints

- 1 <= N <= 2 * 10^5
- S is a string of length N consisting of `0` and `1`
- 1 <= W_i <= 10^9
- N and W_i are integers

## Input

```text
N
S
W_1 W_2 ... W_N
```

## Output

Print the maximum value of f(X) as an integer in a single line.

## Sample Input 1

```text
5
10101
60 45 30 40 80
```

## Sample Output 1

```text
4
```

When X = 50, Takahashi classifies people 2, 3, and 4 as children and people 1 and 5 as adults.
In reality, people 2 and 4 are children and people 1, 3, and 5 are adults, so people 1, 2, 4, and 5
are correct. Thus f(50) = 4.

No X classifies all 5 people correctly, so 4 is the answer.

## Sample Input 2

```text
3
000
1 2 3
```

## Sample Output 2

```text
3
```

For example, X = 10 achieves f(10) = 3. All people may be children or all may be adults.

## Sample Input 3

```text
5
10101
60 50 50 50 60
```

## Sample Output 3

```text
4
```

For example, X = 55 achieves f(55) = 4. Multiple people may share the same weight.
