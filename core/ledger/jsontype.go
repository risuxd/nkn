package ledger

import (
	"github.com/nknorg/nkn/core/contract/program"
	"github.com/nknorg/nkn/core/transaction"
)

type HeaderInfo struct {
	Version          uint32              `json:"version"`
	PrevBlockHash    string              `json:"prevBlockHash"`
	TransactionsRoot string              `json:"transactionsRoot"`
	Timestamp        int64               `json:"timestamp"`
	Height           uint32              `json:"height"`
	ConsensusData    uint64              `json:"consensusData"`
	NextBookKeeper   string              `json:"nextBookKeeper"`
	Signer           string              `json:"signer"`
	Program          program.ProgramInfo `json:"program"`

	Hash string `json:"hash"`
}

type BlocksInfo struct {
	Hash         string                         `json:"hash"`
	Header       *HeaderInfo                    `json:"header"`
	Transactions []*transaction.TransactionInfo `json:"transactions"`
}
