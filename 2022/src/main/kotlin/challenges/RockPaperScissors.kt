package challenges

import Challenge
import java.io.File

class RockPaperScissors : Challenge {
    override fun getName(): String {
        return "Day 2 - Rock Paper Scissors"
    }

    override fun getInputFileName(): String {
        return "day2.txt"
    }

    /**
     * "The first column is what your opponent is going to play:
     *      A for Rock,
     *      B for Paper, and
     *      C for Scissors. The second column--"
     *
     *      X for Rock, Y for Paper, and Z for Scissors
     *
     * The winner of the whole tournament is the player with the highest score.
     * Your total score is the sum of your scores for each round.
     * The score for a single round is the score for the shape you selected
     * (1 for Rock, 2 for Paper, and 3 for Scissors)
     * plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)
     *
     * X means you need to lose,
     * Y means you need to end the round in a draw, and
     * Z means you need to win. Good luck!"
     */
    override fun solve(input: File?): String {
        var score = 0
        input!!.forEachLine { score += getScoreFromLine(it) }

        return score.toString()
    }

    fun getScoreFromLine(line: String): Int {
        val choices = line.split(" ").map { when(it) {
            "A" -> Choice.Rock
            "B" -> Choice.Paper
            "C" -> Choice.Scissors
            "X" -> Outcome.Loss
            "Y" -> Outcome.Draw
            "Z" -> Outcome.Win
            else -> throw Exception("Unknown input.")
        } }
        val choice = executeRound2(opponent = choices[0] as Choice, outcome = choices[1] as Outcome)

        return choice.getScore() + (choices[1] as Outcome).getScore()
    }

    private fun executeRound(us: Choice, opponent: Choice): Outcome {
        return when (us) {
            opponent -> Outcome.Draw
            else -> when {
                (us == Choice.Scissors && opponent == Choice.Paper) ||
                        (us == Choice.Rock && opponent == Choice.Scissors) ||
                        (us == Choice.Paper && opponent == Choice.Rock) -> Outcome.Win

                else -> Outcome.Loss
            }
        }
    }

    private fun executeRound2(opponent: Choice, outcome: Outcome): Choice {
        return when (outcome) {
            Outcome.Draw -> opponent
            Outcome.Win -> when(opponent) {
                Choice.Rock -> Choice.Paper
                Choice.Paper -> Choice.Scissors
                Choice.Scissors -> Choice.Rock
            }
            Outcome.Loss -> when(opponent) {
                Choice.Rock -> Choice.Scissors
                Choice.Paper -> Choice.Rock
                Choice.Scissors -> Choice.Paper
            }
        }
    }

    enum class Outcome {
        Win,
        Draw,
        Loss;

        fun getScore(): Int = when (this) {
            Win -> 6
            Draw -> 3
            Loss -> 0
        }
    }

    enum class Choice {
        Rock, Paper, Scissors;

        fun getScore(): Int = when (this) {
            Rock -> 1
            Paper -> 2
            Scissors -> 3
        }

        companion object {
            fun getFromString(string: String): Choice = when (string) {
                "A", "X" -> Rock
                "B", "Y" -> Paper
                "C", "Z" -> Scissors
                else -> throw Exception("Unknown input.")
            }
        }
    }
}