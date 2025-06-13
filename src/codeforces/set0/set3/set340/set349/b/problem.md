Igor has fallen in love with Tanya. Now Igor wants to show his feelings and write a number on the fence opposite to Tanya's house. Igor thinks that the larger the number is, the more chance to win Tanya's heart he has.

Unfortunately, Igor could only get v liters of paint. He did the math and concluded that digit d requires ad liters of paint. Besides, Igor heard that Tanya doesn't like zeroes. That's why Igor won't use them in his number.

Help Igor find the maximum number he can write on the fence.

Input
The first line contains a positive integer v (0 ≤ v ≤ 106). The second line contains nine positive integers a1, a2, ..., a9 (1 ≤ ai ≤ 105).

Output
Print the maximum number Igor can write on the fence. If he has too little paint for any digit (so, he cannot write anything), print -1.

### ideas
1. 数字越长越大
2. 先考虑最多可以得到多长的字符
3. v / min(a) 这个是最长的
4. 然后在位置1，使用大的数字，如果可以，就使用它，不可以，就继续下一个