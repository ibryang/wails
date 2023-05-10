package main

import (
	_ "embed"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/icons"
	"log"
	"runtime"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var counter int

func clickCount() int {
	counter++
	return counter
}

func main() {
	app := application.New(application.Options{
		Name:        "Systray Demo",
		Description: "A demo of the Systray API",
		Mac: application.MacOptions{
			ActivationPolicy: application.ActivationPolicyAccessory,
		},
	})

	window := app.NewWebviewWindow().Hide()

	systemTray := app.NewSystemTray()
	if runtime.GOOS == "darwin" {
		systemTray.SetTemplateIcon(icons.SystrayMacTemplate)
	}

	myMenu := app.NewMenu()
	myMenu.Add("Hello World!").OnClick(func(ctx *application.Context) {
		println("Hello World!")
		//		app.InfoDialog().SetTitle("Hello World!").SetMessage("Hello World!").Show()
	})
	subMenu := myMenu.AddSubmenu("Submenu")
	subMenu.Add("Click me!").OnClick(func(ctx *application.Context) {
		println("Click me!")
		//		ctx.ClickedMenuItem().SetLabel("Clicked!")
	})
	myMenu.AddSeparator()
	myMenu.Add("Quit").OnClick(func(ctx *application.Context) {
		app.Quit()
	})

	systemTray.SetMenu(myMenu)

	systemTray.OnClick(func() {
		window.SetTitle(fmt.Sprintf("Clicked %d times", clickCount()))
		window.Show().Focus()
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
