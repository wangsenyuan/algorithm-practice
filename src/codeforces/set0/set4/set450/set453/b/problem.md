### Problem

A sequence of positive integers $b_i$ is harmony if and only if for every two elements of the sequence their greatest common divisor equals 1. According to an ancient book, the key of the chest is a harmony sequence $b_i$ which minimizes the following expression:

$$\sum_{i=1}^{n} |a_i - b_i|$$

You are given sequence $a_i$, help Princess Twilight to find the key.

### Input

The first line contains an integer $n$ $(1 \le n \le 100)$ — the number of elements of the sequences $a$ and $b$. The next line contains $n$ integers $a_1, a_2, \ldots, a_n$ $(1 \le a_i \le 30)$.

### Output

Output the key — sequence $b_i$ that minimizes the sum described above. If there are multiple optimal sequences, you can output any of them.

### Examples

**Input**

```text
5
1 1 1 1 1
```

**Output**

```text
1 1 1 1 1
```

**Input**

```text
5
1 6 4 2 8
```

**Output**

```text
1 5 3 1 8
```


### ideas
1. a[i] <= 30, 那么b[i] >= 60 肯定不好（替换成1，反而得到了更好的结果) => b[i] < 60
2. 2，3，5,7,11,13,17,23,29 这些数字只能最多出现一次
3. 其他的都是1.那就简单了