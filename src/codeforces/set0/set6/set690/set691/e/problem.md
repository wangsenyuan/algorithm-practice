You are given n integers $a_1, a_2, ..., a_n$.

A sequence of integers $x_1, x_2, ..., x_k$ is called a "xor-sequence" if:
- For every $1 ≤ i ≤ k - 1$, the number of ones in the binary representation of $x_i ⊕ x_{i+1}$ is a multiple of 3
- For all $1 ≤ i ≤ k$, $x_i$ is one of the given integers

The symbol $⊕$ is used for the binary exclusive or (XOR) operation.

How many "xor-sequences" of length k exist? Output the answer modulo $10^9 + 7$.

**Note:** If $a = [1, 1]$ and $k = 1$, then the answer is 2, because you should consider the two 1's from array a as different elements.

### Input
The first line contains two integers n and k $(1 ≤ n ≤ 100, 1 ≤ k ≤ 10^{18})$ — the number of given integers and the length of the "xor-sequences".

The second line contains n integers $a_i$ $(0 ≤ a_i ≤ 10^{18})$.

### Output
Print the only integer c — the number of "xor-sequences" of length k modulo $10^9 + 7$.

### Examples

**Input 1:**
```
5 2
15 1 2 4 8
```

**Output 1:**
```
13
```

**Input 2:**
```
5 1
15 1 2 4 8
```

**Output 2:**
```
5
```
