import challenges.SupplyStacks
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import utils.Stack

class SupplyStacksTest {
    private val hashMap = mutableMapOf<Int, Stack<Char>>()
    private val game = SupplyStacks()

    // Initialize empty stacks
    @BeforeEach
    fun init() {
        (1..9).forEach { hashMap[it] = Stack() }

        val lines = arrayOf(
            "[C]         [S] [H]",
            "[F] [B]     [C] [S]     [W]",
            "[B] [W]     [W] [M] [S] [B]",
            "[L] [H] [G] [L] [P] [F] [Q]",
            "[D] [P] [J] [F] [T] [G] [M] [T]",
            "[P] [G] [B] [N] [L] [W] [P] [W] [R]",
            "[Z] [V] [W] [J] [J] [C] [T] [S] [C]",
            "[S] [N] [F] [G] [W] [B] [H] [F] [N]",
        )

        lines.forEach { game.handleStackLine(hashMap, it) }
    }

    @Test
    fun testInsertion() {
        assertEquals(hashMap[1]!!.top(), 'C')
        assertEquals(hashMap[2]!!.top(), 'B')
        assertEquals(hashMap[3]!!.top(), 'G')
        assertEquals(hashMap[4]!!.top(), 'S')
        assertEquals(hashMap[5]!!.top(), 'H')
        assertEquals(hashMap[6]!!.top(), 'S')
        assertEquals(hashMap[7]!!.top(), 'W')
        assertEquals(hashMap[8]!!.top(), 'T')
        assertEquals(hashMap[9]!!.top(), 'R')
    }

    @Test
    fun testMove() {
        game.handleMoveLine(hashMap, "move 3 from 1 to 3")
        assertEquals(hashMap[3]!!.top(), 'B')
    }
}