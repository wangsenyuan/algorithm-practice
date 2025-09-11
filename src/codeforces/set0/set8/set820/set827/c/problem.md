# Problem: DNA Evolution

## Description

Everyone knows that DNA strands consist of nucleotides. There are four types of nucleotides: "A", "T", "G", and "C". A DNA strand is a sequence of nucleotides. Scientists decided to track the evolution of a rare species, whose DNA strand was initially string s.

Evolution of the species is described as a sequence of changes in the DNA. Every change is a change of some nucleotide. For example, the following change can happen in DNA strand "AAGC": the second nucleotide can change to "T" so that the resulting DNA strand is "ATGC".

Scientists know that some segments of the DNA strand can be affected by unknown infections. They can represent an infection as a sequence of nucleotides. Scientists are interested in whether there are any changes caused by some infections. Thus, they sometimes want to know the value of impact of some infection on some segment of the DNA. This value is computed as follows:

Let the infection be represented as a string e, and let scientists be interested in the DNA strand segment starting from position l to position r, inclusive.

A prefix of the string eee... (i.e., the string that consists of infinitely many repeats of string e) is written under the string s from position l to position r, inclusive.

The value of impact is the number of positions where the letter of string s coincides with the letter written under it.

Being a developer, Innokenty is interested in bioinformatics also, so the scientists asked him for help. Innokenty is busy preparing VK Cup, so he decided to delegate the problem to the competitors. Help the scientists!

## Input

The first line contains the string s (1 ≤ |s| ≤ 10⁵) that describes the initial DNA strand. It consists only of capital English letters "A", "T", "G", and "C".

The next line contains a single integer q (1 ≤ q ≤ 10⁵) — the number of events.

After that, q lines follow, each describing one event. Each of the lines has one of two formats:

1. `1 x c`, where x is an integer (1 ≤ x ≤ |s|), and c is a letter "A", "T", "G", or "C", which means that there is a change in the DNA: the nucleotide at position x is now c.

2. `2 l r e`, where l, r are integers (1 ≤ l ≤ r ≤ |s|), and e is a string of letters "A", "T", "G", and "C" (1 ≤ |e| ≤ 10), which means that scientists are interested in the value of impact of infection e on the segment of DNA strand from position l to position r, inclusive.

## Output

For each scientists' query (second type query), print a single integer in a new line — the value of impact of the infection on the DNA.

## Examples

### Example 1
**Input:**
```
ATGCATGC
4
2 1 8 ATGC
2 2 6 TTT
1 4 T
2 2 6 TA
```

**Output:**
```
8
2
4
```

### Example 2
**Input:**
```
GAGTTGTTAA
6
2 3 4 TATGGTG
1 1 T
1 6 G
2 5 9 AGTAATA
1 10 G
2 2 6 TTGT
```

**Output:**
```
0
3
1
```

## Note

Consider the first example:

- In the first query of the second type, all characters coincide, so the answer is 8
- In the second query, we compare string "TTTTT..." with the substring "TGCAT". There are two matches
- In the third query, after the DNA change, we compare string "TATAT..." with substring "TGTAT". There are 4 matches

### ideas
1. 先不考虑修改的情况，如何计算x = s[l...r]和e的匹配度？
2. 考虑A的位置, 从e的角度考虑，就是 e[1], e[2] = 'A', 那么就是 由 i % len(e) = 1, i % len(e) = 2 组成的位置，都应该是A
3. 然后考虑x中，对应位置是不是A
4. 但是这样子，肯定会出超时，得加速一下， 考虑 dp[k][j][A, T, G, C][i] = 当 i % k = j时, 且s[i] = A, T, G, C时的前缀和
5. 那么dp[k][j][A, T, G, C]可以变成BIT， 共有 10 * 9 * 4 个树
6. 修改也能搞