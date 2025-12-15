Some time ago Mister B detected a strange signal from space, which he started to study.

After some transformation the signal turned out to be a permutation `p` of length `n` or its cyclic shift. For further investigation Mister B needs some basis, so he decided to choose a cyclic shift of this permutation which has the minimum possible deviation.

Let's define the deviation of a permutation `p` as .

Find a cyclic shift of permutation `p` with minimum possible deviation. If there are multiple solutions, print any of them.

Let's denote id `k` (`0 ≤ k < n`) of a cyclic shift of permutation `p` as the number of right shifts needed to reach this shift, for example:

- `k = 0`: shift `p1, p2, ... pn`
- `k = 1`: shift `pn, p1, ... pn - 1`
- ...
- `k = n - 1`: shift `p2, p3, ... pn, p1`

## Input

First line contains single integer `n` (`2 ≤ n ≤ 10^6`) — the length of the permutation.

The second line contains `n` space-separated integers `p1, p2, ..., pn` (`1 ≤ pi ≤ n`) — the elements of the permutation. It is guaranteed that all elements are distinct.

## Output

Print two integers: the minimum deviation of cyclic shifts of permutation `p` and the id of such shift. If there are multiple solutions, print any of them.

## Examples

### Input

3  
1 2 3

### Output

0 0

### Input

3  
2 3 1

### Output

0 1

### Input

3  
3 2 1

### Output

2 1

## Note

In the first sample test the given permutation `p` is the identity permutation, that's why its deviation equals `0`, the shift id equals `0` as well.

In the second sample test the deviation of `p` equals to `4`, the deviation of the 1st cyclic shift `(1, 2, 3)` equals `0`, the deviation of the 2nd cyclic shift `(3, 1, 2)` equals to `4`, the optimal is the 1st cyclic shift.

In the third sample test the deviation of `p` equals to `4`, the deviation of the 1st cyclic shift `(1, 3, 2)` equals to `2`, the deviation of the 2nd cyclic shift `(2, 1, 3)` also equals to `2`, so the optimal are both 1st and 2nd cyclic shifts.


### ideas
1. 规律是什么？
2. 如果 p[j] = i, 且j < i, 那么当增加k的时候，i - p[j]会越来越小
3. 反过来也是成立的, p[j] = i, 且 i < j, 那么增加k的时候， i - p[j]会越来越大。
4. 所以，要找到这样的pair， 如果j < i 和 j > i的对，如果小的更多，增加k，如果大的更多，减小k
5. 和距离有关系