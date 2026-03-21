# Problem

Fox Ciel has some flowers: `r` red flowers, `g` green flowers and `b` blue flowers. She wants to use these flowers to make several bouquets. There are 4 types of bouquets:

- To make a **red bouquet**, it needs 3 red flowers.
- To make a **green bouquet**, it needs 3 green flowers.
- To make a **blue bouquet**, it needs 3 blue flowers.
- To make a **mixing bouquet**, it needs 1 red, 1 green and 1 blue flower.

Help Fox Ciel to find the maximal number of bouquets she can make.

## Input

The first line contains three integers `r`, `g` and `b` (`0 <= r, g, b <= 10^9`) — the number of red, green and blue flowers.

## Output

Print the maximal number of bouquets Fox Ciel can make.

## Examples

### Example 1

Input

```text
3 6 9
```

Output

```text
6
```

### Example 2

Input

```text
4 4 4
```

Output

```text
4
```

### Example 3

Input

```text
0 0 0
```

Output

```text
0
```

## Note

- In test case 1, we can make 1 red bouquet, 2 green bouquets and 3 blue bouquets.
- In test case 2, we can make 1 red, 1 green, 1 blue and 1 mixing bouquet.
