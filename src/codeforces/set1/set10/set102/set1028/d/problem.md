Let's consider a simplified version of order book of some stock. The order book is a list of orders (offers) from people that want to buy or sell one unit of the stock, each order is described by direction (BUY or SELL) and price.

At every moment of time, every SELL offer has higher price than every BUY offer.

In this problem no two ever existed orders will have the same price.

The lowest-price SELL order and the highest-price BUY order are called the best offers, marked with black frames on the picture below.

The presented order book says that someone wants to sell the product at price 12 and it's the best SELL offer because the other two have higher prices. The best BUY offer has price 10.

There are two possible actions in this orderbook:

- Somebody adds a new order of some direction with some price.
- Somebody accepts the best possible SELL or BUY offer (makes a deal). It's impossible to accept not the best SELL or BUY offer (to make a deal at worse price). After someone accepts the offer, it is removed from the orderbook forever.

It is allowed to add new BUY order only with prices less than the best SELL offer (if you want to buy stock for higher price, then instead of adding an order you should accept the best SELL offer). Similarly, one couldn't add a new SELL order with price less or equal to the best BUY offer. For example, you can't add a new offer "SELL 20" if there is already an offer "BUY 20" or "BUY 25" — in this case you just accept the best BUY offer.

You have a damaged order book log (in the beginning there are no orders in book). Every action has one of the two types:

- "ADD $p$" denotes adding a new order with price $p$ and unknown direction. The order must not contradict with orders still not removed from the order book.
- "ACCEPT $p$" denotes accepting an existing best offer with price $p$ and unknown direction.

The directions of all actions are lost. Information from the log isn't always enough to determine these directions. Count the number of ways to correctly restore all ADD action directions so that all the described conditions are satisfied at any moment. Since the answer could be large, output it modulo $10^9 + 7$. If it is impossible to correctly restore directions, then output $0$.

## Input

The first line contains an integer $n$ ($1 \le n \le 363304$) — the number of actions in the log.

Each of the next $n$ lines contains a string "ACCEPT" or "ADD" and an integer $p$ ($1 \le p \le 308983066$), describing an action type and price.

All ADD actions have different prices. For ACCEPT action it is guaranteed that the order with the same price has already been added but has not been accepted yet.

## Output

Output the number of ways to restore directions of ADD actions modulo $10^9 + 7$.

## Examples

### Input
```
6
ADD 1
ACCEPT 1
ADD 2
ACCEPT 2
ADD 3
ACCEPT 3
```

### Output
```
8
```

### Input
```
4
ADD 1
ADD 2
ADD 3
ACCEPT 2
```

### Output
```
2
```

### Input
```
7
ADD 1
ADD 2
ADD 3
ADD 4
ADD 5
ACCEPT 3
ACCEPT 5
```

### Output
```
0
```

## Note

In the first example each of orders may be BUY or SELL.

In the second example the order with price $1$ has to be BUY order, the order with the price $3$ has to be SELL order.


### ideas
1. 假设目前有n个价格，且best sell价格是i，那么best buy价格 = i - 1
2. 如果 add p, 如果p < best sell, 但是 > best buy, 那么这个是有效的，且它有可能是sell，也有可能是buy
3. 如果 p > best sell, 那么它只能是sell， p < best buy, 那么它只能是buy
4. accept p => 如果accept的是sell，那么sell的最低价格变成新的（但是buy不变）
5. 是不是正难则反呢？
6. 最后剩余的价格列表是确定的，
7. 然后反过来处理某个accept p，那么这个p有可能是sell、也有可能是buy，是个best
8. *2, 如果反过来add p。。 好像也不大对
9. dp[i] 表示到目前为止，best sell = price[i]的方案数
10. add p, 如果 p > price[i] => dp[i] 不变
11.        如果 p < price[i] 且 p > price[i-1] => dp[pos[p]] += dp[i]
12.            p < price[i-1] => dp[i] 不变
13. accept p => 只保留 dp[pos[p]] dp[pos[p1]] p1 > p
14.       所以，accept的时候，有个清0的操作，是个range update
15. 不大对～