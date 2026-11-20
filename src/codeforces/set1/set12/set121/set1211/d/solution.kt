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

private fun solve(a: Int, b: Int, k: Int, r: IntArray): Int {
    r.sort()
    val todo = HashMap<Int, MutableList<Int>>()
    var i = 0
    while (i < r.size) {
        val j = i
        while (i < r.size && r[i] == r[j]) i++
        val cnt = i - j
        val prev = todo[r[j] / k]
        if (prev == null || r[j] % k != 0) {
            todo[r[j]] = mutableListOf(cnt)
        } else {
            prev.add(cnt)
            todo[r[j]] = prev
            todo.remove(r[j] / k)
        }
    }

    var ans = 0
    for (arr in todo.values) ans += play(arr, a, b)
    return ans
}

private fun play(arr: List<Int>, a: Int, b: Int): Int {
    val w = arr.max()
    val dp = IntArray(w + 1)
    val ndp = IntArray(w + 1)
    for (i in 1 until arr.size) {
        val v = arr[i]
        var x = 0
        while (x <= v) {
            val x1 = x / b * a
            if (x1 > arr[i - 1]) break
            ndp[v - x] = maxOf(ndp[v - x], dp[x1] + x / b)
            x += b
        }
        for (t in w downTo 0) {
            if (t + 1 <= w) ndp[t] = maxOf(ndp[t], ndp[t + 1])
            dp[t] = ndp[t]
        }
        ndp.fill(0)
    }
    return dp[0]
}

fun main() {
    val fs = FastScanner()
    val n = fs.nextInt()
    val a = fs.nextInt()
    val b = fs.nextInt()
    val k = fs.nextInt()
    val r = IntArray(n) { fs.nextInt() }
    println(solve(a, b, k, r))
}
