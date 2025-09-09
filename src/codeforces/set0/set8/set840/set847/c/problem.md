# Problem C: Regular Bracket Sequence

## Problem Statement

Recall that a bracket sequence is considered **regular** if it is possible to insert symbols '+' and '1' into it so that the result is a correct arithmetic expression. 

For example:
- `"(()())"` is regular because we can get the correct arithmetic expression by inserting symbols '+' and '1': `"((1+1)+(1+1))"`
- The following sequences are regular: `"()()()"`, `"(())"`, and `"()"`
- The following sequences are **not** regular: `")("`, `"(()"`, and `"())(()"`

In this problem, you are given two integers `n` and `k`. Your task is to construct a regular bracket sequence consisting of round brackets with length `2·n` such that the total sum of nesting of all opening brackets equals exactly `k`.

The **nesting** of a single opening bracket equals the number of pairs of brackets in which the current opening bracket is embedded.

### Example
In the sequence `"()(())"`:
- The nesting of the first opening bracket equals 0
- The nesting of the second opening bracket equals 0  
- The nesting of the third opening bracket equals 1

So the total sum of nestings equals 1.

## Input
The first line contains two integers `n` and `k` (`1 ≤ n ≤ 3·10^5`, `0 ≤ k ≤ 10^18`) — the number of opening brackets and the needed total nesting.

## Output
Print the required regular bracket sequence consisting of round brackets.

If there is no solution, print `"Impossible"` (without quotes).

## Examples

### Example 1
**Input:**
```
3 1
```

**Output:**
```
()(())
```

### Example 2
**Input:**
```
4 6
```

**Output:**
```
(((())))
```

### Example 3
**Input:**
```
2 5
```

**Output:**
```
Impossible
```

## Notes
- The first example is examined in the problem statement.
- In the second example, the answer is `"(((())))"`. The nesting of the first opening bracket is 0, the nesting of the second is 1, the nesting of the third is 2, and the nesting of the fourth is 3. So the total sum of nestings equals `0 + 1 + 2 + 3 = 6`.
- In the third example, it is impossible to construct a regular bracket sequence because the maximum possible total sum of nestings for two opening brackets equals 1. This total sum of nestings is obtained for the sequence `"(())"`.

### ideas
1. 右括号一共有n个
2. 假设它们都嵌套起来， ((((()))))
3. 那么总的sum = (0 + n - 1) * n / 2 
4. 如果 sum < k  => impossible
5. 否则肯定是可以的