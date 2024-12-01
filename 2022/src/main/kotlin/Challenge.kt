import java.io.File

interface Challenge {
    fun getName(): String

    fun getInputFileName(): String
    fun solve(input: File?): String
}