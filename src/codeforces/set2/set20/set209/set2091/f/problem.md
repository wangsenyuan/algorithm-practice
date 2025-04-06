The visitors of the IT Campus "NEIMARK" are not only strong programmers but also physically robust individuals! Some practice swimming, some rowing, and some rock climbing!

Master Igor is a prominent figure in the local rock climbing community. One day, he went on a mountain hike to ascend one of the peaks. As an experienced climber, Igor decided not to follow the established trails but to use his skills to climb strictly vertically.

Igor found a rectangular vertical section of the mountain and mentally divided it into 𝑛
 horizontal levels. He then split each level into 𝑚
 segments using vertical partitions. Upon inspecting these segments, Igor discovered convenient protrusions that can be grasped (hereafter referred to as holds). Thus, the selected part of the mountain can be represented as an 𝑛×𝑚
 rectangle, with some cells containing holds.

Being an experienced programmer, Igor decided to count the number of valid routes. A route is defined as a sequence of distinct holds. A route is considered valid if the following conditions are satisfied:

The first hold in the route is located on the very bottom level (row 𝑛
);
The last hold in the route is located on the very top level (row 1
);
Each subsequent hold is not lower than the previous one;
At least one hold is used on each level (i.e., in every row of the rectangle);
At most two holds are used on each level (since Igor has only two hands);
Igor can reach from the current hold to the next one if the distance between the centers of the corresponding sections does not exceed Igor's arm span.
Igor's arm span is 𝑑
, which means he can move from one hold to another if the Euclidean distance between the centers of the corresponding segments does not exceed 𝑑
. The distance between sections (𝑖1,𝑗1
) and (𝑖2,𝑗2
) is given by (𝑖1−𝑖2)2+(𝑗1−𝑗2)2‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾√
.

Calculate the number of different valid routes. Two routes are considered different if they differ in the list of holds used or in the order in which these holds are visited.

### ideas
1. d * d >= (ax - bx) * (ax - bx) + (ay - by) * (ay - by)
2. 每行最多使用两个hold，最少使用一个hold
3. 假设对于每一行dp[i][j] 表示第i行，第一次到达第j列时的计数
4. 然后在第i行，从j可以到达它前后，距离d的hold
5. 这样的计算完后，fp[i][j]表示第i行，且挺在第j列的情况
6. 然后计算dp[i+1][?], 这时候是个fp的访问查询， 可以使用双指针
