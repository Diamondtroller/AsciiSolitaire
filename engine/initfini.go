package engine

import (
	"fmt"
	"os"

	"github.com/bennicholls/burl-E/reximage"
	"github.com/gdamore/tcell"
	"golang.org/x/text/encoding/charmap"
)

var spriteCache = make(map[string]CellSprite)

// Init Initializes game
func (g *Game) Init() {
	var err error
	(*g).Screen, err = tcell.NewScreen()
	aPanic(err)
	err = (*g).Screen.Init()
	aPanic(err)
	(*g).KeyBindings = make(map[tcell.Key]func())
	(*g).KeyBindingAssignment()
	//(*g).MouseBindings = make(map[tcell.ButtonMask]func(*tcell.EventMouse))
	(*g).Objects = make([]GameObject, 0)
	(*g).Run = true

}

//KeyBindingAssignment Reads file and assigns key to the functions
func (g *Game) KeyBindingAssignment() {
	(*g).KeyBindings[tcell.KeyESC] = func() { (*g).Run = false }
}

//InitSprite initializes sprites
func InitSprite(name string) CellSprite {
	if cachedData, exists := spriteCache[name]; exists {
		return cachedData
	}
	var CellData CellSprite
	spriteFolder := "Sprites"
	var rex reximage.ImageData
	var err error
	rex, err = reximage.Import(spriteFolder + string(os.PathSeparator) + name + ".xp")
	if err == nil {
		fmt.Println("No error while importing", name)
	}
	aPanic(err)
	var i, j int
	var runeColourF, runeColourB tcell.Color
	var runeStyle tcell.Style
	CellData = make(CellSprite, rex.Height)
	for i := range CellData {
		CellData[i] = make([]Cell, rex.Width)
	}
	for n, cell := range rex.Cells {
		i = n % rex.Width
		j = n / rex.Width
		runeColourB = tcell.NewRGBColor((int32)(cell.R_b), (int32)(cell.G_b), (int32)(cell.B_b))
		runeStyle = runeStyle.Background(runeColourB)
		runeColourF = tcell.NewRGBColor((int32)(cell.R_f), (int32)(cell.G_f), (int32)(cell.B_f))
		runeStyle = runeStyle.Foreground(runeColourF)
		CellData[j][i].glyph = charmap.CodePage437.DecodeByte(byte(cell.Glyph))
		CellData[j][i].form = runeStyle
	}
	return CellData
}

//Fini Finalizes game
func (g *Game) Fini() {
	(*g).Screen.Fill(' ', tcell.StyleDefault)
	(*g).Screen.Clear()
	(*g).Screen.Fini()
}
