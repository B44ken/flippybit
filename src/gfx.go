package main

import "fmt"

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func resetScreen() {
	fmt.Print("\033c")
}

func drawData(bots []Bot, score, launchCode int) {
	for _, b := range bots {
		fmt.Printf("(%d %d %s) ", b.x, b.y, toHex(b.code))
	}
	fmt.Println(score, toHex(launchCode), toBin(launchCode))
}

func drawScreen(bots []Bot, score, launchCode int) {
	clearScreen()
	drawData(bots, score, launchCode)
	fmt.Println(gfxSeperator)
	for i := 0; i < 16; i++ {
		// todo: draw bots
		fmt.Println(gfxSides)
	}
	fmt.Println(gfxSeperator)
	fmt.Println(gfxWords)
	// todo: print values
	fmt.Println(gfxSeperator)
}
