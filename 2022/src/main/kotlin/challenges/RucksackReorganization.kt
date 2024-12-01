package challenges

import Challenge
import java.io.File

class RucksackReorganization : Challenge {
    override fun getName(): String {
        return "Day 3 - Rucksack Reorganization"
    }

    override fun getInputFileName(): String {
        return "day3.txt"
    }

    private var priorityRange = mutableListOf<Char>()

    override fun solve(input: File?): String {
        /**
         * To help prioritize item rearrangement, every item type can be converted to a priority:
         *
         * Lowercase item types a through z have priorities 1 through 26.
         * Uppercase item types A through Z have priorities 27 through 52.
         * In the above example, the priority of the item type that appears in both compartments
         * of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.
         *
         * Find the item type that appears in both compartments of each rucksack.
         * What is the sum of the priorities of those item types?
         */

        priorityRange.addAll(('a'..'z').toList())
        priorityRange.addAll(('A'..'Z').toList())

        var sum = 0;
        var block = mutableListOf<Set<Char>>()

        input!!.forEachLine {
            block.add(it.toCharArray().toSet())

            if (block.size == 3) {
                sum += getPriority(block[0].intersect(block[1]).intersect(block[2]).toList()[0])
                block = mutableListOf()
            }
        }

        return sum.toString()
    }

    fun getPriority(char: Char): Int {
        return priorityRange.indexOf(char) + 1
    }

}