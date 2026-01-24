# Problem C

Astronaut Natasha arrived on Mars. She knows that the Martians are very poor aliens. To ensure a better life for the Mars citizens, their emperor decided to take tax from every tourist who visited the planet. Natasha is the inhabitant of Earth, therefore she had to pay the tax to enter the territory of Mars.

There are 𝑛 banknote denominations on Mars: the value of 𝑖-th banknote is 𝑎𝑖. Natasha has an infinite number of banknotes of each denomination.

Martians have 𝑘 fingers on their hands, so they use a number system with base 𝑘. In addition, the Martians consider the digit 𝑑 (in the number system with base 𝑘) divine. Thus, if the last digit in Natasha's tax amount written in the number system with the base 𝑘 is 𝑑, the Martians will be happy. Unfortunately, Natasha does not know the Martians' divine digit yet.

Determine for which values 𝑑 Natasha can make the Martians happy.

Natasha can use only her banknotes. Martians don't give her change.

## Input

The first line contains two integers 𝑛 and 𝑘 (1≤𝑛≤100000, 2≤𝑘≤100000) — the number of denominations of banknotes and the base of the number system on Mars.

The second line contains 𝑛 integers 𝑎₁,𝑎₂,…,𝑎ₙ (1≤𝑎ᵢ≤10⁹) — denominations of banknotes on Mars.

All numbers are given in decimal notation.

## Output

On the first line output the number of values 𝑑 for which Natasha can make the Martians happy.

In the second line, output all these values in increasing order.

Print all numbers in decimal notation.

## Examples

### Example 1

**Input:**
```
2 8
12 20
```

**Output:**
```
2
0 4
```

### Example 2

**Input:**
```
3 10
10 20 30
```

**Output:**
```
1
0
```

## Note

Consider the first test case. It uses the octal number system.

- If you take one banknote with the value of 12, you will get 14₈ in octal system. The last digit is 4₈.
- If you take one banknote with the value of 12 and one banknote with the value of 20, the total value will be 32. In the octal system, it is 40₈. The last digit is 0₈.
- If you take two banknotes with the value of 20, the total value will be 40, this is 50₈ in the octal system. The last digit is 0₈.

No other digits other than 0₈ and 4₈ can be obtained. Digits 0₈ and 4₈ could also be obtained in other ways.

The second test case uses the decimal number system. The nominals of all banknotes end with zero, so Natasha can give the Martians only the amount whose decimal notation also ends with zero.


### ideas
1. 就是说，要能用a[?]的组合，得到d？
2. let a[i] %= k
3. 就是看最后能用哪些d < k可以被表示出来
4. 如果dp[1] = true, 那么所有的都可以组合出来
5. 如果 gcd(d, k) = 1, 那么就可以获得dp[1]
6. 如果是[4, 6] k = 10