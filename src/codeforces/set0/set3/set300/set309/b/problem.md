# Advertising Banner

Advertising has become part of our routine. And now, in the era of progressive technologies, we need your ideas to make advertising better!

In this problem, we'll look at a simplified version of context advertising. You've got a text, consisting of exactly **n** words. A standard advertising banner has exactly **r** lines, each line can contain at most **c** characters.

The potential customer always likes it when they can see lots of advertising, so you should determine which maximum number of consecutive words from the text can be written on the banner. Single words in one line of the banner should be separated by spaces. You are allowed to insert more than one space at once.

- You are **not allowed** to break the words; each word must occupy exactly one line in the banner.
- You **cannot change** the word order; if you read the banner text consecutively, from top to bottom and from left to right, you should get some consecutive part of the advertisement text.

More formally:
Let all words be indexed from 1 to n in the order in which they occur in the advertisement text. Then you have to choose all words, starting from some i-th one and ending with some j-th one (1 ≤ i ≤ j ≤ n), so that all of them could be written on the banner. There must be as many words as possible. See the samples for clarifications.

---

## Input

- The first input line contains three integers:
  `n r c`
  where
  - 1 ≤ n, r, c ≤ 10⁶
  - r × c ≤ 10⁶
- The next line contains a text, consisting of n words.
  - The words consist only of lowercase English letters and are not empty.
  - The words in the lines are separated by single spaces.
  - The total number of characters in all words doesn't exceed 5·10⁶.

---

## Output

- Print at most **r** lines, in each line print at most **c** characters — the optimal advertisement banner.
- If there are multiple advertisement banners, print any of them.

---

**Note:**
Some lines of the banner can be empty. You are allowed not to print such lines.

## ideas

- 选择一段连续的 words，将它们放置到 (r × c) 的区域内。
- 满足两个 word 之间需要一个空格（或者换行）。
- 如果 s[i:j] 可以放入一行中，那么：
  ```text
  dp[i] = dp[j+1] + j - i + 1
  ```
- 也就是说对于 i 要找到最大的 j，使得 s[i:j] 能放入一行中。
- 条件为：
  ```text
  tot_length(s[i:j]) + j - i <= c
  sum[j] - sum[i-1] + j - i <= c
  sum[j] + j - (sum[i-1] + i) <= c
  ```
- 不用计数，直接使用 next 可以吗？`next[i] = j + 1`
- 因为选择起点后，开始放置的时候，是贪心地放（能放就要放进去，否则新起一行）。
- 所以 next 是足够的，最后再检查哪个 `next[i][h]` 最大即可。