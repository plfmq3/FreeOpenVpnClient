package common

import "github.com/gotk3/gotk3/gtk"

func DisableWindowDeinit(window *gtk.Window) error {
	window.Connect("delete-event", func() bool { //Prevent GTK from deinitializing window
		window.Hide()
		return true
	})
	return nil
}
