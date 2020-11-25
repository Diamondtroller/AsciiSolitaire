package main

import (
	"os"
	"time"

	"github.com/Diamondtroller/Asciicker/engine"
)

func main() {
	go2Root()
	var game engine.Game
	game.Init()
	defer game.Fini()
	game.AnimationSpeed = 15 * time.Millisecond
	var Chuck engine.GameObject
	Chuck.X = 10
	Chuck.Y = 10
	Chuck.InitSprite("Chuck")
	game.Objects = append(game.Objects, Chuck)
	go game.EventLoop()
	go game.ScreenLoop()
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
	//var path string
	//path, err = os.Getwd()
	//err = os.Stat("main")
	_, err = os.Stat("Sprites")
	for ; os.IsNotExist(err); _, err = os.Stat("Sprites") {
		os.Chdir("..")
	} /*
		path, err = os.Executable()
		aPanic(err)
		path, err = filepath.EvalSymlinks(path)
		aPanic(err)

		err = os.Chdir(path)*/
	aPanic(err)
}

