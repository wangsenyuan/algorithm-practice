# Problem C: Fibonacci GCD

There are less than 60 years left till the 900-th birthday anniversary of a famous Italian mathematician Leonardo Fibonacci. Of course, such important anniversary needs much preparations.

Dima is sure that it'll be great to learn to solve the following problem by the Big Day:

## Problem Statement

You're given a set A, consisting of numbers l, l + 1, l + 2, ..., r. Let's consider all its k-element subsets. For each such subset, let's find the largest common divisor of Fibonacci numbers with indexes determined by the subset elements. Among all found common divisors, Dima is interested in the largest one.

## Fibonacci Numbers Definition

Dima asked to remind you that Fibonacci numbers are elements of a numeric sequence, where:
- F₁ = 1
- F₂ = 1  
- Fₙ = Fₙ₋₁ + Fₙ₋₂ for n ≥ 3

Dima has more than half a century ahead to solve the given task, but you only have two hours. Count the residue from dividing the sought largest common divisor by m.

## Input

The first line contains four space-separated integers m, l, r and k (1 ≤ m ≤ 10⁹; 1 ≤ l < r ≤ 10¹²; 2 ≤ k ≤ r - l + 1).

**Note:** Please, do not use the %lld specifier to read or write 64-bit integers in С++. It is preferred to use cin, cout streams or the %I64d specifier.

## Output

Print a single integer — the residue from dividing the sought greatest common divisor by m.

## Examples

### Example 1
**Input:**
```
10 1 8 2
```

**Output:**
```
3
```

### Example 2
**Input:**
```
10 1 8 3
```

**Output:**
```
1
```

## Solution

### Key Insight
At first, let's prove the statement: **GCD(Fₙ, Fₘ) = F₍GCD(n,m)₎**.

### Mathematical Proof
Let's express Fₙ₊ₖ using Fₙ and Fₖ. We'll get the formula: 
**Fₙ₊ₖ = Fₖ · Fₙ₊₁ + Fₖ₋₁ · Fₙ**, which is easy to prove by induction.

Then use the derived formula and notice, that **GCD(Fₙ₊ₖ, Fₙ) = GCD(Fₖ, Fₙ)**.

Now you are to notice an analogy with Euclidean algorithm and to understand, that we've got necessary equality for GCD of two Fibonacci numbers.

### Algorithm Approach
So, our current task is to find in the given set subset of k (or at least of k) elements with maximal possible GCD. To be exactly, to find this GCD.

Let the answer be equal to q. Then ⌊r/q⌋ - ⌈l/q⌉ + 1 ≥ k (1) must be true.

Notice, that for each summand from left part of inequality O(√r) segments exist, in which its value is constant. Moreover, we can find all these segments and values in O(√r). To be more precise, we are interested in such q, that in the point q - 1 value of at least one summand changes (obviously, increases). There are also O(√r) such values. Go over all of them and try to use each of them as the answer (i.e., check inequality (1) for each of them), and choose maximum from all satisfying numbers. The answer always exists, as q = 1 is true for any input.

So, we've found index of required Fibonacci number. The number itself can be calculated by matrix exponentiation.

### Implementation

### Complexity
O(√r + log m)