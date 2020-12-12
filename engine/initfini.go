package engine

import (
	"fmt"
	"os"

	"github.com/bennicholls/burl-E/reximage"
	"github.com/gdamore/tcell"
	"golang.org/x/text/encoding/charmap"
)

// Init Initializes game
func (g *Game) Init() {
	var err error
	(*g).Screen, err = tcell.NewScreen()
	aPanic(err)
	//tcell.RegisterEncoding("IBM Code Page 437", char.CodePage437)
	(*g).Screen.Init()
	(*g).KeyBindings = make(map[tcell.Key]func())
	(*g).KeyBindingAssignment()
	(*g).MouseBindings = make(map[tcell.ButtonMask]func(*tcell.EventMouse))
	(*g).Objects = make([]GameObject, 0)
	//(*g).funcChan = make(chan func(), 8)
	(*g).Run = true

}

//KeyBindingAssignment Reads file and assigns key to the functions
func (g *Game) KeyBindingAssignment() {
	(*g).KeyBindings[tcell.KeyESC] = func() { (*g).Run = false }
	(*g).KeyBindings[tcell.KeyUpLeft] = func() { (*g).Player.X--; (*g).Player.Y-- }
	(*g).KeyBindings[tcell.KeyUpRight] = func() { (*g).Player.X++; (*g).Player.Y-- }
	(*g).KeyBindings[tcell.KeyDownLeft] = func() { (*g).Player.X--; (*g).Player.Y++ }
	(*g).KeyBindings[tcell.KeyDownRight] = func() { (*g).Player.X++; (*g).Player.Y++ }
	(*g).KeyBindings[tcell.KeyRight] = func() { (*g).Player.X++ }
	(*g).KeyBindings[tcell.KeyLeft] = func() { (*g).Player.X-- }
	(*g).KeyBindings[tcell.KeyUp] = func() { (*g).Player.Y-- }
	(*g).KeyBindings[tcell.KeyDown] = func() { (*g).Player.Y++ }
}

//InitSprite initializes sprites
func (thisObj *GameObject) InitSprite(name string) {
	spriteFolder := "Sprites"
	var rex reximage.ImageData
	var err error
	rex, err = reximage.Import(spriteFolder + string(os.PathSeparator) + name + ".xp")
	if err == nil {
		fmt.Println("No error importingwhile importing", name)
	}
	aPanic(err)
	var i, j int
	var runeColourF, runeColourB tcell.Color
	var runeStyle tcell.Style
	(*thisObj).CellData = make(CellSprite, rex.Height)
	for i := range (*thisObj).CellData {
		(*thisObj).CellData[i] = make([]Cell, rex.Width)
	}
	for n, cell := range rex.Cells {
		i = n % rex.Width
		j = n / rex.Width
		runeColourB = tcell.NewRGBColor((int32)(cell.R_b), (int32)(cell.G_b), (int32)(cell.B_b))
		runeStyle = runeStyle.Background(runeColourB)
		runeColourF = tcell.NewRGBColor((int32)(cell.R_f), (int32)(cell.G_f), (int32)(cell.B_f))
		runeStyle = runeStyle.Foreground(runeColourF)
		(*thisObj).CellData[j][i].glyph = charmap.CodePage437.DecodeByte(byte(cell.Glyph))
		(*thisObj).CellData[j][i].form = runeStyle
	}
}

//Fini Finalizes game
func (g *Game) Fini() {
	(*g).Screen.Fini()
}
