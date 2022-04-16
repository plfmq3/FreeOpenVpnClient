package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/plfmq3/FreeOpenVpnClient/config"

	. "github.com/plfmq3/FreeOpenVpnClient/ui/windows"
)

type WindowManager struct {
	builder *gtk.Builder
	windows map[string]Window
}

func NewWindowManager(b *gtk.Builder) (*WindowManager, error) {
	cfg := config.GetConfig()
	m := new(WindowManager)
	m.windows = make(map[string]Window)
	m.builder = b
	err := m.builder.AddFromFile(cfg.DataFolder + cfg.UiFile)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (wm *WindowManager) CreateWindows(ids ...string) error {
	for _, item := range ids {
		window, err := NewWindow(wm.builder, item)
		if err != nil {
			return err
		}
		wm.windows[item] = *window
	}
	return nil
}

func (wm *WindowManager) InitAllWindows() error {
	err := InitCrossLogic(wm.windows)
	if err != nil {
		return err
	}
	for _, item := range wm.windows {
		err := InitUiLogic(&item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (wm *WindowManager) GetWindow(id string) (*Window, error) {
	window, exists := wm.windows[id]
	if !exists {
		return nil, fmt.Errorf("window \"%s\" doesn't exist", id)
	}
	return &window, nil
}
