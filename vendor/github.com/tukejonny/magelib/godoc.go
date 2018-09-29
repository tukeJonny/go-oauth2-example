package magelib

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/magefile/mage/sh"
)

// GoDoc settings
const (
	goDocHost = "localhost"
	goDocPort = 6060
)

var (
	goDocURL = fmt.Sprintf(`http://%s:%d/pkg/%s/?m=all`, goDocHost, goDocPort, ProjectPkg)
)

func openBrowser() {
	switch runtime.GOOS {
	case "linux":
		sh.RunV("xdg-open", goDocURL)
	case "windows":
		sh.RunV("rundll32", "url.dll,FileProtocolHandler", goDocURL)
	case "darwin":
		sh.RunV("open", goDocURL)
	default:
		panic(errors.New("We can't open the browser!"))
	}
}

// Godoc ... Run godoc server.
func Godoc() error {
	go func() {
		time.Sleep(5 * time.Second)
		openBrowser()
	}()
	return sh.RunV("godoc", "-v", "-http=:6060")
}
