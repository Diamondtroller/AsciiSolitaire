package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"../engine"
	//char "golang.org/x/text/encoding/charmap"
	//textencoding "golang.org/x/text/encoding"
)

func main() {
	//var err error
	go2Root()
	str, _ := os.Getwd()
	fmt.Println("The path is: ", str)
	//var Screen tcell.Screen
	var game engine.Game
	var Chuck engine.GameObject
	Chuck.X = 10
	Chuck.Y = 10
	Chuck.InitSprite("Chuck")
	t := time.Now()
	//tcell.RegisterEncoding("CodePage437", char.CodePage437)
	game.Init()
	defer game.Screen.Fini()
	/*if Screen, err = tcell.NewScreen(); err == nil {
		Screen.Init()
		defer Screen.Fini()
	}*/
	//aPanic(err)
	//trigger := tcell.NewEventKey(tcell.KeyEscape, ' ', tcell.ModNone)
	//Screen, err = tcell.NewScreen()
	//var ev tcell.Event
	for {
		//Screen.Fill(' ', tcell.StyleDefault)
		//ev = Screen.PollEvent()
		if time.Since(t) > 15*time.Second {
			break
		}
		game.Draw(&Chuck)
		game.Screen.Show()
	}
}

func aPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func go2Root() {
	var err error
	var path string
	path, err = os.Executable()
	aPanic(err)
	path, err = filepath.EvalSymlinks(path)
	aPanic(err)
	if _, err := os.Stat(path + string(os.PathSeparator) + "main"); os.IsNotExist(err) {
		err = os.Chdir(".." + string(os.PathSeparator) + path)
	} else {
		err = os.Chdir(path)
	}
	aPanic(err)
}
