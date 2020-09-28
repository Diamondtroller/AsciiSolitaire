package engine

import (
	//"fmt"

	"github.com/gdamore/tcell"
	//"github.com/bennicholls/burl-E/reximage"
)

//Draw draws stuff
func Draw(GameScreen tcell.Screen, tObj *GameObject) {
	var runeX, runeY int
	var runeColourF, runeColourB tcell.Color
	var runeStyle tcell.Style
	for i, cell := range tObj.Sprite.Cells {
		//fmt.Println(i)
		runeX = tObj.X + i%tObj.Sprite.Width
		runeY = tObj.Y + (i / tObj.Sprite.Width)
		runeColourB = tcell.NewRGBColor(int32(cell.R_b), int32(cell.G_b), int32(cell.B_b))
		runeStyle = runeStyle.Background(runeColourB)
		runeColourF = tcell.NewRGBColor(int32(cell.R_f), int32(cell.G_f), int32(cell.B_f))
		runeStyle = runeStyle.Foreground(runeColourF)
		GameScreen.SetContent(runeX, runeY, rune(cell.Glyph), nil, runeStyle)
	}
}
