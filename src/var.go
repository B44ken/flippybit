package main

type Bot struct {
	x    int
	y    int
	age  int
	code int
}

// todo: replace with unicode seperators
const gfxSides = "|                                    |"
const gfxSeperator = "+------------------------------------+"
const gfxWords = "| HEX  | BINARY    | SCORE  | BEST   |"
const gfxGameOver = "|          | GAME OVER |             |"
const gfxGameOverSep = "|          +-----------+             |"
const gameWidth = 36
const gameHeight = 20
const tickRate = 8
const dropEvery = tickRate * 1.5

// needs work?
func botProbability(botNumInt, timeInt int) float64 {
	botNum := float64(botNumInt)
	time := float64(timeInt)
	if botNum == 0 {
		return 0.2
	}
	return (0.012 / botNum) + (time * 0.0002)
}
