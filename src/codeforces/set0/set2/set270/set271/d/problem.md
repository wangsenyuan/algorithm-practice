# Problem D: Good Substrings

You've got string s, consisting of small English letters. Some of the English letters are good, the rest are bad.

A substring s[l...r] (1 ≤ l ≤ r ≤ |s|) of string s = s₁s₂...s_{|s|} (where |s| is the length of string s) is string sₗs_{l+1}...sᵣ.

The substring s[l...r] is **good**, if among the letters sₗ, s_{l+1}, ..., sᵣ there are at most k bad ones (look at the sample's explanation to understand it more clearly).

Your task is to find the number of distinct good substrings of the given string s. Two substrings s[x...y] and s[p...q] are considered distinct if their content is different, i.e. s[x...y] ≠ s[p...q].

## Input

The first line of the input is the non-empty string s, consisting of small English letters, the string's length is at most 1500 characters.

The second line of the input is the string of characters "0" and "1", the length is exactly 26 characters. If the i-th character of this string equals "1", then the i-th English letter is good, otherwise it's bad. That is, the first character of this string corresponds to letter "a", the second one corresponds to letter "b" and so on.

The third line of the input consists a single integer k (0 ≤ k ≤ |s|) — the maximum acceptable number of bad characters in a good substring.

## Output

Print a single integer — the number of distinct good substrings of string s.

## Examples

### Example 1

**Input:**
```
ababab
01000000000000000000000000
1
```

**Output:**
```
5
```

**Explanation:** In the first example there are following good substrings: "a", "ab", "b", "ba", "bab".

### Example 2

**Input:**
```
acbacbacaa
00000000000000000000000000
2
```

**Output:**
```
8
```

**Explanation:** In the second example there are following good substrings: "a", "aa", "ac", "b", "ba", "c", "ca", "cb".

### ideas
1. 对于l...来说，找到最大的r， s[l...r]是good的，比较容易（就是在s[l...r+1]中有k+1个bad字母）
2. 那么所有s[l..i], i <= r 的就都是good的
3. 貌似用trie是可以的