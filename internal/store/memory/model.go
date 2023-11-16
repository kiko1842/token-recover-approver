package memory

import (
	"encoding/json"

	"github.com/bnb-chain/airdrop-service/pkg/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Account struct {
	Address       sdk.AccAddress `json:"address"`
	AccountNumber int64          `json:"account_number"`
	SummaryCoins  sdk.Coins      `json:"summary_coins,omitempty"`
	Coins         sdk.Coins      `json:"coins,omitempty"`
	FrozenCoins   sdk.Coins      `json:"frozen_coins,omitempty"`
	LockedCoins   sdk.Coins      `json:"locked_coins,omitempty"`
}

// Proofs is a list of account to merkle proof
type Proofs []*Proof

// Proof is a merkle proof of an account
type Proof struct {
	Address sdk.AccAddress `json:"address"`
	Coin    sdk.Coin       `json:"coin"`
	Proof   [][]byte       `json:"proof"`
}

func (p *Proof) UnmarshalJSON(data []byte) error {
	var source = struct {
		Address sdk.AccAddress `json:"address"`
		Coin    sdk.Coin       `json:"coin"`
		Proof   []string       `json:"proof"`
	}{}

	err := json.Unmarshal(data, &source)
	if err != nil {
		return err
	}

	p.Address = source.Address
	p.Coin = source.Coin
	p.Proof = util.MustDecodeHexArrayToBytes(source.Proof)

	return nil
}
