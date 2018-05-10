package content

import (
	"context"
	"log"

	"github.com/icstglobal/go-icst/metadata"
	"github.com/icstglobal/go-icst/transaction"

	"github.com/icstglobal/go-icst/chain"
)

// Publisher is interface for user to publish a skill as a service
type Publisher struct {
	c chain.Chain
	s metadata.Store
}

func NewPublisher(chain chain.Chain, store metadata.Store) *Publisher {
	return &Publisher{chain, store}
}

func (p *Publisher) Pub(ctx context.Context, sender []byte, data map[string]interface{}) (*transaction.ContractTransaction, error) {
	//field names of this struct should match the args of underlying smart contract's constructor
	contractData := struct {
		PHash      string
		PPublisher []byte
		PPlatform  []byte
		PPrice     uint32
		PRatio     uint8
	}{
		PHash:      data["PHash"].(string),
		PPublisher: data["PPublisher"].([]byte),
		PPlatform:  data["PPlatform"].([]byte),

		PPrice: data["PPrice"].(uint32),
		PRatio: data["PRatio"].(uint8),
	}
	trans, err := p.c.NewContract(ctx, sender, string(chain.ContentContractType), contractData)
	if err != nil {
		log.Println("failed to publish content contract, ", err)
		return nil, err
	}
	//TODO: save metadata
	return trans, nil
}
