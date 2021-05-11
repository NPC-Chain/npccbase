package types

import (
	go_amino "github.com/tendermint/go-amino"
)

func RegisterCodec(cdc *go_amino.Codec) {
	cdc.RegisterConcrete(&AccAddress{}, "npccbase/types/AccAddress", nil)
	cdc.RegisterConcrete(&ValAddress{}, "npccbase/types/ValAddress", nil)
	cdc.RegisterConcrete(&ConsAddress{}, "npccbase/types/ConsAddress", nil)
	cdc.RegisterInterface((*Address)(nil), nil)
}
