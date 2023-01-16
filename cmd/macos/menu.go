package main

import (
	"fmt"
	"regexp"

	"github.com/getlantern/systray"
	autootp "github.com/mniak/auto-otp"
)

func showMenu(sendKeysChan chan<- string, menuEntries []autootp.MenuEntry) {
	systray.Run(
		func() { // tray init
			initMenu(sendKeysChan, menuEntries)
		},
		func() { // tray exit
		},
	)
}

func initMenu(sendKeysChan chan<- string, entries []autootp.MenuEntry) {
	systray.SetTitle("Auto OTP")
	subtitle := systray.AddMenuItem("Click to type OTP", "")
	subtitle.Disable()
	for _, entry := range entries {
		menuItem := systray.AddMenuItem(fmt.Sprintf("%s - Loading...", entry.Title), "")
		menuItem.Disable()
		go func(mi *systray.MenuItem, e autootp.MenuEntry) {
			code := e.Code()
			mi.SetTitle(fmt.Sprintf("%s - %s", e.Title, formatCode(code)))
			mi.Enable()
			for {
				<-mi.ClickedCh
				sendKeysChan <- code
			}
		}(menuItem, entry)
	}
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

var re = regexp.MustCompile(`(\d{3})(\d)`)

func formatCode(code string) string {
	return re.ReplaceAllString(code, "$1-$2")
}
