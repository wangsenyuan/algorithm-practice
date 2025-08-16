# Problem: Rebel Fleet Escape

## Problem Description

The Rebel fleet is on the run. It consists of **m** ships currently gathered around a single planet. Just a few seconds ago, the vastly more powerful Empire fleet has appeared in the same solar system, and the Rebels will need to escape into hyperspace.

In order to spread the fleet, the captain of each ship has independently come up with the coordinate to which that ship will jump. In the obsolete navigation system used by the Rebels, this coordinate is given as the value of an arithmetic expression of the form `(a+b)/c`.

To plan the future of the resistance movement, Princess Heidi needs to know, for each ship, how many ships are going to end up at the same coordinate after the jump. You are her only hope!

## Input

The first line of the input contains a single integer **m** (1 ≤ m ≤ 200,000) – the number of ships.

The next **m** lines describe one jump coordinate each, given as an arithmetic expression. An expression has the form `(a+b)/c`. Namely, it consists of:
- An opening parenthesis `(`
- A positive integer **a** of up to two decimal digits
- A plus sign `+`
- A positive integer **b** of up to two decimal digits
- A closing parenthesis `)`
- A slash `/`
- A positive integer **c** of up to two decimal digits

## Output

Print a single line consisting of **m** space-separated integers. The **i**-th integer should be equal to the number of ships whose coordinate is equal to that of the **i**-th ship (including the **i**-th ship itself).

## Example

### Input
```
4
(99+98)/97
(26+4)/10
(12+33)/15
(5+1)/7
```

### Output
```
1 2 2 1
```

### Explanation
In the sample testcase:
- Ship 1: `(99+98)/97 = 197/97 ≈ 2.03` → coordinate 2
- Ship 2: `(26+4)/10 = 30/10 = 3` → coordinate 3
- Ship 3: `(12+33)/15 = 45/15 = 3` → coordinate 3
- Ship 4: `(5+1)/7 = 6/7 ≈ 0.86` → coordinate 0

The second and third ships both end up at coordinate 3, so the output shows `1 2 2 1`.

## Note

This problem has only two versions – easy and hard.
