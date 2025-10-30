# Problem B: Free Market

John Doe has recently found a "Free Market" in his city — that is the place where you can exchange some of your possessions for other things for free.

John knows that his city has n items in total (each item is unique). You can bring any number of items to the market and exchange them for any other one. Note that each item is one of a kind and that means that you cannot exchange set {a, b} for set {v, a}. However, you can always exchange set x for any set y, unless there is item p, such that p occurs in x and p occurs in y.

For each item, John knows its value ci. John's sense of justice doesn't let him exchange a set of items x for a set of items y, if s(x) + d < s(y) (s(x) is the total price of items in the set x).

During one day John can exchange only one set of items for something else. Initially, he has no items. John wants to get a set of items with the maximum total price. Find the cost of such set and the minimum number of days John can get it in.

## Input

The first line contains two space-separated integers n, d (1 ≤ n ≤ 50, 1 ≤ d ≤ 10^4) — the number of items on the market and John's sense of justice value, correspondingly. The second line contains n space-separated integers ci (1 ≤ ci ≤ 10^4).

## Output

Print two space-separated integers: the maximum possible price in the set of items John can get and the minimum number of days needed to get such set.

## Examples

### Example 1

**Input:**

```text
3 2
1 3 10
```

**Output:**

```text
4 3
```

### Example 2

**Input:**

```text
3 5
1 2 3
```

**Output:**

```text
6 2
```

### Example 3

**Input:**

```text
10 10000
10000 9999 1 10000 10000 10000 1 2 3 4
```

**Output:**

```text
50010 6
```

## Note

In the first sample John can act like this:

- Take the first item (1 - 0 ≤ 2).
- Exchange the first item for the second one (3 - 1 ≤ 2).
- Take the first item (1 - 0 ≤ 2).


### anaylysis by AI
Reading the full problem for context, then explaining that line.


[1 tool called]

## Exchange mechanism

- You can trade one set of items for another set.
- Sets don’t need to be the same size.

## Critical constraint: No overlapping items

> "you cannot exchange set {a, b} for set {v, a}"

This means the two sets cannot share any items. In this example, both sets contain item `a`, so this exchange is invalid.

- Invalid examples:
  - `{a, b}` → `{v, a}` (both contain `a`)
  - `{x, y, z}` → `{y, w}` (both contain `y`)
  - `{item1}` → `{item1, item2}` (both contain `item1`)

- Valid examples:
  - `{a, b}` → `{v, w}` (no overlap)
  - `{x, y}` → `{z}` (no overlap)
  - `{a}` → `{b, c, d}` (no overlap)

## Formal statement

> "you can always exchange set x for any set y, unless there is item p, such that p occurs in x and p occurs in y"

- You can exchange set `x` for set `y` if they have no items in common.
- You cannot exchange them if there exists any item `p` that appears in both sets.

## Why this constraint exists

It prevents reusing items in the exchange and makes the model consistent.

Example:
- If `{a}` → `{b, c}` were allowed with overlap, and `{b}` → `{a}` were also allowed, you could create items without giving anything up, which breaks the model.

## In the problem context

Given:
- You start with no items.
- You can take items from the market (exchange `{}` → `{item}`) if the fairness rule is satisfied.
- With overlap disallowed, each exchange is a true swap of distinct items.
- 

This sets up the problem: determine the maximum value set you can obtain and the minimum number of exchanges (days) needed.


### ideas
1. s(x) + d >= s(y) => s(y) - s(x) <= d
2. 也就是交换前后的价值，不能超过d
3. n <= 50
4. n如果更小，比如 n <= 20的时候，应该可以用bitmask dp
5. 但是50, 肯定是不行的～
6. dp[price] = mask ?
7. 不大对～
8. 假设要获得item i， 如果它的价值 <= d, 那么是可以直接获取的
9. 如果 > d, 那么可以看看，通过一路的交换能否获取（这样子可以计算出一条交换路径）
10. 在第一次交换的时候，肯定要使用最接近d的那个，假设是s1
11. 然后第二次，要使用最接近 s1 + d 的那个 s2
12. 大概有点想法了，应该是 dp[x] 表示能够选择的集合，它的sum = s2 - s1 = x <= d的那个
13. 且它不能选中给定的集合