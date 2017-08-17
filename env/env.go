package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type MyCoins struct {
	CurrencyId string  `json:"id"`
	Amount     float32 `json:"amount"`
}

func LoadConfig(configFile string) ([]MyCoins, error) {
	mycoins := make([]MyCoins, 0)

	jsonData, err := readFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonData, &mycoins)

	return mycoins, err
}

func readFile(configFile string) ([]byte, error) {
	// Read config file
	jsonBytes, err := ioutil.ReadFile(configFile)

	return jsonBytes, err
}
