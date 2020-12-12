package engine

import (
	"time"

	"github.com/gdamore/tcell"
	//char "golang.org/x/text/encoding/charmap"
)

//DrawSprite Adds sprites to tcell screen
func (g *Game) DrawSprite(tObj *GameObject) {

	for i, cellrow := range (*tObj).CellData {
		for j, cell := range cellrow {
			(*g).Screen.SetContent(tObj.X+j, tObj.Y+i, cell.glyph, nil, cell.form)
		}
		//fmt.Println(i)
		//b, _ := char.CodePage437.EncodeRune(rune(cell.Glyph))

		//(*g).Screen.SetContent(runeX, runeY, cRune, nil, runeStyle)
		//Screen.SetContent(runeX, runeY, rune(cell.Glyph), nil, runeStyle)
	}
}

//DrawText Adds text to tcell screen buffer
func (g *Game) DrawText(x, y int, text string, s tcell.Style) {
	for i, letter := range text {
		(*g).Screen.SetContent(x+i, y, letter, nil, s)
	}
}

//DrawRect Adds rectangle to tcell screen buffer
func (g *Game) DrawRect(x, y, w, h int, r rune, s tcell.Style) {
	for i := 0; i < w*h; i++ {
		(*g).Screen.SetContent(x+(i%w), y+(i/w), r, nil, s)
	}
}

//ObjectLoop Draws all added objects to the screen
func (g *Game) ObjectLoop() {
	for _, Obj := range (*g).Objects {
		(*g).DrawSprite(&Obj)
	}
}

//RenderScreenLoop  Renders entire screen with sprites in loop
func (g *Game) RenderScreenLoop() {
	for (*g).Run {
		(*g).Screen.Fill(g.BGglyph, g.BGstyle)
		(*g).ObjectLoop()
		(*g).Screen.Show()
		time.Sleep((*g).AnimationSpeed)
	}
}

// //ScreenLoop  Renders entire screen with sprites in loop
// func (g *Game) ScreenLoop() {
// 	// var action func()
// 	timeChan := make(chan time.Time, 5)
// 	go (*g).renderSpeed(timeChan)
// 	timeChan <- time.Now()
// 	for (*g).Run {
// 		(*g).ObjectLoop()
// 		<-(*g).funcChan
// 		// ActionLoop:
// 		// 	for {
// 		// 		select {
// 		// 		case action = <-(*g).funcChan:
// 		// 			action()
// 		// 		default:
// 		// 			break ActionLoop
// 		// 		}
// 		// 	}
// 		(*g).Screen.Show()
// 		if (*g).Debug {
// 			timeChan <- time.Now()
// 		}
// 	}
// }
