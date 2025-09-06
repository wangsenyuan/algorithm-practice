# Problem E: Mushroom Forest

Ralph is going to collect mushrooms in the Mushroom Forest.

There are m directed paths connecting n trees in the Mushroom Forest. On each path grow some mushrooms. When Ralph passes a path, he collects all the mushrooms on the path. The Mushroom Forest has a magical fertile ground where mushrooms grow at a fantastic speed. New mushrooms regrow as soon as Ralph finishes mushroom collection on a path. More specifically, after Ralph passes a path the i-th time, there regrow i mushrooms less than there was before this pass. That is, if there is initially x mushrooms on a path, then Ralph will collect x mushrooms for the first time, x - 1 mushrooms the second time, x - 1 - 2 mushrooms the third time, and so on. However, the number of mushrooms can never be less than 0.

For example, let there be 9 mushrooms on a path initially. The number of mushrooms that can be collected from the path is 9, 8, 6 and 3 when Ralph passes by from first to fourth time. From the fifth time and later Ralph can't collect any mushrooms from the path (but still can pass it).

Ralph decided to start from the tree s. How many mushrooms can he collect using only described paths?

## Input

The first line contains two integers n and m (1 ≤ n ≤ 10^6, 0 ≤ m ≤ 10^6), representing the number of trees and the number of directed paths in the Mushroom Forest, respectively.

Each of the following m lines contains three integers x, y and w (1 ≤ x, y ≤ n, 0 ≤ w ≤ 10^8), denoting a path that leads from tree x to tree y with w mushrooms initially. There can be paths that lead from a tree to itself, and multiple paths between the same pair of trees.

The last line contains a single integer s (1 ≤ s ≤ n) — the starting position of Ralph.

## Output

Print an integer denoting the maximum number of the mushrooms Ralph can collect during his route.

## Examples

### Example 1
**Input:**
```
2 2
1 2 4
2 1 4
1
```

**Output:**
```
16
```

### Example 2
**Input:**
```
3 3
1 2 4
2 3 3
1 3 8
1
```

**Output:**
```
8
```

## Note

In the first sample Ralph can pass three times on the circle and collect 4 + 4 + 3 + 3 + 1 + 1 = 16 mushrooms. After that there will be no mushrooms for Ralph to collect.

In the second sample, Ralph can go to tree 3 and collect 8 mushrooms on the path from tree 1 to tree 3.


### ideas
1. 这个是有一个有向边组成的图；
2. 考虑最简单的情况，没有环的情况，那么就是这些边的sum
3. 考虑环的情况，那么就可以取值到0为止
4. 所以，可以用新的dag去处理