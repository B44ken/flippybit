package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// todo: make these global
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
	runCmd("stty -F '/dev/tty' -icanon min 1")
	defer runCmd("stty sane")

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

func runCmd(cmd string) {
	instance := exec.Command("bash", "-c", cmd)
	instance.Run()
}

// reset terminal on ctrl c (re-buffer)
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

// todo: write high score to disk?
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
