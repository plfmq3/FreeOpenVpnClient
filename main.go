package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Wifx/gonetworkmanager"
	//"github.com/godbus/dbus/v5"
)

func main() {

	nm, err := gonetworkmanager.NewNetworkManager()
	if err != nil {
		log.Fatal("Couldn't initialize NetworkManager", err.Error())
	}
	var nmVersion string
	nmVersion, err = nm.GetPropertyVersion()
	if err != nil {
		log.Fatal("couldn't get version")
	}

	fmt.Println("Network Manager Version: " + nmVersion)

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

	}
}
