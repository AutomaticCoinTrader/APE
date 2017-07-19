package main

import (
	"github.com/AutomaticCoinTrader/ACT/algorithm"
	"github.com/AutomaticCoinTrader/ACT/exchange"
	"github.com/AutomaticCoinTrader/ACT/notifier"
	"github.com/AutomaticCoinTrader/ACT/configurator"
	"github.com/pkg/errors"
	"log"
	"path"
	"fmt"
)

const (
	algorithmName string = "ape"
)

type TradeConfig struct {
}

type ArbitrageTradeConfig struct {
}

type Config struct {
	Trade          *TradeConfig          `json:"trade"          yaml:"trade"          toml:"trade"`
	ArbitrageTrade *ArbitrageTradeConfig `json:"arbitrageTrade" yaml:"arbitrageTrade" toml:"arbitrageTrade"`
}

type APE struct {
	name           string
	config         *TradeConfig
}

func (a *APE) GetName() (string) {
	return a.name
}

func (a *APE) Initialize(tradeContext exchange.TradeContext, notifier *notifier.Notifier) (error) {
	return nil
}

func (a *APE) Update(tradeContext exchange.TradeContext, notifier *notifier.Notifier) (error) {
	// trade
	log.Printf("trade of ape")
	return nil
}

func (a *APE) Finalize(tradeContext exchange.TradeContext, notifier *notifier.Notifier) (error) {
	return nil
}

func newAPE(configDir string) (algorithm.TradeAlgorithm, error) {
	configFilePathPrefix := path.Join(configDir, algorithmName)
	cf, err := configurator.NewConfigurator(configFilePathPrefix)
	if err != nil {
		return nil, errors.Errorf("can not create configurator (config file path prefix = %v)", configFilePathPrefix)
	}
	config := new(Config)
	err = cf.Load(config)
	if err != nil {
		return nil, errors.Errorf("can not load config (config file path prefix = %v)", configFilePathPrefix)
	}
	return &APE{
		name:           algorithmName,
		config:         config.Trade,
	}, nil
}

type ArbitrageAPE struct {
	name           string
	config         *ArbitrageTradeConfig
}

func (a *ArbitrageAPE) GetName() (string) {
	return a.name
}

func (a *ArbitrageAPE) Initialize(exchanges map[string]exchange.Exchange, notifier *notifier.Notifier) (error) {
	return nil
}

func (a *ArbitrageAPE) Update(exchanges map[string]exchange.Exchange, notifier *notifier.Notifier) (error) {
	// arbitrage trade
	log.Printf("arbitrage trade of ape")
	return nil
}

func (a *ArbitrageAPE) Finalize(exchanges map[string]exchange.Exchange, notifier *notifier.Notifier) (error) {
	return nil
}

func newArbitrageAPE(configDir string) (algorithm.ArbitrageTradeAlgorithm, error) {
	configFilePathPrefix := path.Join(configDir, algorithmName)
	cf, err := configurator.NewConfigurator(configFilePathPrefix)
	if err != nil {
		return nil, errors.Errorf("can not create configurator (config file path prefix = %v)", configFilePathPrefix)
	}
	config := new(Config)
	err = cf.Load(config)
	if err != nil {
		return nil, errors.Errorf("can not load config (config file path prefix = %v)", configFilePathPrefix)
	}
	return &ArbitrageAPE{
		name:           algorithmName,
		config:         config.ArbitrageTrade,
	}, nil
}

func init() {
	fmt.Println("plugin init")
}

func GetRegistrationInfo() (string, algorithm.TradeAlgorithmNewFunc, algorithm.ArbitrageTradeAlgorithmNewFunc) {
	return algorithmName, newAPE, newArbitrageAPE
}


