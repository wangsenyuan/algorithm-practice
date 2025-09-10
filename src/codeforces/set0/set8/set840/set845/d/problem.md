# Problem D - Road Signs

Polycarp has just attempted to pass the driving test. He ran over the straight road with the signs of four types.

## Road Sign Types

- **Speed limit**: this sign comes with a positive integer number — maximal speed of the car after the sign (cancel the action of the previous sign of this type)
- **Overtake is allowed**: this sign means that after some car meets it, it can overtake any other car
- **No speed limit**: this sign cancels speed limit if any (car can move with arbitrary speed after this sign)
- **No overtake allowed**: some car can't overtake any other car after this sign

Polycarp goes past the signs consequentially, each new sign cancels the action of all the previous signs of it's kind (speed limit/overtake). It is possible that two or more "no overtake allowed" signs go one after another with zero "overtake is allowed" signs between them. It works with "no speed limit" and "overtake is allowed" signs as well.

In the beginning of the ride overtake is allowed and there is no speed limit.

## Event Types

You are given the sequence of events in chronological order — events which happened to Polycarp during the ride. There are events of following types:

1. Polycarp changes the speed of his car to specified (this event comes with a positive integer number)
2. Polycarp's car overtakes the other car
3. Polycarp's car goes past the "speed limit" sign (this sign comes with a positive integer)
4. Polycarp's car goes past the "overtake is allowed" sign
5. Polycarp's car goes past the "no speed limit"
6. Polycarp's car goes past the "no overtake allowed"

It is guaranteed that the first event in chronological order is the event of type 1 (Polycarp changed the speed of his car to specified).

## Problem Statement

After the exam Polycarp can justify his rule violations by telling the driving instructor that he just didn't notice some of the signs. What is the minimal number of signs Polycarp should say he didn't notice, so that he would make no rule violations from his point of view?

## Input

The first line contains one integer number n (1 ≤ n ≤ 2·10^5) — number of events.

Each of the next n lines starts with integer t (1 ≤ t ≤ 6) — the type of the event.

An integer s (1 ≤ s ≤ 300) follows in the query of the first and the third type (if it is the query of first type, then it's new speed of Polycarp's car, if it is the query of third type, then it's new speed limit).

It is guaranteed that the first event in chronological order is the event of type 1 (Polycarp changed the speed of his car to specified).

## Output

Print the minimal number of road signs Polycarp should say he didn't notice, so that he would make no rule violations from his point of view.

## Examples

### Example 1

**Input:**
```
11
1 100
3 70
4
2
3 120
5
3 120
6
1 150
4
3 300
```

**Output:**
```
2
```

### Example 2

**Input:**
```
5
1 100
3 200
2
4
5
```

**Output:**
```
0
```

### Example 3

**Input:**
```
7
1 20
2
6
4
6
6
2
```

**Output:**
```
2
```

## Note

- In the first example Polycarp should say he didn't notice the "speed limit" sign with the limit of 70 and the second "speed limit" sign with the limit of 120.
- In the second example Polycarp didn't make any rule violation.
- In the third example Polycarp should say he didn't notice both "no overtake allowed" that came after "overtake is allowed" sign.

### ideas
1. 有点乱
2. 假设在现在这个顺序下，P违反了x条规则，然后删除最少的sign，以使的P没有违反规则
3. 假设目前的速度是s，但是目前的限速是s1, s1 < s， 那么P违反了规则
4. 如果目前不能超车，但是P超车了，那么P也违反了规则
5. 如果现在超车了，但是那么就必须找到最近的允许超车的sign，然后把这之后的不允许超车的sign给取消掉
6. 如果现在的速度是s, 那么必须找到最近的限速 si >= s，或者是最近的，没有限速的标志
7. 那么这样子，不就是贪心就可以了吗？
8. 貌似只有限速这个，需要快速的找到最近的超过s的位置（可以把没有限速，当作限速inf来处理）
9. 这个貌似可以用stack来处理？