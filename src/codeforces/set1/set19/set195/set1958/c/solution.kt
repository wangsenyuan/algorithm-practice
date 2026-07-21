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
}

// Need pieces summing to k from splits of 2^n. Each level down to the lowest
// set bit of k costs one cut, so the answer is n - trailingZeros(k).
private fun solve(n: Int, k: Long): Int = n - k.countTrailingZeroBits()

fun main() {
    val fs = FastScanner()
    val t = fs.nextInt()
    val out = StringBuilder()
    repeat(t) {
        val n = fs.nextInt()
        val k = fs.nextLong()
        out.append(solve(n, k)).append('\n')
    }
    print(out.toString())
}
