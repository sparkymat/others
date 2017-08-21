package event

import (
	"os"

	termbox "github.com/nsf/termbox-go"
)

func Handler(ch chan termbox.Event) {
	for {
		ev := <-ch
		handleEvent(ev)
	}
}

func handleEvent(ev termbox.Event) {
	switch ev.Type {
	case termbox.EventKey:
		if ev.Key == termbox.KeyEsc {
			termbox.Close()
			os.Exit(0)
		} else if ev.Key == termbox.KeySpace {
		}
	}
}
