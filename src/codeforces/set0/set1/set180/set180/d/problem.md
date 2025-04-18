Everything got unclear to us in a far away constellation Tau Ceti. Specifically, the Taucetians choose names to their children in a very peculiar manner.

Two young parents abac and bbad think what name to give to their first-born child. They decided that the name will be the permutation of letters of string s. To keep up with the neighbours, they decided to call the baby so that the name was lexicographically strictly larger than the neighbour's son's name t.

On the other hand, they suspect that a name tax will be introduced shortly. According to it, the Taucetians with lexicographically larger names will pay larger taxes. That's the reason abac and bbad want to call the newborn so that the name was lexicographically strictly larger than name t and lexicographically minimum at that.

The lexicographical order of strings is the order we are all used to, the "dictionary" order. Such comparison is used in all modern programming languages to compare strings. Formally, a string p of length n is lexicographically less than string q of length m, if one of the two statements is correct:

n < m, and p is the beginning (prefix) of string q (for example, "aba" is less than string "abaa"),
p1 = q1, p2 = q2, ..., pk - 1 = qk - 1, pk < qk for some k (1 ≤ k ≤ min(n, m)), here characters in strings are numbered starting from 1.
Write a program that, given string s and the heighbours' child's name t determines the string that is the result of permutation of letters in s. The string should be lexicographically strictly more than t and also, lexicographically minimum.

### ideas
1. 从后往前处理，如果能够在位置i处，放置一个比t[i]大的字符，那么后面的部分，可以放置剩余的字符组成的最小字符
2. 但是，同时，必须保证i的前面，能够使用s中的字符组成出来