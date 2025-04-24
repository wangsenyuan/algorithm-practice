In the Main Berland Bank n people stand in a queue at the cashier, everyone knows his/her height hi, and the heights of the other people in the queue. Each of them keeps in mind number ai — how many people who are taller than him/her and stand in queue in front of him.

After a while the cashier has a lunch break and the people in the queue seat on the chairs in the waiting room in a random order.

When the lunch break was over, it turned out that nobody can remember the exact order of the people in the queue, but everyone remembers his number ai.

Your task is to restore the order in which the people stood in the queue if it is possible. There may be several acceptable orders, but you need to find any of them. Also, you need to print a possible set of numbers hi — the heights of people in the queue, so that the numbers ai are correct.

### ideas
  1. 如果某个人的a[i] = 0， 要么他排在第一位，要么他是最高的
  2. 按照a[i]升序处理，如果a[i] = x, 但是前面没有x个人时，那么就没有答案。否则，可以让h[i] 正好比前面x个人小1
  3. 