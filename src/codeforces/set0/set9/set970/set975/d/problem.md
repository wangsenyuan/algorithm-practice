# Problem D

Ghosts live in harmony and peace, they travel the space without any purpose other than scare whoever stands in their way.

There are 𝑛 ghosts in the universe, they move in the 𝑂𝑋𝑌 plane, each one of them has its own velocity that does not change in time: 𝑉→ = 𝑉ₓ𝑖→ + 𝑉ᵧ𝑗→ where 𝑉ₓ is its speed on the 𝑥-axis and 𝑉ᵧ is on the 𝑦-axis.

A ghost 𝑖 has experience value 𝐸𝑋ᵢ, which represent how many ghosts tried to scare him in his past. Two ghosts scare each other if they were in the same cartesian point at a moment of time.

As the ghosts move with constant speed, after some moment of time there will be no further scaring (what a relief!) and the experience of ghost kind 𝐺𝑋 = ∑ᵢ₌₁ⁿ 𝐸𝑋ᵢ will never increase.

Tameem is a red giant, he took a picture of the cartesian plane at a certain moment of time 𝑇, and magically all the ghosts were aligned on a line of the form 𝑦 = 𝑎⋅𝑥 + 𝑏. You have to compute what will be the experience index of the ghost kind 𝐺𝑋 in the indefinite future, this is your task for today.

Note that when Tameem took the picture, 𝐺𝑋 may already be greater than 0, because many ghosts may have scared one another at any moment between [−∞, 𝑇].

## Input

The first line contains three integers 𝑛, 𝑎 and 𝑏 (1 ≤ 𝑛 ≤ 2⋅10⁵, 1 ≤ |𝑎| ≤ 10⁹, 0 ≤ |𝑏| ≤ 10⁹) — the number of ghosts in the universe and the parameters of the straight line.

Each of the next 𝑛 lines contains three integers 𝑥ᵢ, 𝑉ₓᵢ, 𝑉ᵧᵢ (−10⁹ ≤ 𝑥ᵢ ≤ 10⁹, −10⁹ ≤ 𝑉ₓᵢ, 𝑉ᵧᵢ ≤ 10⁹), where 𝑥ᵢ is the current 𝑥-coordinate of the 𝑖-th ghost (and 𝑦ᵢ = 𝑎⋅𝑥ᵢ + 𝑏).

It is guaranteed that no two ghosts share the same initial position, in other words, it is guaranteed that for all (𝑖,𝑗) 𝑥ᵢ ≠ 𝑥ⱼ for 𝑖 ≠ 𝑗.

## Output

Output one line: experience index of the ghost kind 𝐺𝑋 in the indefinite future.

## Examples

**Example 1**

Input:
```
4 1 1
1 -1 -1
2 1 1
3 1 1
4 -1 -1
```

Output:
```
8
```

**Example 2**

Input:
```
3 1 0
-1 1 0
0 0 -1
1 -1 -2
```

Output:
```
6
```

**Example 3**

Input:
```
3 1 0
0 0 0
1 0 0
2 0 0
```

Output:
```
0
```

## Note

There are four collisions (1,2,𝑇−0.5), (1,3,𝑇−1), (2,4,𝑇+1), (3,4,𝑇+0.5), where (𝑢,𝑣,𝑡) means a collision happened between ghosts 𝑢 and 𝑣 at moment 𝑡. At each collision, each ghost gained one experience point, this means that 𝐺𝑋 = 4⋅2 = 8.

In the second test, all points will collide when 𝑡 = 𝑇+1.


### ideas
1. 因为历史也要计算，所以和时间没有关系。
2. GX = 2 * 相交的对数 = 2 * (n * (n - 1) / 2 - 不相交的对数)
3. 不相交分两种情况，在运动的情况下，处于平行上下的，或者不运动的，不在所在直线上的
4. 所有不运动的，都不会相交