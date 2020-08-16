package events

import (
	"log"
	"net/http"

	"github.com/Chutka/golang-vk-api/request"
	"github.com/Chutka/golang-vk-api/serialize"
)

type LongPollResponse struct {
	Ts      string            `json:"ts"`
	Updates []EventWithObject `json:"updates"`
}

type EventWithObject struct {
	Type    string      `json:"type"`
	GroupID int32       `json:"group_id"`
	Object  interface{} `json:"object"`
}

const (
	TmpFileName = "tmpFileName.gob"
)

func getEventModel(event *EventWithObject) (*EventWithObject, error) {
	var err error = nil
	switch event.Type {
	case NewMessageEventName:
		newMessage := NewMessageModel{}
		err = serialize.MapToStruct(event.Object, &newMessage)
		event.Object = newMessage
		break
	}
	return event, err
}

func GetLongPollResponse(resp *http.Response) (*LongPollResponse, error) {
	longPollResponse := &LongPollResponse{}
	err := request.ParseResponse(resp, &longPollResponse)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for i, event := range longPollResponse.Updates {
		newEvent, modelErr := getEventModel(&event)
		if modelErr != nil {
			log.Fatal(modelErr)
			return nil, modelErr
		}
		longPollResponse.Updates[i] = *newEvent
	}
	return longPollResponse, nil
}
