package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	handleQuit()
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	rand.NewSource(time.Now().UnixNano())

	key := make([]byte, 1)
	listenForKey(&key)

	var bots []Bot
	for i := 0; i < 10; i++ {
		bots = append(bots, newRandomBot())
	}
	fmt.Println(bots)

	var launchCode int64
	for {
		if key[0] != 0 {
			launchCode = makeLaunchCode(launchCode, key[0])
			_, bots = filterBotMatch(launchCode, bots)
			fmt.Println(launchCode, "\n", bots)
			key[0] = 0
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
		exec.Command("reset").Run()
		os.Exit(0)
	}()
}

func listenForKey(keys *[]byte) {
	go func() {
		for {
			os.Stdin.Read(*keys)
		}
	}()
}
