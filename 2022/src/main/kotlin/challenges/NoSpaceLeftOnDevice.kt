package challenges

import Challenge
import java.io.File as IOFile

class NoSpaceLeftOnDevice : Challenge {
    override fun getName(): String {
        return "Day 7 - No Space Left On Device"
    }

    override fun getInputFileName(): String {
        return "day7.txt"
    }

    override fun solve(input: IOFile?): String {
        val root = Node(null, "/", 0)
        var curDir = root

        input!!.forEachLine {
            val arguments = it.split(" ")

            when {
                it.startsWith("$") && arguments[1] == "cd" -> curDir = when (arguments[2]) {
                    "/" -> root
                    ".." -> curDir.parent
                    else -> {
                        val found = curDir.findByName(arguments[2])
                        found
                    }
                } ?: curDir

                it.startsWith("dir") -> {
                    curDir.addEdge(Node(curDir, arguments[1], 0))
                }

                !it.startsWith("$") -> curDir.addEdge(Node(curDir, arguments[1], arguments[0].toLong()))
            }
        }

//        return getSumOfAcceptableDirectories(root, 100_000L).toString()
        return getSmallestAcceptableDirectorySize(root, 70_000_000L, 30_000_000L).toString()
    }

    private fun getSumOfAcceptableDirectories(rootDir: Node, acceptableSize: Long): Long {
        val directoriesWithAcceptedSize = mutableListOf<Node>()

        fun getDirectoriesWithAcceptableSize(node: Node, acceptableSize: Long) {
            val weight = node.weightSum()
            if (weight <= acceptableSize) {
                directoriesWithAcceptedSize.add(node)
            }

            node.edges.forEach { getDirectoriesWithAcceptableSize(it, acceptableSize) }
        }

        getDirectoriesWithAcceptableSize(rootDir, acceptableSize)

        return directoriesWithAcceptedSize.sumOf { it.weightSum() }
    }

    private fun getSmallestAcceptableDirectorySize(rootDir: Node, totalDiskSize: Long, requiredFreeSpace: Long): Long {
        val directories = mutableListOf<Node>()

        fun collectDirectoriesByMinimalSize(node: Node, minimalSize: Long) {
            val weight = node.weightSum()
            if (weight >= minimalSize) {
                directories.add(node)
            }

            node.edges.forEach { collectDirectoriesByMinimalSize(it, minimalSize) }
        }

        val unusedSpace = totalDiskSize - rootDir.weightSum()
        val spaceToFreeUp = requiredFreeSpace - unusedSpace

        collectDirectoriesByMinimalSize(rootDir, spaceToFreeUp)

        return directories.filter { it.isDir() }.map { it.weightSum() }.minOf { it }
    }

    class Node(val parent: Node?, private val name: String, private val weight: Long, val edges: MutableList<Node> = mutableListOf()) {
        fun findByName(name: String): Node? {
            if (this.name == name) return this
            return this.edges.findLast { it.name == name }
        }

        fun isDir(): Boolean = weight == 0L

        fun addEdge(node: Node) {
            edges.add(node)
        }

        fun weightSum(): Long {
            return edges.sumOf { it.weight + it.weightSum() }
        }

        fun println(depth: Int) {
            println("\t".repeat(depth) + " $name (${if( weight > 0) "file, size=$weight" else "dir"})")
            this.edges.forEach { it.println(depth + 1) }
        }
    }
}

