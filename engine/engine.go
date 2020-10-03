package engine

import (
	"time"

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
		Run            bool
		Debug          bool
		AnimationSpeed time.Duration
		Objects        []GameObject
		Screen         tcell.Screen
		KeyBindings    map[tcell.Key]func()
		MouseBindings  map[tcell.ButtonMask]func(*tcell.EventMouse)
	}
)

//EventLoop Handles all events
func (g *Game) EventLoop() {
	var ev tcell.Event
	var action func()
	(*g).KeyBindings[tcell.KeyESC] = func() { (*g).Run = false }
	(*g).KeyBindings[tcell.KeyF3] = func() { (*g).Debug = !(*g).Debug; (*g).renderSpeedClear() }
	for {
		ev = (*g).Screen.PollEvent()
		switch event := ev.(type) {
		case *tcell.EventKey:
			action = (*g).KeyBindings[event.Key()]
			if action != nil {
				action()
			}
			/*switch event.Key() {

				case tcell.KeyEscape, tcell.KeyEnter:
					(*g).Run = false
					return
				case tcell.KeyCtrlL:
					(*g).Screen.Sync()
			}*/
		case *tcell.EventResize:
			(*g).Screen.Sync()
		}
	}
}

// func (Objs *[]GameObject) Append (someObj GameObject) {

// }

func aPanic(e error) {
	if e != nil {
		panic(e)
	}
}
