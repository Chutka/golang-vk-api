package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseResponse(resp *http.Response, i interface{}) error {
	byteStr, errRead := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response: %s\n", byteStr)
	if errRead != nil {
		log.Fatal(errRead)
		return errRead
	}
	// fmt.Println("RESPONSE: ", string(byteStr))
	errUnmarshal := json.Unmarshal(byteStr, &i)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
		return errUnmarshal
	}
	return nil
}

func GetRequest(api string, i interface{}) error {
	resp, errReq := http.Get(api)
	if errReq != nil {
		log.Fatal(errReq)
		return errReq
	}
	defer resp.Body.Close()

	errParse := ParseResponse(resp, &i)
	if errParse != nil {
		log.Fatal(errParse)
		return errParse
	}
	return nil
}
