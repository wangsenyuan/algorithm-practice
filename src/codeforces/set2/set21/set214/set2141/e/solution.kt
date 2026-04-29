import java.io.BufferedInputStream
import java.lang.StringBuilder

private const val MOD = 998244353

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

    fun nextToken(): String {
        var c = readByte()
        while (c <= 32) c = readByte()
        val sb = StringBuilder()
        while (c > 32) {
            sb.append(c.toChar())
            c = readByte()
        }
        return sb.toString()
    }
}

private fun add(a: Int, b: Int): Int {
    var x = a + b
    if (x >= MOD) x -= MOD
    return x
}

private fun sub(a: Int, b: Int): Int = add(a, MOD - b)

private fun mul(a: Int, b: Int): Int = ((a.toLong() * b) % MOD).toInt()

private fun check(s: String): Boolean {
    if (s[0] == '1') return false
    for (i in 1 until s.length) {
        if (s[i] == '0') return false
    }
    return true
}

private fun solve(s: String): Int {
    var res = 1
    for (c in s) {
        if (c == '?') {
            res = mul(res, 2)
        }
    }
    if (check(s)) {
        res = sub(res, 1)
    }
    return res
}

fun main() {
    val fs = FastScanner()
    val tc = fs.nextInt()
    val out = StringBuilder()
    repeat(tc) {
        val s = fs.nextToken()
        out.append(solve(s)).append('\n')
    }
    print(out.toString())
}
