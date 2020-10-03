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
		(*g).DrawText(1, 1, fmt.Sprintf("FPS: %.3f", 1.0/average), st)
		(*g).DrawText(1, 2, fmt.Sprintf("LOOP TIME: %.6fs", average), st)
		(*g).DrawText(12, 3, fmt.Sprintf("%.6fms", (1000.0*average)), st)
		average = 0
	}
}

func (g *Game) renderSpeedClear() {
	(*g).DrawText(1, 1, "         ", tcell.StyleDefault)
	(*g).DrawText(1, 2, "         ", tcell.StyleDefault)
	(*g).DrawText(12, 3, "     ", tcell.StyleDefault)
}
