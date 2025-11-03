### Description
A super computer has been built in the Turtle Academy of Sciences. The computer consists of n × m × k CPUs. The architecture is a parallelepiped of size n × m × k, split into 1 × 1 × 1 cells, each cell containing exactly one CPU. Thus, each CPU can be identified by three numbers: the layer number (1…n), the line number (1…m) and the column number (1…k).

During operation, CPUs can send messages according to the turtle scheme: CPU (x, y, z) can send messages to CPUs (x + 1, y, z), (x, y + 1, z) and (x, y, z + 1) (if they exist). There is no feedback: CPUs (x + 1, y, z), (x, y + 1, z) and (x, y, z + 1) cannot send messages to CPU (x, y, z).

Over time some CPUs broke down and stopped working. Such CPUs cannot send messages, receive messages, or serve as intermediaries in transmitting messages. We say that CPU (a, b, c) controls CPU (d, e, f) if there is a chain of CPUs (xᵢ, yᵢ, zᵢ), such that (x₁ = a, y₁ = b, z₁ = c), (xₚ = d, yₚ = e, zₚ = f) (here and below p is the length of the chain) and the CPU in the chain with number i (i < p) can send messages to CPU i + 1.

Turtles are concerned about the denial-proofness of the communication system between the remaining CPUs. They want to know the number of critical CPUs. A CPU (x, y, z) is critical if turning it off disrupts some control, that is, if there exist two CPUs distinct from (x, y, z): (a, b, c) and (d, e, f), such that (a, b, c) controls (d, e, f) before (x, y, z) is turned off and stops controlling it after the turn-off.

### Input
- The first line contains three integers n, m and k (1 ≤ n, m, k ≤ 100) — the dimensions of the Super Computer.

- Then n blocks follow, describing the current state of the processors. The blocks correspond to the layers of the Super Computer in order from 1 to n. Each block consists of m lines, k characters in each — an m × k table. The z-th character of the y-th line of block x corresponds to CPU (x, y, z). Character "1" corresponds to a working CPU and character "0" corresponds to a malfunctioning one. The blocks are separated by exactly one empty line.

### Output
Print a single integer — the number of critical CPUs, i.e., those for which turning only this CPU off will disrupt some control.

### Examples

#### Example 1
Input
```
2 2 3
000
000

111
111
```
Output
```
2
```

#### Example 2
Input
```
3 3 3
111
111
111

111
111
111

111
111
111
```
Output
```
19
```

#### Example 3
Input
```
1 1 10
0101010101
```
Output
```
0
```

### Notes
In the first sample the whole first layer of CPUs is malfunctioning. In the second layer, when CPU (2, 1, 2) is turned off, it disrupts the control by CPU (2, 1, 3) over CPU (2, 1, 1), and when CPU (2, 2, 2) is turned off, it disrupts the control over CPU (2, 2, 3) by CPU (2, 2, 1).

In the second sample all processors except for the corner ones are critical.

In the third sample there is not a single processor controlling another processor, so the answer is 0.


### ideas
1. 一个立方体，正常情况下，一个cell始终可以往它的上、右、后三个方向传递信息
2. 最后一层比较好判断的，从最后一行开始，如果它是1，且它的两边都是1，那么它是critical的（原来的时候，它左边的可以，通过它，control右边的）
3. 或者它上下是1
4. 如果一个(x, y, z)它的上、右、后三个位置，有一个1，且它的下、左、前三个位置有一个1，那么它是critical的
5. 