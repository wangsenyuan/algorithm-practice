#include <algorithm>
#include <iostream>
#include <numeric>
#include <vector>
using namespace std;

struct Rect {
  int x1, y1, x2, y2;
};

int main() {
  ios::sync_with_stdio(false);
  cin.tie(nullptr);

  int n;
  cin >> n;
  vector<Rect> rects(n);
  int mx = 0, my = 0;
  for (int i = 0; i < n; ++i) {
    cin >> rects[i].x1 >> rects[i].y1 >> rects[i].x2 >> rects[i].y2;
    mx = max(mx, rects[i].x2);
    my = max(my, rects[i].y2);
  }

  vector<vector<int>> area(mx + 1, vector<int>(my + 1));
  vector<vector<int>> end_at(mx + 1, vector<int>(my + 1, -1));
  for (int i = 0; i < n; ++i) {
    end_at[rects[i].x2][rects[i].y2] = i;
  }

  for (int x = 0; x <= mx; ++x) {
    for (int y = 0; y <= my; ++y) {
      if (x > 0) area[x][y] += area[x - 1][y];
      if (y > 0) area[x][y] += area[x][y - 1];
      if (x > 0 && y > 0) area[x][y] -= area[x - 1][y - 1];
      int id = end_at[x][y];
      if (id != -1) {
        area[x][y] += (x - rects[id].x1) * (y - rects[id].y1);
      }
    }
  }

  vector<vector<int>> up(mx + 1, vector<int>(my + 1));
  vector<vector<int>> rg(mx + 1, vector<int>(my + 1));
  for (int x = 0; x <= mx; ++x) {
    for (int y = 0; y <= my; ++y) {
      up[x][y] = y;
      rg[x][y] = x;
    }
  }

  auto get_up = [&](int x, int y) {
    int r = y;
    while (up[x][r] != r) r = up[x][r];
    while (y != r) {
      int ny = up[x][y];
      up[x][y] = r;
      y = ny;
    }
    return r;
  };

  auto get_rg = [&](int x, int y) {
    int r = x;
    while (rg[r][y] != r) r = rg[r][y];
    while (x != r) {
      int nx = rg[x][y];
      rg[x][y] = r;
      x = nx;
    }
    return r;
  };

  vector<int> ord(n);
  iota(ord.begin(), ord.end(), 0);
  sort(ord.begin(), ord.end(), [&](int i, int j) {
    return rects[i].x1 > rects[j].x1;
  });
  for (int id : ord) {
    rg[rects[id].x1][rects[id].y1] = get_rg(rects[id].x2, rects[id].y1);
  }

  sort(ord.begin(), ord.end(), [&](int i, int j) {
    return rects[i].y1 > rects[j].y1;
  });
  for (int id : ord) {
    up[rects[id].x1][rects[id].y1] = get_up(rects[id].x1, rects[id].y2);
  }

  auto collect = [&](int x1, int y1, int x2, int y2) {
    vector<int> res;
    for (int x = x1; x <= x2; ++x) {
      for (int y = y1; y <= y2; ++y) {
        if (end_at[x][y] != -1) res.push_back(end_at[x][y] + 1);
      }
    }
    sort(res.begin(), res.end());
    return res;
  };

  for (const Rect& r : rects) {
    int w = max(r.x2 - r.x1, r.y2 - r.y1);
    while (w <= r.x2 && w <= r.y2) {
      int x1 = r.x2 - w;
      int y1 = r.y2 - w;
      int covered = area[r.x2][r.y2] - area[r.x2][y1] - area[x1][r.y2] + area[x1][y1];
      if (covered == w * w && get_up(x1, y1) >= r.y2 && get_rg(x1, y1) >= r.x2) {
        vector<int> ans = collect(x1 + 1, y1 + 1, r.x2, r.y2);
        cout << "YES " << ans.size() << '\n';
        for (int i = 0; i < (int)ans.size(); ++i) {
          if (i) cout << ' ';
          cout << ans[i];
        }
        cout << '\n';
        return 0;
      }
      if (covered < w * w) break;
      ++w;
    }
  }

  cout << "NO\n";
  return 0;
}
