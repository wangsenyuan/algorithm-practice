Valera considers a number beautiful, if it equals 2k or -2k for some integer k (k ≥ 0). Recently, the math teacher asked Valera to represent number n as the sum of beautiful numbers. As Valera is really greedy, he wants to complete the task using as few beautiful numbers as possible.

Help Valera and find, how many numbers he is going to need. In other words, if you look at all decompositions of the number n into beautiful summands, you need to find the size of the decomposition which has the fewest summands.

Input
The first line contains string s (1 ≤ |s| ≤ 106), that is the binary representation of number n without leading zeroes (n > 0).

Output
Print a single integer — the minimum amount of beautiful numbers that give a total of n.


### ideas
1. 从例子看，这个还挺麻烦的
2. 假设从高位到低位，在满足高位的情况下，使用了最少的操作数
3. 现在考虑当前位，当前位的状态，貌似和高位的操作有关系
4. 如果高位操作了1次+，对当前位没有影响？
5. 如果高位操作了1次-，那么当前位似乎就反转了
6. 如果是两次，又反转了...
7. 

### sample
10000000
  -10000
01110000
    -100
 1101100
      +1
 1101101
