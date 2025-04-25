Hackers are once again trying to create entertaining phrases using the output of neural networks. This time, they want to obtain an array of strings 𝑎
 of length 𝑛
.

Initially, they have an array 𝑐
 of length 𝑛
, filled with blanks, which are denoted by the symbol ∗
. Thus, if 𝑛=4
, then initially 𝑐=[∗,∗,∗,∗]
.

The hackers have access to 𝑚
 neural networks, each of which has its own version of the answer to their request – an array of strings 𝑏𝑖
 of length 𝑛
.

The hackers are trying to obtain the array 𝑎
 from the array 𝑐
 using the following operations:

Choose a neural network 𝑖
, which will perform the next operation on the array 𝑐
: it will select a random blank, for example, at position 𝑗
, and replace 𝑐𝑗
 with 𝑏𝑖,𝑗
.
For example, if the first neural network is chosen and 𝑐=[∗,«like»,∗]
, and 𝑏1=[«I»,«love»,«apples»]
, then after the operation with the first neural network, 𝑐
 may become either [«I»,«like»,∗]
 or [∗,«like»,«apples»]
.

Choose position 𝑗
 and replace 𝑐𝑗
 with a blank.
Unfortunately, because of the way hackers access neural networks, they will only be able to see the modified array 𝑐
 after all operations are completed, so they will have to specify the entire sequence of operations in advance.

However, the random behavior of the neural networks may lead to the situation where the desired array is never obtained, or obtaining it requires an excessive number of operations.

Therefore, the hackers are counting on your help in choosing a sequence of operations that will guarantee the acquisition of array 𝑎
 in the minimum number of operations.

More formally, if there exists a sequence of operations that can guarantee obtaining array 𝑎
 from array 𝑐
, then among all such sequences, find the one with the minimum number of operations, and output the number of operations in it.

If there is no sequence of operations that transforms array 𝑐
 into array 𝑎
, then output −1
.
