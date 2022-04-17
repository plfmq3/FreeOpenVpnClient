package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"

	"github.com/plfmq3/FreeOpenVpnClient/config"
	. "github.com/plfmq3/FreeOpenVpnClient/ui/windows"
	"github.com/plfmq3/FreeOpenVpnClient/web"
)

func InitUiLogic(w Window) error {
	cfg := config.GetConfig()
	fmt.Println(w.GetId())
	switch w.GetId() {
	case cfg.MainWindow:
		return initMainWindow(w)
	case cfg.SettingsWindow:
		return initSettingsWindow(w)
	case cfg.CaptchaWindow:
		return initCaptchaWindow(w)
	default:
		fmt.Printf("WARN: %s has no logic block\r\n", w.GetId())
	}
	return nil
}

func InitCrossLogic(m map[string]Window) error {
	cfg := config.GetConfig()
	mainWindow := m[cfg.MainWindow]
	settingsWindow := m[cfg.SettingsWindow]
	captchaWindow := m[cfg.CaptchaWindow]
	btn, e := mainWindow.GetSubitem(cfg.SettingsButton)
	if e != nil {
		return e
	}
	settingsButton := btn.(*gtk.Button)
	settingsButton.Connect("clicked", func() {
		settingsWindow.Show()
	})

	btn, e = mainWindow.GetSubitem(cfg.ConnectButton)
	if e != nil {
		return e
	}
	connectButton := btn.(*gtk.Button)
	connectButton.Connect("clicked", func() {
		captchaWindow.Show()
	})

	return nil
}

func initMainWindow(w Window) error {
	fmt.Println("init main window hit!")
	cfg := config.GetConfig()
	obj, err := w.GetSubitem(cfg.LoaderContainer)
	if err != nil {
		return err
	}
	loaderContainer := obj.(*gtk.Grid)
	loaderContainer.Hide()

	obj, err = w.GetSubitem(cfg.VpnList)
	if err != nil {
		return err
	}
	list := obj.(*gtk.ListBox)
	web := web.NewWeb()
	installedWidgets := []*gtk.ListBoxRow{}

	w.GetGTKWindow().Connect("show", func() {
		fmt.Println("show win")
		web.PullData()

		data := web.GetData()
		for _, widget := range installedWidgets { //remove all widgets inside vpn list
			list.Remove(widget)
		}
		installedWidgets = []*gtk.ListBoxRow{}
		for _, vpn := range *data {

			listboxrow, _ := gtk.ListBoxRowNew()
			box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
			listboxrow.Add(box)
			list.Add(listboxrow)
			installedWidgets = append(installedWidgets, listboxrow)
			name, _ := gtk.LabelNew(vpn.Name)
			status, _ := gtk.LabelNew(vpn.Status)
			image, _ := gtk.ImageNew()
			image.SetFromPixbuf(vpn.Image)

			box.PackStart(image, false, false, 10)
			box.PackStart(name, false, false, 10)
			box.PackEnd(status, false, false, 10)

		}
		list.ShowAll()
	})

	return nil
}

func initSettingsWindow(w Window) error {
	fmt.Println("init settings window hit!")
	fmt.Println(w.GetId())
	cfg := config.GetConfig()
	obj, err := w.GetSubitem(cfg.SaveSettingsButton)
	if err != nil {
		return err
	}
	saveButton := obj.(*gtk.Button)
	saveButton.Connect("clicked", func() {
		fmt.Println("hi", w.GetId())

		w.GetGTKWindow().Close()
	})
	return nil
}

func initCaptchaWindow(w Window) error {
	//cfg := config.GetConfig()
	fmt.Println("init captcha window hit!")
	return nil
}
