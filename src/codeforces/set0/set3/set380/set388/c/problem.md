# Fox and Card Game

## Problem Description

Fox Ciel is playing a card game with her friend Fox Jiro. There are `n` piles of cards on the table. And there is a positive integer on each card.

The players take turns and Ciel takes the first turn. In Ciel's turn she takes a card from the **top** of any non-empty pile, and in Jiro's turn he takes a card from the **bottom** of any non-empty pile. Each player wants to maximize the total sum of the cards he took. The game ends when all piles become empty.

Suppose Ciel and Jiro play optimally, what is the score of the game?

## Input

The first line contain an integer `n` (1 ≤ n ≤ 100). 

Each of the next `n` lines contains a description of the pile: the first integer in the line is `si` (1 ≤ si ≤ 100) — the number of cards in the i-th pile; then follow `si` positive integers `c1, c2, ..., ck, ..., csi` (1 ≤ ck ≤ 1000) — the sequence of the numbers on the cards listed from top of the current pile to bottom of the pile.

## Output

Print two integers: the sum of Ciel's cards and the sum of Jiro's cards if they play optimally.

## ideas
1. 假设目前所有的牌堆都有一张牌，那么策略很明显，都是取当前最大的那张牌
2. 假设排堆中有两张牌，似乎就有点复杂了。比如 （5, 4), (3, 2)
3. 假如A选择了第一堆，取走了5，那么第二个人没有选择，只能取走4；如果A选择了3，第二个人的策略，取走2似乎也没有问题
4. 即使取走了4，对方也能逼迫他在后续步骤中选择2
5. 所以，对于A来说，似乎能够决定B的步骤
6. 假设，所有的c都是偶数，那么只有一种（稳定）策略，就是A拿走前半段，B拿走后半段
7. 现在考虑奇数的部分，那么就是剩余为1的那些堆的策略