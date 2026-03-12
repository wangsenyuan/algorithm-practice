#include <algorithm>
#include <cstdio>
#include <string>
#include <vector>

using namespace std;

class FastScanner {
private:
    static constexpr size_t BUFSIZE = 1 << 20;
    char buf[BUFSIZE];
    size_t idx = 0;
    size_t size = 0;

    inline char nextChar() {
        if (idx >= size) {
            size = fread(buf, 1, BUFSIZE, stdin);
            idx = 0;
            if (size == 0) {
                return 0;
            }
        }
        return buf[idx++];
    }

public:
    int nextInt() {
        char c = nextChar();
        while (c <= ' ' && c != 0) {
            c = nextChar();
        }
        int sign = 1;
        if (c == '-') {
            sign = -1;
            c = nextChar();
        }
        int x = 0;
        while (c > ' ') {
            x = x * 10 + (c - '0');
            c = nextChar();
        }
        return x * sign;
    }
};

struct Node {
    int left;
    int right;
    int sum;
};

class PersistentSegmentTree {
private:
    vector<Node> nodes;

public:
    explicit PersistentSegmentTree(size_t capacity) {
        nodes.reserve(capacity);
        nodes.push_back({0, 0, 0});
    }

    int cloneNode(int from) {
        nodes.push_back(nodes[from]);
        return static_cast<int>(nodes.size()) - 1;
    }

    int update(int node, int l, int r, int pos) {
        int cur = cloneNode(node);
        if (l == r) {
            nodes[cur].sum = 1;
            return cur;
        }
        int mid = (l + r) >> 1;
        if (pos <= mid) {
            nodes[cur].left = update(nodes[cur].left, l, mid, pos);
        } else {
            nodes[cur].right = update(nodes[cur].right, mid + 1, r, pos);
        }
        nodes[cur].sum = nodes[nodes[cur].left].sum + nodes[nodes[cur].right].sum;
        return cur;
    }

    int query(int node, int l, int r, int ql, int qr) const {
        if (node == 0) {
            return 0;
        }
        if (ql == l && qr == r) {
            return nodes[node].sum;
        }
        int mid = (l + r) >> 1;
        int res = 0;
        if (ql <= mid) {
            res += query(nodes[node].left, l, mid, ql, min(qr, mid));
        }
        if (mid < qr) {
            res += query(nodes[node].right, mid + 1, r, max(ql, mid + 1), qr);
        }
        return res;
    }
};

static inline int decode(int z, int x, int n) {
    return (z - 1 + x) % n + 1;
}

int main() {
    FastScanner fs;

    int n = fs.nextInt();
    vector<int> p(n + 1);
    for (int i = 1; i <= n; ++i) {
        p[i] = fs.nextInt();
    }

    vector<int> pos(n + 1);
    for (int i = 1; i <= n; ++i) {
        int v = fs.nextInt();
        pos[v] = i;
    }

    int depth = 0;
    while ((1 << depth) < n) {
        ++depth;
    }
    size_t capacity = 1ULL * n * (depth + 1) + 5;
    PersistentSegmentTree pst(capacity);

    vector<int> roots(n + 1, 0);
    for (int i = 1; i <= n; ++i) {
        roots[i] = pst.update(roots[i - 1], 1, n, pos[p[i]]);
    }

    int m = fs.nextInt();
    int x = 0;
    string out;
    out.reserve(static_cast<size_t>(m) * 4);

    for (int i = 0; i < m; ++i) {
        int a = fs.nextInt();
        int b = fs.nextInt();
        int c = fs.nextInt();
        int d = fs.nextInt();

        int l1 = decode(a, x, n);
        int r1 = decode(b, x, n);
        if (l1 > r1) {
            swap(l1, r1);
        }

        int l2 = decode(c, x, n);
        int r2 = decode(d, x, n);
        if (l2 > r2) {
            swap(l2, r2);
        }

        int ans = pst.query(roots[r1], 1, n, l2, r2) - pst.query(roots[l1 - 1], 1, n, l2, r2);
        x = ans + 1;

        out += to_string(ans);
        out.push_back('\n');
    }

    fwrite(out.data(), 1, out.size(), stdout);
    return 0;
}
