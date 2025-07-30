# Problem D: Dexterina and Womandark's Nim Game

## Problem Description

Dexterina and Womandark have been arch-rivals since they've known each other. Since both are super-intelligent teenage girls, they've always been trying to solve their disputes in a peaceful and nonviolent way. After god knows how many different challenges they've given to one another, their score is equal and they're both desperately trying to best the other in various games of wits. This time, Dexterina challenged Womandark to a game of Nim.

Nim is a two-player game in which players take turns removing objects from distinct heaps. On each turn, a player must remove at least one object, and may remove any number of objects from a single heap. The player who can't make a turn loses. By their agreement, the sizes of piles are selected randomly from the range [0, x]. Each pile's size is taken independently from the same probability distribution that is known before the start of the game.

Womandark is coming up with a brand new and evil idea on how to thwart Dexterina's plans, so she hasn't got much spare time. She, however, offered you some tips on looking fabulous in exchange for helping her win in Nim. Your task is to tell her what is the probability that the first player to play wins, given the rules as above.

## Input

The first line of the input contains two integers n (1 ≤ n ≤ 10^9) and x (1 ≤ x ≤ 100) — the number of heaps and the maximum number of objects in a heap, respectively.

The second line contains x + 1 real numbers, given with up to 6 decimal places each: P(0), P(1), ..., P(x). Here, P(i) is the probability of a heap having exactly i objects at the start of a game. It's guaranteed that the sum of all P(i) is equal to 1.

## Output

Output a single real number, the probability that the first player wins. The answer will be judged as correct if it differs from the correct answer by at most 10^-6.

## Example

### Input
```
2 2
0.500000 0.250000 0.250000
```

### Output
```
0.62500000
```

### ideas
1. first player wins when xor(a[i]) > 0
2. 所以，应该计算 xor(a[i]) = 0 的概率？
3. 一共n堆，然后 x+1个数字
4. 假设有w[0], w[1], ... w[x]个数
5. 那么偶数的部分，可以去掉，剩下的就有奇数的部分， 假设其中, a, b, c, d 是奇数个
6. 且xor(a, b, c, d) = 0
7. 所以，这里先选中一些数，使的它们的xor = 0
8. 假设选中了m个数，那么剩余 (n - m) % 2 = 0, 那么剩余 x+1个数，只要每个选择偶数次
9. 这m个数的概率可以算的， dp[m][s] 表示选中m个数，且xor = s（选中的数都是奇数次）的概率
10. 但是剩余的概率要怎么计算呢？
11. 那么要抛出2个1的概率是多少呢？
12. 0也要特殊处理