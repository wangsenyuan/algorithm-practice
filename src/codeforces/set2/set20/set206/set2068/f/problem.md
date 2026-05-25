# F. Mascot Naming

[Problem link](https://codeforces.com/problemset/problem/2068/F)

time limit per test: 2 seconds

memory limit per test: 1024 megabytes

input: standard input

output: standard output

When organizing a big event, organizers often handle side tasks outside their
expertise. For example, the chief judge of EUC 2025 must find a name for the
event's official mascot while satisfying certain constraints:

- The name must include specific words as subsequences, such as the event name
  and location. You are given the list `s_1, s_2, ..., s_n` of the `n` required
  words.
- The name must not contain as a subsequence the name `t` of last year's mascot.

Please help the chief judge find a valid mascot name or determine that none
exists.

A string `x` is a subsequence of a string `y` if `x` can be obtained from `y` by
erasing some characters, at any positions, while keeping the remaining
characters in the same order. For example, `abc` is a subsequence of `axbycz`
but not of `acbxyz`.

## Input

The first line contains an integer `n` (`1 <= n <= 200000`) -- the number of
words that shall appear as subsequences.

The `i`-th of the following `n` lines contains the string `s_i`
(`1 <= |s_i| <= 200000`, `s_i` consists of lowercase English letters) -- the
`i`-th word in the list of words that shall appear as subsequences. The total
length of these `n` words is at most `200000`, i.e.,
`|s_1| + |s_2| + ... + |s_n| <= 200000`.

The last line contains the string `t` (`1 <= |t| <= 200000`, `t` consists of
lowercase English letters) -- the name of last year's mascot.

## Output

Print `YES` if there is a valid name for the mascot. Otherwise, print `NO`.

If there is a valid name, on the next line print a valid name. The string you
print must have length at most `1000000` and must consist of lowercase English
letters. One can prove that if a valid name for the mascot exists, then there is
one satisfying these additional constraints.

If there are multiple solutions, print any of them.

## Examples

### Input 1

```text
2
porto
euc
prague
```

### Output 1

```text
YES
poretuco
```

### Input 2

```text
6
credit
debit
money
rich
bank
capitalism
trap
```

### Output 2

```text
YES
moncrdebditeychankpitalism
```

### Input 3

```text
2
axiom
choice
io
```

### Output 3

```text
NO
```

### Input 4

```text
4
aaa
aab
abb
bbb
ba
```

### Output 4

```text
YES
aaabbb
```

## Note

In the first sample, the words that must appear as subsequences are `porto` and
`euc`, while the word that must not appear as a subsequence is `prague`. There
are many valid names for the mascot, e.g. `poretuco` or
`xxxpppeortoucyyyy`.

If `poretuco` is chosen as the name of the mascot, observe that `porto` and
`euc` are subsequences, while `prague` is not, as desired.

The string `poretuc` is not a valid mascot name because it does not contain the
word `porto` as a subsequence. Also the string `poretucoague` is not a valid
mascot name because it contains `prague` as a subsequence.


### editorial

This problem requires constructing a string that contains each of the given strings 𝑠1,𝑠2,…,𝑠𝑛
 as subsequences but does not contain 𝑡
 as a subsequence.

Key Observation. If 𝑡
 is a subsequence of any 𝑠𝑖
, then any string containing 𝑠𝑖
 as a subsequence must also contain 𝑡
. In this case, the answer is 𝙽𝙾
. Otherwise, we can construct a valid string using a greedy approach.

Checking for Subsequences. Before solving the main problem, we describe a simple method to check if a string 𝑏
 is a subsequence of another string 𝑎
. We iterate through 𝑎
, trying to match its characters to 𝑏
 in order:

If the first characters of 𝑎
 and 𝑏
 match, remove the first character from both.
Otherwise, remove only the first character from 𝑎
.
Repeat until 𝑎
 is empty.
If 𝑏
 becomes empty, it was a subsequence of 𝑎
; otherwise, it was not.
Constructing the Answer. We build the answer string ans
 by appending characters one by one. Initially, ans
 is empty. We repeat the following steps until all 𝑠𝑖
 are empty:

If there exists a character 𝑐
 such that at least one 𝑠𝑖
 starts with 𝑐
 and 𝑡
 does not start with 𝑐
, append 𝑐
 to ans
 and remove the first character from all 𝑠𝑖
 that start with 𝑐
.
Otherwise, all 𝑠𝑖
 start with the same character as 𝑡
. Append this character to ans
 and remove the first character from all 𝑠𝑖
 and 𝑡
.
At the end of this process, ans
 contains all 𝑠𝑖
 as subsequences.

Let us understand why 𝑡
 is not a subsequence of ans
. Let 𝑠𝑖
 be the last string to become empty during the process. Consider the steps taken for 𝑠𝑖
 and 𝑡
 alone. The sequence of character deletions mirrors the subsequence checking algorithm. Since 𝑡
 is not a subsequence of 𝑠𝑖
, it remains non-empty at the end, ensuring that 𝑡
 is not a subsequence of ans
 either.

Thus, if 𝑡
 is not a subsequence of any 𝑠𝑖
, the algorithm constructs a valid ans
, and the answer is 𝚈𝙴𝚂
.

The solution described is linear in the total length of all strings given in input.

