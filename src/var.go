package main

type Bot struct {
	x    int
	y    int
	age  int
	code int
}

/*  THE GAME
	+--------------------------------+
	|                                |
	|                                |
	|          +49+                  |
	|                                |
	|                                |
	|                                |
	|                                |
	|                                |
	|                                |
	|                                |
	|                                |
	|                                |
	|                                |
	|                                |
	|                                |
	|                                |
	|                                |
	+--------------------------------+
	| HEX | BINARY   | SCORE | BEST  |
	| 00  | 00000000 | 0000  | 0000  |
	+--------------------------------+
*/

// 'graphics' in the most liberal way possible
// todo: replace with unicode seperators
const gfxSides = "|                                    |"
const gfxSeperator = "+------------------------------------+"
const gfxWords = "| HEX  | BINARY    | SCORE  | BEST   |"
const gameWidth = 36
const gameHeight = 20
const tickRate = 8
const dropEvery = tickRate * 1.5