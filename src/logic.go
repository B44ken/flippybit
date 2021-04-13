package main

import "math/rand"

type Bot struct {
	x    int
	y    int
	code int64
}

func makeLaunchCode(cur int64, key byte) int64 {
	key -= 49
	k := int64(key)
	if k > 8 {
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
	code := rand.Int63n(255)
	return Bot{x, y, code}
}

func filterBotMatch(launchCode int64, bots []Bot) (Bot, []Bot) {
	var match Bot
	var noMatch []Bot
	for _, b := range bots {
		if b.code == launchCode {
			match = b
		} else {
			noMatch = append(noMatch, b)
		}
	}
	return match, noMatch
}
