Yaroslav has an array p = p1, p2, ..., pn (1 ≤ pi ≤ n), consisting of n distinct integers. Also, he has m queries:

Query number i is represented as a pair of integers li, ri (1 ≤ li ≤ ri ≤ n).
The answer to the query li, ri is the number of pairs of integers q, w (li ≤ q, w ≤ ri) such that pq is the divisor of pw.
Help Yaroslav, answer all his queries.

### ideas
1. 假设处理x, 那么2 * x, 3 * x, ... i * x的位置是知道的
2. 如果有一个查询区间l...r 包含 x...i*x, 那么这个查询的结果+1
3. 那么也就是任何查询，包含区间min(pos[x], pos[i*x]), max(pos[x], pos[i*x])的查询+1
4. 有点像二维空间内的点？
5. 假设有一个查询，包含了左端点l，那么不考虑r的时候，都可以+1
6. 考虑r的时候，那么就是r后面的区间+1
7. 假设查询区间，按照l从小到大排列
8. 另外一个按照r从小到大排列
9. 对于这样的一对（x, y), 所有在x前面的+1，所有x前面的r的排列-1
10. 这个地方进行不下去了.
11. 假设是个二维空间。一维表示l，第二维表示r
12. 那么对于(x, y), 只考虑x时，所以查询区间l在x的前面的部分，都可以+1
13. 所有区间在y后面的部分，也+1
14. 但是那些，l在x后面，或者r在y前面的部分，要怎么处理呢？