package main

import (
	"fmt"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func resetScreen() {
	fmt.Print("\033c")
}

// dump bot data to terminal
func drawData(bots []Bot) {
	for _, b := range bots {
		fmt.Printf("(%d %d %d %s) ", b.x, b.age, b.y, toHex(b.code))
	}
}

func drawScreen(bots []Bot, score, launchCode int, gameOver bool) {
	clearScreen()
	fmt.Println(gfxSeperator)
	drawMain(bots, gameOver)
	fmt.Println(gfxSeperator)
	fmt.Println(gfxWords)
	drawVals(launchCode, score)
	fmt.Println(gfxSeperator)
}

func drawVals(launchCode, score int) {
	hex := pad(launchCode, 16, 2)
	bin := pad(launchCode, 2, 8)
	scr := pad(score, 10, 5)
	fmt.Printf("| %s   | %s  | %s  | 00000  |\n", hex, bin, scr)
}

func drawMain(bots []Bot, gameOver bool) {
	for i := 0; i < gameHeight; i++ {
		var line = []byte(gfxSides)
		for _, b := range bots {
			if b.y == i {
				code := toHex(b.code)
				line[b.x] = '<'
				line[b.x+1] = code[0]
				line[b.x+2] = code[1]
				line[b.x+3] = '>'
			}
		}
		if (i == 8 || i == 10) && gameOver {
			fmt.Println(gfxGameOverSep)
		} else if i == 9 && gameOver {
			fmt.Println(gfxGameOver)
		} else {
			fmt.Println(string(line))
		}
	}
}
