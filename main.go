package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"godesktop/configs"
	"godesktop/cryptocurrency"
	"godesktop/sys"
	"godesktop/utils/log"
)

const version = "0.0.3 Gas"

var configuration *configs.ViperConfiguration

func init() {

	configuration = configs.NewConfiguration()
	configuration.Init()

	debug := configuration.GetBool("debug")
	log.Init(debug)
	configuration.Set("isPaused", false)

	log.Println("----------------------------------------------")
	log.Warn("ArbitrageX Starting.... version: " + version)
	log.Println("----------------------------------------------")

}

func main() {
	log.Println("ArbitrageX " + version)
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	pipe := &sys.Pipe{}
	sys.Conf = configuration

	cryptocurrency.SyncConversionRates(configuration)

	// run the application
	app := wails.CreateApp(&wails.AppConfig{
		Width:  1600,
		Height: 960,
		Title:  "GoDesktop v1.0",
		JS:     js,
		CSS:    css,
		Colour: "#333333",
	})
	app.Bind(pipe)
	app.Run()
}
