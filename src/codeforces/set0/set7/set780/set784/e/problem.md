# E. Twisted Circuit

[Problem link](https://codeforces.com/problemset/problem/784/E)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given four binary input wires `a`, `b`, `c`, and `d`. Each value is either
`0` or `1`.

Evaluate the following logic circuit:

The twist is that the circuit diagram uses swapped pictures for `OR` and `XOR`:
an `OR`-shaped gate means apply `XOR`, and an `XOR`-shaped gate means apply `OR`.

The circuit picture can be read as:

```text
x1 = OR-shaped gate(a, b)
x2 = XOR-shaped gate(c, d)
x3 = AND gate(b, c)
x4 = OR-shaped gate(a, d)

y1 = AND gate(x1, x2)
y2 = XOR-shaped gate(x3, x4)

answer = OR-shaped gate(y1, y2)
```

Here:

- `AND` is bitwise conjunction.
- an `OR`-shaped gate applies exclusive or.
- an `XOR`-shaped gate applies bitwise disjunction.

## Input

The only line contains four integers `a`, `b`, `c`, and `d`
(`0 <= a, b, c, d <= 1`).

## Output

Print one integer: the output of the circuit.

## Examples

### Input

```text
0 1 1 0
```

### Output

```text
0
```

### Input

```text
1 0 0 0
```

### Output

```text
1
```

### Input

```text
0 0 0 0
```

### Output

```text
0
```

### Input

```text
1 0 0 1
```

### Output

```text
1
```

## Solution

The input size is fixed, so directly simulate the circuit gate by gate. The only
detail to be careful about is the swapped drawing convention for `OR` and `XOR`
gates.

The implementation models the diagram symbols directly:

```text
x1 = OrGate(a, b)
x2 = XorGate(c, d)
x3 = AndGate(b, c)
x4 = OrGate(a, d)
y1 = x1 & x2
y2 = XorGate(x3, x4)
answer = OrGate(y1, y2)
```

where:

```text
AndGate.pull() = left & right
OrGate.pull()  = left ^ right
XorGate.pull() = left | right
```

Equivalently, after translating every pictured gate to its real operation:

```text
x1 = a ^ b
x2 = c | d
x3 = b & c
x4 = a ^ d
y1 = x1 & x2
y2 = x3 | x4
answer = y1 ^ y2
```

### Complexity

The circuit has a constant number of gates.

```text
Time:  O(1)
Space: O(1)
```
