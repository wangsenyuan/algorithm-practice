# Problem G: Smith's Escape

## Story

Smith wakes up at the side of a dirty, disused bathroom, his ankle chained to pipes. Next to him is tape-player with a hand-written message "Play Me". He finds a tape in his own back pocket. After putting the tape in the tape-player, he sees a key hanging from a ceiling, chained to some kind of a machine, which is connected to the terminal next to him. After pressing a Play button a rough voice starts playing from the tape:

> "Listen up Smith. As you can see, you are in pretty tough situation and in order to escape, you have to solve a puzzle.
>
> You are given N strings which represent words. Each word is of the maximum length L and consists of characters 'a'-'e'. You are also given M strings which represent patterns. Pattern is a string of length ≤ L and consists of characters 'a'-'e' as well as the maximum 3 characters '?'. Character '?' is an unknown character, meaning it can be equal to any character 'a'-'e', or even an empty character. For each pattern find the number of words that matches with the given pattern. After solving it and typing the result in the terminal, the key will drop from the ceiling and you may escape. Let the game begin."

**Help Smith escape.**

## Problem Description

You are given N strings which represent words. Each word is of the maximum length L and consists of characters 'a'-'e'. You are also given M strings which represent patterns. Pattern is a string of length ≤ L and consists of characters 'a'-'e' as well as the maximum 3 characters '?'. Character '?' is an unknown character, meaning it can be equal to any character 'a'-'e', or even an empty character. For each pattern find the number of words that matches with the given pattern.

## Input

The first line of input contains two integers N and M (1 ≤ N ≤ 100,000, 1 ≤ M ≤ 5,000), representing the number of words and patterns respectively.

The next N lines represent each word, and after those N lines, following M lines represent each pattern. Each word and each pattern has a maximum length L (1 ≤ L ≤ 50). Each pattern has no more than three characters '?'. All other characters in words and patterns are lowercase English letters from 'a' to 'e'.

## Output

Output contains M lines and each line consists of one integer, representing the number of words that match the corresponding pattern.

## Example

### Input
```
3 1
abc
aec
ac
a?c
```

### Output
```
3
```

## Note

If we switch '?' with 'b', 'e' and with empty character, we get 'abc', 'aec' and 'ac' respectively.


## ideas
1. 有点像一个状态机
2. 最多3个？,每个问号可以表示6个状态，那么共有 6 * 6 * 6中组合（216）, 5000 * 216 = 1e6个pattern
3. 但是又出现一个新的问题, a??b 有可能匹配 ab 两次, acb两次, 