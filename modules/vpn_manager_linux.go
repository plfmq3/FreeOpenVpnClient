package modules

import (
	"log"

	"github.com/Wifx/gonetworkmanager"
	"github.com/godbus/dbus/v5"
	//"github.com/godbus/dbus/v5"
)

type VpnManagerLinux struct {
	networkManager *gonetworkmanager.NetworkManager
}

type update struct {
	Type string
	Code int
}

type VpnState struct {
	body string
}

func NewVpnStateFromSignal(signal *dbus.Signal) *VpnState {
	s := new(VpnState)
	return s
}

func NewVpnManagerLinux() *VpnManagerLinux {
	m := new(VpnManagerLinux)
	nm, err := gonetworkmanager.NewNetworkManager()
	if err != nil {
		log.Fatal("Couldn't initialize NetworkManager", err.Error())
	}
	m.networkManager = &nm

	return m
}

func (m *VpnManagerLinux) Init() {

}

func Subscribe() {

}

func UnSubscribe() {

}
