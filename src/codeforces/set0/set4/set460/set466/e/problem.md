There are n employees working in company "X" (let's number them from 1 to n for convenience). Initially the employees didn't have any relationships among each other. On each of m next days one of the following events took place:

- either employee y became the boss of employee x (at that, employee x didn't have a boss before);
- or employee x gets a packet of documents and signs them; then he gives the packet to his boss. The boss signs the documents and gives them to his boss and so on (the last person to sign the documents sends them to the archive);
- or comes a request of type "determine whether employee x signs certain documents".

Your task is to write a program that will, given the events, answer the queries of the described type. At that, it is guaranteed that throughout the whole working time the company didn't have cyclic dependencies.

### Input

The first line contains two integers n and m $(1 \le n, m \le 10^5)$ — the number of employees and the number of events.

Each of the next m lines contains the description of one event (the events are given in the chronological order). The first number of the line determines the type of event t $(1 \le t \le 3)$.

- If $t = 1$, then next follow two integers x and y $(1 \le x, y \le n)$ — numbers of the company employees. It is guaranteed that employee x doesn't have the boss currently.
- If $t = 2$, then next follow integer x $(1 \le x \le n)$ — the number of the employee who got a document packet.
- If $t = 3$, then next follow two integers x and i $(1 \le x \le n; 1 \le i \le$ [number of packets that have already been given]) — the employee and the number of the document packet for which you need to find out information. The document packets are numbered started from 1 in the chronological order.

It is guaranteed that the input has at least one query of the third type.

### Output

For each query of the third type print "YES" if the employee signed the document package and "NO" otherwise. Print all the words without the quotes.

### Example

#### Input
```
4 9
1 4 3
2 4
3 3 1
1 2 3
2 2
3 1 2
1 3 1
2 2
3 1 3
```

#### Output
```
YES
NO
YES
```

### ideas
1. 有点像DSU，但似乎又不是～
2. 操作1将y变成x的boss，操作2的时候，沿着x的路径往上传递，
3. 假设先搞出最终的树（这样把结构就确定了）
4. 然后对于操作2，需要知道它在哪个地方停下来；
5. 对于操作3，只要放过来计算i，停在哪个位置，就可以知道x是否sign过
6. 关键是如何知道在哪个地方停下来
7. 操作2，它的操作序列是知道的，就是要知道，离它最近的，且操作序列比它大的位置（貌似可以用一个segment tree来处理？）
8. 