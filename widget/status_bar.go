package widget

import (
	"os"
	"time"

	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan"
)

type StatusBar struct {
	spartan.View

	Left   uint32
	Right  uint32
	Top    uint32
	Bottom uint32
}

func (sb *StatusBar) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	width := right - left + 1
	height := bottom - top + 1

	if width == 0 || height == 0 {
		return
	}

	sb.Left = left
	sb.Right = right
	sb.Top = top
	sb.Bottom = bottom

	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			termbox.SetCell(int(i), int(j), ' ', sb.ForegroundColor, sb.BackgroundColor)
		}
	}

	sb.drawCurrentPath(left, top, right-10, bottom)
	sb.drawClocK(right-9, top, right, bottom)
}

func (sb *StatusBar) drawCurrentPath(left uint32, top uint32, right uint32, bottom uint32) {
	currentPath, err := os.Getwd()
	if err != nil {
		return
	}

	width := right - left + 1
	length := uint32(len(currentPath))

	if length > width {
		length = width
	}

	for i := left; i < (left + length); i++ {
		termbox.SetCell(int(i), int(top+(bottom-top)/2), rune(currentPath[i-left]), sb.ForegroundColor, sb.BackgroundColor)
	}
}

func (sb *StatusBar) drawClocK(left uint32, top uint32, right uint32, bottom uint32) {
	t := time.Now()
	timeText := t.Format(" 15:04:05")

	width := right - left + 1
	length := uint32(len(timeText))

	if length > width {
		length = width
	}

	for i := left; i < (left + length); i++ {
		termbox.SetCell(int(i), int(top+(bottom-top)/2), rune(timeText[i-left]), sb.ForegroundColor, sb.BackgroundColor)
	}
}

func (sb *StatusBar) OnTick() {
	sb.Draw(sb.Left, sb.Top, sb.Right, sb.Bottom)
	termbox.Flush()
}
