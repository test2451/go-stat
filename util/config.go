package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pancakeswap/pancake-statas/common"
)

type Config struct {
	StatasDBConfig *DBConfig    `json:"statas_db_config"`
	ChainConfig    *ChainConfig `json:"chain_config"`
	LogConfig      *LogConfig   `json:"log_config"`
	AlertConfig    *AlertConfig `json:"alert_config"`
	ServerConfig   ServerConfig `json:"server_config"`
}

func (cfg *Config) Validate() {
	cfg.StatasDBConfig.Validate()
	cfg.ChainConfig.Validate()
	cfg.LogConfig.Validate()
	cfg.AlertConfig.Validate()
}

type AlertConfig struct {
	TelegramBotId  string `json:"telegram_bot_id"`
	TelegramChatId string `json:"telegram_chat_id"`

	BlockUpdateTimeout int64 `json:"block_update_timeout"`
}

func (cfg *AlertConfig) Validate() {
	if cfg.BlockUpdateTimeout <= 0 {
		panic(fmt.Sprintf("block_update_timeout should be larger than 0"))
	}
}

type DBConfig struct {
	Dialect string `json:"dialect"`
	DBPath  string `json:"db_path"`
}

func (cfg *DBConfig) Validate() {
	if cfg.Dialect != common.DBDialectMysql && cfg.Dialect != common.DBDialectSqlite3 {
		panic(fmt.Sprintf("only %s and %s supported", common.DBDialectMysql, common.DBDialectSqlite3))
	}
	if cfg.DBPath == "" {
		panic("db path should not be empty")
	}
}

type ChainConfig struct {
	BSCStartHeight    int64    `json:"bsc_start_height"`
	BSCProvider       string   `json:"bsc_provider"`
	BSCConfirmNum     int64    `json:"bsc_confirm_num"`
	BSCFetchInterval  int64    `json:"bsc_fetch_interval"`
	SwapFactory       string   `json:"swap_factory"`
	CertificatedPairs []string `json:"certificated_pairs"`
	SynupPools        []string `json:"synup_pools"`
}

func (cfg *ChainConfig) Validate() {
	if cfg.BSCStartHeight < 0 {
		panic("bsc_start_height should not be less than 0")
	}
	if cfg.BSCProvider == "" {
		panic("bsc_provider should not be empty")
	}
	if cfg.BSCConfirmNum <= 0 {
		panic("bsc_confirm_num should be larger than 0")
	}
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

func (cfg *LogConfig) Validate() {
	if cfg.UseFileLogger {
		if cfg.Filename == "" {
			panic("filename should not be empty if use file logger")
		}
		if cfg.MaxFileSizeInMB <= 0 {
			panic("max_file_size_in_mb should be larger than 0 if use file logger")
		}
		if cfg.MaxBackupsOfLogFiles <= 0 {
			panic("max_backups_off_log_files should be larger than 0 if use file logger")
		}
	}
}

type ServerConfig struct {
	ListenAddr string `json:"listen_addr"`
}

func ParseConfigFromFile(filePath string) *Config {
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}
	return &config
}

func ParseConfigFromJson(content string, config *Config) *Config {
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		panic(err)
	}
	return config
}
