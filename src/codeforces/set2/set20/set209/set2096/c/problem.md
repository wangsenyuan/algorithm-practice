You are the proud leader of a city in Ancient Berland. There are 𝑛2
 buildings arranged in a grid of 𝑛
 rows and 𝑛
 columns. The height of the building in row 𝑖
 and column 𝑗
 is ℎ𝑖,𝑗
.

The city is beautiful if no two adjacent by side buildings have the same height. In other words, it must satisfy the following:

There does not exist a position (𝑖,𝑗)
 (1≤𝑖≤𝑛
, 1≤𝑗≤𝑛−1
) such that ℎ𝑖,𝑗=ℎ𝑖,𝑗+1
.
There does not exist a position (𝑖,𝑗)
 (1≤𝑖≤𝑛−1
, 1≤𝑗≤𝑛
) such that ℎ𝑖,𝑗=ℎ𝑖+1,𝑗
.
There are 𝑛
 workers at company A, and 𝑛
 workers at company B. Each worker can be hired at most once.

It costs 𝑎𝑖
 coins to hire worker 𝑖
 at company A. After hiring, worker 𝑖
 will:

Increase the heights of all buildings in row 𝑖
 by 1
. In other words, increase ℎ𝑖,1,ℎ𝑖,2,…,ℎ𝑖,𝑛
 by 1
.
It costs 𝑏𝑗
 coins to hire worker 𝑗
 at company B. After hiring, worker 𝑗
 will:

Increase the heights of all buildings in column 𝑗
 by 1
. In other words, increase ℎ1,𝑗,ℎ2,𝑗,…,ℎ𝑛,𝑗
 by 1
.
Find the minimum number of coins needed to make the city beautiful, or report that it is impossible.