package audittool

import (
	"context"
	"log"
)

// PubSubMessage Struct para recibir datos de pubsub
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// GetIamMessage recibe un mensaje de audit de infra
func GetIamMessage(ctx context.Context, m PubSubMessage) error {
	log.Printf("Recibido mensaje de PubSub: %s", m.Data)

	return nil
}
