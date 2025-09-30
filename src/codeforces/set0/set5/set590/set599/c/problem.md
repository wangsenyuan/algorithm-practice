# Problem C: Sand Castles

## Problem Statement

One day Squidward, Spongebob and Patrick decided to go to the beach. Unfortunately, the weather was bad, so the friends were unable to ride waves. However, they decided to spend their time building sand castles.

At the end of the day there were $n$ castles built by friends. Castles are numbered from $1$ to $n$, and the height of the $i$-th castle is equal to $h_i$. When friends were about to leave, Squidward noticed that castles are not ordered by their height, and this looks ugly. Now friends are going to reorder the castles in a way to obtain that condition $h_i \leq h_{i+1}$ holds for all $i$ from $1$ to $n-1$.

## Sorting Process

Squidward suggested the following process of sorting castles:

1. Castles are split into blocks — groups of consecutive castles. Therefore the block from $i$ to $j$ will include castles $i, i+1, \ldots, j$. A block may consist of a single castle.
2. The partitioning is chosen in such a way that every castle is a part of exactly one block.
3. Each block is sorted independently from other blocks, that is the sequence $h_i, h_{i+1}, \ldots, h_j$ becomes sorted.
4. The partitioning should satisfy the condition that after each block is sorted, the sequence $h_i$ becomes sorted too. This may always be achieved by saying that the whole sequence is a single block.

Even Patrick understands that increasing the number of blocks in partitioning will ease the sorting process. Now friends ask you to count the maximum possible number of blocks in a partitioning that satisfies all the above requirements.

## Input

- The first line of the input contains a single integer $n$ ($1 \leq n \leq 100,000$) — the number of castles Spongebob, Patrick and Squidward made from sand during the day.
- The next line contains $n$ integers $h_i$ ($1 \leq h_i \leq 10^9$). The $i$-th of these integers corresponds to the height of the $i$-th castle.

## Output

Print the maximum possible number of blocks in a valid partitioning.

## Examples

### Example 1
**Input:**
```
3
1 2 3
```

**Output:**
```
3
```

### Example 2
**Input:**
```
4
2 1 3 2
```

**Output:**
```
2
```