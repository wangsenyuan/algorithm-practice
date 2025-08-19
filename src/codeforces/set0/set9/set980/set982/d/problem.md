# Problem D: Shark Locations

## Problem Description

For a long time, scientists have been studying the behavior of sharks. Sharks, like many other species, alternate between short movements within a certain location and long movements between locations.

Max is a young biologist. For $n$ days, he watched a specific shark and recorded the distance the shark traveled each day. All the distances are distinct. Max now wants to determine how many locations the shark visited.

He made the following assumption: there exists an integer $k$ such that:
- If the shark traveled a distance **strictly less than** $k$ on a given day, then it did not change its location
- If the shark traveled a distance **greater than or equal to** $k$ on a given day, then it was changing its location on that day

Note that it is possible for the shark to change locations on several consecutive days, as long as the distance traveled each day is at least $k$.

The shark never returns to the same location after moving away from it. Therefore, in the sequence of $n$ days, we can find consecutive non-empty segments where the shark traveled distances less than $k$ each day. Each such segment corresponds to one location.

Max wants to choose $k$ such that the lengths of all such segments are equal.

## Task

Find an integer $k$ that maximizes the number of locations while ensuring that:
1. The shark was in each location for the same number of days
2. The number of locations is maximum possible satisfying the first condition
3. $k$ is the smallest possible value satisfying the first and second conditions

## Input

- **First line**: A single integer $n$ ($1 \leq n \leq 10^5$) — the number of days
- **Second line**: $n$ distinct positive integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^9$) — the distance traveled each day

## Output

Print a single integer $k$ that satisfies all the conditions above.

## Examples

### Example 1

**Input:**
```
8
1 2 7 3 4 8 5 6
```

**Output:**
```
7
```

**Explanation:**
The shark travels within a location on days 1 and 2 (first location), then on days 4 and 5 (second location), then on days 7 and 8 (third location). There are three locations in total.

### Example 2

**Input:**
```
6
25 1 2 3 14 36
```

**Output:**
```
2
```

**Explanation:**
The shark only moves within a location on the 2nd day, so there is only one location.

### ideas
1. 对于给定的k，它分割出来的 < k 的段，每段的大小要一致
2. 