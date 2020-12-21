package engine

import (
	"time"

	"github.com/gdamore/tcell"
)

type (
	//Cell Tcell compatible cell struct
	Cell struct {
		glyph rune
		form  tcell.Style
	}
	//CellSprite 2 dimensional array of cells to store rexpaint sprites
	CellSprite [][]Cell
	//GameObject is struct for all game objects one screen
	GameObject struct {
		X, Y     int
		CellData CellSprite
		Hidden   bool
		DrawFunc func(g *Game)
	}
	//Game conatins all stuff for main game run
	Game struct {
		Run            bool
		Debug          bool
		AnimationSpeed time.Duration
		Objects        []GameObject
		Screen         tcell.Screen
		BGglyph        rune
		BGstyle        tcell.Style
		KeyBindings    map[tcell.Key]func()
		//MouseBindings  map[tcell.ButtonMask]func(*tcell.EventMouse)
	}
)

//EventLoop Handles all events
func (g *Game) EventLoop() {
	var ev tcell.Event
	var action func()
	// (*g).KeyBindings[tcell.KeyCtrlD] = func() { (*g).Debug = !(*g).Debug; (*g).funcChan <- (*g).DebugClear }
	// (*g).KeyBindings[tcell.KeyCtrlL] = func() { (*g).funcChan <- (*g).Screen.Sync }
	// (*g).KeyBindings[tcell.KeyCtrlE] = func() { (*g).funcChan <- (*g).Screen.Clear }
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
			(*g).Screen.Sync()
		case *tcell.EventKey:
			if action = (*g).KeyBindings[event.Key()]; action != nil {
				action()
			}
		}
	}
}

func aPanic(e error) {
	if e != nil {
		panic(e)
	}
}
