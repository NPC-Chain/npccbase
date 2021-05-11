package baseabci

import (
	"github.com/NPC-Chain/npccbase/account"
	"github.com/NPC-Chain/npccbase/consensus"
	"github.com/NPC-Chain/npccbase/keys"
	"github.com/NPC-Chain/npccbase/txs"
	"github.com/NPC-Chain/npccbase/types"
	go_amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

func MakenpccbaseCodec() *go_amino.Codec {

	var cdc = go_amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	RegisterCodec(cdc)

	return cdc
}

func RegisterCodec(cdc *go_amino.Codec) {
	txs.RegisterCodec(cdc)
	account.RegisterCodec(cdc)
	keys.RegisterCodec(cdc)
	consensus.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
}
