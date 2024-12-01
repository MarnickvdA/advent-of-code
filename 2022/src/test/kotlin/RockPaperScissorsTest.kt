import challenges.RockPaperScissors
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class RockPaperScissorsTest {

    @Test
    fun testAllScenarios() {
        val game = RockPaperScissors();

        assertEquals(game.getScoreFromLine("A X"), 4) // 1 + 3
        assertEquals(game.getScoreFromLine("A Y"), 8) // 2 + 6
        assertEquals(game.getScoreFromLine("A Z"), 3) // 3 + 0
        assertEquals(game.getScoreFromLine("B X"), 1) // 1 + 0
        assertEquals(game.getScoreFromLine("B Y"), 5) // 2 + 3
        assertEquals(game.getScoreFromLine("B Z"), 9) // 3 + 6
        assertEquals(game.getScoreFromLine("C X"), 7) // 1 + 6
        assertEquals(game.getScoreFromLine("C Y"), 2) // 2 + 0
        assertEquals(game.getScoreFromLine("C Z"), 6) // 3 + 3
    }
}