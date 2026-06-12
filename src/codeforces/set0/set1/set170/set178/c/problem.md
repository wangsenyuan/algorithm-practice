# C2. Smart Beaver and Resolving Collisions

[Problem link](https://codeforces.com/problemset/problem/178/C2)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: stdin

output: stdout

The Smart Beaver from ABBYY has a lot of hobbies. One of them is constructing efficient
hash tables. One of the most serious problems in hash tables is resolving
collisions, and the Beaver decided to explore it in detail.

The hash table has `h` cells numbered from `0` to `h - 1`. Objects are added to
and removed from it. Every object has a unique identifier and a hash value — an
integer from `0` to `h - 1`. When an object is added, if the cell for its hash
value is free, the object goes there. Otherwise there is a collision. When an
object is deleted, its cell becomes empty.

The Beaver uses **linear probing** with step `m`:

- let the hash value be `t`;
- if cell `t` is occupied, try `(t + m) mod h`, then `(t + 2m) mod h`, then
  `(t + 3m) mod h`, and so on;
- place the object in the first free cell found.

It is guaranteed that the input always allows every addition to succeed.

Operation `a mod b` means the remainder when `a` is divided by `b`.

You are given a sequence of operations, each either adding an object or deleting
one. When adding an object, the algorithm above performs a sequence of table
calls. Calls to already occupied cells are called **dummy** calls. If the object
is finally placed in cell `(t + i * m) mod h` (`i >= 0`), then exactly `i` dummy
calls were performed.

Your task is to compute the total number of dummy calls for the whole sequence.
Deletions perform no dummy calls. The table is initially empty.

## Solution summary

Starting from any cell `t`, the probe sequence `t, t+m, t+2m, ... (mod h)` only
ever visits cells that are congruent modulo `g = gcd(h, m)`. This splits the table
into `g` independent **cycles**, each of length `h / g`. Within one cycle the
probe order visits its members in a fixed cyclic order, so an addition just needs
the first free cell at-or-after the hash position along that cycle.

Implementation (see [solution.go](solution.go)):

1. **Cycle decomposition + relabeling.** Walk `u = (u + m) % h` from each
   unvisited cell to enumerate every cycle. Lay the cycles out one after another
   into a single linear array of length `h`; record for each cell `u` its global
   position `pos[u]`, its index inside its cycle `posInCycle[u]`, and which cycle
   it belongs to. Each cycle thus occupies a contiguous block, so "next free along
   the cycle" becomes "next free index in the block (with wrap-around)."

2. **Segment tree over positions (range-min).** A free cell stores its own
   position index; an occupied cell stores `inf`. A min-query over a range returns
   the smallest free position in it, or `inf` if the range is full.

3. **Add `(id, hash)`.** Query `[pos[hash], blockEnd)` for the first free cell.
   - If found at `free`, the dummy-call count is `free - pos[hash]`.
   - Otherwise add the steps to the end of the cycle, then wrap and query
     `[blockStart, pos[hash])`, adding `free - blockStart`.
   Mark `free` occupied and remember `history[id] = free`.

4. **Delete `id`.** Mark `history[id]` free again (store its own position).

Each operation is `O(log h)`, so the whole sequence runs in `O(n log h)` —
fast enough for the `h, n <= 2 * 10^5` subtask. The answer can exceed 32 bits,
so accumulate dummy calls in a 64-bit integer.