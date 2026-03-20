// CF 246E — distinct names among k-th descendants (forest).
// Memory: O(n) for tree + per-depth segment trees (sum of sizes = n) + O(n) maps along DFS path.

#include <algorithm>
#include <functional>
#include <iostream>
#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

using namespace std;

struct SegTree {
    int n;
    vector<int> t;

    explicit SegTree(int sz) : n(sz) {
        if (n <= 0) n = 1;
        t.assign(4 * n, 0);
    }

    void update(int p, int v, int i = 0, int l = 0, int r = -1) {
        if (r < 0) r = n - 1;
        if (l == r) {
            t[i] = v;
            return;
        }
        int mid = (l + r) >> 1;
        if (p <= mid) {
            update(p, v, i * 2 + 1, l, mid);
        } else {
            update(p, v, i * 2 + 2, mid + 1, r);
        }
        t[i] = t[i * 2 + 1] + t[i * 2 + 2];
    }

    int query(int L, int R, int i = 0, int l = 0, int r = -1) {
        if (r < 0) r = n - 1;
        if (L > R) return 0;
        if (L == l && R == r) return t[i];
        int mid = (l + r) >> 1;
        int res = 0;
        if (L <= mid) {
            res += query(L, min(R, mid), i * 2 + 1, l, mid);
        }
        if (mid < R) {
            res += query(max(L, mid + 1), R, i * 2 + 2, mid + 1, r);
        }
        return res;
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n;
    if (!(cin >> n)) return 0;
    vector<string> names(n);
    vector<vector<int>> adj(n);
    vector<int> roots;
    roots.reserve(n);
    for (int i = 0; i < n; ++i) {
        int p;
        cin >> names[i] >> p;
        if (p == 0) {
            roots.push_back(i);
        } else {
            --p;
            adj[p].push_back(i);
        }
    }

    vector<int> depthCnt;
    function<void(int, int)> dfs1 = [&](int u, int d) {
        if ((int)depthCnt.size() == d) {
            depthCnt.push_back(0);
        }
        depthCnt[d]++;
        for (int v : adj[u]) {
            dfs1(v, d + 1);
        }
    };
    for (int r : roots) {
        dfs1(r, 0);
    }

    int D = (int)depthCnt.size();
    vector<SegTree> depthTrs;
    depthTrs.reserve(D);
    for (int c : depthCnt) {
        depthTrs.emplace_back(c);
    }
    vector<unordered_map<string, int>> namePos(D);
    for (int d = 0; d < D; ++d) {
        namePos[d].reserve(depthCnt[d] * 2 + 1);
    }

    int m;
    cin >> m;
    vector<pair<int, int>> queries(m);
    vector<vector<int>> qsAt(n);
    for (int i = 0; i < m; ++i) {
        int v, k;
        cin >> v >> k;
        --v;
        queries[i] = {v, k};
        qsAt[v].push_back(i);
    }

    vector<int> depthCurPos(D, 0);
    vector<int> ans(m, 0);

    function<void(int, int)> dfs2 = [&](int u, int d) {
        vector<int> keep;
        keep.reserve(qsAt[u].size());
        for (int qi : qsAt[u]) {
            int k = queries[qi].second;
            if (d + k < D) {
                keep.push_back(depthCurPos[d + k]);
            } else {
                keep.push_back(-1);
            }
        }

        auto it = namePos[d].find(names[u]);
        if (it != namePos[d].end()) {
            depthTrs[d].update(it->second, 0);
        }
        namePos[d][names[u]] = depthCurPos[d];
        depthTrs[d].update(depthCurPos[d], 1);
        depthCurPos[d]++;

        for (int v : adj[u]) {
            dfs2(v, d + 1);
        }

        for (size_t j = 0; j < qsAt[u].size(); ++j) {
            if (keep[j] < 0) continue;
            int qi = qsAt[u][j];
            int k = queries[qi].second;
            int dd = d + k;
            ans[qi] = depthTrs[dd].query(keep[j], depthCurPos[dd] - 1);
        }
    };

    for (int r : roots) {
        dfs2(r, 0);
    }

    for (int i = 0; i < m; ++i) {
        cout << ans[i] << '\n';
    }
    return 0;
}
