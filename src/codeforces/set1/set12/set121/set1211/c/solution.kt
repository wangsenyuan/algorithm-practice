import java.io.BufferedInputStream

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
}

private fun solve(k: Long, days: Array<LongArray>): Long {
    var lo = 0L
    var hi = 0L
    for (cur in days) {
        lo += cur[0]
        hi += cur[1]
    }
    if (lo > k || hi < k) return -1

    var cost = 0L
    for (cur in days) cost += cur[0] * cur[2]
    var need = k - lo
    days.sortBy { it[2] }
    for (cur in days) {
        if (need == 0L) break
        val take = minOf(need, cur[1] - cur[0])
        need -= take
        cost += take * cur[2]
    }
    return cost
}

fun main() {
    val fs = FastScanner()
    val n = fs.nextLong().toInt()
    val k = fs.nextLong()
    val days = Array(n) { longArrayOf(fs.nextLong(), fs.nextLong(), fs.nextLong()) }
    println(solve(k, days))
}
