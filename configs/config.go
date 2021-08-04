package configs

import (
	"strings"

	"github.com/spf13/viper"
)

// type CoinconvApiOptions struct {
// 	URL     string
// 	Timeout time.Duration
// }

// Options struct for store application options
type CoinconvApiOptions struct {
	CoinMarketKey string
	URL           string
}

func NewViper(configName string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(configName)
	v.AddConfigPath(".")

	setDefaults()

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

func NewMainConfig(v *viper.Viper) CoinconvApiOptions {
	return CoinconvApiOptions{
		CoinMarketKey: v.GetString("apiKey.key"),
		URL:           v.GetString("coinmarket.url"),
	}
}

func setDefaults() {
	viper.SetDefault("apiKey.key", "aac8340f-4ae2-4f14-80b0-f11fe74fde51")
	viper.SetDefault("coinmarket.url", "https://pro-api.coinmarketcap.com")
}
