package main

import (
	"math/rand"
	"strconv"
	"time"
)

func seedRandom() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// pad string given length and base
func pad(n, base, length int) string {
	str := strconv.FormatInt(int64(n), base)
	for len(str) != length {
		str = "0" + str
	}
	return str
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

// make a new bot from random values if probability check is met
func newRandomBot() Bot {
	// - 4 + 1 to avoid screen edges
	x := rand.Intn(gameWidth-4) + 1
	code := rand.Intn(255)
	return Bot{x, 0, 0, code}
}

// run through all the bots and destroy all the ones with matching codes
func filterBotMatch(launchCode int, bots []Bot) ([2]int, []Bot) {
	matchCoords := [2]int{-1, -1}
	var noMatch []Bot
	for _, b := range bots {
		if b.code == launchCode && matchCoords[0] == -1 {
			matchCoords = [2]int{b.x, b.age}
		} else {
			noMatch = append(noMatch, b)
		}
	}
	return matchCoords, noMatch
}

func gameTick(bots []Bot, ticks, score int) []Bot {
	bots = dropBots(bots, score)
	if botProbability(len(bots), ticks) > rand.Float64() {
		bots = append(bots, newRandomBot())
	}
	return bots
}

// let the bots fall by one every tickRate
func dropBots(bots []Bot, score int) []Bot {
	var newBots []Bot
	for _, b := range bots {
		b.age += 1
		if b.age == dropEvery {
			b.age = 0
			b.y += 1
		}
		newBots = append(newBots, b)
		if b.y == gameHeight {
			quit(false, score)
		}
	}
	return newBots
}
