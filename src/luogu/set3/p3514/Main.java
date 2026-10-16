import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        StringTokenizer tok = new StringTokenizer(reader.readLine());
        int n = Integer.parseInt(tok.nextToken());
        int m = Integer.parseInt(tok.nextToken());
        String s = reader.readLine().trim();

        int[] left = new int[2 * n + 1];
        int[] right = new int[2 * n + 1];
        int[] best = {-1, -1};
        int sum = 0;
        for (int i = 0; i < n; i++) {
            sum += s.charAt(i) == 'W' ? 1 : 2;
            left[sum] = 0;
            right[sum] = i;
            best[sum % 2] = Math.max(best[sum % 2], sum);
        }
        sum = 0;
        for (int i = n - 1; i >= 0; i--) {
            sum += s.charAt(i) == 'W' ? 1 : 2;
            left[sum] = i;
            right[sum] = n - 1;
            best[sum % 2] = Math.max(best[sum % 2], sum);
        }

        int[] ansL = new int[2 * n + 1];
        int[] ansR = new int[2 * n + 1];
        if (best[0] >= 0) {
            fill(s, best[0], left[best[0]], right[best[0]], ansL, ansR);
        }
        if (best[1] >= 0) {
            fill(s, best[1], left[best[1]], right[best[1]], ansL, ansR);
        }

        StringBuilder out = new StringBuilder();
        for (int i = 0; i < m; i++) {
            int x = Integer.parseInt(reader.readLine().trim());
            if (x > 2 * n || ansL[x] == 0) {
                out.append("NIE\n");
            } else {
                out.append(ansL[x]).append(' ').append(ansR[x]).append('\n');
            }
        }
        System.out.print(out);
    }

    static void fill(String s, int sum, int l, int r, int[] ansL, int[] ansR) {
        while (l <= r) {
            ansL[sum] = l + 1;
            ansR[sum] = r + 1;
            if (sum <= 1) {
                return;
            }
            if (s.charAt(l) == 'T') {
                l++;
            } else if (s.charAt(r) == 'T') {
                r--;
            } else {
                l++;
                r--;
            }
            sum -= 2;
        }
    }
}
