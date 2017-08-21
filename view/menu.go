package view

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
)

type Menu struct {
	spartan.View
}

func (view Menu) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			termbox.SetCell(int(i), int(j), ' ', view.ForegroundColor, view.BackgroundColor)
		}
	}
}

func (view Menu) OnTick() {
}
