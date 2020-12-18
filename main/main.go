package main

import (
	"os"
	"time"

	tr "github.com/Diamondtroller/AsciiSolitaire/trexsolitaire"
)

// var (
// 	game engine.Game
// )

func main() {
	game := tr.InitGame()

	defer game.Fini()

	go game.EventLoop()
	go game.RenderScreenLoop()
	for game.Run {
		time.Sleep(200 * time.Millisecond)
	}
}

func aPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func go2Root() {
	var err error
	_, err = os.Stat("Sprites")
	for ; os.IsNotExist(err); _, err = os.Stat("Sprites") {
		os.Chdir("..")
	}
	aPanic(err)
}
