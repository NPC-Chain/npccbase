package txs

import (
	"github.com/NPC-Chain/npccbasetypes"
	"github.com/QOSGroup/qbase/types"
	go_amino "github.com/tendermint/go-amino"
)

func RegisterCodec(cdc *go_amino.Codec) {
	cdc.RegisterConcrete(&QcpTxResult{}, "npccbase/txs/qcpresult", nil)
	cdc.RegisterConcrete(&Signature{}, "npccbase/txs/signature", nil)
	cdc.RegisterConcrete(&TxStd{}, "npccbase/txs/stdtx", nil)
	cdc.RegisterConcrete(&TxQcp{}, "npccbase/txs/qcptx", nil)
	cdc.RegisterInterface((*ITx)(nil), nil)
	cdc.RegisterInterface((*types.Tx)(nil), nil)
}
