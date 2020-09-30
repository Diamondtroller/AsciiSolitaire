package engine

import (
	"github.com/bennicholls/burl-E/reximage"
)

type (
	//GameObject is struct for all game objects one screen
	GameObject struct {
		X, Y   int
		Sprite reximage.ImageData
	}
	//Game conatins all stuff for main game run
)

// Init Initializes game
/*func Init() (Screen tcell.Screen) {
	var err error
	Screen, err = tcell.NewScreen()
	aPanic(err)
	//tcell.RegisterEncoding("IBM Code Page 437", char.CodePage437)
	return
}*/

//InitSprite initializes sprites
func (thisObj *GameObject) InitSprite(name string) {
	spritePath := "..\\Sprites\\"
	var tmp reximage.ImageData
	var err error
	tmp, err = reximage.Import(spritePath + name + ".xp")
	aPanic(err)
	thisObj.Sprite = tmp
	return
}

func aPanic(e error) {
	if e != nil {
		panic(e)
	}
}
