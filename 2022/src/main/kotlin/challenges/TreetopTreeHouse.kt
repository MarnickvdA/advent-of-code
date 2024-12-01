package challenges

import Challenge
import java.io.File
import kotlin.math.max

class TreetopTreeHouse : Challenge {
    override fun getName(): String {
        return "Day 8 - Treetop Tree House"
    }

    override fun getInputFileName(): String {
        return "day8.txt"
    }

    override fun solve(input: File?): String {
        val lines = input!!.readLines()
        val h = lines.size
        val w = lines.first().length

        /**
         * M by N Matrix of heights of trees. 0 is shortest, 9 is tallest.
         */
        val matrix = Array(lines.size) { i ->
            IntArray(lines.first().length) { j -> lines[i][j].toString().toInt() }
        }

        var maxScenicScore = 0

        // Don't iterate over the borders, since they are visible.
        for (y in 1 until h - 1) {
            for (x in 1 until w - 1) {
                val height = matrix[y][x]

                var left = 0
                var right = 0
                var up = 0
                var down = 0

                for (i in 1..x) {
                    left += 1

                    if (height <= matrix[y][x - i]) break
                }

                for (i in 1 until (w - x)) {
                    right += 1

                    if (height <= matrix[y][x + i]) break
                }

                for (i in 1..y) {
                    up += 1

                    if (height <= matrix[y - i][x]) break
                }

                for (i in 1 until (h - y)) {
                    down += 1

                    if (height <= matrix[y + i][x]) break
                }

                val scenicScore = left * right * up * down
                maxScenicScore = maxOf(maxScenicScore, scenicScore)
            }
        }

        return maxScenicScore.toString()
    }
}