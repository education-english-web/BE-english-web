package pusherclient

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/pusher/pusher-http-go/v5"

	"github.com/education-english-web/BE-english-web/pkg/uuidstring/ulid"
)

const triggerChunkSize = 5000

type pusherClient struct {
	client pusher.Client
}

var (
	once       = sync.Once{}
	onceClient *pusherClient
)

// TODO: implement this function
//func Setup(cfg config.Pusher) error {
//	once.Do(func() {
//		onceClient = &pusherClient{
//			client: pusher.Client{
//				AppID:                     cfg.AppID,
//				Key:                       cfg.Key,
//				Secret:                    cfg.Secret,
//				Cluster:                   cfg.Cluster,
//				Secure:                    cfg.Secure,
//				EncryptionMasterKeyBase64: cfg.EncryptionMasterKeyBase64, // openssl rand -base64 32
//			},
//		}
//	})
//
//	return nil
//}

func Get() PusherClient {
	return onceClient
}

func (pc *pusherClient) AuthorizePrivateChannel(params []byte) ([]byte, error) {
	return pc.client.AuthorizePrivateChannel(params)
}

func (pc *pusherClient) Trigger(channelID, event string, data map[string]interface{}) error {
	if strings.HasPrefix(event, "chunked-") {
		return pc.triggerChunked(channelID, event, data)
	}

	return pc.client.Trigger(channelID, event, data)
}

func (pc *pusherClient) TriggerMultiChannels(channelIDs []string, event string, data map[string]interface{}) error {
	for i := range channelIDs {
		if err := pc.client.Trigger(channelIDs[i], event, data); err != nil {
			return err
		}
	}

	return nil
}

func (pc *pusherClient) triggerChunked(channelID, event string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	msgID := ulid.New().GetUUID()

	for i := 0; i*triggerChunkSize < len(jsonData); i++ {
		start := i * triggerChunkSize
		end := start + triggerChunkSize

		if end > len(jsonData) {
			end = len(jsonData)
		}

		chunk := string(jsonData[start:end])
		final := end >= len(jsonData)

		if err := pc.client.Trigger(channelID, event, map[string]interface{}{
			"id":    msgID,
			"index": i,
			"chunk": chunk,
			"final": final,
		}); err != nil {
			return fmt.Errorf("failed to trigger Pusher: %w", err)
		}
	}

	return nil
}
