# Codeforces 387C - George and Number

## Problem Description

George is a cat, so he really likes to play. Most of all he likes to play with his array of positive integers b. During the game, George modifies the array by using special changes. 

Let's mark George's current array as b1, b2, ..., b|b| (record |b| denotes the current length of the array). Then one change is a sequence of actions:

1. Choose two distinct indexes i and j (1 ≤ i, j ≤ |b|; i ≠ j), such that bi ≥ bj.
2. Get number v = concat(bi, bj), where concat(x, y) is a number obtained by adding number y to the end of the decimal record of number x. For example, concat(500, 10) = 50010, concat(2, 2) = 22.
3. Add number v to the end of the array. The length of the array will increase by one.
4. Remove from the array numbers with indexes i and j. The length of the array will decrease by two, and elements of the array will become re-numbered from 1 to current length of the array.

George played for a long time with his array b and received from array b an array consisting of exactly one number p. Now George wants to know: what is the maximum number of elements array b could contain originally? 

Help him find this number. Note that originally the array could contain only positive integers.

## Input

The first line of the input contains a single integer p (1 ≤ p < 10^100000). It is guaranteed that number p doesn't contain any leading zeroes.

## Output

Print an integer — the maximum number of elements array b could contain originally.

## ideas
1. 考虑最后的数p, 它肯定是两个数 a ++ b, 且 a >= b, 
2. 如果a/b很接近，更优？还是b很小更优呢？
3. 如果b > 0, 似乎把b作为一个单独的数更好