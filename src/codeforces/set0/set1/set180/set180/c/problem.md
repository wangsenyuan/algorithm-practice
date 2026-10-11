# C. Letter

[Problem link](https://codeforces.com/problemset/problem/180/C)

**Contest:** [Codeforces Round 116 (Div. 2)](https://codeforces.com/contest/180)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

Patrick wrote a message in which some letters are lowercase and some are uppercase. A message is fancy if every uppercase letter stands to the left of every lowercase letter: zero or more uppercase letters, then zero or more lowercase letters.

In one action he flips the case of one letter. No other edits are allowed. Find the minimum number of actions that make the message fancy.

## Input

The only line contains a non-empty string of uppercase and lowercase English letters. Its length does not exceed `10^5`.

## Output

Print one integer, the minimum number of case flips.

## Example

### Input 1

```text
PRuvetSTAaYA
```

### Output 1

```text
5
```

### Input 2

```text
OYPROSTIYAOPECHATALSYAPRIVETSTASYA
```

### Output 2

```text
0
```

### Input 3

```text
helloworld
```

### Output 3

```text
0
```

## Status

I/O and official sample cases are in place. `solve` is left as a TODO.
