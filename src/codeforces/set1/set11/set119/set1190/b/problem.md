Tokitsukaze and CSL are playing a little game of stones.

In the beginning, there are 𝑛
 piles of stones, the 𝑖
-th pile of which has 𝑎𝑖
 stones. The two players take turns making moves. Tokitsukaze moves first. On each turn the player chooses a nonempty pile and removes exactly one stone from the pile. A player loses if all of the piles are empty before his turn, or if after removing the stone, two piles (possibly empty) contain the same number of stones. Supposing that both players play optimally, who will win the game?

Consider an example: 𝑛=3
 and sizes of piles are 𝑎1=2
, 𝑎2=3
, 𝑎3=0
. It is impossible to choose the empty pile, so Tokitsukaze has two choices: the first and the second piles. If she chooses the first pile then the state will be [1,3,0]
 and it is a good move. But if she chooses the second pile then the state will be [2,2,0]
 and she immediately loses. So the only good move for her is to choose the first pile.

Supposing that both players always take their best moves and never make mistakes, who will win the game?

Note that even if there are two piles with the same number of stones at the beginning, Tokitsukaze may still be able to make a valid first move. It is only necessary that there are no two piles with the same number of stones after she moves.

### ideas
1. interesting~
2. 如果a中存在0，那么最后最后一个人，就不能把某一堆取完（也就是不能取1的那堆）否则马上输掉了
3. 同理，如果存在1，那么他就不能取2的那堆。。。。
4. 所以，在某个状态，所有的数字，都最多出现一次。所以唯一能够move的步骤就是在两个接近的中间进行移动
5. 如果，同时存在（2, 2), 那貌似是可以移动的
6. 如果存在一对，那么先动手的，肯定能赢（如果只存在这一对，对方的操作，完全被先手控制了）
7. 但是如果是3个2呢？
8. 比如 (3, 3, 3) -> (2, 3, 3) -> (1, 3, 3) —> (0, 3, 3) (先手被转换了)
9.     (1,3,3) -> (1, 2, 3) -> (0, 2, 3) -> (0, 1, 3) -> (0, 1, 2) (还是对方赢了)
10. 感觉还是要看，目前的这些数，最多可以出现移动多少次
11. 排序后，和位置的差值，但是要把那些不能移动的排除掉