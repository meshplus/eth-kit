package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type Log struct {
	Address     common.Address `json:"address" gencodec:"required"`
	Topics      []common.Hash  `json:"topics" gencodec:"required"`
	Data        hexutil.Bytes  `json:"data" gencodec:"required"`
	BlockNumber hexutil.Uint64 `json:"blockNumber"`
	TxHash      common.Hash    `json:"transactionHash" gencodec:"required"`
	TxIndex     hexutil.Uint   `json:"transactionIndex"`
	BlockHash   common.Hash    `json:"blockHash"`
	Index       hexutil.Uint   `json:"logIndex"`
	Removed     bool           `json:"removed"`
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

func BytesToAddress(b []byte) common.Address {
	return common.BytesToAddress(b)
}

func BytesToHash(b []byte) common.Hash {
	return common.BytesToHash(b)
}

func Uint64(i uint64) hexutil.Uint64 {
	return hexutil.Uint64(i)
}

func Uint(i uint64) hexutil.Uint {
	return hexutil.Uint(i)
}
