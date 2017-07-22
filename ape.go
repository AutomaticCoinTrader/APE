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

type tradeConfig struct {
}

type arbitrageTradeConfig struct {
}

type config struct {
	Trade          *tradeConfig          `json:"trade"          yaml:"trade"          toml:"trade"`
	ArbitrageTrade *arbitrageTradeConfig `json:"arbitrageTrade" yaml:"arbitrageTrade" toml:"arbitrageTrade"`
}

type ape struct {
	name           string
	config         *tradeConfig
}

func (a *ape) GetName() (string) {
	return a.name
}

func (a *ape) Initialize(tradeContext exchange.TradeContext, notifier *notifier.Notifier) (error) {
	return nil
}

func (a *ape) Update(tradeContext exchange.TradeContext, notifier *notifier.Notifier) (error) {
	// trade
	log.Printf("trade of ape")
	return nil
}

func (a *ape) Finalize(tradeContext exchange.TradeContext, notifier *notifier.Notifier) (error) {
	return nil
}

func newAPE(configDir string) (algorithm.TradeAlgorithm, error) {
	configFilePathPrefix := path.Join(configDir, algorithmName)
	cf, err := configurator.NewConfigurator(configFilePathPrefix)
	if err != nil {
		return nil, errors.Errorf("can not create configurator (config file path prefix = %v)", configFilePathPrefix)
	}
	conf := new(config)
	err = cf.Load(conf)
	if err != nil {
		return nil, errors.Errorf("can not load config (config file path prefix = %v)", configFilePathPrefix)
	}
	return &ape{
		name:           algorithmName,
		config:         conf.Trade,
	}, nil
}

type arbitrageAPE struct {
	name           string
	config         *arbitrageTradeConfig
}

func (a *arbitrageAPE) GetName() (string) {
	return a.name
}

func (a *arbitrageAPE) Initialize(exchanges map[string]exchange.Exchange, notifier *notifier.Notifier) (error) {
	return nil
}

func (a *arbitrageAPE) Update(exchanges map[string]exchange.Exchange, notifier *notifier.Notifier) (error) {
	// arbitrage trade
	log.Printf("arbitrage trade of ape")
	return nil
}

func (a *arbitrageAPE) Finalize(exchanges map[string]exchange.Exchange, notifier *notifier.Notifier) (error) {
	return nil
}

func newArbitrageAPE(configDir string) (algorithm.ArbitrageTradeAlgorithm, error) {
	configFilePathPrefix := path.Join(configDir, algorithmName)
	cf, err := configurator.NewConfigurator(configFilePathPrefix)
	if err != nil {
		return nil, errors.Errorf("can not create configurator (config file path prefix = %v)", configFilePathPrefix)
	}
	conf := new(config)
	err = cf.Load(conf)
	if err != nil {
		return nil, errors.Errorf("can not load config (config file path prefix = %v)", configFilePathPrefix)
	}
	return &arbitrageAPE{
		name:           algorithmName,
		config:         conf.ArbitrageTrade,
	}, nil
}

func init() {
	fmt.Println("plugin init")
}

func GetRegistrationInfo() (string, algorithm.TradeAlgorithmNewFunc, algorithm.ArbitrageTradeAlgorithmNewFunc) {
	return algorithmName, newAPE, newArbitrageAPE
}
