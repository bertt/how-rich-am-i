package main

import (
	"fmt"
	"github.com/mc388/how-rich-am-i/crawler"
	"github.com/mc388/how-rich-am-i/env"
	"log"
	"os"
)

func main() {
	if (len(os.Args) != 2) {
		log.Fatal("Wrong usage: how-rich-am-i <path/to/coins.json>")
	}

	coinsConfigFile := os.Args[1]

	myCoins, err := env.LoadConfig(coinsConfigFile)

	if err != nil {
		log.Fatal(err)
	}

	api := crawler.NewApi()

	var total float32
	total = 0.0
	fmt.Printf("|%-12s|%15s|%15s|%15s|%15s|\n",
		"Name",
		"Price in €",
		"% Change (7d)",
		"My coin amount",
		"My amount in €",
	)
	fmt.Println("------------------------------------------------------------------------------")
	for _, myCoin := range myCoins {
		coin := api.GetCurrencyData(myCoin.CurrencyId)
		fmt.Printf("|%-12s|%15.8f|%15.2f|%15.8f|%15.2f|\n",
			coin.Name,
			coin.PriceEUR,
			coin.PercentChange7d,
			myCoin.Amount,
			myCoin.Amount*coin.PriceEUR)
		total = total + myCoin.Amount*coin.PriceEUR
	}

	fmt.Println("------------------------------------------------------------------------------")
	fmt.Printf("\nI have %.2f € in crypto currency\n", total)
}
