In an Italian village, a hungry mouse starts at vertex st
 on a given tree∗
 with 𝑛
 vertices.

Given a permutation 𝑝
 of length 𝑛
†
, there are 𝑛
 steps. For the 𝑖
-th step:

A tempting piece of Parmesan cheese appears at vertex 𝑝𝑖
. If the mouse is currently at vertex 𝑝𝑖
, it will stay there and enjoy the moment. Otherwise, it will move along the simple path to vertex 𝑝𝑖
 by one edge.
Your task is to find such a permutation so that, after all 𝑛
 steps, the mouse inevitably arrives at vertex en
, where a trap awaits.

Note that the mouse must arrive at en
 after all 𝑛
 steps, though it may pass through en
 earlier during the process.

### ideas
1. 那么有没有可能没有答案呢？如果n是偶数，但是 st -> en 的距离是奇数，那么没有答案
2. 可以证明，假设在经过x步， 那么可以证明x必须是奇数步（经过奇数的数字）
3. 这是因为，没经过一步，st->en的距离的奇偶性会改变，不管是远离en，还是靠近en
4. 但是偶数个n，mouse必须走n步，是偶数次，肯定是不行的
5. 但是、但是，mouse是可以在某一步停在一个位置（似乎就能改变奇偶性了）
6. 特别的，如果在最后一步的前面，Jerry就正好在en，如果最后一个数字就是en，那么就可以做到？
7. 所以，如果st->en的奇偶性不一样，必须有（至少有一个位置，jerry正好停在那里）
8. 所以还有点复杂
9. 考虑st = en的情况（其他的情况可以归结到这一种情况）
10. 假设剩余的操作数是偶数（还剩偶数个数）
11. 如果有一个子树的大小超过了一半，那么就没有答案了。
12. 否则可以使用贪心的方式，先进入最大的子树，然后再选择第二大的子树， and so on
13. 如果是奇数个，那么最后一个最好选择en
14. 然后考虑怎么到达st -> en
15. 如果st != en, 那么就用这个路径一直访问到en
16. 如果en是叶子节点，也是有问题的，到达en时，就剩一个子节点了。如果没有用完，就回不去了
17. 假设st是中点（没有子节点的size超过1半）那么是否可以在st这里消耗到，可以一次性到达en的时候，再运动？
18. 所以这个思路不大对
19. 还是应该在st的附近基本上用完，然后st变成了一个叶子节点，然后移动到下一个节点
20. 改变奇偶性，靠en的位置
21. 这个题目不止1700分呐～