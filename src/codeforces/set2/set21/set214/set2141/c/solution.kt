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

private fun solve(n: Int): List<String> {
    val res = ArrayList<String>()
    var d = 1
    for (l in 1..n) {
        if (d == 1) {
            var r = l
            while (r <= n) {
                res.add("pushback a[${r - 1}]")
                res.add("min")
                r++
            }
            d = -1
        } else {
            var r = n
            while (r > l) {
                res.add("min")
                res.add("popback")
                r--
            }
            res.add("min")
            d = 1
        }

        if (l < n) {
            res.add("popfront")
        }
    }
    return res
}

fun main() {
    val fs = FastScanner()
    val n = fs.nextInt()
    val ans = solve(n)
    val out = StringBuilder()
    out.append(ans.size).append('\n')
    for (cmd in ans) {
        out.append(cmd).append('\n')
    }
    print(out.toString())
}

