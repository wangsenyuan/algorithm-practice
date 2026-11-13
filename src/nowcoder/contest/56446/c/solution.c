/* Nowcoder contest 56446 C
 * GCC 10: gcc -std=c11 -O2 solution.c
 */

#include <stdio.h>
#include <stdlib.h>

int main(void) {
    int n, q;
    if (scanf("%d %d", &n, &q) != 2) return 0;

    long long *diff0 = calloc((size_t)n + 2, sizeof(long long));
    long long *diff1 = calloc((size_t)n + 2, sizeof(long long));
    if (!diff0 || !diff1) return 0;

    for (int k = 0; k < q; ++k) {
        int x, y;
        scanf("%d %d", &x, &y);
        /* a[j] += max(y-|x-j|, 0)，正贡献在 [x-y+1, x+y-1]
         * [l, x] 上是 j+(y-x)，斜率 +1
         * [x+1, r] 上是 -j+(x+y)，斜率 -1 */
        int l = x - y + 1;
        if (l < 1) l = 1;
        int r = x + y - 1;
        if (r > n) r = n;

        diff1[l]++;
        diff1[x + 1]--;
        diff0[l] += (long long)l - 1 + y - x;
        diff0[x + 1] -= y;
        if (x < r) {
            diff1[x + 1]--;
            diff1[r + 1]++;
            diff0[x + 1] += y;
            diff0[r + 1] += (long long)r - x - y;
        }
    }

    long long diff = 0;
    for (int i = 1; i <= n; ++i) {
        diff0[i] += diff0[i - 1];
        diff1[i] += diff1[i - 1];
        diff += diff1[i];
        if (i > 1) printf(" ");
        printf("%lld", diff0[i] + diff);
    }
    printf("\n");

    free(diff0);
    free(diff1);
    return 0;
}
