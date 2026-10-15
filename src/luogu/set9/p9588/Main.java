import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.IOException;
import java.util.StringTokenizer;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder out = new StringBuilder();
        for (int v : drive(reader)) {
            out.append(v).append('\n');
        }
        System.out.print(out);
    }

    static int[] drive(BufferedReader reader) throws IOException {
        StringTokenizer tok = new StringTokenizer(reader.readLine());
        tok.nextToken();
        int q = Integer.parseInt(tok.nextToken());
        int[][] ops = new int[q][2];
        for (int i = 0; i < q; i++) {
            tok = new StringTokenizer(reader.readLine());
            int op = Integer.parseInt(tok.nextToken());
            ops[i][0] = op;
            if (op != 4) {
                ops[i][1] = Integer.parseInt(tok.nextToken());
            }
        }
        return solve(ops);
    }

    static int[] solve(int[][] ops) {
        int m = ops.length;
        int[] nums = new int[m + 1];
        long[] pref = new long[m + 1];
        int blocks = 0;
        SegTree st = new SegTree(m + 1);
        long del = 0;
        int gone = 0;
        int[] ans = new int[m];
        int answers = 0;

        for (int[] cur : ops) {
            switch (cur[0]) {
                case 1 -> {
                    int x = cur[1];
                    blocks++;
                    nums[blocks] = x;
                    pref[blocks] = pref[blocks - 1] + x;
                    st.update(blocks, x);
                }
                case 2 -> {
                    del += cur[1];
                    while (gone + 1 <= blocks && pref[gone + 1] <= del) {
                        gone++;
                        st.update(gone, 0);
                    }
                }
                case 3 -> {
                    long pos = del + cur[1];
                    int k = lowerBound(pref, blocks + 1, pos);
                    ans[answers++] = (int) (pos - pref[k - 1]);
                }
                default -> ans[answers++] = st.get(1, blocks + 1);
            }
        }
        return java.util.Arrays.copyOf(ans, answers);
    }

    static int lowerBound(long[] pref, int len, long pos) {
        int lo = 0;
        int hi = len;
        while (lo < hi) {
            int mid = (lo + hi) >>> 1;
            if (pref[mid] >= pos) {
                hi = mid;
            } else {
                lo = mid + 1;
            }
        }
        return lo;
    }

    static final class SegTree {
        final int n;
        final int[] t;

        SegTree(int n) {
            this.n = n;
            this.t = new int[n << 1];
        }

        void update(int p, int v) {
            p += n;
            t[p] = v;
            for (p >>= 1; p > 0; p >>= 1) {
                t[p] = Math.max(t[p << 1], t[p << 1 | 1]);
            }
        }

        int get(int l, int r) {
            int res = 0;
            for (l += n, r += n; l < r; l >>= 1, r >>= 1) {
                if ((l & 1) == 1) {
                    res = Math.max(res, t[l++]);
                }
                if ((r & 1) == 1) {
                    res = Math.max(res, t[--r]);
                }
            }
            return res;
        }
    }
}
