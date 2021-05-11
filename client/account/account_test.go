package account

import (
	"fmt"
	"testing"

	"github.com/NPC-Chain/npccbase/baseabci"
	"github.com/NPC-Chain/npccbase/client/context"
	"github.com/NPC-Chain/npccbase/example/basecoin/app"
)

func TestGetAccountFromBech32Addr(t *testing.T) {

	ctx := context.NewCLIContext()

	cdc := baseabci.MakenpccbaseCodec()
	app.RegisterCodec(cdc)

	ctx = ctx.WithCodec(cdc)
	ctx = ctx.WithNodeIP("192.168.1.224")

	addr := "qos1mhvraeml8pjtm8fscyl7nmmrk2y372jpaw5sux"

	acc, err := GetAccountFromBech32Addr(ctx, addr)

	fmt.Println(err)

	ctx.PrintResult(acc)
}
