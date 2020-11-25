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
		funcChan       chan func()
		KeyBindings    map[tcell.Key]func()
		MouseBindings  map[tcell.ButtonMask]func(*tcell.EventMouse)
	}
)

//EventLoop Handles all events
func (g *Game) EventLoop() {
	var ev tcell.Event
	var action func()
	(*g).KeyBindings[tcell.KeyESC] = func() { (*g).Run = false }
	(*g).KeyBindings[tcell.KeyCtrlD] = func() { (*g).Debug = !(*g).Debug; (*g).funcChan <- (*g).DebugClear }
	(*g).KeyBindings[tcell.KeyCtrlL] = func() { (*g).funcChan <- (*g).Screen.Sync }
	(*g).KeyBindings[tcell.KeyCtrlE] = func() { (*g).funcChan <- (*g).Screen.Clear }
	for {
		ev = (*g).Screen.PollEvent()
		switch event := ev.(type) {
		/*switch event.Key() {

			case tcell.KeyEscape, tcell.KeyEnter:
				(*g).Run = false
				return
			case tcell.KeyCtrlL:
				(*g).Screen.Sync()
		}*/
		case *tcell.EventResize:
			(*g).funcChan <- (*g).Screen.Sync
		case *tcell.EventKey:
			action = (*g).KeyBindings[event.Key()]
			if action != nil {
				action()
			}
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
