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
	go2Root()

	game := tr.InitGame()

	defer game.Fini()

	//var Chuck engine.GameObject

	//Chuck.CellData = engine.InitSprite("Chuck")

	//game.Objects = append(game.Objects, Chuck)

	//game.Player = &game.Objects[0]
	go game.EventLoop()
	go game.RenderScreenLoop()
	for game.Run {
		time.Sleep(100 * time.Millisecond)
		//game.Screen.Clear()
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
