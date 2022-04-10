package ui

import (
	"fmt"
	"io/ioutil"

	"github.com/getlantern/systray"
	"github.com/plfmq3/FreeOpenVpnClient/config"
)

type Tray struct {
	exitChannel    *chan struct{}
	showWinChannel *chan struct{}
}

func NewTray() *Tray {
	t := new(Tray)
	go systray.Run(t.onReady, t.onExit)
	return t
}

func (t *Tray) onReady() {
	cfg := config.GetConfig()
	systray.SetIcon(getIcon(cfg.DataFolder + cfg.Icon))

	showWin := systray.AddMenuItem("Show main window", "Shows the vpn list")
	t.showWinChannel = &showWin.ClickedCh

	quit := systray.AddMenuItem("Quit", "Quit the whole app")
	t.exitChannel = &quit.ClickedCh

}

func (t *Tray) GetChannels() (exitChannel *chan struct{}, showWinChannel *chan struct{}) {
	return t.exitChannel, t.showWinChannel
}

func (t *Tray) onExit() {
	return
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
