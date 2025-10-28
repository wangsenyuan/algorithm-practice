# Problem C

People in the Tomskaya region like magic formulas very much. You can see some of them below.

Imagine you are given a sequence of positive integer numbers p1, p2, ..., pn. Let's write down some magic formulas:

```
Q = (p1 + p2 + ... + pn) mod 2^32
```

Here, "mod" means the operation of taking the residue after dividing.

The expression `x ⊕ y` means applying the bitwise XOR (excluding "OR") operation to integers x and y. The given operation exists in all modern programming languages. For example, in languages C++ and Java it is represented by "^", in Pascal — by "xor".

People in the Tomskaya region like magic formulas very much, but they don't like to calculate them! Therefore you are given the sequence p, calculate the value of Q.

## Input

The first line of the input contains the only integer n (1 ≤ n ≤ 10^6). The next line contains n integers: p1, p2, ..., pn (0 ≤ pi ≤ 2·10^9).

## Output

The only line of output should contain a single integer — the value of Q.

## Examples

### Input
```
3
1 2 3
```

### Output
```
3
```