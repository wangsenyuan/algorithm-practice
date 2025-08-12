# Problem C: Correct Bracket Sequences

## Problem Description

As you know, Vova has recently become a new shaman in the city of Ultima Thule. So, he has received the shaman knowledge about the correct bracket sequences. The shamans of Ultima Thule have been using lots of different types of brackets since prehistoric times. A bracket type is a positive integer. The shamans define a correct bracket sequence as follows:

1. **An empty sequence is a correct bracket sequence.**
2. **If {a₁, a₂, ..., aₗ} and {b₁, b₂, ..., bₖ} are correct bracket sequences, then sequence {a₁, a₂, ..., aₗ, b₁, b₂, ..., bₖ} (their concatenation) also is a correct bracket sequence.**
3. **If {a₁, a₂, ..., aₗ} is a correct bracket sequence, then sequence {v, a₁, a₂, ..., aₗ, -v} also is a correct bracket sequence, where v (v > 0) is an integer.**

### Examples of Correct Bracket Sequences
- **{1, 1, -1, 2, -2, -1}** is a correct bracket sequence
- **{3, -3}** is a correct bracket sequence
- **{2, -3}** is **NOT** a correct bracket sequence

## The Problem

After Vova became a shaman, he learned the most important correct bracket sequence {x₁, x₂, ..., xₙ}, consisting of n integers. As sequence x is the most important, Vova decided to encrypt it just in case.

### Encryption Process
The encryption consists of two sequences:
1. **Sequence {p₁, p₂, ..., pₙ}** contains types of brackets, where pᵢ = |xᵢ| (1 ≤ i ≤ n)
2. **Sequence {q₁, q₂, ..., qₜ}** contains t integers — some positions (possibly, not all of them), which had negative numbers in sequence {x₁, x₂, ..., xₙ}

### Current Situation
Unfortunately, Vova forgot the main sequence. But he was lucky enough to keep the encryption: sequences {p₁, p₂, ..., pₙ} and {q₁, q₂, ..., qₜ}. 

**Help Vova restore sequence x by the encryption.** If there are multiple sequences that correspond to the encryption, restore any of them. If there are no such sequences, you should tell so.

## Input Format

- **Line 1:** Integer n (1 ≤ n ≤ 10⁶)
- **Line 2:** n integers: p₁, p₂, ..., pₙ (1 ≤ pᵢ ≤ 10⁹)
- **Line 3:** Integer t (0 ≤ t ≤ n), followed by t distinct integers q₁, q₂, ..., qₜ (1 ≤ qᵢ ≤ n)

The numbers in each line are separated by spaces.

## Output Format

- Print a single string "NO" (without quotes) if Vova is mistaken and a suitable sequence {x₁, x₂, ..., xₙ} doesn't exist.
- Otherwise, print:
  - **Line 1:** "YES" (without quotes)
  - **Line 2:** n integers x₁, x₂, ..., xₙ where |xᵢ| = pᵢ and x_{qⱼ} < 0

If there are multiple sequences that correspond to the encryption, you are allowed to print any of them.

## Examples

### Example 1
**Input:**
```
2
1 1
0
```

**Output:**
```
YES
1 -1
```

### Example 2
**Input:**
```
4
1 1 1 1
1 3
```

**Output:**
```
YES
1 1 -1 -1
```

### Example 3
**Input:**
```
3
1 1 1
0
```

**Output:**
```
NO
```

### Example 4
**Input:**
```
4
1 2 2 1
2 3 4
```

**Output:**
```
YES
1 2 -2 -1
```


## ideas
1. 有一些简单的case，可以先处理; 如果n不是偶数 => -1
2. 如果t超过了 n/2 => -1
3. 