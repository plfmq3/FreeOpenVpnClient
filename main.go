package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/plfmq3/FreeOpenVpnClient/ui"
	//"github.com/godbus/dbus/v5"
)

func main() {
	gtk.Init(nil)

	tray := ui.NewTray()
	_ = tray
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	go gtkLoop()

	windowManager, e := ui.NewWindowManager(b)
	if e != nil {
		log.Fatal(e)
	}
	_ = windowManager

	windowManager.InitWindows()

	exitChan, showMainWinChan := tray.GetChannels()

	for {
		select {
		case <-*exitChan:
			fmt.Println("exit")
			return
		case <-*showMainWinChan:
			windowManager.GetWindows().MainWindow.Show()
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
