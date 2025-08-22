# Problem B: Tank Message Transmission

## Problem Description

The Smart Beaver from ABBYY started cooperating with the Ministry of Defence. Now they train soldiers to move armoured columns. The training involves testing a new type of tanks that can transmit information. To test the new type of tanks, the training has a special exercise, its essence is as follows.

## Exercise Rules

Initially, the column consists of **n** tanks sequentially numbered from 1 to n in the order of position in the column from its beginning to its end. During the whole exercise, exactly **n** messages must be transferred from the beginning of the column to its end.

### Message Transfer Process

Transferring one message is as follows:
- The tank that goes first in the column transmits the message to some tank in the column
- The tank which received the message sends it further down the column
- The process is continued until the last tank receives the message
- It is possible that not all tanks in the column will receive the message — it is important that the last tank in the column should receive the message

### Tank Movement

After the last tank (tank number n) receives the message:
- It moves to the beginning of the column and sends another message to the end of the column in the same manner
- When the message reaches the last tank (tank number n-1), that tank moves to the beginning of the column and sends the next message to the end of the column
- This continues until the exercise is completed when the tanks in the column return to their original order, that is, immediately after tank number 1 moves to the beginning of the column

### Tank Order Changes

If the tanks were initially placed in the column in the order 1, 2, ..., n:
- After the first message: order changes to n, 1, ..., n-1
- After the second message: order changes to n-1, n, 1, ..., n-2
- And so on...

## Tank Characteristics

The tanks are constructed in a very peculiar way. The tank with number i is characterized by one integer **ai**, which is called the **message receiving radius** of this tank.

## Message Transmission Rules

Transferring a message between two tanks takes **one second**, however, not always one tank can transmit a message to another one.

Consider two tanks in the column:
- First tank is the i-th in the column counting from the beginning
- Second tank is the j-th in the column, with tank number x

The first tank can transmit a message to the second tank if:
- i < j **AND**
- i ≥ j - ax

## Objective

The Ministry of Defense (and soon the Smart Beaver) faced the question of how to organize the training efficiently. The exercise should be finished as quickly as possible. We'll neglect the time that the tanks spend on moving along the column, since improving the tanks' speed is not a priority for this training.

You are given the number of tanks, as well as the message receiving radii of all tanks. You must help the Smart Beaver and organize the transferring of messages in a way that makes the total transmission time of all messages as small as possible.

## Input

- **First line**: integer **n** — the number of tanks in the column
- **Next n lines**: each contains one integer **ai** (1 ≤ ai ≤ 250000, 1 ≤ i ≤ n) — the message receiving radii of the tanks in the order from tank 1 to tank n

### Test Groups

- **First group**: 2 ≤ n ≤ 300
- **Second group**: 2 ≤ n ≤ 10000  
- **Third group**: 2 ≤ n ≤ 250000

## Output

Print a single integer — the minimum possible total time of transmitting the messages.

## Examples

### Example 1

**Input:**
```
3
2
1
1
```

**Output:**
```
5
```

### Example 2

**Input:**
```
5
2
2
2
2
2
```

**Output:**
```
10
```

## Notes

**Example 1 Explanation:**
- Original order: 1, 2, 3
- First message: 1→2→3 (2 seconds)
- Tank 3 moves to beginning: order becomes 3, 1, 2
- Second message: 3→1→2 (2 seconds)
- Tank 2 moves to beginning: order becomes 2, 3, 1
- Third message: 2→1 (1 second, since tank 1's radius is large enough)
- Tanks return to original order: 1, 2, 3
- **Total time: 5 seconds**

**Example 2 Explanation:**
- All five tanks are the same
- Sending a single message takes 2 seconds
- **Total time: 10 seconds**

## Programming Notes

Please, do not use the `%lld` specifier to read or write 64-bit integers in C++. It is preferred to use the `cin`, `cout` streams or the `%I64d` specifier.


### ideas
1. 先模拟一下， 
2. 假设目前的头部是l, 那么尾巴 = (l - 1) （也有可能是n）
3. 如果 a[l-1] >= n - 1 那么可以在一秒内直接完成
4. 如果不能，那么l应该尽可能的传递消息到某个i, a[i] >= pos[i] - 1
5. 传递消息，应该尽可能的贪心
6. 感觉咋要从后面开始算呢？
7. 假设处理r， 那么就是找到 i >= r - a[r], 中间的最小的 为止j, 其中(i - a[i])最小
8. 也就是说 i - a[i] 越小越好（但是如果 i - a[i] <= i - n 又不需要那么小
9. 