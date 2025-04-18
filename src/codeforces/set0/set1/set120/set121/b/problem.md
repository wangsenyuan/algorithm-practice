Petya loves lucky numbers. Everybody knows that lucky numbers are positive integers whose decimal representation contains only the lucky digits 4 and 7. For example, numbers 47, 744, 4 are lucky and 5, 17, 467 are not.

Petya has a number consisting of n digits without leading zeroes. He represented it as an array of digits without leading zeroes. Let's call it d. The numeration starts with 1, starting from the most significant digit. Petya wants to perform the following operation k times: find the minimum x (1 ≤ x < n) such that dx = 4 and dx + 1 = 7, if x is odd, then to assign dx = dx + 1 = 4, otherwise to assign dx = dx + 1 = 7. Note that if no x was found, then the operation counts as completed and the array doesn't change at all.

You are given the initial number as an array of digits and the number k. Help Petya find the result of completing k operations.

### ideas
1. 假设 s[i,i+1] = 47, i是奇数，i+1就是偶数， 44
2.                  , i是偶数，...         77
3. 477，怎么变化呢？ 477 =>(1是奇数) 447 =>(2是偶数) 477 （这里变成一个loop了）
4.  所以，操作后，必须要立刻知道，下一个47的位置，如果这个位置就是i+1(i是奇数时)或者i-1(i是偶数时), 那么就无法继续了
5.  4477 => 4777 => 4477 
6.  否则的话，肯定是后面的某个位置
7.  也就是说，虽然k很大，但是确实可以模拟的。模拟次数，不超过n/2次