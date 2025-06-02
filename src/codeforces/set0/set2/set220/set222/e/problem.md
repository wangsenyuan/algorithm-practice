# Problem: Martian DNA Chains

Recently a top secret mission to Mars has taken place. As a result, scientists managed to obtain some information about the Martian DNA. Now we know that any Martian DNA contains at most **m** different nucleotides, numbered from 1 to **m**. Special characteristics of the Martian DNA prevent some nucleotide pairs from following consecutively in this chain. For example, if the nucleotide 1 and nucleotide 2 cannot follow consecutively in the Martian DNA, then the chain of nucleotides `[1, 2]` is not a valid chain of Martian DNA, but the chain of nucleotides `[2, 1]` can be a valid chain (if there is no corresponding restriction).

The number of nucleotide pairs that can't follow in the DNA chain consecutively is **k**.

The needs of gene research required information about the quantity of correct **n**-long chains of the Martian DNA. Your task is to write a program that will calculate this value.

---

## Input
- The first line contains three space-separated integers: `n`, `m`, `k`  
  (1 ≤ n ≤ 10¹⁵, 1 ≤ m ≤ 52, 0 ≤ k ≤ m²).
- Next **k** lines contain two characters each, without a space between them, representing a forbidden nucleotide pair. The first character represents the first nucleotide in the forbidden pair, the second character represents the second nucleotide.

  - The nucleotides with assigned numbers from 1 to 26 are represented by English alphabet letters from "a" to "z" (1 is an "a", 2 is a "b", ..., 26 is a "z").
  - Nucleotides with assigned numbers from 27 to 52 are represented by English alphabet letters from "A" to "Z" (27 is an "A", 28 is a "B", ..., 52 is a "Z").

- It is guaranteed that each forbidden pair occurs at most once in the input.
- It is guaranteed that nucleotide's numbers in all forbidden pairs cannot be more than **m**.
- Note that order is important in nucleotide pairs.

---

## Output
Print a single integer — the sought number modulo `1000000007` (10⁹ + 7).


## ideas
1. 没有限制时，pow(m, n)
2. 考虑限制的情况，假设在a的后面，不能跟b
3. mat[x][y] = 1， 表示x的后面可以是1，else 0
4. 然后pow(mat, n-1) * vector(1)