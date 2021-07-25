package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var bots []Bot
	var matched [2]int
	var key byte
	globalStop := false
	score := 0
	ticks := 0

	listenForKey(&key)
	seedRandom()
	handleQuit(&globalStop)

	// unbuffer terminal for char-by-char input
	exec.Command("stty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-echo").Run()
	defer exec.Command("stty", "echo").Run()

	bots = append(bots, newRandomBot())

	launchCode := 0

	go func() {
		for {
			if globalStop {
				continue
			}
			ticks++
			drawScreen(bots, score, launchCode, false)
			// definitely a better way than manually refreshing
			time.Sleep(time.Second / tickRate)
			bots = gameTick(bots, ticks/tickRate, score)
		}
	}()

	for {
		if key != 0 {
			launchCode = makeLaunchCode(launchCode, key)
			matched, bots = filterBotMatch(launchCode, bots)
			if matched != [2]int{-1, -1} {
				launchCode = 0
				score++
			}
			key = 0
		}
	}
}

// reset terminal on ctrl c (re-buffer)
// todo: fix?
func handleQuit(globalStop *bool) {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	signal.Ignore(syscall.SIGTERM)
	go func() {
		<-sig
		*globalStop = true
		quit(true, 0)
	}()
}

// todo: write high score to disk
func quit(wasManual bool, score int) {
	exec.Command("stty", "sane").Run()
	if !wasManual {
		drawScreen(make([]Bot, 0), score, 0, true)
	}
	os.Exit(0)
}

func listenForKey(key *byte) {
	go func() {
		for {
			keys := make([]byte, 1)
			os.Stdin.Read(keys)
			*key = keys[0]
		}
	}()
}
