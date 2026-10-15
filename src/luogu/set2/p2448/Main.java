import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;
import java.util.HashMap;
import java.util.StringTokenizer;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        System.out.println(drive(reader));
    }

    static long drive(BufferedReader reader) throws IOException {
        StringTokenizer tok = new StringTokenizer(reader.readLine());
        int k = Integer.parseInt(tok.nextToken());
        int[][] swaps = new int[k][2];
        for (int i = 0; i < k; i++) {
            tok = new StringTokenizer(reader.readLine());
            swaps[i][0] = Integer.parseInt(tok.nextToken());
            swaps[i][1] = Integer.parseInt(tok.nextToken());
        }
        return solve(swaps);
    }

    static long solve(int[][] swaps) {
        int[] seen = new int[swaps.length * 2];
        int p = 0;
        for (int[] swap : swaps) {
            seen[p++] = swap[0];
            seen[p++] = swap[1];
        }
        Arrays.sort(seen);
        int n = 0;
        for (int v : seen) {
            if (n == 0 || seen[n - 1] != v) {
                seen[n++] = v;
            }
        }

        HashMap<Integer, Integer> pos = new HashMap<>(n * 2);
        int[] arr = new int[n];
        for (int i = 0; i < n; i++) {
            pos.put(seen[i], i);
            arr[i] = i;
        }
        for (int[] swap : swaps) {
            int i = pos.get(swap[0]);
            int j = pos.get(swap[1]);
            int tmp = arr[i];
            arr[i] = arr[j];
            arr[j] = tmp;
        }

        Fenwick tr = new Fenwick(n + 2);
        long ans = 0;
        long placed = 0;
        for (int rank : arr) {
            ans += placed - tr.pre(rank);
            placed++;
            tr.update(rank, 1);
        }
        for (int i = 0; i < n; i++) {
            int rank = arr[i];
            if (rank == i) {
                continue;
            }
            int lo = i;
            int hi = rank;
            if (rank < i) {
                lo = rank;
                hi = i;
            }
            ans += (long) seen[hi] - seen[lo] - 1 - (hi - lo - 1);
        }
        return ans;
    }

    static class Fenwick {
        final long[] f;

        Fenwick(int n) {
            f = new long[n];
        }

        void update(int i, long v) {
            for (i++; i < f.length; i += i & -i) {
                f[i] += v;
            }
        }

        long pre(int i) {
            long sum = 0;
            for (i++; i > 0; i -= i & -i) {
                sum += f[i];
            }
            return sum;
        }
    }
}
