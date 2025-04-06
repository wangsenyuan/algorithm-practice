Let 𝑡𝑖=𝑘1+𝑘2+...+𝑘𝑖−1
, 𝑏𝑖=𝑎𝑖−𝑡𝑖
.

We can rewrite the condition 𝑎𝑖+1>=𝑎𝑖+𝑘𝑖
, using array 𝑏
:

𝑎𝑖+1>=𝑎𝑖+𝑘𝑖

𝑎𝑖+1−𝑘𝑖>=𝑎𝑖

𝑎𝑖+1−𝑘𝑖−𝑘𝑖−1−...−𝑘1>=𝑎𝑖−𝑘𝑖−1−...−𝑘1

𝑎𝑖+1−𝑡𝑖+1>=𝑎𝑖−𝑡𝑖

𝑏𝑖+1>=𝑏𝑖

Let's calculate arrays 𝑡
 and 𝑏
.

So as 𝑎𝑖=𝑏𝑖+𝑡𝑖
, in order to get sum in subarray of 𝑎
, we can sum corresponding sums in 𝑏
 and 𝑡
.

Now let's find out what happens with 𝑏
 after addition 𝑥
 to position 𝑖
. 𝑏𝑖
 increases exactly on 𝑥
. Then, if 𝑏𝑖+1<𝑏𝑖
, 𝑏𝑖+1
 becomes equal to 𝑏𝑖
, and so on for 𝑖+2
, 𝑖+3
, ..., 𝑛
. Note that array 𝑏
 is always sorted and each addition sets value 𝑏𝑖+𝑥
 in half-interval [𝑖,𝑝𝑜𝑠)
, where 𝑝𝑜𝑠
 - the lowest index such as 𝑏𝑝𝑜𝑠>=𝑏𝑖+𝑥

To handle these modifications, let's build segment tree on array 𝑏
 with operation "set value on a segment", which stores sum and maximum in every vertex. The only problem is how to find 𝑝𝑜𝑠
. This can be done with descending along the segment tree. If the maximum in the left son of current vertex is bigger or equal that 𝑏𝑖+𝑥
, we go to the left son, otherwise we go the right son.

BONUS: solve it with modifications of elements of 𝑘
.