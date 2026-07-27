import java.io.BufferedInputStream
import java.lang.StringBuilder

private const val MOD = 1_000_000_007L

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
        var value = 0L
        while (c > 32) {
            value = value * 10 + (c - '0'.code)
            c = readByte()
        }
        return value
    }

    fun next(): String {
        var c = readByte()
        while (c <= 32) c = readByte()
        val result = StringBuilder()
        while (c > 32) {
            result.append(c.toChar())
            c = readByte()
        }
        return result.toString()
    }
}

private fun gcd(a: Long, b: Long): Long {
    var x = a
    var y = b
    while (y != 0L) {
        val next = x % y
        x = y
        y = next
    }
    return x
}

private fun solve(n: Int, c: Long, s: String): Long {
    if (s.first() != '1' || s.last() != '1') {
        return -1L
    }

    var answer = 1L
    var remainingC = c

    for (k in 1 until n) {
        val factor = if (s[k - 1] == '1') 2L else (k - 1).toLong()
        answer = answer * factor % MOD
        remainingC /= gcd(remainingC, factor)
    }

    return if (remainingC == 1L) -1L else answer
}

fun main() {
    val fs = FastScanner()
    val t = fs.nextLong().toInt()
    val out = StringBuilder()

    repeat(t) {
        val n = fs.nextLong().toInt()
        val c = fs.nextLong()
        val s = fs.next()
        out.append(solve(n, c, s)).append('\n')
    }

    print(out.toString())
}
