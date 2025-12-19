### Problem

Alex doesn't like boredom. That's why whenever he gets bored, he comes up with games. One long winter evening he came up with a game and decided to play it.

Given a sequence $a$ consisting of $n$ integers. The player can make several steps. In a single step he can choose an element of the sequence (let's denote it $a_k$) and delete it, at that all elements equal to $a_k + 1$ and $a_k - 1$ also must be deleted from the sequence. That step brings $a_k$ points to the player.

Alex is a perfectionist, so he decided to get as many points as possible. Help him.

### Input

The first line contains integer $n$ $(1 \le n \le 10^5)$ that shows how many numbers are in Alex's sequence.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ $(1 \le a_i \le 10^5)$.

### Output

Print a single integer — the maximum number of points that Alex can earn.

### Examples

**Input**

```text
2
1 2
```

**Output**

```text
2
```

**Input**

```text
3
1 2 3
```

**Output**

```text
4
```

**Input**

```text
9
1 2 1 3 2 2 2 2 3
```

**Output**

```text
10
```

### Note

Consider the third test example. At first step we need to choose any element equal to $2$. After that step our sequence looks like this $[2, 2, 2, 2]$. Then we do $4$ steps, on each step we choose any element equals to $2$. In total we earn $10$ points.
