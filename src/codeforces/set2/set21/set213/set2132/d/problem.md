# D. From 1 to Infinity

Consider the infinite digit string formed by writing all positive integers in order: `123456789101112131415…`

Vadim **truncates** this string after exactly `k` digits (the rest is discarded). Find the **sum of digits** of the resulting length-`k` prefix.

## Input

Each test file has several test cases. The first line contains an integer `t` (`1 ≤ t ≤ 2 · 10^4`) — the number of test cases.

Each test case is one line with a single integer `k` — how many digits remain (`1 ≤ k ≤ 10^15`).

## Output

For each test case, print one integer: the sum of all digits in that prefix of length `k`.

## Example

**Input**

```text
6
5
10
13
29
1000000000
1000000000000000
```

**Output**

```text
15
46
48
100
4366712386
4441049382716054
```

## Note

In the first test case, the prefix is `12345`.

In the second test case, the prefix is `1234567891`.

In the third test case, the prefix is `1234567891011`.
