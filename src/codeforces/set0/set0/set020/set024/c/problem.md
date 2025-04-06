You are given the following points with integer coordinates on the plane: M0, A0, A1, ..., An - 1, where n is odd number. Now we define the following infinite sequence of points Mi: Mi is symmetric to Mi - 1 according  (for every natural number i). Here point B is symmetric to A according M, if M is the center of the line segment AB. Given index j find the point Mj.

### ideas
1. 假设M[n]计算出来后， 要计算M[n+1], 那么 M[n]..M[n+1] 关于A[0]对称
2. M[0] = (x0, y0)
3. M[1] = (2 * a0 - x0, 2 * b0 - y0)
4. M[2] = (2 * a1 - x1, 2 * b1 - y1) = (2 * a1 - 2 * a0 + x0, 2 * b1 - 2 * b0 + y0)
5. M[3] = ... = (2 * a2 - 2 * a1 + 2 * a0 - x0, 2 * b2 - 2 * b1 + 2 * b0 - y0)
6. M[i][0] = 2 * ai - 2 * a[i-1] + 2 * a[i-2] ... + ... +/- x0
7. M[i][1] = 2 * bi - 2 * b[i-1] + ... +/- y0
8. M[n][0] = 2 * a(n-1) + ....+ 2 * a[0] - x0
9. M[n][1] = 2 * b[n-1] + .... - y0
10. M[n+1][0] = 2 * a[0] - (.... + 2 * a[0] - x0) (这里 a[0]就没有了)
11. 。。。
12. M[2 * n][0] = x0 （也就是说，循环两次后， M[2 * n] = M[0]
13. 也就是说，计算出前 2 * n 个， 然后 M[j] = M[j % (2 * n)]