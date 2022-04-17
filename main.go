package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/plfmq3/FreeOpenVpnClient/config"
	"github.com/plfmq3/FreeOpenVpnClient/ui"
	//"github.com/godbus/dbus/v5"
)

func main() {
	gtk.Init(nil)
	tray := ui.NewTray()
	_ = tray
	b, err := gtk.BuilderNew()
	checkFatal(err)

	go gtkLoop()

	windowManager, e := ui.NewWindowManager(b)
	checkFatal(e)

	cfg := config.GetConfig()

	e = windowManager.CreateWindows(cfg.MainWindow, cfg.SettingsWindow, cfg.CaptchaWindow)

	checkFatal(e)
	mainWindow, e := windowManager.GetWindow(cfg.MainWindow)
	checkFatal(e)

	e = mainWindow.CreateSubitems(cfg.LoaderContainer, cfg.LoaderText,
		cfg.SettingsButton, cfg.ConnectButton, cfg.VpnList)
	checkFatal(e)

	settingsWindow, _ := windowManager.GetWindow(cfg.SettingsWindow)

	e = settingsWindow.CreateSubitems(cfg.SolverSwitch, cfg.SolverUrl,
		cfg.TcpButton, cfg.UdpButton, cfg.SaveSettingsButton)
	checkFatal(e)

	captchaWindow, _ := windowManager.GetWindow(cfg.CaptchaWindow)

	e = captchaWindow.CreateSubitems(cfg.LoaderContainer, cfg.LoaderText,
		cfg.SettingsButton, cfg.ConnectButton)
	checkFatal(e)

	e = windowManager.InitAllWindows()
	checkFatal(e)

	exitChan, showMainWinChan := tray.GetChannels()

	for {
		select {
		case <-*exitChan:
			fmt.Println("exit")
			return
		case <-*showMainWinChan:
			mainWindow.Show()
		}
	}

	/* nm, err := gonetworkmanager.NewNetworkManager()
		if err != nil {
			log.Fatal("Couldn't initialize NetworkManager", err.Error())
		}
		var nmVersion string
		nmVersion, err = nm.GetPropertyVersion()
		if err != nil {
			log.Fatal("couldn't get version")
		}

		fmt.Println("Network Manager Version: " + nmVersion)
	5
		c, e := gonetworkmanager.NewVpnConnection("Netherlands_freeopenvpn_tcp")
		if e != nil {
			log.Fatal(e)
		}

		fmt.Println(c.GetPropertyVpnState())
		for {
			e := <-nm.Subscribe()
			if strings.Contains(strings.ToLower(e.Name), "vpn") {
				fmt.Println(e.Name, e.Sender, e.Path, e.Body)
			}

		} */
}

func gtkLoop() {
	gtk.Main()
}

func checkFatal(e error) {
	if e != nil {
		fmt.Println(e)
		log.Fatal(e)
	}
}
