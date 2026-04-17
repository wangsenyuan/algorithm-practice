# Problem

You are given a boolean function of three variables, defined by its truth table.
Find an expression of minimum length that is equal to this function.

The expression may use:

- AND: `&` (ASCII 38)
- OR: `|` (ASCII 124)
- NOT: `!` (ASCII 33)
- variables: `x`, `y`, `z` (ASCII 120-122)
- parentheses: `(` and `)` (ASCII 40 and 41)

If there are multiple minimum-length expressions, output the lexicographically smallest one.

Operators have standard precedence:

1. `!` (highest)
2. `&`
3. `|` (lowest)

The expression must satisfy the grammar:

```text
E ::= E '|' T | T
T ::= T '&' F | F
F ::= '!' F | '(' E ')' | 'x' | 'y' | 'z'
```

## Input

The first line contains one integer `n` — the number of functions (`1 <= n <= 10000`).

Each of the next `n` lines contains a binary string of length `8`, describing one truth table.
For position `j` (`0 <= j < 8`), the character gives the function value at:

- `x = (j >> 2) & 1`
- `y = (j >> 1) & 1`
- `z = (j >> 0) & 1`

## Output

Print `n` lines.
The `i`-th line should contain a minimum-length expression for the `i`-th function.
If several expressions have minimum length, print the lexicographically smallest one.
Expressions must satisfy the grammar and contain no spaces.

## Example

### Input

```text
4
00110011
00000111
11110000
00011111
```

### Output

```text
y
(y|z)&x
!x
x|y&z
```

## Note

The truth table for the second function:

## Solution Idea

The key observation is that there are only `3` variables, so there are only:

- `2^3 = 8` input assignments

Therefore a boolean function is completely determined by its 8-bit truth table, and the total number of possible functions is only:

- `2^8 = 256`

So instead of searching over infinitely many formulas, we can do dynamic programming over these `256` functions.

### 1. Represent each formula by an 8-bit mask

For every expression, compute its truth table as an 8-bit integer:

- bit `j` is the value of the expression on assignment `j`
- where
  - `x = (j >> 2) & 1`
  - `y = (j >> 1) & 1`
  - `z = (j >> 0) & 1`

For example:

- `x` corresponds to mask `11110000`
- `y` corresponds to mask `00110011`
- `z` corresponds to mask `01010101`

Thus the problem becomes:

- for each mask from `0` to `255`,
- find the shortest expression producing it,
- and among equal lengths choose the lexicographically smallest one.

### 2. Why one DP table is not enough

Because of operator precedence.

The grammar is:

```text
E ::= E '|' T | T
T ::= T '&' F | F
F ::= '!' F | '(' E ')' | 'x' | 'y' | 'z'
```

So the same function may need different textual forms depending on context.

Example:

- `x|y` is fine as a full expression
- but if we want to use it inside `&`, we need `(x|y)`

Therefore we store three best strings for each mask:

- `bestE[mask]`: best expression of grammar type `E`
- `bestT[mask]`: best expression of grammar type `T`
- `bestF[mask]`: best expression of grammar type `F`

This is the standard way to handle precedence cleanly.

### 3. Initial states

Variables are valid factors:

- `bestF[mask(x)] = "x"`
- `bestF[mask(y)] = "y"`
- `bestF[mask(z)] = "z"`

### 4. Transitions

We repeatedly relax the following rules.

#### Promotion

Every factor is also a term, and every term is also an expression:

- `F -> T`
- `T -> E`

So:

- `bestT[m]` can be improved by `bestF[m]`
- `bestE[m]` can be improved by `bestT[m]`

#### Parentheses

If we already have an expression `E`, then:

- `"(" + E + ")"` is a factor

So:

- `bestF[m]` can be improved by parenthesizing `bestE[m]`

#### Negation

If `f` is a factor, then:

- `"!" + f` is also a factor

and its truth table is bitwise negation:

- new mask = `~m` on 8 bits

#### AND

From the grammar:

- `T ::= T '&' F`

So if:

- left side is a term with mask `a`
- right side is a factor with mask `b`

then:

- `left & right` is a term with mask `a & b`

#### OR

From the grammar:

- `E ::= E '|' T`

So if:

- left side is an expression with mask `a`
- right side is a term with mask `b`

then:

- `left | right` is an expression with mask `a | b`

### 5. Comparing candidates

Whenever we derive a new candidate string for the same mask and same grammar level, we keep the better one using:

1. shorter length first
2. lexicographically smaller string if lengths are equal

### 6. Why repeated relaxation works

There are only:

- `256` masks
- `3` grammar levels

So only `768` DP states exist.

Every successful relaxation makes a state strictly better:

- shorter, or
- same length but lexicographically smaller

Hence the process stabilizes quickly.

### 7. Final answer

After precomputation, for each input truth table:

1. convert the 8-character string into a mask
2. print `bestE[mask]`

We use `bestE` because the final printed formula must be a complete expression.

## Summary

The whole problem reduces to:

- DP on all 256 possible truth tables,
- with 3 layers (`E`, `T`, `F`) to respect precedence,
- and transitions for `!`, `&`, `|`, plus parentheses.

Because the state space is tiny, this brute-force-style DP is completely feasible.
