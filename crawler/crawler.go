package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type Api struct {
	Url      string
	Currency string
}

type Coin struct {
	CurrencyId string  `json:"id"`
	Price      float32 `json:"price_eur,string"`
	LastUpdate int     `json:"last_updated,string"`
}

func NewApi() Api {
	api := Api{}
	api.Url = "https://api.coinmarketcap.com/v1/ticker/"
	api.Currency = "EUR"
	return api
}

func (a *Api) GetCurrentPrice(coinId string) float32 {
	url := a.Url + coinId + "/?convert=" + a.Currency

	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	coins := make([]Coin, 0)

	err = json.Unmarshal(body, &coins)

	return coins[0].Price
}
