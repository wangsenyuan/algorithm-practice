# Problem E

Each evening Roma plays online poker on his favourite website. The rules of poker on this website are a bit strange: there are always two players in a hand, there are no bets, and the winner takes 1 virtual bourle from the loser.

Last evening Roma started to play poker. He decided to spend no more than k virtual bourles — he will stop immediately if the number of his loses exceeds the number of his wins by k. Also Roma will leave the game if he wins enough money for the evening, i.e. if the number of wins exceeds the number of loses by k.

Next morning Roma found a piece of paper with a sequence on it representing his results. Roma doesn't remember the results exactly, and some characters in the sequence are written in a way such that it's impossible to recognize this character, so Roma can't recall whether he won k bourles or he lost.

The sequence written by Roma is a string s consisting of characters W (Roma won the corresponding hand), L (Roma lost), D (draw) and ? (unknown result). Roma wants to restore any valid sequence by changing all ? characters to W, L or D. The sequence is called valid if all these conditions are met:

- In the end the absolute difference between the number of wins and loses is equal to k;
- There is no hand such that the absolute difference before this hand was equal to k.

Help Roma to restore any such sequence.

## Input

The first line contains two numbers n (the length of Roma's sequence) and k (1 ≤ n, k ≤ 1000).

The second line contains the sequence s consisting of characters W, L, D and ?. There are exactly n characters in this sequence.

## Output

If there is no valid sequence that can be obtained from s by replacing all ? characters by W, L or D, print NO.

Otherwise print this sequence. If there are multiple answers, print any of them.

## Examples

### Example 1

**Input:**

```text
3 2
L??
```

**Output:**

```text
LDL
```

### Example 2

**Input:**

```text
3 1
W??
```

**Output:**

```text
NO
```

### Example 3

**Input:**

```text
20 5
?LLLLLWWWWW?????????
```

**Output:**

```text
WLLLLLWWWWWWWWLWLWDW
```


### ideas
1. dp[i][j] = 到第i步时的，得分为j，是否yes
2. 也是一个图问题