package widget

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
)

type CommandArea struct {
	spartan.View
}

func (ca CommandArea) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			termbox.SetCell(int(i), int(j), ' ', ca.ForegroundColor, ca.BackgroundColor)
		}
	}

	ca.drawWaitingState(left, top, right, bottom)
}

func (ca CommandArea) drawWaitingState(left uint32, top uint32, right uint32, bottom uint32) {
	msgText := "[Press Space]"

	x := left + (right-left)/2 - uint32(len(msgText)/2)
	y := top + (bottom-top)/2

	for i := x; i < (x + uint32(len(msgText))); i++ {
		termbox.SetCell(int(i), int(y), rune(msgText[i-x]), ca.ForegroundColor, ca.BackgroundColor)
	}
}

func (ca CommandArea) OnTick() {
}
