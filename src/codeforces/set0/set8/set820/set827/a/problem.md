# Problem: String Reconstruction

## Description

Ivan had a string s consisting of small English letters. However, his friend Julia decided to make fun of him and hid the string s. Ivan preferred making a new string to finding the old one.

Ivan knows some information about the string s. Namely, he remembers that string tᵢ occurs in string s at least kᵢ times or more, and he also remembers exactly kᵢ positions where the string tᵢ occurs in string s: these positions are xᵢ,₁, xᵢ,₂, ..., xᵢ,ₖᵢ. He remembers n such strings tᵢ.

You are to reconstruct the lexicographically minimal string s such that it fits all the information Ivan remembers. Strings tᵢ and string s consist of small English letters only.

## Input

The first line contains a single integer n (1 ≤ n ≤ 10⁵) — the number of strings Ivan remembers.

The next n lines contain information about the strings. The i-th of these lines contains:
- A non-empty string tᵢ
- A positive integer kᵢ, which equals the number of times the string tᵢ occurs in string s
- kᵢ distinct positive integers xᵢ,₁, xᵢ,₂, ..., xᵢ,ₖᵢ in increasing order — positions where occurrences of the string tᵢ in the string s start

**Constraints:**
- The sum of lengths of strings tᵢ doesn't exceed 10⁶
- 1 ≤ xᵢ,ⱼ ≤ 10⁶
- 1 ≤ kᵢ ≤ 10⁶
- The sum of all kᵢ doesn't exceed 10⁶
- The strings tᵢ can coincide

It is guaranteed that the input data is not self-contradictory, and thus at least one answer always exists.

## Output

Print the lexicographically minimal string that fits all the information Ivan remembers.

## Examples

### Example 1
**Input:**
```
3
a 4 1 3 5 7
ab 2 1 5
ca 1 4
```

**Output:**
```
abacaba
```

### Example 2
**Input:**
```
1
a 1 3
```

**Output:**
```
aaa
```

### Example 3
**Input:**
```
3
ab 1 1
aba 1 3
ab 2 3 5
```

**Output:**
```
ababab
```