package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/mc388/how-rich-am-i/crawler"
	"github.com/mc388/how-rich-am-i/env"
)

func main() {
	coinsConfigFile := flag.String("file", "coins.json", "path to your coins.json file")
	flag.Parse()

	myCoins, err := env.LoadConfig(*coinsConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	printCoins(myCoins)
}

func printCoins(myCoins []env.MyCoins) {
	api := crawler.NewApi()

	var totalUSD float32
	totalUSD = 0.0

	fmt.Printf("|%-15s|%-6s|%15s|%15s|%15s|%15s|%15s||%15s|%15s|\n",
		"Name",
		"Symbol",
		"Price in $",
		"Price in BTC",
		"% Change (1h)",
		"% Change (24h)",
		"% Change (7d)",
		"My coin amount",
		"My amount in $",
	)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------")
	for _, myCoin := range myCoins {
		coin := api.GetCurrencyData(myCoin.CurrencyId)

		percentChange1h := getColoredOutput(coin.PercentChange1h)
		percentChange24h := getColoredOutput(coin.PercentChange24h)
		percentChange7d := getColoredOutput(coin.PercentChange7d)

		fmt.Printf("|%-15s|%-6s|%15.8f|%15.8f|%15s|%15s|%15s||%15.8f|%15.2f|\n",
			coin.Name,
			coin.Symbol,
			coin.PriceUSD,
			coin.PriceBTC,
			percentChange1h,
			percentChange24h,
			percentChange7d,
			myCoin.Amount,
			myCoin.Amount*coin.PriceUSD)
		totalUSD = totalUSD + myCoin.Amount*coin.PriceUSD
	}

	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("\nI have €%.2f in crypto currency!\n", totalUSD*0.8)
}

func getColoredOutput(percentChange float32) string {
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	coloredOutput := fmt.Sprintf("%15.2f%%", percentChange)

	coloredOutput = green(percentChange)
	if percentChange == 0.0 {
		coloredOutput = yellow(percentChange)
	} else if percentChange < 0.0 {
		coloredOutput = red(percentChange)
	}

	return fmt.Sprintf("%24s", coloredOutput)
}
