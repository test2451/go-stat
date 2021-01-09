package model

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type BlockLog struct {
	Id         int64
	Chain      string
	BlockHash  string `gorm:"not null;index:block_hash"`
	ParentHash string `gorm:"not null;index:block_parent_hash"`
	Height     int64  `gorm:"not null;index:block_height"`
	BlockTime  int64
	CreateTime int64
}

func (BlockLog) TableName() string {
	return "block_log"
}

func (l *BlockLog) BeforeCreate() (err error) {
	l.CreateTime = time.Now().Unix()
	return nil
}

type TxStatus int

const (
	TxStatusInit      TxStatus = 0
	TxStatusConfirmed TxStatus = 1
)

type TxEventLog struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`

	ContractAddress string  `gorm:"not null;index:tx_event_contract_addr"`
	Amount0         float64 `gorm:"not null" sql:"type:decimal(28,18);"`
	Amount1         float64 `gorm:"not null" sql:"type:decimal(28,18);"`

	Status       TxStatus `gorm:"not null;index:tx_event_status"`
	TxHash       string   `gorm:"not null;index:tx_event_tx_hash"`
	BlockHash    string   `gorm:"not null"`
	BlockTime    int64    `gorm:"not null;index:tx_event_block_time"`
	Height       int64    `gorm:"not null;index:tx_event_tx_height"`
	ConfirmedNum int64
}

func (TxEventLog) TableName() string {
	return "tx_event_log"
}

func (l *TxEventLog) BeforeCreate() (err error) {
	l.CreatedAt = time.Now()
	l.ContractAddress = strings.ToLower(l.ContractAddress)
	l.TxHash = strings.ToLower(l.TxHash)
	l.BlockHash = strings.ToLower(l.BlockHash)
	return nil
}

type Result24Hour struct {
	ContractAddress string
	TotalAmount0    float64
	TotalAmount1    float64
}

func GetLast24HourTotalAccount(db *gorm.DB) ([]Result24Hour, error) {
	blockLog := BlockLog{}
	err := db.Order("height desc").First(&blockLog).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	res := make([]Result24Hour, 0)
	dbIns := db.Table("tx_event_log").Select("contract_address, sum(amount0) as total_amount0, sum(amount1) as total_amount1").Group("contract_address").Where("block_time > ?", blockLog.BlockTime-time.Duration(24*time.Hour).Milliseconds()/1000).Find(&res)
	return res, dbIns.Error
}
