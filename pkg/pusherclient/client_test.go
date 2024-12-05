package pusherclient

import (
	"testing"

	"github.com/pusher/pusher-http-go/v5"
	"github.com/stretchr/testify/assert"

	"github.com/education-english-web/BE-english-web/app/config"
)

func TestGet(t *testing.T) {
	t.Run("pusherclient.Get success", func(t *testing.T) {
		cfg := config.Pusher{
			AppID:                     "app_id",
			Key:                       "key",
			Secret:                    "secret",
			Cluster:                   "cluster",
			Secure:                    true,
			EncryptionMasterKeyBase64: "6Jlfv99Ex7rebcWdn7U7a77Gqb+/3ZowXrutTFo0k8w=",
		}
		_ = Setup(cfg)
		wantClient := &pusherClient{
			client: pusher.Client{
				AppID:                     cfg.AppID,
				Key:                       cfg.Key,
				Secret:                    cfg.Secret,
				Cluster:                   cfg.Cluster,
				Secure:                    cfg.Secure,
				EncryptionMasterKeyBase64: cfg.EncryptionMasterKeyBase64,
			},
		}
		gotClient := Get()

		assert.Equal(t, wantClient, gotClient)
	})
}

func Test_pusherclient_AuthorizePrivateChannel(t *testing.T) {
	cfg := config.Pusher{
		AppID:                     "app_id",
		Key:                       "key",
		Secret:                    "secret",
		Cluster:                   "cluster",
		Secure:                    true,
		EncryptionMasterKeyBase64: "6Jlfv99Ex7rebcWdn7U7a77Gqb+/3ZowXrutTFo0k8w=",
	}
	_ = Setup(cfg)
	gotClient := Get()
	_, _ = gotClient.AuthorizePrivateChannel([]byte{})
}

func Test_pusherclient_Trigger(t *testing.T) {
	cfg := config.Pusher{
		AppID:                     "app_id",
		Key:                       "key",
		Secret:                    "secret",
		Cluster:                   "cluster",
		Secure:                    true,
		EncryptionMasterKeyBase64: "6Jlfv99Ex7rebcWdn7U7a77Gqb+/3ZowXrutTFo0k8w=",
	}
	_ = Setup(cfg)
	gotClient := Get()
	_ = gotClient.Trigger("channel_id", "my_event", map[string]interface{}{"message": "my message detail"})
	_ = gotClient.Trigger("channel_id", "chunked-my_event", map[string]interface{}{"message": "my message detail"})
}

func Test_pusherclient_TriggerMultiChannels(t *testing.T) {
	cfg := config.Pusher{
		AppID:                     "app_id",
		Key:                       "key",
		Secret:                    "secret",
		Cluster:                   "cluster",
		Secure:                    true,
		EncryptionMasterKeyBase64: "6Jlfv99Ex7rebcWdn7U7a77Gqb+/3ZowXrutTFo0k8w=",
	}
	_ = Setup(cfg)
	gotClient := Get()
	_ = gotClient.TriggerMultiChannels([]string{"channel_id"}, "my_event", map[string]interface{}{"message": "my message detail"})
}
