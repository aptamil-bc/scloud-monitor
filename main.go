package main

import (
	"fmt"

	"github.com/shinecloudfoundation/scloud-monitor/constant"
	"github.com/shinecloudfoundation/scloud-monitor/daemon"
	"github.com/shinecloudfoundation/scloud-monitor/node"
	"github.com/shinecloudfoundation/shinecloudnet/app"
	sdk "github.com/shinecloudfoundation/shinecloudnet/types"
)

func main() {
	cdc := app.MakeCodec()

	rpcNode := node.NewNode(constant.ChainID, constant.NodeRpc, cdc)

	latestHeight, err := rpcNode.GetLatestHeight()
	if err != nil {
		fmt.Println(err.Error())
	}

	vals, err := rpcNode.QueryValiatorList(1, 100, sdk.BondStatusBonded)
	if err != nil {
		fmt.Println(err.Error())
	}

	go daemon.MonitorDaemon(latestHeight, rpcNode, vals)

	select {}
}
