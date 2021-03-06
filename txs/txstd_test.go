package txs

import (
	"testing"

	"github.com/NPC-Chain/npccbase/context"
	"github.com/NPC-Chain/npccbase/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func TestTxStd_GetSigners(t *testing.T) {

	txStd := TxStd{}

	require.Panics(t, func() {
		txStd.GetSigners()
	})

	txStd.ITxs = []ITx{&mockITX{0, false}}
	require.Equal(t, 0, len(txStd.GetSigners()))

	txStd.ITxs = []ITx{&mockITX{1, false}}
	require.Equal(t, 1, len(txStd.GetSigners()))

	txStd.ITxs = []ITx{&mockITX{2, false}}
	require.Equal(t, 2, len(txStd.GetSigners()))

	txStd.ITxs = []ITx{&mockITX{2, true}}
	require.Equal(t, 1, len(txStd.GetSigners()))

	txStd.ITxs = []ITx{&mockITX{3, true}}
	require.Equal(t, 2, len(txStd.GetSigners()))

	txStd.ITxs = []ITx{&mockITX{3, false}}
	require.Equal(t, 3, len(txStd.GetSigners()))

}

type mockITX struct {
	signerCount int
	hasDup      bool
}

func (m *mockITX) ValidateData(ctx context.Context) error {
	return nil
}

func (m *mockITX) Exec(ctx context.Context) (result types.Result, crossTxQcp *TxQcp) {
	return types.Result{}, nil
}

func (m *mockITX) GetSigner() []types.AccAddress {

	if m.signerCount == 0 {
		return nil
	}

	if m.signerCount == 1 {
		return []types.AccAddress{getAddress()}
	}

	if m.signerCount == 2 {
		fAddr := getAddress()
		if m.hasDup {
			return []types.AccAddress{fAddr, fAddr}
		} else {
			return []types.AccAddress{fAddr, getAddress()}
		}
	}

	if m.signerCount == 3 {
		fAddr := getAddress()
		if m.hasDup {
			return []types.AccAddress{fAddr, fAddr, getAddress()}
		} else {
			return []types.AccAddress{fAddr, getAddress(), getAddress()}
		}
	}

	return nil
}

func (m *mockITX) CalcGas() types.BigInt {
	return types.BigInt{}
}

func (m *mockITX) GetGasPayer() types.AccAddress {
	return nil
}

func (m *mockITX) GetSignData() []byte {
	return nil
}

var _ ITx = (*mockITX)(nil)

func getAddress() types.AccAddress {
	key := ed25519.GenPrivKey()
	pub := key.PubKey()
	addr := types.AccAddress(pub.Address())
	return addr
}
