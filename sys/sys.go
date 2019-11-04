package sys

import (
	"github.com/wailsapp/wails"
	"godesktop/configs"
	"godesktop/cryptocurrency"
	"godesktop/models"
	"godesktop/utils/log"
)

// Pipe .
type Pipe struct {
	log     *wails.CustomLogger
	runtime *wails.Runtime
}

// Runtime .
var Runtime *wails.Runtime

// Conf .
var Conf *configs.ViperConfiguration

// PriceFor .
func (p *Pipe) PriceFor(symbol string) models.Price {

	log.Printf("Got a price request from the frontend -> get price for %s", symbol)
	SendToast("here is a toast...")
	return models.Price{
		Symbol:   symbol,
		PriceUSD: cryptocurrency.GetRate(symbol),
	}
}

// SendToast  ...
func SendToast(message string) {
	Runtime.Events.Emit("toast", message)
}

// WailsInit .
func (p *Pipe) WailsInit(runtime *wails.Runtime) error {
	log.Println("======== wails init called ===========")

	p.log = runtime.Log.New("Pipe")
	Runtime = runtime

	return nil
}

// WailsShutdown .
func (p *Pipe) WailsShutdown() {
	log.Println("wails shutdown initiated...")
}
