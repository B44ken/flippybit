package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
	"fmt"
)

func main() {
	var bots []Bot
	var matched [2]int
	var key byte
	globalStop := false
	score := 0

	listenForKey(&key)
	seedRandom()
	handleQuit(&globalStop)

	// unbuffer terminal for char-by-char input
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	// todo: remove & auto place bots
	for i := 0; i < 2; i++ {
		bots = append(bots, newRandomBot())
	}

	var launchCode int

	go func() {
		for {
			if globalStop {
				continue
			}
			drawScreen(bots, score, launchCode)
			// there's definitely a better way than manually refreshing
			time.Sleep(time.Second / tickRate) 
			bots = dropBots(bots)
		}
	}()

	for {
		if key != 0 {
			launchCode = makeLaunchCode(launchCode, key)
			matched, bots = filterBotMatch(launchCode, bots)
			if matched != [2]int{-1, -1} {
				launchCode = 0
			}
			key = 0
		}
	}
}

// reset terminal on ctrl c (re-buffer)
func handleQuit(globalStop *bool) {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	signal.Ignore(syscall.SIGTERM)
	go func() {
		<-sig
		*globalStop = true
		quit(true)
	}()
}
	
func quit(wasManual bool) {
	exec.Command("stty", "sane").Run()
	if wasManual {
		fmt.Println("game over!")
		// todo: final score, etc.
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
