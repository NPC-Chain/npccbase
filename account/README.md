# Account
version:
v0.1

date:
2021.05.10

## Introduction：


Account provides basic account information and related methods
### Account interface
    
```
type Account interface {
	GetAddress() types.AccAddress
	SetAddress(addr types.AccAddress) error
	GetPublicKey() crypto.PubKey
	SetPublicKey(pubKey crypto.PubKey) error
	GetNonce() int64
	SetNonce(nonce int64) error
}
```

### BaseAccount
```
type BaseAccount struct {
	AccountAddress types.AccAddress `json:"account_address"` // account address
	Publickey      crypto.PubKey    `json:"public_key"`      // public key
	Nonce          int64            `json:"nonce"`           // identifies tx_status of an account
}
