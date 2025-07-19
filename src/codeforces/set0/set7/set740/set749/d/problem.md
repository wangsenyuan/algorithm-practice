# Auction Problem

## Problem Description

There are n people taking part in auction today. The rules of auction are classical. There were n bids made, though it's not guaranteed they were from different people. It might happen that some people made no bids at all.

Each bid is defined by two integers (ai, bi), where:
- ai is the index of the person who made this bid
- bi is the size of the bid

Bids are given in chronological order, meaning bi < bi+1 for all i < n. Moreover, a participant never makes two bids in a row (no one updates his own bid), i.e. ai ≠ ai+1 for all i < n.

Now you are curious with the following question: **who (and which bid) will win the auction if some participants were absent?** Consider that if someone was absent, all his bids are just removed and no new bids are added.

**Important Note:** If during this imaginary exclusion of some participants it happens that some of the remaining participants makes a bid twice (or more times) in a row, only the first of these bids is counted. For better understanding, take a look at the samples.

You have several questions in your mind, compute the answer for each of them.

## Input

- **Line 1:** An integer n (1 ≤ n ≤ 200,000) — the number of participants and bids
- **Lines 2 to n+1:** Two integers ai and bi (1 ≤ ai ≤ n, 1 ≤ bi ≤ 10^9, bi < bi+1) — the number of participant who made the i-th bid and the size of this bid
- **Line n+2:** An integer q (1 ≤ q ≤ 200,000) — the number of questions you have in mind
- **Lines n+3 to n+q+2:** An integer k (1 ≤ k ≤ n), followed by k integers lj (1 ≤ lj ≤ n) — the number of people who are not coming in this question and their indices. It is guaranteed that lj values are different for a single question

**Constraints:**
- The sum of k over all questions won't exceed 200,000

## Output

For each question, print two integers — the index of the winner and the size of the winning bid. If there is no winner (there are no remaining bids at all), print two zeroes.

## Examples

### Example 1

**Input:**
```
6
1 10
2 100
3 1000
1 10000
2 100000
3 1000000
3
1 3
2 2 3
2 1 2
```

**Output:**
```
2 100000
1 10
3 1000
```

**Explanation:**
- **Question 1:** Participant 3 is absent, so the sequence becomes: `1 10`, `2 100`, `1 10000`, `2 100000`. Participant 2 wins with bid 100,000.
- **Question 2:** Participants 2 and 3 are absent, so the sequence becomes: `1 10`, `1 10000`. Participant 1 wins with bid 10 (not 10,000, as no one increases their own bid).
- **Question 3:** Participants 1 and 2 are absent, so the sequence becomes: `3 1000`, `3 1000000`. Participant 3 wins with bid 1,000.

### Example 2

**Input:**
```
3
1 10
2 100
1 1000
2
2 1 2
2 2 3
```

**Output:**
```
0 0
1 10
```

**Explanation:**
- **Question 1:** Participants 1 and 2 are absent, leaving no bids. Output: `0 0`
- **Question 2:** Participants 2 and 3 are absent, leaving only: `1 10`. Participant 1 wins with bid 10.

## ideas
1. 考虑一棵树，每个节点，返回对应区间的winer，已经size
2. 似乎还不大够，winer和size是最后的那个出价的人，如果前面的区间也是他，如果中间没有其他人出价，那么应该放弃后面那次出价
3. 这样子肯定是不对的
4. 