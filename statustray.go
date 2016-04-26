package main

import (
	"fmt"
	"github.com/cvtienhoven/tray_plugin"
	"github.com/cvtienhoven/tray_sensu_plugin"
	"github.com/koweblomke/statustray/trayicons"
	"github.com/koweblomke/trayhost"
	"runtime"
	"time"
)

// Refer to documentation at http://github.com/cratonica/trayhost for generating this
func main() {
	// EnterLoop must be called on the OS's main thread
	runtime.LockOSThread()

	go func() {
		// Run your application/server code in here. Most likely you will
		// want to start an HTTP server that the user can hit with a browser
		// by clicking the tray icon.

		// Be sure to call this to link the tray icon to the target url
		url := "http://sensu.dev.intra.tkppensioen.nl"
		trayhost.SetUrl(url)
		for {
			config := tray_plugin.Config{{"url", "http://sensu.dev.intra.tkppensioen.nl:4567/results"}}
			trayhost.SetIcon(tray_sensu_plugin.GetStatus(config))
			time.Sleep(30 * time.Second)
		}
	}()

	icons := [][]byte{trayicons.OkIcon, trayicons.WarningIcon, trayicons.AlertIcon}
	// Enter the host system's event loop
	trayhost.EnterLoop("Status Tray", icons)

	// This is only reached once the user chooses the Exit menu item
	fmt.Println("Exiting")
}
