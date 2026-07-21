import java.io.BufferedInputStream
import java.lang.StringBuilder
import kotlin.math.max

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

// Process each contiguous run of broken steps independently.
// Min days for a run of length L is ceil(L/2). Pairing adjacent steps costs
// 2*(a+b); a single costs a. Optimal effort for a run is 2*sum(a), minus the
// max element on the cheaper parity when L is odd (one unpaired single).
private fun solve(a: IntArray): Long {
    var res = 0L
    val n = a.size
    var i = 0
    while (i < n) {
        if (a[i] == 0) {
            i++
            continue
        }
        val j = i
        var sum = 0L
        var mx = 0
        while (i < n && a[i] > 0) {
            sum += a[i]
            if ((i and 1) == (j and 1)) {
                mx = max(mx, a[i])
            }
            i++
        }
        res += 2L * sum
        if (((i - j) and 1) == 1) {
            res -= mx.toLong()
        }
    }
    return res
}

fun main() {
    val fs = FastScanner()
    val t = fs.nextInt()
    val out = StringBuilder()
    repeat(t) {
        val n = fs.nextInt()
        val a = IntArray(n) { fs.nextInt() }
        out.append(solve(a)).append('\n')
    }
    print(out.toString())
}
