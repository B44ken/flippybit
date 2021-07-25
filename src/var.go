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
const gameWidth = 36
const gameHeight = 20
const tickRate = 8
const dropEvery = tickRate * 1.5
const botProbability = 0.02

var DEBUG = ""

// bot prob. as a function of existing bots and time?
// func botProbability(botNum time int) float64 { }
