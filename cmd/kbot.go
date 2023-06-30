/*
Copyright © 2023 NAME HERE <ivan.voloboyev@gmail.com>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/hirosassa/zerodriver"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
	// MetricsHost exporter host:port
	MetricsHost = os.Getenv("METRICS_HOST")
)

// Initialize OpenTelemetry
func initMetrics(ctx context.Context) {

	// Create a new OTLP Metric gRPC exporter with the specified endpoint and options
	exporter, _ := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint(MetricsHost),
		otlpmetricgrpc.WithInsecure(),
	)

	// Define the resource with attributes that are common to all metrics.
	// labels/tags/resources that are common to all metrics.
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(fmt.Sprintf("kbot_%s", appVersion)),
	)

	// Create a new MeterProvider with the specified resource and reader
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resource),
		sdkmetric.WithReader(
			// collects and exports metric data every 10 seconds.
			sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(10*time.Second)),
		),
	)

	// Set the global MeterProvider to the newly created MeterProvider
	otel.SetMeterProvider(mp)
}

func pmetrics(ctx context.Context, payload string) {
	// Get the global MeterProvider and create a new Meter with the name "kbot_crypto_counter"
	meter := otel.GetMeterProvider().Meter("kbot_crypto_counter")

	// Get or create an Int64Counter instrument with the name "kbot_crypto_counter_<payload>"
	counter, _ := meter.Int64Counter(fmt.Sprintf("kbot_crypto_counter_%s", payload))

	// Add a value of 1 to the Int64Counter
	counter.Add(ctx, 1)
}

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
		logger := zerodriver.NewProductionLogger()

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			logger.Fatal().Str("Error", err.Error()).Msg("Please check TELE_TOKEN")
			return
		} else {
			logger.Info().Str("Version", appVersion).Msg("kbot started")

		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			logger.Info().Str("Payload", m.Text()).Msg(m.Message().Payload)

			payload := m.Message().Payload
			pmetrics(context.Background(), payload)

			switch payload {
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
	logger := zerodriver.NewProductionLogger()
	logger.Info().Str("Currency", currency).Msg("Getting price for currency")
	coincapApiUrl := "https://api.coincap.io/v2/assets/" + currency + "/markets?limit=10"

	client := http.Client{}

	req, err := http.NewRequest("GET", coincapApiUrl, nil)

	if err != nil {
		logger.Fatal().Str("Error", err.Error()).Msg("Error creating request")
		return "Error creating request"
	}

	res, err := client.Do(req)

	if err != nil {
		logger.Fatal().Str("Error", err.Error()).Msg("Error sending request")
		return "Error sending request"
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)

	if err != nil {
		logger.Fatal().Str("Error", err.Error()).Msg("Error reading response body")
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
	ctx := context.Background()
	initMetrics(ctx)
	rootCmd.AddCommand(kbotCmd)
}
