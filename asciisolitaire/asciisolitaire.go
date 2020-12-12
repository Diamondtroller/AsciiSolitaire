package asciisolitaire

import (
	"container/list"

	"github.com/Diamondtroller/AsciiSolitaire/engine"
)

type (
	/*Card Stores card data
	has Open boolean value to know if player has discovered card
	and sprite of the card*/
	Card struct {
		Open   bool
		sprite engine.CellSprite
	}
	//Dstack Card stack that has cards shifted cards down
	Dstack list.List
	//Stack Card stack where cards are stacked one on top of the other
	Stack list.List
)

/*Init initializes Ascii Solitaire:
* Initializes game on engine level
* Sets all needed variables for engine.game
 */
func Init() (game *engine.Game) {
	game.Init()
	return
}

//InitBoard Initializes soloitaire board to place cards on
func InitBoard(g *engine.Game) {}

//InitCards generates all cards
func InitCards(g *engine.Game) {}
