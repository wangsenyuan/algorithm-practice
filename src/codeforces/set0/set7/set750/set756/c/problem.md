# Problem Description

Nikita has a stack. A stack in this problem is a data structure that supports two operations:
- **push(x)**: puts an integer x on the top of the stack
- **pop()**: deletes the top integer from the stack (the last added). If the stack is empty, then the operation pop() does nothing.

Nikita made m operations with the stack but forgot them. Now Nikita wants to remember them. He remembers them one by one, on the i-th step he remembers an operation he made pi-th. In other words, he remembers the operations in order of some permutation p₁, p₂, ..., pₘ. After each step Nikita wants to know what is the integer on the top of the stack after performing the operations he have already remembered, in the corresponding order. Help him!

## Input

The first line contains the integer m (1 ≤ m ≤ 10⁵) — the number of operations Nikita made.

The next m lines contain the operations Nikita remembers. The i-th line starts with two integers pᵢ and tᵢ (1 ≤ pᵢ ≤ m, tᵢ = 0 or tᵢ = 1) — the index of operation he remembers on the step i, and the type of the operation:
- tᵢ equals 0, if the operation is pop()
- tᵢ equals 1, if the operation is push(x)

If the operation is push(x), the line also contains the integer xᵢ (1 ≤ xᵢ ≤ 10⁶) — the integer added to the stack.

It is guaranteed that each integer from 1 to m is present exactly once among integers pᵢ.

## Output

Print m integers. The integer i should equal the number on the top of the stack after performing all the operations Nikita remembered on the steps from 1 to i. If the stack is empty after performing all these operations, print -1.