package web

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gotk3/gotk3/gdk"
)

type Web struct {
	VpnData *[]VpnEntry
	parser  *Parser
}

func NewWeb() *Web {
	w := new(Web)
	w.VpnData = new([]VpnEntry)
	w.parser = NewParser()
	return w
}

var (
	BaseUrl = "https://www.freeopenvpn.org/"
)

func (w *Web) GetData() *[]VpnEntry {
	return w.VpnData
}

func (w *Web) PullData() error {
	cli := http.Client{}
	resp, e := cli.Get(BaseUrl)
	if e != nil {
		return e
	}

	r, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}
	w.parser.ParseHTML(string(r))
	fmt.Println()
	w.VpnData = w.parser.Extract()

	return nil
}
func DownloadFile(url string) (*[]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, e := ioutil.ReadAll(resp.Body)

	return &b, e
}

type VpnEntry struct {
	Image  *gdk.Pixbuf
	Name   string
	Status string
	Url    string
}
