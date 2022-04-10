package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/plfmq3/FreeOpenVpnClient/config"

	. "github.com/plfmq3/FreeOpenVpnClient/ui/windows"
)

type WindowManager struct {
	builder        *gtk.Builder
	mainWindow     *MainWindow
	settingsWindow *SettingsWindow
	captchaWindow  *CaptchaWindow
}

type Windows struct {
	*MainWindow
	*SettingsWindow
	*CaptchaWindow
}

func NewWindowManager(b *gtk.Builder) (*WindowManager, error) {
	cfg := config.GetConfig()
	m := new(WindowManager)
	m.builder = b

	err := m.builder.AddFromFile(cfg.DataFolder + cfg.UiFile)
	if err != nil {
		return nil, err
	}

	return m, nil

}

func (wm *WindowManager) InitWindows() error {

	mainWindow, e := NewMainWindow(wm.builder)
	if e != nil {
		return e
	}
	wm.mainWindow = mainWindow
	return nil
}

func (wm *WindowManager) GetWindows() Windows {
	return Windows{wm.mainWindow, wm.settingsWindow, wm.captchaWindow}
}
