package main

import (
	"math/rand"
	"strconv"
	"time"
)

func seedRandom() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// make 8 digit binary string
func toBin(n int) string {
	base := strconv.FormatInt(int64(n), 2)
	for {
		if len(base) == 8 {
			return base
		}
		base = "0" + base
	}
}

// make 2 digit hex string
func toHex(n int) string {
	base := strconv.FormatInt(int64(n), 16)
	if len(base) == 1 {
		base = "0" + base
	}
	return base
}

// flip a bit based on keyboard key (1-8)
func makeLaunchCode(cur int, key byte) int {

	key -= 49
	key = 7 - key
	if key > 7 {
		return cur
	}
	if key == 0 {
		return cur ^ 1
	}
	return cur ^ (1 << key)
}

func newRandomBot() Bot {
	y := gameHeight
	x := rand.Intn(gameWidth)
	code := rand.Intn(255)
	return Bot{x, y, code}
}

func filterBotMatch(launchCode int, bots []Bot) ([2]int, []Bot) {
	matchCoords := [2]int{-1, -1}
	var noMatch []Bot
	for _, b := range bots {
		if b.code == launchCode && matchCoords[0] == -1 {
			matchCoords = [2]int{b.x, b.y}
		} else {
			noMatch = append(noMatch, b)
		}
	}
	return matchCoords, noMatch
}
