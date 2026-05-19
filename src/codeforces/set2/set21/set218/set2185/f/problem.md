# F. BattleCows

https://codeforces.com/problemset/problem/2185/F

## Problem

Farmer John has `2^n` cows in a line. Cow `i` has skill `a_i`.

Each cow starts as a stack containing only itself. The skill of a stack is the
bitwise XOR of the skills of all cows in that stack.

The tournament repeats until only one stack remains:

1. Every odd-positioned stack fights the stack immediately to its right.
2. The stack with larger skill wins. If the two skills are equal, the left stack
   wins.
3. The winning stack jumps on top of the losing stack, and the line of stacks is
   compressed.

Farmer John has `q` independent potions. Potion `i` is tested by giving it to
cow `b_i`, which temporarily changes that cow's skill to `c_i`, then running the
tournament.

For each potion trial, output how many cows are above cow `b_i` in the final
stack. After each trial, the cow's skill returns to its original value.


## Solution

Think of the tournament as a complete binary tree over the original array. A
leaf is one cow, and an internal node represents the fight between its left and
right halves. The XOR value of any stack is just the XOR of the leaves in that
node.

Build a segment tree `tr` where `tr[i]` is the XOR of the interval represented
by node `i`. A potion query only changes one cow `pos`, so every interval that
contains `pos` has its XOR changed by:

```text
delta = old_value ^ new_value
```

All other intervals keep their original XOR. Therefore a query can be answered
by walking from the root to `pos` without rebuilding the tree.

At a node `[l, r]` with left child `[l, mid]` and right child `[mid+1, r]`:

* If `pos` is in the left child, the left stack's current skill is
  `tr[left] ^ delta`, and the right stack's skill is `tr[right]`.
  If the left stack loses, every cow in the right stack ends above `pos`, so add
  `r - mid`.
* If `pos` is in the right child, the left stack's skill is `tr[left]`, and the
  right stack's current skill is `tr[right] ^ delta`.
  If the right stack loses, every cow in the left stack ends above `pos`, so add
  `mid - l + 1`.

The tie rule says the left stack wins, so the comparisons are:

```text
left side loses:  (tr[left] ^ delta) < tr[right]
right side loses: tr[left] >= (tr[right] ^ delta)
```

After deciding the current fight, recurse into the child that contains `pos`.
The accumulated count is the answer for this potion.

The segment tree build is `O(2^n)`. Each query follows one root-to-leaf path, so
it costs `O(n)`, which is enough because `n <= 18`.
