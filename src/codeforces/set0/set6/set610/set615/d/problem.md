Ayrat has number n, represented as it's prime factorization pi of size m, i.e. $n = p_1 \cdot p_2 \cdot ... \cdot p_m$. Ayrat got secret
information that the product of all divisors of n taken modulo $10^9 + 7$ is the password to the secret data base. Now he wants to
calculate this value.

## Input

The first line of the input contains a single integer m ($1 \le m \le 200000$) — the number of primes in factorization of n.

The second line contains m primes numbers $p_i$ ($2 \le p_i \le 200000$).

## Output

Print one integer — the product of all divisors of n modulo $10^9 + 7$.

## Examples

### Example 1

**Input**
```
2
2 3
```

**Output**
```
36
```

### Example 2

**Input**
```
3
2 3 2
```

**Output**
```
1728
```

## Note

In the first sample $n = 2 \cdot 3 = 6$. The divisors of 6 are 1, 2, 3 and 6, their product is equal to $1 \cdot 2 \cdot 3 \cdot 6 = 36$.

In the second sample $2 \cdot 3 \cdot 2 = 12$. The divisors of 12 are 1, 2, 3, 4, 6 and 12. $1 \cdot 2 \cdot 3 \cdot 4 \cdot 6 \cdot 12 = 1728$.


### ideas
1. n = 6, p = [2, 3]
2. n = 2 * 2 * 2 * 3 * 3 * 5
3. 考虑 [2 ** 1] 出现了 (其他的数的组合数) = [1, 3, 5, 9, 15, 45] = [3 * 2]
4.     [2 ** 2] 出现了 6次
5.     [2 ** 3] 出现了 6次
6. = [2 ** 1 * 2 ** 2 * 2 ** 3] ** 6
7. = 2 ** ((1 + 2 + 3)) ** 6
8. 回到 n = 6
9. 4 * 9 = 36
10. n = 12
11. pow(2, 1 + 2) ** 2 = 8 * 8 = 64
12. pow(3, 1) ** 3 = 27  


### AI

当指数 `b` 非常大（比如超过1000位）时，我们需要用到**费马小定理**或**欧拉定理**来降低指数的大小。

## 核心思想

### 1. 费马小定理（当 mod 是素数时）

如果 `p` 是素数，且 `gcd(a, p) = 1`，则：

$$a^{p-1} \equiv 1 \pmod{p}$$

因此：

$$a^b \equiv a^{b \bmod (p-1)} \pmod{p}$$

### 2. 欧拉定理（更一般的情况）

如果 `gcd(a, mod) = 1`，则：

$$a^{\phi(mod)} \equiv 1 \pmod{mod}$$

因此：

$$a^b \equiv a^{b \bmod \phi(mod)} \pmod{mod}$$

其中 $\phi(mod)$ 是欧拉函数。

## Go 代码实现

```go
package main

import (
    "fmt"
    "math/big"
)

const MOD = 1000000007

// 使用费马小定理计算 pow(a, b, mod)，其中 b 可以非常大
// 当 mod 是素数时，a^b ≡ a^(b mod (mod-1)) (mod mod)
func powBigExponent(a int64, b *big.Int, mod int64) int64 {
    if mod == 1 {
        return 0
    }
    
    // 费马小定理：a^(p-1) ≡ 1 (mod p)
    // 所以 a^b ≡ a^(b mod (p-1)) (mod p)
    modMinus1 := big.NewInt(mod - 1)
    
    // 计算 b mod (mod - 1)
    bReduced := new(big.Int).Mod(b, modMinus1)
    
    // 将 bReduced 转换为 int64（现在它足够小了）
    bSmall := bReduced.Int64()
    
    // 使用快速幂计算
    return fastPow(a, bSmall, mod)
}

// 快速幂：计算 a^b mod m
func fastPow(a, b, m int64) int64 {
    result := int64(1)
    base := a % m
    
    for b > 0 {
        if b&1 == 1 {
            result = (result * base) % m
        }
        base = (base * base) % m
        b >>= 1
    }
    
    return result
}

// 示例：处理超大指数
func main() {
    // 例子1：计算 2^(10^100) mod (10^9 + 7)
    a := int64(2)
    
    // 创建一个非常大的指数：10^100
    b := new(big.Int)
    b.SetString("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 10)
    
    result := powBigExponent(a, b, MOD)
    fmt.Printf("2^(10^100) mod %d = %d\n", MOD, result)
    
    // 例子2：在当前问题中的应用
    // 假设需要计算 某个数^(divisor_count) mod (10^9 + 7)
    // 而 divisor_count 可能非常大
    
    // 假设 divisor_count 是一个很大的数
    divisorCount := new(big.Int)
    divisorCount.SetString("123456789012345678901234567890", 10)
    
    base := int64(12)
    result2 := powBigExponent(base, divisorCount, MOD)
    fmt.Printf("%d^%s mod %d = %d\n", base, divisorCount.String(), MOD, result2)
}
```

## 在当前问题中的应用

对于这道题（计算所有因子的乘积），当我们需要计算 $p^{exp}$ 其中 `exp` 可能很大时：

```go
// 计算除数个数（可能很大）
func calculateDivisorCount(primeFreq map[int]int) *big.Int {
    count := big.NewInt(1)
    for _, freq := range primeFreq {
        count.Mul(count, big.NewInt(int64(freq + 1)))
    }
    return count
}

// 使用费马小定理计算结果
func solve(primes []int) int64 {
    // 统计每个质数的频率
    freq := make(map[int]int)
    for _, p := range primes {
        freq[p]++
    }
    
    // 计算除数个数（可能非常大）
    divisorCount := calculateDivisorCount(freq)
    
    result := int64(1)
    const MOD = 1000000007
    
    for prime, f := range freq {
        // 计算这个质数的总贡献指数
        totalExp := big.NewInt(int64(f * (f + 1) / 2))
        totalExp.Mul(totalExp, new(big.Int).Div(divisorCount, big.NewInt(int64(f + 1))))
        
        // 使用费马小定理计算 prime^totalExp mod MOD
        contribution := powBigExponent(int64(prime), totalExp, MOD)
        result = (result * contribution) % MOD
    }
    
    return result
}
```

## 关键点总结

1. **对于 $10^9 + 7$ 这样的素数模**：使用 $b \bmod (10^9 + 6)$ 来降低指数
2. **使用 `math/big` 包**：处理超大数字的模运算
3. **组合使用**：先用 `big.Int` 做模运算降低指数，再用普通的快速幂计算结果

这样即使指数有1000位，也能高效计算出结果！