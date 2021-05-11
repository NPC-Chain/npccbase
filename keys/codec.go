package keys

import (
	go_amino "github.com/tendermint/go-amino"
)

func RegisterCodec(cdc *go_amino.Codec) {
	cdc.RegisterInterface((*Info)(nil), nil)
	cdc.RegisterConcrete(&localInfo{}, "npccbase/keys/localInfo", nil)
	cdc.RegisterConcrete(&offlineInfo{}, "npccbase/keys/offlineInfo", nil)
	cdc.RegisterConcrete(&importInfo{}, "npccbase/keys/importInfo", nil)
}
