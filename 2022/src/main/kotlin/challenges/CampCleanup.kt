package challenges

import Challenge
import java.io.File

class CampCleanup: Challenge {
    override fun getName(): String {
        return "Day 4 - Camp Cleanup"
    }

    override fun getInputFileName(): String {
        return "day4.txt"
    }

    override fun solve(input: File?): String {
        var count = 0
        input!!.forEachLine {
            val ranges = it.split(",").map { l -> l.split("-").map { s -> s.toInt() } }
            val range1: Set<Int> = (ranges[0][0]..ranges[0][1]).toList().toSet()
            val range2: Set<Int> = (ranges[1][0]..ranges[1][1]).toList().toSet()
            val intSize = range1.intersect(range2).size

            if (intSize > 0) {
                count++
            }
        }

        return count.toString()
    }

}