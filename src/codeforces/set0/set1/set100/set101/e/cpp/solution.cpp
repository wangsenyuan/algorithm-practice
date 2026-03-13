#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n, m, p;
    cin >> n >> m >> p;

    vector<int> x(n), y(m);
    for (int i = 0; i < n; ++i) {
        cin >> x[i];
    }
    for (int i = 0; i < m; ++i) {
        cin >> y[i];
    }

    vector<int> arr1(n), arr2(n), arr3(n);
    const int neg = -(1 << 30);
    auto f = [&](int i, int j) -> int {
        return (x[i] + y[j]) % p;
    };

    auto dp = [&](int r1, int c1, int r2, int c2, int k) -> vector<int>& {
        for (int i = r1; i <= r2; ++i) {
            arr1[i] = neg;
            arr2[i] = neg;
        }
        arr1[r1] = f(r1, c1);
        int prev_lo = r1;
        int prev_hi = r1;

        for (int diag = r1 + c1 + 1; diag <= k; ++diag) {
            int lo = max(r1, diag - c2);
            int hi = min(r2, diag - c1);
            for (int i = lo; i <= hi; ++i) {
                int j = diag - i;
                arr2[i] = f(i, j);
                int best_here = neg;
                if (i > r1 && prev_lo <= i - 1 && i - 1 <= prev_hi) {
                    best_here = max(best_here, arr1[i - 1]);
                }
                if (j > c1 && prev_lo <= i && i <= prev_hi) {
                    best_here = max(best_here, arr1[i]);
                }
                arr2[i] += best_here;
            }
            for (int i = prev_lo; i <= prev_hi; ++i) {
                arr1[i] = neg;
            }
            for (int i = lo; i <= hi; ++i) {
                arr1[i] = arr2[i];
                arr2[i] = neg;
            }
            prev_lo = lo;
            prev_hi = hi;
        }
        return arr1;
    };

    auto fp = [&](int r1, int c1, int r2, int c2, int k) -> vector<int>& {
        for (int i = r1; i <= r2; ++i) {
            arr3[i] = neg;
            arr2[i] = neg;
        }
        arr3[r2] = f(r2, c2);
        int prev_lo = r2;
        int prev_hi = r2;

        for (int diag = r2 + c2 - 1; diag >= k; --diag) {
            int lo = max(r1, diag - c2);
            int hi = min(r2, diag - c1);
            for (int i = lo; i <= hi; ++i) {
                int j = diag - i;
                arr2[i] = f(i, j);
                int best_here = neg;
                if (i < r2 && prev_lo <= i + 1 && i + 1 <= prev_hi) {
                    best_here = max(best_here, arr3[i + 1]);
                }
                if (j < c2 && prev_lo <= i && i <= prev_hi) {
                    best_here = max(best_here, arr3[i]);
                }
                arr2[i] += best_here;
            }
            for (int i = prev_lo; i <= prev_hi; ++i) {
                arr3[i] = neg;
            }
            for (int i = lo; i <= hi; ++i) {
                arr3[i] = arr2[i];
                arr2[i] = neg;
            }
            prev_lo = lo;
            prev_hi = hi;
        }
        return arr3;
    };

    string path(n + m - 1, '\0');

    auto build = [&](auto&& self, int r1, int c1, int r2, int c2) -> int {
        int sum = 0;
        if (r1 == r2) {
            for (int j = c1; j < c2; ++j) {
                path[r1 + j + 1] = 'S';
                sum += f(r1, j + 1);
            }
            return sum;
        }
        if (c1 == c2) {
            for (int i = r1; i < r2; ++i) {
                path[i + c1 + 1] = 'C';
                sum += f(i + 1, c1);
            }
            return sum;
        }

        int k = (r1 + c1 + r2 + c2) / 2;
        vector<int>& d1 = dp(r1, c1, r2, c2, k);
        vector<int>& d2 = fp(r1, c1, r2, c2, k);

        int lo = max(r1, k - c2);
        int hi = min(r2, k - c1);

        int best_score = neg;
        int best_row = r1;
        for (int i = lo; i <= hi; ++i) {
            int j = k - i;
            int cur = d1[i] + d2[i] - f(i, j);
            if (cur > best_score) {
                best_score = cur;
                best_row = i;
            }
        }

        int i = best_row;
        int j = k - best_row;
        sum += self(self, r1, c1, i, j);
        sum += self(self, i, j, r2, c2);
        return sum;
    };

    int best = f(0, 0) + build(build, 0, 0, n - 1, m - 1);
    cout << best << '\n';
    cout << path.substr(1) << '\n';
    return 0;
}
