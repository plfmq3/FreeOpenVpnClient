package config

import "sync"

type Config struct {
	DataFolder string
	Icon       string
	UiFile     string
	Selectors
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		selectors := new(Selectors)
		selectors.MainWindow = "mainWindow"
		selectors.LoaderContainer = "loaderContainer"
		selectors.LoaderText = "loaderText"
		selectors.SettingsButton = "settingsButton"
		selectors.VpnList = "vpnServerList"
		selectors.ConnectButton = "connectButton"

		selectors.SettingsWindow = "settings"
		selectors.SolverSwitch = "solverSwitch"
		selectors.SolverUrl = "solverUrl"
		selectors.TcpButton = "tcpButton"
		selectors.UdpButton = "udpButton"
		selectors.SaveSettingsButton = "saveButton"

		selectors.CaptchaWindow = "captcha"
		selectors.CaptchaHolder = "captchaHolder"
		selectors.CaptchaConnectButton = "connectBtn"
		selectors.CaptchaCancelButton = "cancelBtn"

		cfg = new(Config)
		cfg.Selectors = *selectors
		cfg.DataFolder = "data/"
		cfg.Icon = "icon.ico"
		cfg.UiFile = "ui.glade"
	})
	return cfg
}

type Selectors struct {
	MainWindow      string
	LoaderContainer string
	LoaderText      string
	SettingsButton  string
	VpnList         string
	ConnectButton   string

	SettingsWindow     string
	SolverSwitch       string
	SolverUrl          string
	TcpButton          string
	UdpButton          string
	SaveSettingsButton string

	CaptchaWindow        string
	CaptchaHolder        string
	CaptchaEntry         string
	CaptchaConnectButton string
	CaptchaCancelButton  string
}
