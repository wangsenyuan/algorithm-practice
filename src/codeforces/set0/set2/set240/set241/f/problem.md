# Problem F

The Old City is a rectangular city represented as an m × n grid of blocks. This city contains many buildings, straight two-way streets and junctions. Each junction and each building is exactly one block. All the streets have width of one block and are either vertical or horizontal. There is a junction on both sides of each street. We call two blocks adjacent if and only if they share a common side. No two blocks of different streets are adjacent and no two junctions are adjacent.

There is an annual festival and as a part of it, The Old Peykan follows a special path in the city. This path starts from a block in a street, continues with many junctions and ends in a block of some street. For each street block, we know how much time it takes for the Old Peykan to go from this block to an adjacent block. Also the Old Peykan can go from each junction to its adjacent street blocks in one minute. Of course Old Peykan can't go to building blocks.

We know the initial position of the Old Peykan and the sequence of junctions that it passes to reach its destination. After passing all the junctions and reaching the destination, it will stay there forever. Your task is to find out where will the Old Peykan be k minutes after it starts moving. Consider that The Old Peykan always follows the shortest path that passes through the given sequence of junctions and reaches the destination.

Note that the Old Peykan may visit some blocks more than once.

## Input

The first line of input contains three integers m, n and k (3 ≤ m, n ≤ 100, 1 ≤ k ≤ 10⁵). Next m lines are representing the city's map. Each of them contains n characters, each character is a block:

- Character "#" represents a building.
- Digits "1", "2", ..., "9" represent a block of a street and this digit means the number of minutes it takes for the Old Peykan to pass this block.
- Characters "a", "b", ..., "z" means that this block is a junction and this character is its name. All the junction names are unique.

Consider that all blocks have the coordinates: the j-th in the i-th line have coordinates (i, j) (1 ≤ i ≤ m, 1 ≤ j ≤ n).

The (m + 2)th line contains two integers rₛ and cₛ (1 ≤ rₛ ≤ m, 1 ≤ cₛ ≤ n), string s and another two integers rₑ and cₑ (1 ≤ rₑ ≤ m, 1 ≤ cₑ ≤ n). The path starts from block (rₛ, cₛ), continues through junctions in the order that is specified by s and will end in block (rₑ, cₑ). Length of s is between 1 and 1000.

It's guaranteed that string s denotes a correct path from the start position to the end position and string s doesn't contain two consecutive equal letters. Also start position (rₛ, cₛ) and the end position (rₑ, cₑ) are street blocks.

## Output

In a single line print two integers r_f and c_f — (r_f, c_f) being the position of the Old Peykan after exactly k minutes.

## Examples

**Input:**
```
3 10 12
##########
#z1a1111b#
##########
2 3 ab 2 8
```

**Output:**
```
2 8
```

**Input:**
```
10 3 5
###
#w#
#1#
#a#
#1#
#1#
#1#
#1#
#b#
###
3 2 abababababababab 6 2
```

**Output:**
```
8 2
```

**Input:**
```
3 10 6
##########
#z1a1311b#
##########
2 3 ab 2 8
```

**Output:**
```
2 7
```

## Example Explanations

**Movement rules recap:** Moving *from* a street block with digit d to an adjacent block costs d minutes. Moving *from* a junction to an adjacent street block costs 1 minute. After reaching the destination the Peykan stays there forever.

### Example 1

Row 2 of the grid: `#z1a1111b#` (positions 1–10).

| Position | (2,2)=z | (2,3)=1 | (2,4)=a | (2,5)=1 | (2,6)=1 | (2,7)=1 | (2,8)=1 | (2,9)=b |
|----------|---------|---------|---------|---------|---------|---------|---------|---------|

- Start at (2,3)='1'. Path: a → b. End: (2,8).
- t=0: at (2,3). Move from street '1' → costs 1 min.
- t=1: at (2,4)=junction 'a'. Move from junction → costs 1 min.
- t=2: at (2,5)='1'. Move from street '1' → costs 1 min.
- t=3: at (2,6)='1'. Move from street '1' → costs 1 min.
- t=4: at (2,7)='1'. Move from street '1' → costs 1 min.
- t=5: at (2,8)='1'. Move from street '1' → costs 1 min.
- t=6: at (2,9)=junction 'b'. Now head to destination (2,8). Costs 1 min.
- t=7: at (2,8)='1'. Destination reached. Stays forever.
- k=12 → answer is **(2, 8)**.

### Example 2

The grid is a vertical corridor (column 2): w(2,2), 1(3,2), a(4,2), 1(5,2), 1(6,2), 1(7,2), 1(8,2), b(9,2).

- Start at (3,2)='1'. Path: abababababababab. End: (6,2).
- t=0: at (3,2). Move from street '1' → costs 1 min.
- t=1: at (4,2)=junction 'a'. Move from junction toward 'b' → costs 1 min.
- t=2: at (5,2)='1'. Costs 1 min.
- t=3: at (6,2)='1'. Costs 1 min.
- t=4: at (7,2)='1'. Costs 1 min.
- t=5: at (8,2)='1'.
- k=5 → answer is **(8, 2)**.

(The Peykan is still on its way from 'a' to 'b' the first time when time runs out.)

### Example 3

Row 2 of the grid: `#z1a1311b#` — same as Example 1 but position (2,6) is now '3' instead of '1'.

- Start at (2,3)='1'. Path: a → b. End: (2,8).
- t=0: at (2,3). Move from street '1' → costs 1 min.
- t=1: at (2,4)=junction 'a'. Move from junction → costs 1 min.
- t=2: at (2,5)='1'. Move from street '1' → costs 1 min.
- t=3: at (2,6)='3'. Move from street '3' → costs **3** min.
- t=6: at (2,7)='1'.
- k=6 → answer is **(2, 7)**.

The slow street block '3' at (2,6) delays the Peykan so it only reaches (2,7) at minute 6, compared to minute 4 in Example 1.
