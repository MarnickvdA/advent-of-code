package challenges

import Challenge
import utils.Stack
import java.io.File

class SupplyStacks : Challenge {
    override fun getName(): String {
        return "Day 5 - Supply Stacks"
    }

    override fun getInputFileName(): String {
        return "day5.txt"
    }

    override fun solve(input: File?): String {
        val hashMap = mutableMapOf<Int, Stack<Char>>()

        // Initialize empty stacks
        (1..9).forEach { hashMap[it] = Stack() }

        input!!.forEachLine {
            if (it.startsWith("[")) {
                handleStackLine(hashMap, it)
            }

            // Print the start position of all stacks.
            if (it.startsWith(" 1")) {
                printStacks(hashMap)
            }

            if (it.startsWith("move")) {
                handleMoveLine(hashMap, it)
            }
        }

        return hashMap.map { t -> t.value.top() }.joinToString("")
    }

    private fun printStacks(hashMap: MutableMap<Int, Stack<Char>>) {
        println(hashMap.forEach { h -> println("${h.key}: ${h.value.list()}") })
    }

    fun handleStackLine(hashMap: MutableMap<Int, Stack<Char>>, line: String) {
        val entries = line.chunked(4).map { s -> s[1] }

        // Insert all entries in 'reverse' stack order
        entries.forEachIndexed { i, c -> if (c != ' ') hashMap[i + 1]?.insert(c) }
    }

    fun handleMoveLine(hashMap: MutableMap<Int, Stack<Char>>, line: String) {
        val l = line.split(" ")
        val amount = l[1].toInt()
        val from = l[3].toInt()
        val to = l[5].toInt()

        hashMap[to]!!.putInOrder(hashMap[from]!!.popInOrder(amount))
    }
}