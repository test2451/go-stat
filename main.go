package main

import (
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/pancakeswap/pancake-statas/executor"
	"github.com/pancakeswap/pancake-statas/model"
	"github.com/pancakeswap/pancake-statas/observer"
	"github.com/pancakeswap/pancake-statas/server"
	"github.com/pancakeswap/pancake-statas/statas"
	"github.com/pancakeswap/pancake-statas/util"
)

const (
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
	flagConfigPath         = "config-path"
)

func initFlags() {
	flag.String(flagConfigPath, "", "config path")
	flag.String(flagConfigAwsRegion, "", "aws s3 region")
	flag.String(flagConfigAwsSecretKey, "", "aws s3 secret key")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("bind flags error, err=%s", err))
	}
}

func main() {
	initFlags()

	var config *util.Config
	configFilePath := viper.GetString(flagConfigPath)
	awsSecretKey := viper.GetString(flagConfigAwsSecretKey)
	awsRegion := viper.GetString(flagConfigAwsRegion)

	if configFilePath == "" {
		panic("config-path can't be empty")
	}

	config = util.ParseConfigFromFile(configFilePath)

	if awsSecretKey != "" && awsRegion != "" {

		configContent, err := util.GetSecret(awsSecretKey, awsRegion)
		if err != nil {
			fmt.Printf("get aws config error, err=%s", err.Error())
			return
		}
		config = util.ParseConfigFromJson(configContent, config)
	}

	config.Validate()

	// init logger
	util.InitLogger(*config.LogConfig)
	util.InitTgAlerter(config.AlertConfig)

	reconDb, err := gorm.Open(config.StatasDBConfig.Dialect, config.StatasDBConfig.DBPath)
	if err != nil {
		panic(fmt.Sprintf("open recon db error, err=%s", err.Error()))
	}
	defer reconDb.Close()

	reconDb.AutoMigrate(&model.TxEventLog{}, &model.BlockLog{})

	bscExecutor := executor.NewExecutor(config.ChainConfig.BSCProvider, config.ChainConfig.SwapFactory)

	bscObserver := observer.NewObserver(reconDb, config, bscExecutor)

	reconSvc := statas.NewStatasSvc(reconDb, config, bscExecutor)
	reconSvc.Start()
	bscExecutor.SetInfoQuery(reconSvc)
	bscExecutor.Start()
	bscObserver.Start()

	server := server.NewServer(config, reconSvc)
	go server.Serve()
	select {}
}

