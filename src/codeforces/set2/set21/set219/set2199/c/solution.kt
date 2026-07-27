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

// Empty list => NO; otherwise two rows of equal minimum width.
private fun solve(k: Int): List<String> {
    if (k % 5 == 2 || k % 5 == 4) return emptyList()

    val row0 = StringBuilder()
    val row1 = StringBuilder()
    val m = k / 5

    if (k % 5 > 0) {
        // remainder 1 or 3: repeat (... / *..) then (. / *), optional (. / .)
        repeat(m) {
            row0.append("...")
            row1.append("*..")
        }
        row0.append('.')
        row1.append('*')
        if (k % 5 == 3) {
            row0.append('.')
            row1.append('.')
        }
    } else {
        // remainder 0: repeat (... / .*.)
        repeat(m) {
            row0.append("...")
            row1.append(".*.")
        }
    }
    return listOf(row0.toString(), row1.toString())
}

fun main() {
    val fs = FastScanner()
    val t = fs.nextInt()
    val out = StringBuilder()
    repeat(t) {
        val k = fs.nextInt()
        val ans = solve(k)
        if (ans.isEmpty()) {
            out.append("NO\n")
        } else {
            out.append("YES\n")
            out.append(ans[0].length).append('\n')
            ans.forEach { out.append(it).append('\n') }
        }
    }
    print(out.toString())
}
