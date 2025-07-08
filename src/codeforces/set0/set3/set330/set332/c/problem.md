# Problem C: Student Revenge

## Problem Description

A student's life is fraught with complications. Some Berland University students know this only too well. Having studied for two years, they contracted strong antipathy towards the chairperson of some department. Indeed, the person in question wasn't the kindest of ladies to begin with: prone to reforming groups, banning automatic passes and other mean deeds. At last the students decided that she just can't get away with all this anymore...

The students pulled some strings on the higher levels and learned that the next University directors' meeting is going to discuss **n** orders about the chairperson and accept exactly **p** of them. There are two values assigned to each order:
- **ai**: the number of the chairperson's hairs that turn grey if she obeys the order
- **bi**: the displeasement of the directors if the order isn't obeyed

The students may make the directors pass any **p** orders chosen by them. The students know that the chairperson will obey exactly **k** out of these **p** orders. She will pick the orders to obey in the way that minimizes:
1. First, the directors' displeasement
2. Second, the number of hairs on her head that turn grey

The students want to choose **p** orders in the way that maximizes the number of hairs on the chairperson's head that turn grey. If there are multiple ways to accept the orders, then the students are keen on maximizing the directors' displeasement with the chairperson's actions. Help them.

## Input

The first line contains three integers:
- **n** (1 ≤ n ≤ 10^5): the number of orders the directors are going to discuss
- **p** (1 ≤ p ≤ n): the number of orders to pass
- **k** (1 ≤ k ≤ p): the number of orders to be obeyed by the chairperson

Each of the following **n** lines contains two integers **ai** and **bi** (1 ≤ ai, bi ≤ 10^9), describing the corresponding order.

## Output

Print in an arbitrary order **p** distinct integers — the numbers of the orders to accept so that the students could carry out the revenge. The orders are indexed from 1 to n in the order they occur in the input. If there are multiple solutions, you can print any of them.

## Examples

### Example 1

**Input:**
```
5 3 2
5 6
5 8
1 3
4 3
4 11
```

**Output:**
```
3 1 2
```

**Explanation:** One of the optimal solutions is to pass orders 1, 2, 3. In this case the chairperson obeys orders number 1 and 2. She gets 10 new grey hairs in the head and the directors' displeasement will equal 3. Note that the same result can be achieved with order 4 instead of order 3.

### Example 2

**Input:**
```
5 3 3
10 18
18 17
10 20
20 18
20 18
```

**Output:**
```
2 4 5
```

**Explanation:** The chairperson can obey all the orders, so the best strategy for the students is to pick the orders with the maximum sum of ai values. The chairperson gets 58 new gray hairs and the directors' displeasement will equal 0.


### ideas
1. 题目好绕，提供p个选择，其中p-k个不会被选中，首先要最小化,这个 p - k 个不被选中的b的sum，也就是说会对这p个进行升序排列，选择后k个
2. 当b相等时，按照a倒序排列
3. 按照b升序排列，（相同时按照a降序排列）
4. 前 p - k 个选中后，在后面部分选择k个a最大的
5. 应该是固定后k个，然后选择p-k个
6. 要在后面选择k个（一定要选到），然后前面选择 p - k 个，最大化b的值
7. 后面选择k个最大化