Vitya has learned that the answer for The Ultimate Question of Life, the Universe, and Everything is not the integer 42, but an increasing integer sequence $a_1, \ldots, a_n$. In order to not reveal the secret earlier than needed, Vitya encrypted the answer and obtained the sequence $b_1, \ldots, b_n$ using the following rules:

- $b_1 = a_1$;
- $b_i = a_i \oplus a_{i-1}$ for all $i$ from $2$ to $n$, where $x \oplus y$ is the bitwise XOR of $x$ and $y$.

It is easy to see that the original sequence can be obtained using the rule $a_i = b_1 \oplus \ldots \oplus b_i$.

However, some time later Vitya discovered that the integers $b_i$ in the cipher got shuffled, and it can happen that when decrypted using the rule mentioned above, it can produce a sequence that is not increasing. In order to save his reputation in the scientific community, Vitya decided to find some permutation of integers $b_i$ so that the sequence $a_i = b'_1 \oplus \ldots \oplus b'_i$ is strictly increasing. Help him find such a permutation or determine that it is impossible.

## Input

The first line contains a single integer $n$ ($1 \le n \le 10^5$).

The second line contains $n$ integers $b_1, \ldots, b_n$ ($1 \le b_i < 2^{60}$).

## Output

If there are no valid permutations, print a single line containing "No".

Otherwise in the first line print the word "Yes", and in the second line print integers $b'_1, \ldots, b'_n$ — a valid permutation of integers $b_i$. The unordered multisets $\{b_1, \ldots, b_n\}$ and $\{b'_1, \ldots, b'_n\}$ should be equal, i.e. for each integer $x$ the number of occurrences of $x$ in the first multiset should be equal to the number of occurrences of $x$ in the second multiset. Apart from this, the sequence $a_i = b'_1 \oplus \ldots \oplus b'_i$ should be strictly increasing.

If there are multiple answers, print any of them.

## Examples

### Input
```
3
1 2 3
```

### Output
```
No
```

### Input
```
6
4 7 7 12 31 61
```

### Output
```
Yes
4 12 7 31 7 61
```

## Note

In the first example no permutation is valid.

In the second example the given answer leads to the sequence $a_1 = 4$, $a_2 = 8$, $a_3 = 15$, $a_4 = 16$, $a_5 = 23$, $a_6 = 42$.

## Solution (notes)

Assume we have already found a suitable permutation of all numbers **except** all occurrences of the number $1$.  
When can we insert the $1$'s so that the new arrangement is still good?

For any position where we place a $1$, the XOR of all numbers **before** that $1$ must be even, so there must be an even number of odd numbers before it.

Let $x$ be the number of $1$'s in the input, and $y$ be the number of odd numbers greater than $1$.

- If $x > y + 1$, then in any arrangement there will be a pair of $1$'s with **no** odd numbers between them, so the condition above cannot hold for both of them simultaneously → no valid permutation.
- If $x \le y + 1$, then it is possible to insert the $1$'s into any permutation of the greater numbers: put one $1$ at the start, and then place each remaining $1$ immediately after some greater odd number.

This argument generalizes: for any $k$, we can treat numbers in the range $[2^k, 2^{k+1})$ as “$1$'s” and numbers $\ge 2^{k+1}$ as “greater numbers”.  
It does not matter exactly how we insert the “$1$'s” since the number of available gaps depends only on counts.

Hence, we can:

1. Group the numbers by their highest (leading) bit.
2. Process groups in order of **decreasing** leading bit.
3. For a group with leading bit $k$, let $x$ be its size, and let $y$ be the number of already processed numbers that have bit $k$ set.
   - If $x > y + 1$, there is no answer.
   - Otherwise, we can merge this group into the answer as above.

The complexity of this solution is $\mathcal{O}(n \log A)$, where $A$ is the maximum value in the input.