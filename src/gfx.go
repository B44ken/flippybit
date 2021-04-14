package main

import "fmt"

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func resetScreen() {
	fmt.Print("\033c")
}

func drawScreen(bots []Bot, score, launchCode int64) {
	clearScreen()
	for _, b := range bots {
		fmt.Printf("(%d %d %s) ", b.x, b.y, toHex(b.code))
	}
	fmt.Println(score, toHex(launchCode), toBin(launchCode))
}
