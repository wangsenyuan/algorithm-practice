# Problem C: Quiz

Manao is taking part in a quiz. The quiz consists of n consecutive questions. A correct answer gives one point to the player. The game also has a counter of consecutive correct answers. When the player answers a question correctly, the number on this counter increases by 1. If the player answers a question incorrectly, the counter is reset, that is, the number on it reduces to 0. If after an answer the counter reaches the number k, then it is reset, and the player's score is doubled. Note that in this case, first 1 point is added to the player's score, and then the total score is doubled. At the beginning of the game, both the player's score and the counter of consecutive correct answers are set to zero.

Manao remembers that he has answered exactly m questions correctly. But he does not remember the order in which the questions came. He's trying to figure out what his minimum score may be. Help him and compute the remainder of the corresponding number after division by 1000000009 (10^9 + 9).

## Input

The single line contains three space-separated integers n, m and k (2 ≤ k ≤ n ≤ 10^9; 0 ≤ m ≤ n).

## Output

Print a single integer — the remainder from division of Manao's minimum possible score in the quiz by 1000000009 (10^9 + 9).

## Examples

### Example 1

**Input:**
```
5 3 2
```

**Output:**
```
3
```

**Note:** Sample 1. Manao answered 3 questions out of 5, and his score would double for each two consecutive correct answers. If Manao had answered the first, third and fifth questions, he would have scored as much as 3 points.

### Example 2

**Input:**
```
5 4 2
```

**Output:**
```
6
```

**Note:** Sample 2. Now Manao answered 4 questions. The minimum possible score is obtained when the only wrong answer is to the question 4.

Also note that you are asked to minimize the score and not the remainder of the score modulo 1000000009. For example, if Manao could obtain either 2000000000 or 2000000020 points, the answer is 2000000000 mod 1000000009, even though 2000000020 mod 1000000009 is a smaller number.


## ideas
1. n个问题，回答正确了m个, 要得到最小的score
2. 应该尽量避免连续回答正确k个问题，如果避免不了，那么应该尽早的回答连续k个问题，越早double的越小
3. 假设最后剩余x个问题，回答正确了y个问题, 且没有连续k个问题回答正确
4. 那么有 (k - 1) * t <= y, k * t <= x
5. y <= m => t <=m / (k - 1) and t <= n / k
6. 所以 t = min(m / (k-1), n / k)
7. u = m - (k - 1) * t
8. v = n - k * t
9. 前面u个问题，可以全部正确
10. x = n - u, y = m - u
11. 要满足 y / (k - 1) >= x / k => k * y >= (k - 1) * x
12. (y < x)是成立的