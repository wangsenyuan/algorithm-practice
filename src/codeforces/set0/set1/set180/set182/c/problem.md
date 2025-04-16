And here goes another problem on arrays. You are given positive integer len and array a which consists of n integers a1, a2, ..., an. Let's introduce two characteristics for the given array.

Let's consider an arbitrary interval of the array with length len, starting in position i. Value , is the modular sum on the chosen interval. In other words, the modular sum is the sum of integers on the chosen interval with length len, taken in its absolute value.
Value  is the optimal sum of the array. In other words, the optimal sum of an array is the maximum of all modular sums on various intervals of array with length len.
Your task is to calculate the optimal sum of the given array a. However, before you do the calculations, you are allowed to produce no more than k consecutive operations of the following form with this array: one operation means taking an arbitrary number from array ai and multiply it by -1. In other words, no more than k times you are allowed to take an arbitrary number ai from the array and replace it with  - ai. Each number of the array is allowed to choose an arbitrary number of times.

Your task is to calculate the maximum possible optimal sum of the array after at most k operations described above are completed.

### ideas
1. 对于一个确定的区间l...r （长度为len）
2. 将这个区间按照从小到大排序，假设它头部有x个负数，尾部有y个正数
3. 那么最优的操作，就是把这(少于或等于)x个负数变成正数, 把尾部（少于或等于）y个正数变成负数
4. 剩下就是怎么样的数据结构，可以快速的计算前x（和后y）个数的和
5. 正数、负数分别用两个pq来保存？不大行，没法快速的计算前x个数
6. 用segment tree？