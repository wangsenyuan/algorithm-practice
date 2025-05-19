For an array 𝑎
 of integers let's denote its maximal element as max(𝑎)
, and minimal as min(𝑎)
. We will call an array 𝑎
 of 𝑘
 integers interesting if max(𝑎)−min(𝑎)≥𝑘
. For example, array [1,3,4,3]
 isn't interesting as max(𝑎)−min(𝑎)=4−1=3<4
 while array [7,3,0,4,3]
 is as max(𝑎)−min(𝑎)=7−0=7≥5
.

You are given an array 𝑎
 of 𝑛
 integers. Find some interesting nonempty subarray of 𝑎
, or tell that it doesn't exist.

An array 𝑏
 is a subarray of an array 𝑎
 if 𝑏
 can be obtained from 𝑎
 by deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end. In particular, an array is a subarray of itself.