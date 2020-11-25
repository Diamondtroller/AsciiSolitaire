package engine

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
)

func (g *Game) renderSpeed(rTime chan time.Time) {
	var average, tmp float64
	var t, lt time.Time
	var i int
	samples := 100
	tSlice := make([]float64, samples)
	st := tcell.StyleDefault
	st = st.Foreground(tcell.ColorWhite)
	st = st.Background(tcell.ColorGray)
	lt = time.Now()
	for t = range rTime {
		if !(*g).Debug {
			continue
		}
		i = ((i + 1) % samples)
		tSlice[i] = t.Sub(lt).Seconds()
		lt = t
		for _, tmp = range tSlice {
			average += tmp / float64(samples)
		}
		(*g).funcChan <- func() {
			(*g).DrawText(
				1,
				1,
				fmt.Sprintf("FPS: %.3f  |  LOOP TIME: %.6f s  |  %.6f ms", 1.0/average, average, (1000.0*average)),
				st,
			)
		}
		//close((*g).funcChan)
		average = 0
	}
}

//DebugClear Clears debugging Data
func (g *Game) DebugClear() {
	(*g).DrawRect(1, 1, 40, 1, ' ', tcell.StyleDefault)
}
