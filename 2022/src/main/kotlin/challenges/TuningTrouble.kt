package challenges

import Challenge
import java.io.File


class TuningTrouble : Challenge {
    override fun getName(): String {
        return "Day 6 - Tuning Trouble"
    }

    override fun getInputFileName(): String {
        return "day6.txt"
    }

    override fun solve(input: File?): String {
        val lines = input!!.readLines()

        return getMarker(lines.first(), 14).toString()
    }

    fun getMarker(line: String, distinctCount: Int = 4): Int {
        line.windowed(distinctCount, 1).forEachIndexed { index, s ->
            val size = s.toSet().size
            if(size == distinctCount) {
                return index + distinctCount
            }
        }

        return -1
    }
}