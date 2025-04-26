The Smart Beaver from ABBYY invented a new message encryption method and now wants to check its performance. Checking it manually is long and tiresome, so he decided to ask the ABBYY Cup contestants for help.

A message is a sequence of n integers a1, a2, ..., an. Encryption uses a key which is a sequence of m integers b1, b2, ..., bm (m ≤ n). All numbers from the message and from the key belong to the interval from 0 to c - 1, inclusive, and all the calculations are performed modulo c.

Encryption is performed in n - m + 1 steps. On the first step we add to each number a1, a2, ..., am a corresponding number b1, b2, ..., bm. On the second step we add to each number a2, a3, ..., am + 1 (changed on the previous step) a corresponding number b1, b2, ..., bm. And so on: on step number i we add to each number ai, ai + 1, ..., ai + m - 1 a corresponding number b1, b2, ..., bm. The result of the encryption is the sequence a1, a2, ..., an after n - m + 1 steps.



### ideas
1. 1 2 3 4 5
2. 1 2 3 4
3.   1 2 3 4
4. 1 2 3 4 5 
5. 1 2
6.   1 2
7.     1 2
8.       1 2    