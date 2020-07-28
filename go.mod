module github.com/shinecloudfoundation/scloud-monitor

go 1.13

require (
	github.com/shinecloudfoundation/shinecloudnet v1.2.1
	github.com/tendermint/tendermint v0.32.2
)

replace github.com/tendermint/iavl => github.com/shinecloudfoundation/iavl v0.12.4-shinecloudnet
