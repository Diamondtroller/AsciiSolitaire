package trexsolitaire

import (
	"container/list"
	"time"

	"github.com/Diamondtroller/AsciiSolitaire/engine"
	"github.com/gdamore/tcell"
)

const (
	//diamond = 0
	//club    = 1
	//heart   = 2
	//spade   = 3
	showall      = -1
	showtop      = 0
	boardoffsetX = 0
	boardoffsetY = 2
	cardW        = 13
	cardH        = 18
)

var suit = [4]rune{'♦', '♣', '♥', '♠'}

type (
	/*Card Stores card data
	has Open boolean value to know if player has discovered card
	and sprite of the card*/
	Card struct {
		Open   bool
		sprite engine.CellSprite
	}
	//Stack Card stack that has cards shifted cards down
	Stack struct {
		List     *list.List
		Gameobj  *engine.GameObject
		CardShow int //Number of cards shown in stack
	}
	//Board Board contains all cards stacks of a solitaire game
	Board struct {
		CardPile    Stack
		OpenCards   Stack
		tablestacks [7]Stack
		suitstacks  [4]Stack
	}
)

//InitGame initializes Ascii Solitaire:
//Initializes game on engine level
//Sets all needed variables for engine.game
func InitGame() (g *engine.Game) {
	g = new(engine.Game)
	g.Init()

	g.AnimationSpeed = 1000 * time.Millisecond
	g.BGglyph = '░'
	g.BGstyle = g.BGstyle.Background(tcell.ColorGreen)
	g.BGstyle = g.BGstyle.Foreground(tcell.NewHexColor(0x16950d))
	CreateBoard(g)
	return
}

//CreateStack Creates a stack of cards
func CreateStack(g *engine.Game, X, Y, numOfCards int) (s Stack) {
	s.List = list.New()
	base := engine.InitSprite("place")
	g.Objects = append(g.Objects, engine.GameObject{X: X, Y: Y, CellData: base, Hidden: false})
	s.Gameobj = &g.Objects[len(g.Objects)-1]
	return
}

//CreateBoard Initializes solitaire board to place cards on
func CreateBoard(g *engine.Game) (b *Board) {
	b = new(Board)
	b.CardPile = CreateStack(g, 1, boardoffsetY, showtop)
	b.OpenCards = CreateStack(g, 1, boardoffsetY+cardH+1, 3)
	for i := 0; i < len(b.tablestacks); i++ {
		b.tablestacks[i] = CreateStack(g, 1+(cardW+1)*(i+1), boardoffsetY+1, showall)
		for j := 0; j <= i; j++ {
			b.tablestacks[i].List.Add
		}
	}
	for i := 0; i < len(b.suitstacks); i++ {
		b.suitstacks[i] = CreateStack(g, 1+(cardW+1)*8, boardoffsetY+(cardH+1)*i, showtop)
	}

	return
}

func SliceToLList([])