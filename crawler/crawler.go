package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/mc388/how-rich-am-i/model"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type Api struct {
	Url      string
	Currency string
}

func NewApi() Api {
	api := Api{}
	api.Url = "https://api.coinmarketcap.com/v1/ticker/"
	api.Currency = "EUR"
	return api
}

func (a *Api) GetCurrencyData(coinId string) model.Coin {
	url := a.Url + coinId + "/?convert=" + a.Currency

	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	coins := make([]model.Coin, 0)

	err = json.Unmarshal(body, &coins)

	return coins[0]
}
