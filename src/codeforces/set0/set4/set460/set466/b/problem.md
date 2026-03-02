# Wonder Room

The start of the new academic year brought about the problem of accommodating students into dormitories. One of such dormitories has an a × b square meter wonder room. The caretaker wants to accommodate exactly n students there. But the law says that there must be at least 6 square meters per student in a room (that is, the room for n students must have the area of at least 6n square meters). The caretaker can enlarge any (possibly both) side of the room by an arbitrary positive integer of meters. Help him change the room so that all n students could live in it and the total area of the room was as small as possible.

## Input

The first line contains three space-separated integers n, a and b (1 ≤ n, a, b ≤ 10^9) — the number of students and the sizes of the room.

## Output

Print two lines: the first line contains the final area s of the room; the second line contains a1 and b1 (a ≤ a1, b ≤ b1) — the final sizes of the room. If there are multiple optimal solutions, print any of them.

## Examples

### Example 1

**Input:**

```
3 3 5
```

**Output:**

```
18
3 6
```

### Example 2

**Input:**

```
2 4 4
```

**Output:**

```
16
4 4
```
