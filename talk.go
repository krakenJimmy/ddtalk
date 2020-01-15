package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// TextMsg comment
type TextMsg struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

// SendMsg comment
func SendMsg(content string) {
	requestData := TextMsg{
		MsgType: "text",
	}

	requestData.Text.Content = content
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://oapi.dingtalk.com/robot/send?access_token=6f1369f41eb0b124263ceb458d9cceedd4a85c36c9962ade594af1db92f55fd9", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	fmt.Println(result)
}
