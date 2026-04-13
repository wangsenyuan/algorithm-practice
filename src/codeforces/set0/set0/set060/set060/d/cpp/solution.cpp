// Codeforces 60D — Mushrooms (memory-friendly).
// No per-value adjacency vectors: O(maxA) for pos[] + O(n) DSU.
// Build: g++ -std=c++17 -O2 -pipe -o solution solution.cpp
// (On Codeforces Linux you may add -static -s if you want a static stripped binary.)

#include <algorithm>
#include <iostream>
#include <numeric>
#include <vector>

using namespace std;

static constexpr int kMaxA = 10'000'000;
// Hypotenuse when both legs ≤ kMaxA is ≤ ⌈√2·kMaxA⌉; keep a small margin.
static constexpr int kMaxHyp = 15'000'000;

static int gcd(int a, int b) {
    while (b != 0) {
        int t = a % b;
        a = b;
        b = t;
    }
    return a;
}

struct DSU {
    vector<int> parent;
    vector<unsigned char> rankv;

    explicit DSU(int n) : parent(n), rankv(n, 0) {
        iota(parent.begin(), parent.end(), 0);
    }

    int find(int x) {
        if (parent[x] != x) {
            parent[x] = find(parent[x]);
        }
        return parent[x];
    }

    void unite(int a, int b) {
        a = find(a);
        b = find(b);
        if (a == b) {
            return;
        }
        if (rankv[a] < rankv[b]) {
            swap(a, b);
        }
        parent[b] = a;
        if (rankv[a] == rankv[b]) {
            ++rankv[a];
        }
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n;
    if (!(cin >> n)) {
        return 0;
    }

    vector<int> pos(static_cast<size_t>(kMaxA) + 1, 0);
    for (int i = 0; i < n; ++i) {
        int x;
        cin >> x;
        pos[static_cast<size_t>(x)] = i + 1; // 0 = absent
    }

    DSU dsu(n);

    auto unite_vals = [&](int u, int v) {
        int pu = pos[static_cast<size_t>(u)];
        int pv = pos[static_cast<size_t>(v)];
        if (pu != 0 && pv != 0) {
            dsu.unite(pu - 1, pv - 1);
        }
    };

    for (int m = 2; m * m + 1 <= kMaxHyp; ++m) {
        for (int k = 1; k < m && m * m + k * k <= kMaxHyp; ++k) {
            if (gcd(m, k) != 1 || ((m ^ k) & 1) == 0) {
                continue;
            }
            int leg_a = m * m - k * k;
            int leg_b = 2 * m * k;
            int hyp = m * m + k * k;
            int l1 = min(leg_a, leg_b);
            int l2 = max(leg_a, leg_b);
            if (l1 <= 0 || l1 > kMaxA || l2 > kMaxA) {
                continue;
            }
            if (gcd(l1, l2) != 1 || gcd(l1, hyp) != 1 || gcd(l2, hyp) != 1) {
                continue;
            }
            unite_vals(l1, l2);
            if (hyp <= kMaxA) {
                unite_vals(l1, hyp);
                unite_vals(l2, hyp);
            }
        }
    }

    int ans = 0;
    for (int i = 0; i < n; ++i) {
        if (dsu.find(i) == i) {
            ++ans;
        }
    }
    cout << ans << '\n';
    return 0;
}
