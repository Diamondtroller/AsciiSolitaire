package engine

import (
	"github.com/bennicholls/burl-E/reximage"
	"github.com/gdamore/tcell"
)

type (
	//GameObject is struct for all game objects one screen
	GameObject struct {
		X, Y   int
		Sprite reximage.ImageData
	}
	//Game conatins all stuff for main game run
	Game struct {
		objects []GameObject
		Screen  tcell.Screen
		root    string
	}
)

func aPanic(e error) {
	if e != nil {
		panic(e)
	}
}
