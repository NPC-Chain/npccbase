package keys

import (
	"github.com/NPC-Chain/npccbase/client/context"
	"github.com/spf13/cobra"
	go_amino "github.com/tendermint/go-amino"
)

func listKeysCommand(cdc *go_amino.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all keys",
		Long: `Return a list of all public keys stored by this key manager
along with their associated name and address.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			return runListCmd(cliCtx, cmd, args)
		},
	}
}

func runListCmd(ctx context.CLIContext, cmd *cobra.Command, args []string) error {
	kb, err := GetKeyBase(ctx)
	if err != nil {
		return err
	}

	infos, err := kb.List()
	if err == nil {
		printInfos(ctx, infos)
	}
	return err
}
