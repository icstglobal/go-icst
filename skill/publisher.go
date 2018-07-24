package skill

import (
	"context"
	"log"

	"github.com/icstglobal/go-icst/metadata"
	"github.com/icstglobal/go-icst/transaction"

	"github.com/ethereum/go-ethereum/common"
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

func (p *Publisher) Pub(ctx context.Context, sender []byte, data map[string]interface{}) (*transaction.Transaction, error) {

	if pPublisher, ok := data["pPublisher"]; ok {
		data["pPublisher"] = common.HexToAddress(pPublisher.(string))
	}
	if pPlatform, ok := data["pPlatform"]; ok {
		data["pPlatform"] = common.HexToAddress(pPlatform.(string))
	}
	if pConsumer, ok := data["pConsumer"]; ok {
		data["pConsumer"] = common.HexToAddress(pConsumer.(string))
	}

	trans, err := p.c.NewContract(ctx, sender, string(chain.SkillContractType), data)
	if err != nil {
		log.Println("failed to publish skill contract, ", err)
		return nil, err
	}
	//TODO: save metadata
	return trans, nil
}
