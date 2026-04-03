#include <algorithm>
#include <cstdint>
#include <functional>
#include <iostream>
#include <unordered_set>
#include <vector>

using namespace std;

static const uint64_t MOD[2] = {1000000007ULL, 1000000009ULL};

struct Hash {
  uint64_t h[2]{};

  static Hash newHash(uint64_t x) {
    Hash r;
    r.h[0] = x % MOD[0];
    r.h[1] = x % MOD[1];
    return r;
  }

  Hash add(const Hash &that) const {
    Hash r;
    for (int i = 0; i < 2; i++) {
      r.h[i] = (h[i] + that.h[i]) % MOD[i];
    }
    return r;
  }

  Hash sub(const Hash &that) const {
    Hash r;
    for (int i = 0; i < 2; i++) {
      r.h[i] = (h[i] + MOD[i] - that.h[i]) % MOD[i];
    }
    return r;
  }

  Hash addInt(int x) const {
    Hash r;
    uint64_t ux = static_cast<uint64_t>(static_cast<unsigned int>(x));
    for (int i = 0; i < 2; i++) {
      r.h[i] = (h[i] + ux % MOD[i]) % MOD[i];
    }
    return r;
  }

  Hash mul(const Hash &that) const {
    Hash r;
    for (int i = 0; i < 2; i++) {
      r.h[i] = (h[i] * that.h[i]) % MOD[i];
    }
    return r;
  }

  Hash mulInt(int x) const {
    Hash r;
    uint64_t ux = (uint64_t)x;
    for (int i = 0; i < 2; i++) {
      r.h[i] = (h[i] * (ux % MOD[i])) % MOD[i];
    }
    return r;
  }
};

struct HashEq {
  bool operator()(const Hash &a, const Hash &b) const {
    return a.h[0] == b.h[0] && a.h[1] == b.h[1];
  }
};

struct HashH {
  size_t operator()(const Hash &x) const {
    return x.h[0] ^ (x.h[1] << 1);
  }
};

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(nullptr);

  int n;
  if (!(cin >> n)) {
    return 0;
  }
  vector<int> a(max(0, n - 1));
  for (int i = 0; i < n - 1; i++) {
    cin >> a[i];
  }
  vector<vector<int>> adj(n);
  for (int i = 0; i < n - 1; i++) {
    int u, v;
    cin >> u >> v;
    --u;
    --v;
    adj[u].push_back(v);
    adj[v].push_back(u);
  }

  vector<int> freq(n, 0);
  for (int v : a) {
    freq[v]++;
  }

  if (freq[0] > 1) {
    cout << 0 << "\n";
    return 0;
  }

  vector<Hash> bases(n);
  bases[0] = Hash::newHash(1);
  for (int i = 1; i < n; i++) {
    bases[i] = bases[i - 1].mulInt(n);
  }

  Hash target = Hash::newHash(0);
  for (int i = 0; i < n; i++) {
    if (freq[i] > 0) {
      target = target.add(bases[i].mulInt(freq[i]));
    }
  }

  unordered_set<Hash, HashH, HashEq> valid;
  valid.reserve((size_t)n * 2);
  for (int i = n - 1; i >= 0; i--) {
    valid.insert(target.add(bases[i]));
  }

  vector<Hash> dp(n);

  function<void(int, int)> dfs = [&](int p, int u) {
    for (int v : adj[u]) {
      if (v == p) {
        continue;
      }
      dfs(u, v);
      dp[u] = dp[u].add(dp[v].mulInt(n));
    }
    dp[u] = dp[u].addInt(1);
  };

  dfs(-1, 0);

  vector<int> res;

  function<void(int, int, Hash)> dfs2 = [&](int p, int u, Hash w) {
    Hash cur = dp[u];
    if (p >= 0) {
      cur = cur.add(w.mulInt(n));
    }
    if (valid.find(cur) != valid.end()) {
      res.push_back(u + 1);
    }
    for (int v : adj[u]) {
      if (v == p) {
        continue;
      }
      Hash tmp = cur.sub(dp[v].mulInt(n));
      dfs2(u, v, tmp);
    }
  };

  dfs2(-1, 0, Hash::newHash(0));

  sort(res.begin(), res.end());

  cout << res.size() << "\n";
  if (!res.empty()) {
    for (size_t i = 0; i < res.size(); i++) {
      if (i) {
        cout << ' ';
      }
      cout << res[i];
    }
    cout << "\n";
  }

  return 0;
}
