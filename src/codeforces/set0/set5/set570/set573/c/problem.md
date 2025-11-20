# Problem C: Bear and Drawing

Limak is a little bear who learns to draw. People usually start with houses, fences and flowers but why would bears do it? Limak lives in the forest and he decides to draw a tree.

Recall that tree is a connected graph consisting of n vertices and n - 1 edges.

Limak chose a tree with n vertices. He has infinite strip of paper with two parallel rows of dots. Little bear wants to assign vertices of a tree to some n distinct dots on a paper so that edges would intersect only at their endpoints — drawn tree must be planar. Below you can see one of correct drawings for the first sample test.

Is it possible for Limak to draw chosen tree?

## Input

The first line contains single integer n (1 ≤ n ≤ 10^5).

Next n - 1 lines contain description of a tree. i-th of them contains two space-separated integers ai and bi (1 ≤ ai, bi ≤ n, ai ≠ bi) denoting an edge between vertices ai and bi. It's guaranteed that given description forms a tree.

## Output

Print "Yes" (without the quotes) if Limak can draw chosen tree. Otherwise, print "No" (without the quotes).

## Examples

### Example 1

**Input:**

```text
8
1 2
1 3
1 6
6 4
6 7
6 5
7 8
```

**Output:**

```text
Yes
```

### Example 2

**Input:**

```text
13
1 2
1 3
1 4
2 5
2 6
2 7
3 8
3 9
3 10
4 11
4 12
4 13
```

**Output:**

```text
No
```


### ideas
1. 大概知道这个问题的意思
2. 有两排dots， 然后给这些dots分配标号， 目标是用这些标号（根据edges）去连接
3. 连接后，不出现交叉的edge
4. 如果第二排都是叶子节点， 第一排都是主干节点
5. 叶子节点, 或者离叶子节点一层的，deg <= 3的
6. 其他的都必须是主干节点
7. 主干节点必须依次相连；
8. 好像是ok的