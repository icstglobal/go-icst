package skill

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
	contractData := struct {
		PHash      string
		PPublisher []byte
		PPlatform  []byte
		PConsumer  []byte
		PPrice     uint32
		PRatio     uint8
	}{
		PHash:      "hash",
		PPublisher: data["PPublisher"].([]byte),
		PPlatform:  data["PPlatform"].([]byte),
		PConsumer:  data["PConsumer"].([]byte),

		PPrice: 1,
		PRatio: 50,
	}
	trans, err := p.c.NewContract(ctx, sender, string(chain.SkillContractType), contractData)
	if err != nil {
		log.Println("failed to publish skill contract, ", err)
		return nil, err
	}
	//TODO: save metadata
	return trans, nil
}
