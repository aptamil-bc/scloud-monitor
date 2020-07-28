package daemon

import (
	sdk "github.com/shinecloudfoundation/shinecloudnet/types"
	"github.com/tendermint/tendermint/crypto"
	cmn "github.com/tendermint/tendermint/libs/common"
)

type MonitorValidator struct {
	OperatorAddr    sdk.ValAddress
	ConsensusAddr   cmn.HexBytes
	ConsensusPubKey crypto.PubKey
}
