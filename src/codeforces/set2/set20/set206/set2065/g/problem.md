Skibidus was abducted by aliens of Amog! Skibidus tries to talk his way out, but the Amog aliens don't believe him. To prove that he is not totally capping, the Amog aliens asked him to solve this task:

An integer 𝑥
 is considered a semi-prime if it can be written as 𝑝⋅𝑞
 where 𝑝
 and 𝑞
 are (not necessarily distinct) prime numbers. For example, 9
 is a semi-prime since it can be written as 3⋅3
, and 3
 is a prime number.

Skibidus was given an array 𝑎
 containing 𝑛
 integers. He must report the number of pairs (𝑖,𝑗)
 such that 𝑖≤𝑗
 and lcm(𝑎𝑖,𝑎𝑗)
∗
 is semi-prime.

 ### ideas
 1. if a[i] and a[j] both primes， good
 2. lcm(a[i], a[j]) = a[i] * a[j] / gcd(a[i], a[j])
 3. 所以，如果 a[j] % a[i] = 0, and a[j] / a[i] is semi prime
 4. a[i] = 6, a[j] = 9, lcm(6, 9) = 54 / 3 = 18 = 2 * 9, 不是
 5. 