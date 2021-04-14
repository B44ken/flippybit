package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var score int64

	handleQuit()
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var key byte
	listenForKey(&key)

	var bots []Bot
	for i := 0; i < 10; i++ {
		bots = append(bots, newRandomBot())
	}
	fmt.Println(bots)

	var launchCode int64
	for {
		if key != 0 {
			launchCode = makeLaunchCode(launchCode, key)
			var matched [2]int
			matched, bots = filterBotMatch(launchCode, bots)
			if matched != [2]int{-1, -1} {
				launchCode = 0
			}
			drawScreen(bots, score, launchCode)
			key = 0
		}
	}
}

// reset terminal on ctrl c (re-buffer)
func handleQuit() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	signal.Ignore(syscall.SIGTERM)
	go func() {
		<-sig
		exec.Command("stty", "echo", "cooked").Run()
		time.Sleep(1)
		os.Exit(0)
	}()
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
