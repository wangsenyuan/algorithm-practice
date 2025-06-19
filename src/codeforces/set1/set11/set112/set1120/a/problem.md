# Flower Festival Problem

## Problem Description
At the first holiday in spring, the town Shortriver traditionally conducts a flower festival. Townsfolk wear traditional wreaths during these festivals. Each wreath contains exactly $k$ flowers.

The work material for the wreaths for all $n$ citizens of Shortriver is cut from the longest flowered liana that grew in the town that year. Liana is a sequence $a_1, a_2, ..., a_m$, where $a_i$ is an integer that denotes the type of flower at the position $i$. This year the liana is very long $(m \geq n \cdot k)$, and that means every citizen will get a wreath.

## Machine Operation
Very soon the liana will be inserted into a special cutting machine in order to make work material for wreaths. The machine works in a simple manner:
- It cuts $k$ flowers from the beginning of the liana, then another $k$ flowers and so on
- Each such piece of $k$ flowers is called a workpiece
- The machine works until there are less than $k$ flowers on the liana

## Diana's Schematic
Diana has found a weaving schematic for the most beautiful wreath imaginable. In order to weave it, $k$ flowers must contain flowers of types $b_1, b_2, ..., b_s$, while others can be of any type. If a type appears in this sequence several times, there should be at least that many flowers of that type as the number of occurrences of this flower in the sequence. The order of the flowers in a workpiece does not matter.

## Diana's Opportunity
Diana has a chance to remove some flowers from the liana before it is inserted into the cutting machine. She can remove flowers from any part of the liana without breaking liana into pieces. If Diana removes too many flowers, it may happen so that some of the citizens do not get a wreath.

**Question**: Could some flowers be removed from the liana so that at least one workpiece would conform to the schematic and machine would still be able to create at least $n$ workpieces?

## Input Format
The first line contains four integers $m, k, n$ and $s$ $(1 \leq n,k,m \leq 5\cdot10^5, k \cdot n \leq m, 1 \leq s \leq k)$:
- $m$ = number of flowers on the liana
- $k$ = number of flowers in one wreath
- $n$ = amount of citizens
- $s$ = length of Diana's flower sequence

The second line contains $m$ integers $a_1, a_2, ..., a_m$ $(1 \leq a_i \leq 5\cdot10^5)$ — types of flowers on the liana.

The third line contains $s$ integers $b_1, b_2, ..., b_s$ $(1 \leq b_i \leq 5\cdot10^5)$ — the sequence in Diana's schematic.

## Output Format
If it's impossible to remove some of the flowers so that there would be at least $n$ workpieces and at least one of them fulfills Diana's schematic requirements, output $-1$.

Otherwise:
1. In the first line output one integer $d$ — the number of flowers to be removed by Diana
2. In the next line output $d$ different integers — the positions of the flowers to be removed

If there are multiple answers, print any.

### ideas
1. 选择一段l...r， 在这一段内存在序列b，且 r - l + 1 >= k, l / k + 1 + (n - r) / k >= m
2. 前面分配的，+ 当前段 + 后面可分配的 >= m
3. 所以关键是中间这一段