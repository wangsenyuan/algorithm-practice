 # Problem Description

Ehab has an array `a` of `n` integers. He likes the bitwise-xor operation and he likes to bother Mahmoud so he came up with a problem. He gave Mahmoud `q` queries. In each of them, he gave Mahmoud 2 integers `l` and `x`, and asked him to find the number of subsequences of the first `l` elements of the array such that their bitwise-xor sum is `x`. Can you help Mahmoud answer the queries?

**Note:** A subsequence can contain elements that are not neighboring.

## Input

- The first line contains integers `n` and `q` (1 ≤ n, q ≤ 10^5), the number of elements in the array and the number of queries.
- The next line contains `n` integers a₁, a₂, ..., aₙ (0 ≤ aᵢ < 2^20), the elements of the array.
- The next `q` lines, each contains integers `l` and `x` (1 ≤ l ≤ n, 0 ≤ x < 2^20), representing the queries.

## Output

For each query, output its answer modulo 10^9 + 7 in a newline.