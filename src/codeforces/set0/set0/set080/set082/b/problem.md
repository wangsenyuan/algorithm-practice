Little Vasya likes very much to play with sets consisting of positive integers. To make the game more interesting, Vasya chose n non-empty sets in such a way, that no two of them have common elements.

One day he wanted to show his friends just how interesting playing with numbers is. For that he wrote out all possible unions of two different sets on n·(n - 1) / 2 pieces of paper. Then he shuffled the pieces of paper. He had written out the numbers in the unions in an arbitrary order.

For example, if n = 4, and the actual sets have the following form {1, 3}, {5}, {2, 4}, {7}, then the number of set pairs equals to six. The six pieces of paper can contain the following numbers:

2, 7, 4.
1, 7, 3;
5, 4, 2;
1, 3, 5;
3, 1, 2, 4;
5, 7.
Then Vasya showed the pieces of paper to his friends, but kept the n sets secret from them. His friends managed to calculate which sets Vasya had thought of in the first place. And how about you, can you restore the sets by the given pieces of paper?


### ideas
1. 对于set x，它出现的次数 = n-1 (只会是n-1， 不可能更多)
2. 然后以2为第一个字符，和在一起出现了n-1次的，就是一个set
3. 依次处理就好了