# Problem B: Pipeline

## Problem Description

Vova, the Ultimate Thule new shaman, wants to build a pipeline. As there are exactly $n$ houses in Ultimate Thule, Vova wants the city to have exactly $n$ pipes, each such pipe should be connected to the water supply. A pipe can be connected to the water supply if there's water flowing out of it. Initially Vova has only one pipe with flowing water. Besides, Vova has several splitters.

A **splitter** is a construction that consists of:
- One input (it can be connected to a water pipe)
- $x$ output pipes

When a splitter is connected to a water pipe, water flows from each output pipe. You can assume that the output pipes are ordinary pipes. For example, you can connect water supply to such pipe if there's water flowing out from it. **At most one splitter can be connected to any water pipe.**

> **Note:** The figure shows a 4-output splitter

Vova has one splitter of each kind: with 2, 3, 4, ..., $k$ outputs. Help Vova use the **minimum number of splitters** to build the required pipeline or otherwise state that it's impossible.

**Goal:** Vova needs the pipeline to have exactly $n$ pipes with flowing out water. Note that some of those pipes can be the output pipes of the splitters.

## Input

The first line contains two space-separated integers $n$ and $k$:
- $1 \leq n \leq 10^{18}$
- $2 \leq k \leq 10^9$

> **Important:** Please, do not use the `%lld` specifier to read or write 64-bit integers in C++. It is preferred to use the `cin`, `cout` streams or the `%I64d` specifier.

## Output

Print a single integer — the **minimum number of splitters needed** to build the pipeline. If it is impossible to build a pipeline with the given splitters, print `-1`.

## Examples

### Example 1
**Input:**
```
4 3
```

**Output:**
```
2
```

### Example 2
**Input:**
```
5 5
```

**Output:**
```
1
```

### Example 3
**Input:**
```
8 4
```

**Output:**
```
-1
```


### ideas
1. 假设现在共有w个出口，如果使用一个x口的splitter，那么就是增加 x - 1 个出口
2. 也就是要找到 x1 - 1 + x2 - 1 + ... + xi - 1 = n
3. x1 + x2 + ... xi - i = n
4. x1 + ... + xi = n + i
5. xi >= 2, and xi <= k
6.  假设最后一个使用了一个k头的（但是产生的头多了w个）
7.  那么这个时候，需要把这多出来的w个给合并掉，那么只能去合并w+1 => 1
8.  也就是说把那个w+2的头给舍弃掉，那么必须有 w+2 < n, 且w+2之前被使用到了
9.  不过倒是，尽量使用k是不是更优呢？
10. 假设使用k，那么目前的出口 = k个
11. 