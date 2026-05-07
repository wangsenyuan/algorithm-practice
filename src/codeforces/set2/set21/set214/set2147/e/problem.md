You are given an array of `n` non-negative integers.

You want to answer `q` independent scenarios. In the `i`-th scenario, you are allowed to perform the following operation at most `bi` times:

- Choose an element of the array and increase it by `1`.

Your goal is to maximize the number of bits that are equal to `1` in the bitwise OR of all numbers in the array. Find this number for each scenario.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 10^3`). The description of the test cases follows.

The first line of each test case contains two integers `n` and `q` (`1 <= n, q <= 10^5`) — the size of the array and the number of scenarios.

The second line contains `n` integers `a1, a2, ..., an` (`0 <= ai <= 10^9`) — the elements of the array.

The `i`-th of the next `q` lines contains a single integer `bi` (`0 <= bi <= 10^9`) — the maximum number of operations allowed in the `i`-th scenario.

It is guaranteed that the sum of `n` over all test cases does not exceed `10^5`, and the sum of `q` over all test cases does not exceed `10^5`.

## Output

For each test case, output `q` lines, the `i`-th of them containing a single integer — the maximum possible number of bits equal to `1` in the bitwise OR in the `i`-th scenario.

## Example

### Input

```text
3
1 3
0
0
2
4
2 2
1 3
0
3
2 1
1000000000 1000000000
1000000000
```

### Output

```text
0
1
2
2
3
31
```

## Note

Visualizer link

**First test case**

- In the first scenario, there are no operations, so the answer is the number of `1`-bits in the bitwise OR of the original array. The array is `0`, which has `0` bits set.
- In the second scenario, one way to get `1` bit set in the bitwise OR is to increase `a1` by `1` twice, obtaining a bitwise OR of `2 = (10)_2`. This is optimal with at most two operations.
- In the third scenario, one way to get `2` bits set is to add `1` to `a1` three times, which is optimal. You do not have to use all four allowed operations.

**Second test case**

- In the first scenario, there are no operations, so the answer is the number of `1`-bits in the bitwise OR of the original array, which is `2`.
- In the second scenario, one way to get `3` bits set is to add `1` to `a2` three times, which is optimal.

### ideas
1. 最终的or sum = ???11111 类似这样的结构
2. 

## Solution explanation

Let

`cur = a[0] | a[1] | ... | a[n-1]`.

The bits that are already `1` in `cur` are free. Let their count be:

`bitsCount = popcount(cur)`.

Every operation only increases one array element. The goal for each query budget `b` is to know how many additional missing bits can be made present in the final OR.

The code precomputes a list:

`arr = [(minimum cost, resulting number of set bits)]`.

Then each query only needs a binary search by budget.

## Cost to set one bit

Suppose bit `d` is currently missing from the OR. Then every array element has bit `d = 0`.

For one element `x`, the cheapest way to make its bit `d` become `1` is to increase it to the next number whose bit `d` is set.

Let:

`w = 2^d`.

Only the lower `d` bits matter for this cost. If `x & (w - 1)` is large, then `x` is closer to the next block where bit `d` becomes `1`.

The needed operations are:

`w - (x & (w - 1))`.

So among all elements, we choose the one maximizing:

`x & (w - 1)`.

This is exactly what the code does:

```go
w := 1 << d1
for i, v := range a2 {
	if v&(w-1) > a2[pos]&(w-1) {
		pos = i
	}
}
ops := w - a2[pos]&(w-1)
a2[pos] += ops
```

In Go, `&` has higher precedence than `-`, so the last expression means:

`ops = w - (a2[pos] & (w - 1))`.

## Why process bits from high to low

When we increase a number to set bit `d`, lower bits may change because of carry. For example, turning on bit `2` may clear bits `0` and `1`.

However, after bit `d` is set, later operations for lower bits cannot disturb bit `d`. When a lower missing bit `e < d` is set, the added amount is at most `2^e`, and since bit `e` was `0`, this operation does not carry into higher bits.

Therefore, if we want the final OR to contain a certain missing bit `d` and all useful lower missing bits, the safe order is:

`d, d-1, ..., 0`.

That is why for each candidate highest bit `d`, the code copies the original array and runs:

```go
for d1 := d; d1 >= 0; d1-- {
	if or(a2)&(1<<d1) == 0 {
		...
	}
}
```

Whenever bit `d1` is still absent from the current OR, it pays the minimum cost to make it present.

## What each precomputed pair means

The loop over `d` considers only bits missing from the original OR:

```go
for d := range 31 {
	if (sum>>d)&1 == 1 {
		continue
	}
	...
	bitsCount++
	arr = append(arr, pair{tot, bitsCount})
}
```

For this missing bit `d`, `tot` is the minimum cost needed to make one more originally missing bit available at this level, while also ensuring lower missing bits are repaired after any carry.

The resulting number of set bits is stored as `bitsCount`.

The initial pair:

```go
arr = append(arr, pair{0, bitsCount})
```

means that with zero operations, the answer is just the number of set bits in the original OR.

## Answering queries

For a query budget `q`, we need the best precomputed state whose cost is at most `q`.

The costs in `arr` are generated in increasing bit order and are monotonic, so the code binary-searches the first cost that is too large:

```go
j := sort.Search(len(arr), func(j int) bool {
	return arr[j].first > q
})
ans[i] = arr[j-1].second
```

So `arr[j-1]` is the best reachable state under this budget.

## Complexity

There are only 31 relevant bits because `a_i, b_i <= 10^9`.

For each missing bit `d`, the implementation may scan the array for each lower bit, so preprocessing costs:

`O(31 * 31 * n)`.

Each query is answered by binary search over at most 32 precomputed states:

`O(log 31)`, effectively constant.
