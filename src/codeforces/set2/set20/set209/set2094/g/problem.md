Chimpanzini Bananini stands on the brink of a momentous battle—one destined to bring finality.

For an arbitrary array 𝑏
 of length 𝑚
, let's denote the rizziness of the array to be ∑𝑚𝑖=1𝑏𝑖⋅𝑖=𝑏1⋅1+𝑏2⋅2+𝑏3⋅3+…+𝑏𝑚⋅𝑚
.

Chimpanzini Bananini gifts you an empty array. There are three types of operations you can perform on it.

Perform a cyclic shift on the array. That is, the array [𝑎1,𝑎2,…,𝑎𝑛]
 becomes [𝑎𝑛,𝑎1,𝑎2,…,𝑎𝑛−1].
Reverse the entire array. That is, the array [𝑎1,𝑎2,…,𝑎𝑛]
 becomes [𝑎𝑛,𝑎𝑛−1,…,𝑎1].
Append an element to the end of the array. The array [𝑎1,𝑎2,…,𝑎𝑛]
 becomes [𝑎1,𝑎2,…,𝑎𝑛,𝑘]
 after appending 𝑘
 to the end of the array.
After each operation, you are interested in calculating the rizziness of your array.

Note that all operations are persistent. This means that each operation modifies the array, and subsequent operations should be applied to the current state of the array after the previous operations.

Input
The first line contains an integer 𝑡
 (1≤𝑡≤104
) — the number of test cases.

The first line of the input contains an integer 𝑞
 (1≤𝑞≤2⋅105
) — the number of operations you perform on your array.

The following 𝑞
 lines first contain a single integer 𝑠
 (1≤𝑠≤3
) — the operation type.

If 𝑠=1
, then the cyclic shift operation should be performed.
If 𝑠=2
, then the reversal operation should be performed.
If 𝑠=3
, then the line will contain an additional integer 𝑘
 (1≤𝑘≤106
), denoting the element appended to the back of the array.
It is guaranteed that the sum of 𝑞
 will not exceed 2⋅105
 over all test cases. Additionally, it is guaranteed that the first operation on each test case will be one with 𝑠=3
.

### ideas
1. 操作1的结果, x = a1 * 1 + a2 * 2 + ... an * n
2.            y = an * 1 + a1 * 2, a2 * 3 + ... + a(n-1) * n
3.            x - y = -an - a1 - a2 + ... - a(n-1) + an * n
4.            那么知道x，就可以很容易的计算出y
5.  操作2的结果, x = ...
6.             y = an * 1 + a(n-1) * 2 + ... + a2 * (n - 1) + a1 * n
7.             x + y = (a1 + an) * 1 + (a2 + a(n-1)) * 2 + ... + (a1 + an) * n
8.             x - y = 
9.    操作2，可以直接算出来，然后反转一下
10. 假设每次都维护{x, y}, x = 从前完后， y = 从后往前
11. 操作2，就变成 {y, x}, 
12. 操作1，同时更新 {x, y} (还需要知道s)
13. 操作3，需要更新 {x, y}, s, 还有头尾