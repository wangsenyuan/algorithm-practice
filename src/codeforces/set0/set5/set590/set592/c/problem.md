# Problem

Vector Willman and Array Bolt are the two most famous athletes of Byteforces. They are going to compete in a race with a distance of $L$ meters today.

Willman and Bolt have exactly the same speed, so when they compete the result is always a tie. That is a problem for the organizers because they want a winner.

While watching previous races the organizers have noticed that Willman can perform only steps of length equal to $w$ meters, and Bolt can perform only steps of length equal to $b$ meters. The organizers decided to slightly change the rules of the race. Now, at the end of the racetrack there will be an abyss, and the winner will be declared the athlete who manages to run farther from the starting point of the racetrack (which is not subject to change by any of the athletes).

Note that none of the athletes can run infinitely far: at some moment each will face a point such that one more step would cause them to fall into the abyss. In other words, an athlete will not fall into the abyss if the total length of all his steps is less than or equal to the chosen distance $L$.

Since the organizers are very fair, they are going to set the length of the racetrack as an integer chosen randomly and uniformly in the range from $1$ to $t$ (inclusive). What is the probability that Willman and Bolt tie again today?

## Input

The first line contains three integers $t$, $w$ and $b$ ($1 \le t, w, b \le 5 \cdot 10^{18}$) — the maximum possible length of the racetrack, the length of Willman's steps, and the length of Bolt's steps respectively.

## Output

Print the answer as an **irreducible fraction** in the form `p/q`, following the sample format.

A fraction $p/q$ (with integers $p \ge 0$, $q > 0$) is called irreducible if there is no integer $d > 1$ such that both $p$ and $q$ are divisible by $d$.

## Examples

### Example 1

**Input:**

```
10 3 2
```

**Output:**

```
3/10
```

### Example 2

**Input:**

```
7 1 2
```

**Output:**

```
3/7
```

## Note

In the first sample, Willman and Bolt tie when $1$, $6$, or $7$ is chosen as the length of the racetrack.


## ideas
1. x = lcm(w, b), 只有当两人停在位置 pos = m*x 的地方，两个人才能平局
2. 且 L >= pos and L < pos + min(w, b) (否则) 肯定会有获胜的一方
3. 如果 w = b, result = 1/1
4. 这样的 let R = pos + min(w, b) 
5. R - L = min(w, b) 也就是每段的长度是固定的
6. 那就看有多少段了
7. t / x + 1 (如果t = 1) / x