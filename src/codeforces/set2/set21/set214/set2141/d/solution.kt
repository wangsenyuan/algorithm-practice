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

    fun nextLong(): Long {
        var c = readByte()
        while (c <= 32) c = readByte()
        var sign = 1L
        if (c == '-'.code) {
            sign = -1L
            c = readByte()
        }
        var res = 0L
        while (c > 32) {
            res = res * 10 + (c - '0'.code)
            c = readByte()
        }
        return res * sign
    }

    fun nextInt(): Int = nextLong().toInt()
}

private fun solve(k: Long, a: LongArray): Long {
    val n = a.size
    a.sort()

    var sum = 0L
    for (v in a) sum += v

    val avg = (sum + k) / n
    if (avg < a[n - 1]) return -1L

    var score = 0L
    for (i in 1 until n) {
        var diff = avg - a[i]
        if (a[i] == a[0]) {
            diff--
        }
        score += max(0L, diff)
    }
    return score
}

fun main() {
    val fs = FastScanner()
    val tc = fs.nextInt()
    val out = StringBuilder()
    repeat(tc) {
        val n = fs.nextInt()
        val k = fs.nextLong()
        val a = LongArray(n) { fs.nextLong() }
        out.append(solve(k, a)).append('\n')
    }
    print(out.toString())
}

