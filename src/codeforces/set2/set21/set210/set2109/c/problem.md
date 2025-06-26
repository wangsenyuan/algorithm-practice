# Interactive Problem

## Overview
The interaction for each test case begins by reading the integer 𝑛.

## Available Commands

### 1. Add Command
```
add 𝑦
```
- Adds integer 𝑦 (−10^18 ≤ 𝑦 ≤ 10^18) to 𝑥
- Jury response:
  - "1" if 𝑥 + 𝑦 is within [1, 10^18] (successful)
  - "0" otherwise
- If successful, updates 𝑥 ← 𝑥 + 𝑦

### 2. Multiply Command
```
mul 𝑦
```
- Multiplies 𝑥 by a positive integer 𝑦 (1 ≤ 𝑦 ≤ 10^18)
- Jury response:
  - "1" if 𝑥⋅𝑦 is within [1, 10^18] (successful)
  - "0" otherwise
- If successful, updates 𝑥 ← 𝑥⋅𝑦

### 3. Divide Command
```
div 𝑦
```
- Divides 𝑥 by a positive integer 𝑦 (1 ≤ 𝑦 ≤ 10^18)
- Jury response:
  - "1" if 𝑦 is a divisor of 𝑥 (successful)
  - "0" otherwise
- If successful, updates 𝑥 ← 𝑥/𝑦

### 4. Digit Sum Command
```
digit
```
- Makes 𝑥 equal to the sum of its digits
- Jury always outputs "1"
- Updates 𝑥 ← 𝑆(𝑥)

## Answer Command
```
!
```
- Use when you believe 𝑥 equals 𝑛
- Jury response:
  - "1" if 𝑛 equals 𝑥
  - "-1" otherwise
- Note: Answering does not count toward the 7-command limit

## Important Notes
- Commands are case sensitive
- Maximum 7 commands per test case
- Invalid commands or exceeding command limit results in "-1" response
- Program should terminate after receiving "-1" response
- Remember to flush output after each command:
  - C++: `fflush(stdout)` or `cout.flush()`
  - Java: `System.out.flush()`
  - Python: `sys.stdout.flush()`
  - Rust: `std::io::stdout().flush()`

## Hacks Format
1. First line: single integer 𝑡 (1 ≤ 𝑡 ≤ 5000) — number of test cases
2. For each test case:
   - Two positive integers 𝑛 and 𝑥 (1 ≤ 𝑛,𝑥 ≤ 10^9)
   - 𝑛: unknown integer
   - 𝑥: target value to which it should be made equal

## ideas
1. 如果x + y 在范围内，就更新x, 否则不更新（不更新的时候，可以知道x的一个范围）
2. x + y < 1 => x < 1 - y
3. 如果 x + y (被更新了) => x + y >= 1 and x + y <= 1e18
4. => x <= 1 - y and x >= 1e18 - y
5. x * y 也可以知道一个范围, x * y >= 1 and x * y <= 1e18
6. digit 操作似乎很有用
7. digit(x)一个9位的数，
8. x + y (这个y > 1e9)
9. y的digit是不会和x混合的，假设是u, x的digit sum = v, digit(x + y) = u + v
10. x <= u + v (这里u是知道的, y的digit sum)
11. 再计算一次digit sum (如果成功) u + v >= 10 