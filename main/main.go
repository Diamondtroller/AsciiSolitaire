package main

import (
	"os"
	"time"

	tr "github.com/Diamondtroller/AsciiSolitaire/trexsolitaire"
)

func main() {
	//Initializes game object
	game := tr.InitGame()
	//Defers game closing to the end
	defer game.Fini()

	go game.EventLoop()
	go game.RenderScreenLoop()
	for game.Run {
		time.Sleep(100 * time.Millisecond)
	}
}

func aPanic(e error) {
	if e != nil { //Checks if error is nil if it is panics
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
