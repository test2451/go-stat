package common

import "time"

const (
	ObserverMaxBlockNumber = 10000
	ObservceMaxTxNumber =  100000
	ObserverPruneInterval  = 30 * time.Second
	ObserverAlertInterval  = 100 * time.Second

	RefreshInterval = 300 * time.Second
)

const (
	DBDialectMysql   = "mysql"
	DBDialectSqlite3 = "sqlite3"
)

type BlockAndEventLogs struct {
	Height          int64
	BlockHash       string
	ParentBlockHash string
	BlockTime       int64
	Events          []interface{}
}
