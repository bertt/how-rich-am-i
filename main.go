package main

import (
	"fmt"
	"github.com/mc388/how-rich-am-i/crawler"
	"github.com/mc388/how-rich-am-i/env"
	"log"
	"os"
)

func main() {
	if(len(os.Args) != 2) {
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
	for _, myCoin := range myCoins {
		priceInEur := api.GetCurrentPrice(myCoin.CurrencyId)
		fmt.Printf("|%-12s|%12.8f|%12.2f|\n", myCoin.CurrencyId, myCoin.Amount, myCoin.Amount*priceInEur)
		total = total + myCoin.Amount*priceInEur
	}

	fmt.Println("----------------------------------------")
	fmt.Printf("I have %.2f € in crypto currency\n", total)
}
