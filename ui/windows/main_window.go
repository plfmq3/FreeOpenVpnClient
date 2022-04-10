package windows

import (
	"fmt"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/plfmq3/FreeOpenVpnClient/config"
	"github.com/plfmq3/FreeOpenVpnClient/ui/items"
)

type MainWindow struct {
	window         *gtk.Window
	settingsButton *items.Button
}

func NewMainWindow(b *gtk.Builder) (*MainWindow, error) {
	w := new(MainWindow)
	cfg := config.GetConfig()
	obj, err := b.GetObject(cfg.MainWindow)
	if err != nil {
		return nil, err
	}
	w.window = obj.(*gtk.Window)

	w.window.Connect("delete-event", func() bool {
		w.window.Hide()
		return true
	})
	return w, nil
}

func (w *MainWindow) Show() {
	glib.IdleAdd(func() {
		w.window.Show()
	})
}

func (w *MainWindow) Hide() {
	glib.IdleAdd(func() {
		w.window.Close()
	})
}
