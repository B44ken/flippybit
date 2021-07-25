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
func drawData(bots []Bot, score, launchCode int) {
	for _, b := range bots {
		fmt.Printf("(%d %d %d %s) ", b.x, b.age, b.y, toHex(b.code))
	}
	fmt.Println(score, toHex(launchCode), toBin(launchCode))
}

func drawScreen(bots []Bot, score, launchCode int) {
	clearScreen()
	drawData(bots, score, launchCode)
	fmt.Println(gfxSeperator)
	drawMain(bots)
	fmt.Println(gfxSeperator)
	fmt.Println(gfxWords)
	// todo: print values
	fmt.Println(gfxSeperator)
}

func drawMain(bots []Bot) {
	// todo: move to drawMain() or something
	for i := 0; i < 16; i++ {
		var line = []byte(gfxSides)
		for _, b := range bots {
			if(b.y == i) {
				code := toHex(b.code)
				line[b.x] = '<'
				line[b.x+1] = code[0]
				line[b.x+2] = code[1]
				line[b.x+3] = '>'
			}
		}
		fmt.Println(string(line))
	}
}