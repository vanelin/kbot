/*
Copyright © 2023 NAME HERE <ivan.voloboyev@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("kbot %s started", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env varible. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Print(m.Message().Payload, m.Text())
			playload := m.Message().Payload

			switch playload {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello I'm Kbot %s!", appVersion))
			case "ping":
				err = m.Send(fmt.Sprintln("Pong"))
			case "bitcoin":
				marketMsg := getMarkets("bitcoin")
				err = m.Send(marketMsg)
			case "ethereum":
				marketMsg := getMarkets("ethereum")
				err = m.Send(marketMsg)
			}
			return err
		})

		kbot.Start()
	},
}

type MarketResponse struct {
	Data []MarketData `json:"data"`
}

type MarketData struct {
	ExchangeId    string `json:"exchangeId"`
	BaseId        string `json:"baseId"`
	QuoteId       string `json:"quoteId"`
	QuoteSymbol   string `json:"quoteSymbol"`
	VolumeUsd24Hr string `json:"volumeUsd24Hr"`
	PriceUsd      string `json:"priceUsd"`
	VolumePercent string `json:"volumePercent"`
}

func getMarkets(currency string) string {
	fmt.Println("Getting markets for ", currency)
	coincapApiUrl := "https://api.coincap.io/v2/assets/" + currency + "/markets?limit=20"

	client := http.Client{}

	req, err := http.NewRequest("GET", coincapApiUrl, nil)

	if err != nil {
		log.Printf("Error creating request: %v", err)
		return "Error creating request"
	}

	res, err := client.Do(req)

	if err != nil {
		log.Printf("Error sending request: %v", err)
		return "Error sending request"
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return "Error reading response body"
	}

	var data MarketResponse
	json.Unmarshal(respBody, &data)

	var marketsInfo string
	for _, market := range data.Data {
		marketsInfo += fmt.Sprintf("\n\nExchange: %s\nBase: %s\nQuote: %s\nPrice: $%s\nVolume: %s\nVolume Percent: %s\n", market.ExchangeId, market.BaseId, market.QuoteSymbol, market.PriceUsd, market.VolumeUsd24Hr, market.VolumePercent)
	}
	return marketsInfo
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
