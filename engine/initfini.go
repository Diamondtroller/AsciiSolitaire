package engine

import (
	"fmt"
	"os"

	"github.com/bennicholls/burl-E/reximage"
	"github.com/gdamore/tcell"
)

// Init Initializes game
func (g *Game) Init() {
	var err error
	(*g).Screen, err = tcell.NewScreen()
	aPanic(err)
	//tcell.RegisterEncoding("IBM Code Page 437", char.CodePage437)
	(*g).Screen.Init()
	//return
	(*g).KeyBindings = make(map[tcell.Key]func())
	(*g).MouseBindings = make(map[tcell.ButtonMask]func(*tcell.EventMouse))
	(*g).Objects = make([]GameObject, 0)
	(*g).funcChan = make(chan func(), 8)
	(*g).Run = true

}

//InitSprite initializes sprites
func (thisObj *GameObject) InitSprite(name string) {
	spriteFolder := "Sprites"
	var tmp reximage.ImageData
	var err error
	tmp, err = reximage.Import(spriteFolder + string(os.PathSeparator) + name + ".xp")
	if err == nil {
		fmt.Println("No error importingwhile importing", name)
	}
	aPanic(err)
	thisObj.Sprite = tmp
	return
}

//Fini Finalizes game
func (g *Game) Fini() {
	(*g).Screen.Fini()
}
