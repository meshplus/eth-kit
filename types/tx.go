package types

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

// Transaction types.
const (
	LegacyTxType = iota
	AccessListTxType
)

var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

// TxData is the underlying data of a transaction.
//
// This is implemented by LegacyTx and AccessListTx.
type TxData interface {
	TxType() byte // returns the type ID
	copy() TxData // creates a deep copy and initializes all fields

	GetChainID() *big.Int
	GetAccessList() types.AccessList
	GetData() []byte
	GetGas() uint64
	GetGasPrice() *big.Int
	GetValue() *big.Int
	GetNonce() uint64
	GetTo() *common.Address

	RawSignatureValues() (v, r, s *big.Int)
	setSignatureValues(chainID, v, r, s *big.Int)
}

// AccessListTx is the data of EIP-2930 access list transactions.
type AccessListTx struct {
	ChainID    *big.Int         // destination chain ID
	Nonce      uint64           // nonce of sender account
	GasPrice   *big.Int         // wei per gas
	Gas        uint64           // gas limit
	To         *common.Address  `rlp:"nil"` // nil means contract creation
	Value      *big.Int         // wei amount
	Data       []byte           // contract invocation input data
	AccessList types.AccessList // EIP-2930 access list
	V, R, S    *big.Int         // signature values
}

// LegacyTx is the transaction data of regular Ethereum transactions.
type LegacyTx struct {
	Nonce    uint64          // nonce of sender account
	GasPrice *big.Int        // wei per gas
	Gas      uint64          // gas limit
	To       *common.Address `rlp:"nil"` // nil means contract creation
	Value    *big.Int        // wei amount
	Data     []byte          // contract invocation input data
	V, R, S  *big.Int        // signature values
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *AccessListTx) copy() TxData {
	cpy := &AccessListTx{
		Nonce: tx.Nonce,
		To:    tx.To, // TODO: copy pointed-to address
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are copied below.
		AccessList: make(types.AccessList, len(tx.AccessList)),
		Value:      new(big.Int),
		ChainID:    new(big.Int),
		GasPrice:   new(big.Int),
		V:          new(big.Int),
		R:          new(big.Int),
		S:          new(big.Int),
	}
	copy(cpy.AccessList, tx.AccessList)
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.

func (tx *AccessListTx) TxType() byte                    { return AccessListTxType }
func (tx *AccessListTx) GetChainID() *big.Int            { return tx.ChainID }
func (tx *AccessListTx) protected() bool                 { return true }
func (tx *AccessListTx) GetAccessList() types.AccessList { return tx.AccessList }
func (tx *AccessListTx) GetData() []byte                 { return tx.Data }
func (tx *AccessListTx) GetGas() uint64                  { return tx.Gas }
func (tx *AccessListTx) GetGasPrice() *big.Int           { return tx.GasPrice }
func (tx *AccessListTx) GetValue() *big.Int              { return tx.Value }
func (tx *AccessListTx) GetNonce() uint64                { return tx.Nonce }
func (tx *AccessListTx) GetTo() *common.Address          { return tx.To }

func (tx *AccessListTx) RawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *AccessListTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.ChainID, tx.V, tx.R, tx.S = chainID, v, r, s
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *LegacyTx) copy() TxData {
	cpy := &LegacyTx{
		Nonce: tx.Nonce,
		To:    tx.To, // TODO: copy pointed-to address
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are initialized below.
		Value:    new(big.Int),
		GasPrice: new(big.Int),
		V:        new(big.Int),
		R:        new(big.Int),
		S:        new(big.Int),
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.

func (tx *LegacyTx) TxType() byte                    { return LegacyTxType }
func (tx *LegacyTx) GetChainID() *big.Int            { return deriveChainId(tx.V) }
func (tx *LegacyTx) GetAccessList() types.AccessList { return nil }
func (tx *LegacyTx) GetData() []byte                 { return tx.Data }
func (tx *LegacyTx) GetGas() uint64                  { return tx.Gas }
func (tx *LegacyTx) GetGasPrice() *big.Int           { return tx.GasPrice }
func (tx *LegacyTx) GetValue() *big.Int              { return tx.Value }
func (tx *LegacyTx) GetNonce() uint64                { return tx.Nonce }
func (tx *LegacyTx) GetTo() *common.Address          { return tx.To }

func (tx *LegacyTx) RawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *LegacyTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.V, tx.R, tx.S = v, r, s
}

// deriveChainId derives the chain id from the given v parameter
func deriveChainId(v *big.Int) *big.Int {
	if v.BitLen() <= 64 {
		v := v.Uint64()
		if v == 27 || v == 28 {
			return new(big.Int)
		}
		return new(big.Int).SetUint64((v - 35) / 2)
	}
	v = new(big.Int).Sub(v, big.NewInt(35))
	return v.Div(v, big.NewInt(2))
}

func RecoverPlain(hash []byte, R, S, Vb *big.Int, homestead bool) ([]byte, error) {
	if Vb.BitLen() > 8 {
		return nil, fmt.Errorf("invalid signature")
	}
	V := byte(Vb.Uint64() - 27)
	if !crypto.ValidateSignatureValues(V, R, S, homestead) {
		return nil, fmt.Errorf("invalid signature")
	}
	// encode the signature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, crypto.SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V
	// recover the public key from the signature
	pub, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return nil, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return nil, fmt.Errorf("invalid public key")
	}

	return crypto.Keccak256(pub[1:])[12:], nil
}

// RlpHash encodes x and hashes the encoded bytes.
func RlpHash(x interface{}) *common.Hash {
	var h common.Hash
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	rlp.Encode(sha, x)
	sha.Read(h[:])
	return &h
}

// PrefixedRlpHash writes the prefix into the hasher before rlp-encoding x.
// It's used for typed transactions.
func PrefixedRlpHash(prefix byte, x interface{}) *common.Hash {
	var h common.Hash
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	sha.Write([]byte{prefix})
	rlp.Encode(sha, x)
	sha.Read(h[:])
	return &h
}

// CallArgs represents the arguments for a call.
type CallArgs struct {
	From       *common.Address   `json:"from"`
	To         *common.Address   `json:"to"`
	Gas        *hexutil.Uint64   `json:"gas"`
	GasPrice   *hexutil.Big      `json:"gasPrice"`
	Value      *hexutil.Big      `json:"value"`
	Data       *hexutil.Bytes    `json:"data"`
	AccessList *types.AccessList `json:"accessList"`
}
