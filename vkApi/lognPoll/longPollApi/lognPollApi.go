package longPoll

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Chutka/golang-vk-api/vkApi/api"
	"github.com/Chutka/golang-vk-api/vkApi/lognPoll/events"
)

const (
	MessageNewEvent = "message_new"
)

type VkLongPoll struct {
	Key    string
	Server string
	Ts     string
	Wait   string
}

func NewLongPoll(vkApi *api.VkAPI) (*VkLongPoll, error) {
	resp := &api.GroupsGetLongPollServerModel{}
	err := vkApi.MethodRequest(api.GroupsGetLongPollServer, map[string]string{
		"group_id": vkApi.VkAPIConfig.CommunityID,
	}, &resp)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &VkLongPoll{
		Key:    resp.Response.Key,
		Server: resp.Response.Server,
		Ts:     resp.Response.Ts,
		Wait:   "25",
	}, nil
}

func (longPoll *VkLongPoll) getURL() string {
	return fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%s", longPoll.Server, longPoll.Key, longPoll.Ts, longPoll.Wait)
}

type callback func(typeEvent string, groupId int32, i interface{})

func (longPoll *VkLongPoll) makeReq(cb callback) error {
	resp, errReq := http.Get(longPoll.getURL())
	if errReq != nil {
		log.Fatal(errReq)
		return errReq
	}
	defer resp.Body.Close()

	longPollResp, err := events.GetLongPollResponse(resp)

	if err != nil {
		log.Fatal(err)
		return err
	}
	longPoll.Ts = longPollResp.Ts

	fmt.Printf("%+v\n", longPollResp)
	for _, event := range longPollResp.Updates {
		cb(event.Type, event.GroupID, event.Object)
	}
	return nil
}

/**
Как должно работать ?
1. Делаем запрос longPoll
2. Дожидаемся тайаута или ответа и делаем снова
2.1 пришел ответ -> делаем запрос + формируем ответ на предыдущий
2.2 таймаут -> просто делаем запрос
*/

func (longPoll *VkLongPoll) StartPolling(cb callback) {
	for {
		err := longPoll.makeReq(cb)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
	}
}
