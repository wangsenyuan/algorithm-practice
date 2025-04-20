Happy PMP is freshman and he is learning about algorithmic problems. He enjoys playing algorithmic games a lot.

One of the seniors gave Happy PMP a nice game. He is given two permutations of numbers 1 through n and is asked to convert the first one to the second. In one move he can remove the last number from the permutation of numbers and inserts it back in an arbitrary position. He can either insert last number between any two consecutive numbers, or he can place it at the beginning of the permutation.

Happy PMP has an algorithm that solves the problem. But it is not fast enough. He wants to know the minimum number of moves to convert the first permutation to the second.

### ideas
1. 考虑s[:i]和t[:i]相同
2. 目前要处理最后一个位置，似乎就应该是放置位置i处？
3. 假设和t[i]匹配的数是x, 那么要把x拿出来，必须把x后面的数移出来
4. 移出来的数，一定能排列好吗？
5. 它们的相对顺序肯定是可以保证的
6. 有点想法了