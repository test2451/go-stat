package executor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stretchr/testify/assert"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	eabi "github.com/pancakeswap/pancake-statas/abi"
)

func TestParseSwapEvent(t *testing.T) {
	client, err := ethclient.Dial("wss://bsc-ws-node.nariox.org:443")
	assert.NoError(t, err)
	r, err := client.TransactionReceipt(context.Background(), common.HexToHash("0xb3dfe9a97d08481faa5d3e5afad2d502383d69ced2360fbae87f69206dea0a83"))
	assert.NoError(t, err)
	swapLog := r.Logs[4]
	swapPairAbi, err := abi.JSON(strings.NewReader(eabi.SwappairABI))
	if err != nil {
		panic("marshal abi error")
	}
	swapEvent, err := ParseSwapEvent(&swapPairAbi, swapLog)
	assert.NoError(t, err)
	bz, err := json.Marshal(swapEvent)
	fmt.Println(string(bz))

	eventModel := swapEvent.ToTxLog(swapLog)
	bz, err = json.Marshal(eventModel)
	fmt.Println(string(bz))
}
