package main

import (
	"os"
	"time"

	"../engine"
	"github.com/gdamore/tcell"
)

func main() {
	os.Chdir(programPath())
	var Screen tcell.Screen
	var err error
	var Chuck engine.GameObject
	Chuck.X = 10
	Chuck.Y = 10
	Chuck.InitSprite("Chuck")
	t := time.Now()
	//trigger := tcell.NewEventKey(tcell.KeyEscape, ' ', tcell.ModNone)
	if Screen, err = tcell.NewScreen(); err == nil {
		Screen.Init()
		defer Screen.Fini()
	}
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
