package app

import (
	"time"

	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/others/app/mode"
	"github.com/sparkymat/others/event"
	"github.com/sparkymat/others/view"
	"github.com/sparkymat/spartan"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/size"
)

type OthersApp struct {
	menu                view.Menu
	eventHandlerChannel chan termbox.Event
	mainLayout          spartan.LinearLayout
	spartanApp          spartan.App
	mode                mode.Mode
	ticker              *time.Ticker
}

func New() *OthersApp {
	othersApp := OthersApp{}

	othersApp.spartanApp = spartan.New()

	othersApp.eventHandlerChannel = make(chan termbox.Event)
	othersApp.mode = mode.Menu

	othersApp.mainLayout = spartan.LinearLayout{}
	othersApp.mainLayout.Direction = direction.Vertical
	othersApp.mainLayout.Width = size.MatchParent
	othersApp.mainLayout.Height = size.MatchParent

	othersApp.menu = view.Menu{}
	othersApp.menu.Width = size.MatchParent
	othersApp.menu.Height = size.MatchParent
	othersApp.mainLayout.AddChild(&othersApp.menu)

	othersApp.spartanApp.SetContent(othersApp.mainLayout)

	othersApp.ticker = time.NewTicker(time.Millisecond * 200)

	return &othersApp
}

func (othersApp *OthersApp) Run() {
	go event.Handler(othersApp.eventHandlerChannel)
	go func() {
		for _ = range othersApp.ticker.C {
			othersApp.OnTick()
		}
	}()

	othersApp.spartanApp.Run(othersApp.eventHandlerChannel)
}

func (othersApp *OthersApp) CleanupForTermination() {
	othersApp.ticker.Stop()
}

func (othersApp *OthersApp) OnTick() {
	othersApp.menu.OnTick()
}
