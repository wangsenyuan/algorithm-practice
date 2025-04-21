The beauty of an array 𝑏
 of length 𝑚
 is defined as max(𝑏𝑖⊕𝑏𝑗)
 among all possible pairs 1≤𝑖≤𝑗≤𝑚
, where 𝑥⊕𝑦
 is the bitwise XOR of numbers 𝑥
 and 𝑦
. We denote the beauty value of the array 𝑏
 as 𝑓(𝑏)
.

An array 𝑏
 is called beautiful if 𝑓(𝑏)≥𝑘
.

Recently, Kostya bought an array 𝑎
 of length 𝑛
 from the store. He considers this array too long, so he plans to cut out some beautiful subarray from it. That is, he wants to choose numbers 𝑙
 and 𝑟
 (1≤𝑙≤𝑟≤𝑛
) such that the array 𝑎𝑙…𝑟
 is beautiful. The length of such a subarray will be the number 𝑟−𝑙+1
. The entire array 𝑎
 is also considered a subarray (with 𝑙=1
 and 𝑟=𝑛
).

Your task is to find the length of the shortest beautiful subarray in the array 𝑎
. If no subarray is beautiful, you should output the number −1
.

Input
The first line contains the number of test cases 𝑡
 (1≤𝑡≤104
).

Next, there are 𝑡
 blocks of two lines:

In the first line of the block, there are two integers 𝑛
 and 𝑘
 (1≤𝑛≤2⋅105
, 0≤𝑘≤109
).

In the second line of the block, there is the array 𝑎
 consisting of 𝑛
 integers (0≤𝑎𝑖≤109
).

It is guaranteed that the sum of 𝑛
 across all tests does not exceed 2⋅105
.

Output
For each test case, you need to output a single integer — the minimum length of the segment (𝑙,𝑟)
 for which 𝑓(𝑎𝑙…𝑟)≥𝑘
. If such a segment is not found, you should output −1
.