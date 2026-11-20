import java.io.BufferedInputStream
import java.lang.StringBuilder

private class FastScanner {
    private val input = BufferedInputStream(System.`in`)
    private val buffer = ByteArray(1 shl 16)
    private var len = 0
    private var ptr = 0

    private fun readByte(): Int {
        if (ptr >= len) {
            len = input.read(buffer)
            ptr = 0
            if (len <= 0) return -1
        }
        return buffer[ptr++].toInt()
    }

    fun nextInt(): Int {
        var c = readByte()
        while (c <= 32) c = readByte()
        var sign = 1
        if (c == '-'.code) {
            sign = -1
            c = readByte()
        }
        var res = 0
        while (c > 32) {
            res = res * 10 + (c - '0'.code)
            c = readByte()
        }
        return res * sign
    }
}

private const val INF = 1_000_000_000

private fun solve(n: Int, k: Int, swaps: Array<IntArray>): IntArray {
    val dp = IntArray(n) { INF }
    dp[k - 1] = 0
    for (cur in swaps) {
        val a = cur[0] - 1
        val b = cur[1] - 1
        val w = dp[a]
        val v = dp[b]
        dp[a] = minOf(w + 1, v)
        dp[b] = minOf(w, v + 1)
    }
    for (i in dp.indices) {
        if (dp[i] == INF) dp[i] = -1
    }
    return dp
}

fun main() {
    val fs = FastScanner()
    val n = fs.nextInt()
    val m = fs.nextInt()
    val k = fs.nextInt()
    val swaps = Array(m) { intArrayOf(fs.nextInt(), fs.nextInt()) }
    val ans = solve(n, k, swaps)
    val out = StringBuilder()
    for (i in ans.indices) {
        if (i > 0) out.append(' ')
        out.append(ans[i])
    }
    out.append('\n')
    print(out.toString())
}
