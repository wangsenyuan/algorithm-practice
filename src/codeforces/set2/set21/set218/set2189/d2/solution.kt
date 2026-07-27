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

private fun countPowerOfTwo(value: Long): Int {
    var x = value
    var count = 0
    while (x % 2L == 0L) {
        count++
        x /= 2L
    }
    return count
}

private fun minimumFactor(k: Int, ch: Char): Long {
    if (k == 1 || ch == '1') return 2L
    if (ch == '0') return (k - 1).toLong()
    return minOf(2, k - 1).toLong()
}

private fun solve(n: Int, c: Long, s: String): Long {
    if (s.first() == '0' || s.last() == '0') {
        return -1L
    }

    var answer = 1L
    var remainingC = c
    var productPowerOfTwo = 0

    // Build the smallest possible product.
    // Position n is only forced to 1; it contributes no factor.
    for (k in 1 until n) {
        val factor = minimumFactor(k, s[k - 1])
        answer = answer * factor % MOD
        remainingC /= gcd(remainingC, factor)
        productPowerOfTwo += countPowerOfTwo(factor)
    }

    if (remainingC != 1L) {
        return answer
    }

    // The minimum product is divisible by c. Each useful change at an even
    // position k replaces factor 2 with odd factor k-1, removing one power of 2.
    var changesNeeded = productPowerOfTwo - countPowerOfTwo(c) + 1
    answer = 1L

    for (k in 1 until n) {
        val factor =
            if (s[k - 1] == '?' && k >= 4 && k % 2 == 0 && changesNeeded > 0) {
                changesNeeded--
                (k - 1).toLong()
            } else {
                minimumFactor(k, s[k - 1])
            }
        answer = answer * factor % MOD
    }

    return if (changesNeeded == 0) answer else -1L
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
