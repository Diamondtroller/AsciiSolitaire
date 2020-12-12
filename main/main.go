package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell"

	as "github.com/Diamondtroller/AsciiSolitaire/asciisolitaire"
	"github.com/Diamondtroller/AsciiSolitaire/engine"
)

// var (
// 	game engine.Game
// )

func main() {
	go2Root()

	game := as.Init()

	defer game.Fini()

	game.AnimationSpeed = 70 * time.Millisecond
	game.BGglyph = '░'
	game.BGstyle = game.BGstyle.Background(tcell.ColorGreen)
	game.BGstyle = game.BGstyle.Foreground(tcell.NewHexColor(0x16950d))

	var Chuck engine.GameObject

	Chuck.InitSprite("Chuck")

	game.Objects = append(game.Objects, Chuck)

	game.Player = &game.Objects[0]
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
