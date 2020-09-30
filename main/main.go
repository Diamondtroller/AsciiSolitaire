package main

import (
	"os"
	"time"

	"../engine"
	"github.com/gdamore/tcell"
	//char "golang.org/x/text/encoding/charmap"
	//textencoding "golang.org/x/text/encoding"
)

func main() {
	os.Chdir(programPath())
	var err error
	var Screen tcell.Screen
	//var Game engine.Game
	var Chuck engine.GameObject
	Chuck.X = 10
	Chuck.Y = 10
	Chuck.InitSprite("Chuck")
	t := time.Now()
	//tcell.RegisterEncoding("CodePage437", char.CodePage437)
	if Screen, err = tcell.NewScreen(); err == nil {
		Screen.Init()
		defer Screen.Fini()
	}
	aPanic(err)
	//trigger := tcell.NewEventKey(tcell.KeyEscape, ' ', tcell.ModNone)
	//Screen, err = tcell.NewScreen()

	for {
		//Screen.Fill(' ', tcell.StyleDefault)
		//ev := Screen.PollEvent()
		//fmt.Println(ev)
		if time.Since(t) > 5*time.Second {
			break
		}
		engine.Draw(Screen, &Chuck)
		Screen.Show()
	}
}

func programPath() string {
	path, errPath := os.Executable()
	if errPath != nil {
		panic(errPath)
	}
	return path
}

func aPanic(e error) {
	if e != nil {
		panic(e)
	}
}
