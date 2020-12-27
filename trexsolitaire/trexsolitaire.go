package trexsolitaire

import (
	"container/list"
	"math/rand"
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
	cardW        = 15
	cardH        = 22
	diffY        = 3
)

var (
	suits = [4]rune{'♦', '♣', '♥', '♠'}
	ranks = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
)

type (
	/*Card Stores card data
	has Open boolean value to know if player has discovered card
	and sprite of the card*/
	Card struct {
		Open   bool
		suit   int
		rank   int
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
	rand.Seed(time.Now().UnixNano())
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
//CreatStack(game variable, deck of cards, stack's X, Y coord, cards in stack, cards shown, sprite shown on empty stack)
func CreateStack(g *engine.Game, deck *[]Card, X, Y, cards, cardsShown int, sprite string) (s Stack) {
	s.List = list.New()
	for _, curCard := range (*deck)[0:cards] {
		s.List.PushBack(curCard)
	}
	(*deck) = (*deck)[cards:]
	s.CardShow = cardsShown
	(*g).Objects = append((*g).Objects, engine.GameObject{X: X, Y: Y, CellData: engine.InitSprite(sprite), Hidden: false, DrawFunc: func(g *engine.Game) {
		(&s).DrawStack(g)
	}})
	s.Gameobj = &(*g).Objects[len((*g).Objects)-1]
	return s
}

//CreateBoard Initializes solitaire board to place cards on
func CreateBoard(g *engine.Game) (b *Board) {
	cards := cardSlice()
	b = new(Board)

	for i := 0; i < len(b.tablestacks); i++ {
		b.tablestacks[i] = CreateStack(g, &cards, 1+(cardW+1)*(i+1), boardoffsetY+1, i+1, showall, "place")
	}
	b.CardPile = CreateStack(g, &cards, 1, boardoffsetY, len(cards), showtop, "place")
	b.OpenCards = CreateStack(g, &cards, 1, boardoffsetY+cardH+1, 0, 3, "place")
	for i := 0; i < len(b.suitstacks); i++ {
		b.suitstacks[i] = CreateStack(g, &cards, 1+(cardW+1)*8, boardoffsetY+(cardH+1)*i, 0, showtop, "place")
	}
	return
}

//cardSlice Generates random deck of cards
func cardSlice() []Card { //TODO: Have to improve this function that generates shuffled cards
	slice := make([]Card, 4*13)
	for i := range suits {
		for j := range ranks {
			index := j + i/13
			slice[index].rank = i
			slice[index].suit = j
			slice[index].sprite = engine.InitSprite("place")
			slice[index].Open = false
		}
	}
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

func (s *Stack) DrawStack(g *engine.Game) {
	curCard := s.List.Front()
	if curCard == nil {
		g.DrawSprite(s.Gameobj.X, s.Gameobj.Y, s.Gameobj.CellData)
	} else if s.CardShow == showall {
		for i := 0; i < s.List.Len()-1; i++ {
			curCard.Value.(Card).draw(s.Gameobj.X, s.Gameobj.Y+i*diffY, false, g)
			curCard = curCard.Next()
		}
		s.List.Back().Value.(Card).draw(s.Gameobj.X, s.Gameobj.Y+(s.List.Len()-1)*diffY, true, g)
	} else if s.CardShow == showtop {
		s.List.Back().Value.(Card).draw(s.Gameobj.X, s.Gameobj.Y, true, g)
	}
}

func (c Card) draw(x, y int, last bool, g *engine.Game) {
	if last {
		if c.Open {
			g.DrawSprite(x, y, c.sprite)
		} else {
			g.DrawSprite(x, y, engine.InitSprite("hidden"))
		}
	} else {
		if c.Open {
			g.DrawSprite(x, y, c.sprite[:3])
		} else {
			g.DrawSprite(x, y, engine.InitSprite("hidden")[:3])
		}
	}
}
