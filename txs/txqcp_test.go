package txs

import (
	"fmt"
	"testing"

	"github.com/NPC-Chain/npccbase/account"
	"github.com/NPC-Chain/npccbase/context"
	"github.com/NPC-Chain/npccbase/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/NPC-Chain/npccbase/store"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func newQcpTxResult() (txqcpresult *QcpTxResult) {

	result := types.Result{
		Code: 0,
	}

	txqcpresult = NewQcpTxResult(result, 10, "qcp result info", "")
	var ctx context.Context
	err := txqcpresult.ValidateData(ctx)
	if err != nil {
		fmt.Print("QcpTxResult ValidateData Error")
		return nil
	}

	return
}

func newTxStd(tx ITx) (txstd *TxStd) {
	txstd = NewTxStd(tx, "qsc1", types.NewInt(100))
	signer := txstd.ITxs[0].GetSigner()

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(types.NewKVStoreKey("test"), types.StoreTypeIAVL, db)
	cms.LoadLatestVersion()
	ctx := context.NewContext(cms, abci.Header{}, false, log.NewNopLogger(), nil)
	err := txstd.ValidateBasicData(ctx, true, "qsc1")
	if err != nil {
		return nil
	}

	// no signer, no signature
	if signer == nil {
		txstd.Signature = []Signature{}
		return
	}

	accmapper := account.NewAccountMapper(nil, account.ProtoBaseAccount)
	// 填充 txstd.Signature[]
	for _, sg := range signer {
		prvKey := ed25519.GenPrivKey()
		nonce, err := accmapper.GetNonce(sg)
		if err != nil {
			return nil
		}

		signbyte, errsign := txstd.SignTx(prvKey, int64(nonce), "", ctx.ChainID())
		if signbyte == nil || errsign != nil {
			return nil
		}

		signdata := Signature{
			prvKey.PubKey(),
			signbyte,
			int64(nonce),
		}

		txstd.Signature = append(txstd.Signature, signdata)
	}

	return
}

func TestNewQcpTxResult(t *testing.T) {
	txResult := newQcpTxResult()
	require.NotNil(t, txResult)

	txResult.GetSigner()
	txResult.GetGasPayer()

	gas := txResult.CalcGas().Int64() < 0
	require.Equal(t, gas, false)
}

func TestNewTxStd(t *testing.T) {
	txResult := newQcpTxResult()
	require.NotNil(t, txResult)

	txStd := newTxStd(txResult)
	require.NotNil(t, txStd)

	txtype := txStd.Type()
	require.Equal(t, txtype, "txstd")
}

func TestTxQcp(t *testing.T) {
	txResult := newQcpTxResult()
	require.NotNil(t, txResult)

	txStd := newTxStd(txResult)
	require.NotNil(t, txStd)

	txqcp := NewTxQCP(txStd, "qsc1", "qos", 1, 13452345, 2, false, "")
	require.NotNil(t, txqcp)

	prvkey := ed25519.GenPrivKey()
	signbyte, err := txqcp.SignTx(prvkey)
	require.NotNil(t, signbyte)
	require.Nil(t, err)
	txqcp.Sig = Signature{
		prvkey.PubKey(),
		signbyte,
		txqcp.Sequence,
	}

	err = txqcp.ValidateBasicData(true, "qos")
	require.Nil(t, err)

	data := txqcp.BuildSignatureBytes()
	require.NotNil(t, data)
}
