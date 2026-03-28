Let's define an algorithm to generate a sequence of m+k integers as follows:

First, receive a sequence x of m integers as input. If k=0, terminate immediately and return the sequence x.

Then, select any index 1≤i≤|x| and insert (x_i+1) immediately after the element x_i.

If x contains exactly m+k integers, terminate and return the sequence x. Otherwise, return to the second step.

Alice knows that this algorithm was used by an ancient civilization in order to hide their secrets safely. Alice wants to learn the knowledge that they wanted to hide, but it is not an easy job to infer the input from the output of the algorithm.

For a sequence b of n integers, let us define f(b) as the length of the shortest sequence that could be given as an input for the algorithm to generate b.

Given a sequence a of n integers, please compute the value of the following sum:

∑_{l=1}^{n} ∑_{r=l}^{n} f([a_l, a_{l+1}, …, a_r])

In other words, you must find the sum of f(c) over all subsegments c of a.

A sequence a is a subsegment of a sequence b if a can be obtained from b by the deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end. Two subsegments are considered different if the sets of positions of the deleted elements are different.

## Input

Each test contains multiple test cases. The first line contains the number of test cases t (1≤t≤10^4). The description of the test cases follows.

The first line of each test case contains a single integer n (1≤n≤300000).

The second line of each test case contains n integers a_1, a_2, …, a_n (1≤a_i≤10^9).

It is guaranteed that the sum of n over all test cases does not exceed 300000.

## Output

For each test case, output the answer on a separate line.

## Example

**Input**

```
5
5
1 2 3 4 5
5
1 3 5 7 9
5
1 2 5 6 5
7
1 2 4 5 3 7 8
9
9 8 9 2 3 4 4 5 3
```

**Output**

```
15
35
25
60
78
```

## Note

In the first test case, all subsegments of a=[1,2,3,4,5] can be generated from a sequence of length 1.

In the second test case, all subsegments of a=[1,3,5,7,9] cannot be generated from any shorter sequence than itself.

## algorithm

For a fixed sequence `b`, the value `f(b)` is the number of independent trees in the forest encoded by `b`.

The generation rule

- choose `x_i`
- insert `x_i + 1` immediately after it

means that every element `v` may only appear inside the subtree of some earlier element `v - 1`.
If we scan the final sequence from left to right, the valid parent of the current element is the nearest previous element that is still "open".

This is exactly the same as maintaining a strictly increasing stack for the current block:

1. Before processing `b[i]`, pop while `stack.top >= b[i]`.
2. After these pops, the only possible parent is the new top of the stack.
3. If that top equals `b[i] - 1`, then `b[i]` can be attached to the same tree.
4. Otherwise `b[i]` must start a new tree, so it contributes `1` to `f(b)`, and the previous stack is no longer useful.
   We clear the stack because future elements may only belong to this new tree, not to any older tree.
5. Push `b[i]` into the current stack.

So `f(b)` is just the number of positions where step 4 happens.

## contribution of one position

Now fix the original array `a` and a position `i`.
We want to count in how many subarrays `[l..r]` the element `a[i]` starts a new tree.

For a fixed subarray ending at or after `i`, the element `a[i]` does **not** start a new tree iff, inside the current last block of that subarray, after popping all previous elements `>= a[i]`, the nearest surviving value on the left is exactly `a[i] - 1`.

If such a parent does not exist, `a[i]` starts a new block and all older blocks are irrelevant for future attachments.

So globally, while scanning the whole array left to right, define:

- `p(i)` = the nearest previous index that remains on the current block stack after popping all values `>= a[i]`

Then only `p(i)` can possibly be the parent of `i`.

There are two cases:

1. `p(i)` does not exist, or `a[p(i)] != a[i] - 1`
   Then `a[i]` starts a new block in every subarray containing `i`.
   Its contribution is:

   - number of choices for `l`: `i + 1`
   - number of choices for `r`: `n - i`
   - total: `(i + 1) * (n - i)`

2. `p(i)` exists and `a[p(i)] = a[i] - 1`
   Then `a[i]` can attach exactly in subarrays with `l <= p(i)`.
   If `l > p(i)`, that parent is cut off, so `a[i]` becomes the first element of the last block.
   Therefore the valid left endpoints are:

   - `l = p(i) + 1, p(i) + 2, ..., i`

   which gives `i - p(i)` choices.
   The right endpoint still has `n - i` choices, so the contribution is:

   - `(i - p(i)) * (n - i)`

## implementation

We can compute everything in one left-to-right scan.

For each `a[i]`:

1. Pop from the stack while top value `>= a[i]`.
2. If the stack is non-empty and `stack.top.value + 1 == a[i]`, add:
   - `(n - i) * (i - stack.top.index)`
   and push `(a[i], i)`.
3. Otherwise add:
   - `(n - i) * (i + 1)`
   then clear the stack and start a new block with `(a[i], i)`.

This is `O(n)` per test case because every element is pushed and popped at most once.

## final algorithm

Maintain a stack of pairs `(value, index)` for the current block.

For each position `i` with value `a[i]`:

1. Pop while the stack is not empty and `stack.top.value >= a[i]`.
2. If the stack is not empty and `stack.top.value + 1 == a[i]`, then `a[i]` attaches to that parent.
   Its contribution is:
   - `(n - i) * (i - stack.top.index)`
   Then push `(a[i], i)` onto the stack.
3. Otherwise `a[i]` starts a new block.
   Its contribution is:
   - `(n - i) * (i + 1)`
   Then clear the whole stack and push `(a[i], i)`.

Why does the second formula use `i + 1`?

- if `a[i]` starts a new block, then for every left endpoint `l` with `0 <= l <= i`, the subarray `[l..r]` also makes `a[i]` start a new block
- there are `i + 1` choices for `l`
- there are `n - i` choices for `r`

Why does the first formula use `i - parent`?

- if `parent = stack.top.index`, then `a[i]` can attach only when the subarray starts at some `l <= parent`
- so `a[i]` starts a new block exactly when `parent < l <= i`
- that gives `i - parent` choices for `l`
- again there are `n - i` choices for `r`

Therefore summing these contributions during one left-to-right scan gives the answer in linear time.
