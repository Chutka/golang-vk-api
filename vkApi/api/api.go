package api

import (
	"fmt"
	"log"
	"net/url"

	"github.com/Chutka/golang-vk-api/config"
	"github.com/Chutka/golang-vk-api/request"
)

// List of methods of vk api
const (
	MessagesGetConversationMembers string = "messages.getConversationMembers"
	MessagesGetConversations       string = "messages.getConversations"
	MessagesSend                   string = "messages.send"

	GroupsGetLongPollServer string = "groups.getLongPollServer"
)

type GroupsGetLongPollServerModel struct {
	Response struct {
		Key    string `json:"key"`
		Server string `json:"server"`
		Ts     string `json:"ts"`
	} `json:"response"`
}

type VkAPI struct {
	VkAPIConfig config.VkConfig
}

type params map[string]string

func GetAPI(vkConfig config.VkConfig) *VkAPI {
	return &VkAPI{
		VkAPIConfig: vkConfig,
	}
}

func (api *VkAPI) paramsToQueryString(p params) string {
	queryParams := url.Values{}
	queryParams.Add("v", api.VkAPIConfig.APIVersion)
	queryParams.Add("access_token", api.VkAPIConfig.CommunityAccessToken)
	for k, v := range p {
		queryParams.Add(k, v)
	}
	return queryParams.Encode()
}

func (api *VkAPI) preparedURLWithMethod(method string, p params) string {
	resultS := fmt.Sprintf(
		"%s%s?%s",
		api.VkAPIConfig.APIMethodPath,
		method,
		api.paramsToQueryString(p),
	)
	fmt.Println("URL - ", resultS)
	return resultS
}

func (api *VkAPI) MethodRequest(method string, p params, i interface{}) error {
	err := request.GetRequest(api.preparedURLWithMethod(method, p), i)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
