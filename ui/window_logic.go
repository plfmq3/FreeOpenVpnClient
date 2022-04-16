package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/plfmq3/FreeOpenVpnClient/config"
	. "github.com/plfmq3/FreeOpenVpnClient/ui/windows"
)

func InitUiLogic(w *Window) error {
	cfg := config.GetConfig()
	win_id := *w.GetId()
	switch win_id {
	case cfg.MainWindow:
		return initMainWindow(w)
	case cfg.SettingsWindow:
		return initSettingsWindow(w)
	case cfg.CaptchaWindow:
		return initCaptchaWindow(w)
	default:
		fmt.Printf("WARN: %s has no logic block\r\n", win_id)
	}
	return nil
}

func InitCrossLogic(m map[string]Window) error {
	cfg := config.GetConfig()
	mainWindow := m[cfg.MainWindow]
	settingsWindow := m[cfg.SettingsWindow]
	btn, e := mainWindow.GetSubitem(cfg.SettingsButton)
	if e != nil {
		return e
	}
	settingsButton := btn.(*gtk.Button)
	settingsButton.Connect("clicked", func() {
		settingsWindow.Show()
	})
	return nil
}

func initMainWindow(w *Window) error {
	fmt.Println("init main window hit!")
	cfg := config.GetConfig()
	obj, err := w.GetSubitem(cfg.LoaderContainer)
	if err != nil {
		return err
	}
	loaderContainer := obj.(*gtk.Grid)
	loaderContainer.Hide()
	return nil
}

func initSettingsWindow(w *Window) error {
	fmt.Println("init settings window hit!")
	cfg := config.GetConfig()
	obj, err := w.GetSubitem(cfg.SaveSettingsButton)
	if err != nil {
		return err
	}
	saveButton := obj.(*gtk.Button)
	saveButton.Connect("clicked", func() {
		w.GetGTKWindow().Close()
	})
	return nil
}

func initCaptchaWindow(w *Window) error {
	//cfg := config.GetConfig()
	fmt.Println("init captcha window hit!")
	return nil
}
