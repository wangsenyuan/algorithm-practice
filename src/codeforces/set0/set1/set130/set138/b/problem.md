# Andrey and Numbers

Andrey's favourite number is $n$. Andrey's friends gave him two identical numbers $n$ as a New Year present. He hung them on a wall and watched them adoringly.

Then Andrey got bored from looking at the same number and he started to swap digits first in one, then in the other number, then again in the first number and so on (arbitrary number of changes could be made in each number). At some point it turned out that if we sum the resulting numbers, then the number of zeroes with which the sum will end would be maximum among the possible variants of digit permutations in those numbers.

Given number $n$, can you find the two digit permutations that have this property?

## Input
The first line contains a positive integer $n$ — the original number. The number of digits in this number does not exceed $10^5$. The number is written without any leading zeroes.

## Output
Print two permutations of digits of number $n$, such that the sum of these numbers ends with the maximum number of zeroes. The permutations can have leading zeroes (if they are present, they all should be printed). The permutations do not have to be different. If there are several answers, print any of them.

## ideas
1. 尽量的从后端开始处理（因为后端的会产生进位）
2. 最后安排0（因为0不产生进位，且在有进位的情况下0 + 9 + 1 = 10， 似乎更有利)
3. 也不大对，如果产生足够多的0，那么这些0，肯定能全部移动到尾部去
4. 假设是有0 + 0, 产生的，那么这些，全部移动到最后面，不影响结果
5. 如果是有a+b产生的进位，那么就把它们联动在一起，也是ok的
6. 但是存在一种情况，第一个进位是有6 + 6产生的？
7. 比如说，不存在4的情况下，那么用6+6产生进位，就可以使的 3 + 6变成0
8. 还很麻烦么。
9. 这个数字是可以产生多段的。每段的开始都是 a + b = 10的情况，然后所有进位的情况都可以算在一段里
10. 然后是剩下所有不进位的情况，比如0/0, 1/9, 2/8(这些只能单独开一段)
11. 假设从尾部开始产生了一串的0（通过进位得到的）然后，要使用一次不进位，然后0，然后不进位，然后0
12. 不进位的是[0...4]这些组合（但是它们如果能在进位种被使用掉，似乎也没问题有，因为也就是1换1）
13. 只有0/0似乎是特殊的(0, 0)似乎直接配对更好（因为和9配对，也就是产生1，和0配对也是1）
14. 