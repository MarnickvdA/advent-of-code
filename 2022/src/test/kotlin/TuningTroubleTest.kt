import challenges.TuningTrouble
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions.assertEquals

class TuningTroubleTest {
    @Test
    fun testAllScenarios() {
        val game = TuningTrouble();

        assertEquals(5, game.getMarker("bvwbjplbgvbhsrlpgdmjqwftvncz"))
        assertEquals(19, game.getMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))

    }
}