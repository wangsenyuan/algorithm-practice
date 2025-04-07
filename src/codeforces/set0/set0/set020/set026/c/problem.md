Once Bob decided to lay a parquet floor in his living room. The living room is of size n × m metres. Bob had planks of three types: a planks 1 × 2 meters, b planks 2 × 1 meters, and c planks 2 × 2 meters. Help Bob find out, if it is possible to parquet the living room with such a set of planks, and if it is possible, find one of the possible ways to do so. Bob doesn't have to use all the planks.

Input
The first input line contains 5 space-separated integer numbers n, m, a, b, c (1 ≤ n, m ≤ 100, 0 ≤ a, b, c ≤ 104), n and m — the living room dimensions, a, b and c — amount of planks 1 × 2, 2 × 1 и 2 × 2 respectively. It's not allowed to turn the planks.

### ideas
1. 如果n = 3， m = 3的情况下，似乎是没有解的
2. 如果n是奇数，那么必须有一整行使用了1 * 2的结构（所以要求m必须是偶数）
3. 同理，对于m是奇数页成立
4. 所以 n, m 同是奇数时没有答案
5. 如果 n是奇数（m是偶数）那么应该先使用 1 * 2， 排满一行（最下面，或者最上面）
6. 然后处理 n, m都是偶数的情况
7. 