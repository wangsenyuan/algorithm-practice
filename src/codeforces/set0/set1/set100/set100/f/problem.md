### Problem

You are given a polynomial in the form  
\( p(x) = (x + a_1)(x + a_2)\dots(x + a_n) \).  

You must print it in the expanded standard form  
\( p(x) = x^n + b_1 x^{n-1} + \dots + b_{n-1} x + b_n \).  

Each term must be written as `C*X^K` (for example, `5*X^8`).

Please write the polynomial in the shortest way, skipping unnecessary terms.  
That is, combine like terms and omit any term whose coefficient becomes zero.  
See the samples for clarification.

### Input

- The first line contains an integer `n` \((1 \le n \le 9)\).  
- Each of the next `n` lines contains an integer `a_i` \((-10 \le a_i \le 10)\).

### Output

Print the expanded polynomial in standard form with terms sorted by decreasing powers of `X`.  
The answer for each test is uniquely determined.

### Examples

**Input**
```
2
-1
1
```

**Output**
```
X^2-1
```

**Input**
```
2
1
1
```

**Output**
```
X^2+2*X+1
```