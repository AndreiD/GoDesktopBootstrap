package cryptocurrency

import (
	"github.com/buger/jsonparser"
	"godesktop/configs"
	"godesktop/utils"
	"godesktop/utils/log"
	"time"
)

var rates map[string]float64

func init() {
	rates = make(map[string]float64)
}

// SyncConversionRates updates the rates at an interval
func SyncConversionRates(configuration *configs.ViperConfiguration) {

	go func() {
		ticker := time.NewTicker(120 * time.Second)

		for {
			err := getRates(configuration)
			if err != nil {
				log.Error(err)
			}
			<-ticker.C
		}
	}()

}

func getRates(configuration *configs.ViperConfiguration) error {

	data, err := utils.GetRequest("https://api.nomics.com/v1/currencies/ticker?key=" + configuration.Get("nomicsAPIKey") + "&ids=BTC,ETH,BCH,LTC,XRP&interval=1h&convert=USD")
	if err != nil {
		return err
	}

	_, err = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		priceUSD, _, _, _ := jsonparser.Get(value, "price")
		symbol, _, _, _ := jsonparser.Get(value, "symbol")

		rates[string(symbol)] = utils.GetFloat(string(priceUSD))
	})

	if err != nil {
		return err
	}
	return nil
}

// GetRate - returns the rate of a symbol (ex: BTC)
func GetRate(symbol string) float64 {
	if rates == nil {
		log.Error("Rates are nil!")
		return 0
	}
	rate, ok := rates[symbol]
	if !ok {
		log.Errorf("could not find symbol %s in the rates map", symbol)
		return 0
	}
	return rate
}
