package txs

import (
	"encoding/hex"
	"fmt"

	"github.com/NPC-Chain/npccbase/types"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"
)

// 功能：
type TxQcp struct {
	TxStd       *TxStd    `json:"txstd"`       //TxStd结构
	From        string    `json:"from"`        //qscName
	To          string    `json:"to"`          //qosName
	Sequence    int64     `json:"sequence"`    //发送Sequence
	Sig         Signature `json:"sig"`         //签名
	BlockHeight int64     `json:"blockheight"` //Tx所在block高度
	TxIndex     int64     `json:"txindex"`     //Tx在block的位置
	IsResult    bool      `json:"isresult"`    //是否为Result
	Extends     string    `json:"extends"`     //扩展字段
}

var _ types.Tx = (*TxQcp)(nil)

// Type: just for implements types.Tx
func (tx *TxQcp) Type() string {
	return "txqcp"
}

// 功能：获取 TxQcp 的签名字段
// 返回：字段拼接后的 []byte
func (tx *TxQcp) getSigData() []byte {
	ret := tx.TxStd.getSignData()
	if ret == nil {
		return nil
	}

	ret = append(ret, []byte(tx.From)...)
	ret = append(ret, []byte(tx.To)...)
	ret = append(ret, types.Int2Byte(tx.Sequence)...)
	ret = append(ret, types.Int2Byte(tx.BlockHeight)...)
	ret = append(ret, types.Int2Byte(tx.TxIndex)...)
	ret = append(ret, types.Bool2Byte(tx.IsResult)...)
	ret = append(ret, []byte(tx.Extends)...)
	return ret
}

func (tx *TxQcp) BuildSignatureBytes() []byte {
	return tx.getSigData()
}

func (tx *TxQcp) SignTx(privateKey crypto.PrivKey) (signByte []byte, err error) {
	data := tx.BuildSignatureBytes()
	if data == nil {
		return nil, errors.New("Signature txQcp err!")
	}
	signByte, err = privateKey.Sign(data)
	if err != nil {
		return nil, err
	}

	return
}

// 构建TxQCP结构体
func NewTxQCP(txStd *TxStd, from string, to string, sequence int64,
	blockHeight int64, txIndex int64, isResult bool, extends string) (rTx *TxQcp) {

	rTx = &TxQcp{
		txStd,
		from,
		to,
		sequence,
		Signature{},
		blockHeight,
		txIndex,
		isResult,
		extends,
	}

	return
}

// ValidateBasicData 校验txQcp基础数据是否合法
func (tx *TxQcp) ValidateBasicData(isCheckTx bool, currentChainID string) (err types.Error) {
	// 1. From To Sequence Sig 不为空
	// 2. to == current.chainId

	if tx.From == "" || tx.To == "" || tx.Sequence <= 0 || tx.BlockHeight <= 0 || tx.TxIndex < 0 {
		return types.ErrInternal(fmt.Sprintf("TxQcp's Basic data is not valid. basic data: from: %s , to: %s , seq: %d , height: %d,index:%d ",
			tx.From, tx.To, tx.Sequence, tx.BlockHeight, tx.TxIndex))
	}

	if len(tx.Sig.Signature) == 0 {
		return types.ErrInternal("TxQcp's Signature is empty")
	}

	if tx.Sig.Nonce < 0 {
		return types.ErrInternal("TxQcp's Signature nonce is not valid.")
	}

	if tx.To != currentChainID {
		return types.ErrInternal(fmt.Sprintf("TxQcp's ToChainID is not match current chainID. expect: %s, actual: %s", currentChainID, tx.To))
	}

	return
}

func (tx *TxQcp) GenTxHash() string {
	bz := crypto.Sha256(tx.BuildSignatureBytes())
	return hex.EncodeToString(bz)
}
