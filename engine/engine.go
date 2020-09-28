package engine

import (
	"fmt"

	"github.com/bennicholls/burl-E/reximage"
)

type (
	//GameObject is struct for all game objects one screen
	GameObject struct {
		X, Y   int
		Sprite reximage.ImageData
	}
)

//InitSprite initializes sprites
func (thisObj *GameObject) InitSprite(name string) {
	spritePath := "..\\Sprites\\"
	var tmp reximage.ImageData
	var err error
	tmp, err = reximage.Import(spritePath + name + ".xp")
	if err == nil {
		thisObj.Sprite = tmp
	} else {
		fmt.Println(err)
	}
	return
}
