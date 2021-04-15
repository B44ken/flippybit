package main

type Bot struct {
	x    int
	y    int
	code int
}

var game string = `
+--------------------------------+
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
|                                |
|                                |
|                                |
+--------------------------------+
| HEX | BINARY   | SCORE | BEST  |
| 00  | 00000000 | 0000  | 0000  |
+--------------------------------+
`

// 'graphics' in the most liberal way possible
const gfxSeperator = "+--------------------------------+"
const gfxSides = "|                                |"
const gfxWords = "| HEX | BINARY   | SCORE | BEST  |"
const gameWidth = 32
const gameHeight = 16
const tickRate = 1.4
