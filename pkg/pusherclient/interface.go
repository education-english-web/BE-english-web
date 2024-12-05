package pusherclient

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

type PusherClient interface {
	AuthorizePrivateChannel(params []byte) (response []byte, err error)
	Trigger(channelID string, event string, data map[string]interface{}) error
	TriggerMultiChannels(channelIDs []string, event string, data map[string]interface{}) error
}
