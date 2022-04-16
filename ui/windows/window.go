package windows

import (
	"fmt"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/plfmq3/FreeOpenVpnClient/common"
)

type Window struct {
	window   *gtk.Window
	builder  *gtk.Builder
	id       *string
	subitems map[string]glib.IObject
}

func NewWindow(b *gtk.Builder, windowId string) (*Window, error) {
	w := new(Window)
	w.builder = b
	w.subitems = make(map[string]glib.IObject)
	obj, err := b.GetObject(windowId)
	w.id = &windowId
	if err != nil {
		return nil, err
	}
	w.window = obj.(*gtk.Window)
	common.DisableWindowDeinit(w.window)
	return w, nil
}

func (w *Window) Show() {
	glib.IdleAdd(func() {
		w.window.Show()
	})
}

func (w *Window) Hide() {
	glib.IdleAdd(func() {
		w.window.Close()
	})
}

func (w *Window) GetId() *string {
	return w.id
}

func (w *Window) GetGTKWindow() *gtk.Window {
	return w.window
}

func (w *Window) CreateSubitems(widgetIds ...string) error {
	for _, widgetId := range widgetIds {
		obj, err := w.builder.GetObject(widgetId)
		if err != nil {
			return err
		}
		w.subitems[widgetId] = obj
	}
	return nil
}

func (w *Window) GetSubitem(widgetId string) (glib.IObject, error) {
	widget, exists := w.subitems[widgetId]
	if !exists {
		return nil, fmt.Errorf("subitem \"%s\" doesn't exist", widgetId)
	}
	return widget, nil
}
