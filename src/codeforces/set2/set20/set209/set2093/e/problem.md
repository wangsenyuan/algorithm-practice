You are given an array 𝑎
of length 𝑛
and a number 𝑘
.

A subarray is defined as a sequence of one or more consecutive elements of the array. You need to split the array 𝑎
into 𝑘
non-overlapping subarrays 𝑏1,𝑏2,…,𝑏𝑘
such that the union of these subarrays equals the entire array. Additionally, you need to maximize the value of 𝑥
, which is equal to the minimum MEX(𝑏𝑖)
, for 𝑖∈[1..𝑘]
.

MEX(𝑣)
denotes the smallest non-negative integer that is not present in the array 𝑣
.