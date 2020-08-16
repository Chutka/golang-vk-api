package events

const (
	NewMessageEventName = "message_new"
)

type NewMessageModel struct {
	Message struct {
		Id     int32  `json:"id"`
		Text   string `json:"text"`
		PeerId int32  `json:"peer_id"`
		FromId int32  `json:"from_id"`
	} `json:"message"`
}
