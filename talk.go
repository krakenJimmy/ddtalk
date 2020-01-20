package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const token = "6f1369f41eb0b124263ceb458d9cceedd4a85c36c9962ade594af1db92f55fd9"

// Word 名言结构体
type Word struct {
	Code   string `json:"code"`
	Date   string `json:"date"`
	Ciba   string `json:"ciba"`
	CibaEN string `json:"ciba-en"`
	ImgURL string `json:"imgurl"`
}

// TextMsg 普通文本消息
type TextMsg struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

// LinkMsg 链接消息
type LinkMsg struct {
	MsgType string `json:"msgtype"`
	Link    struct {
		MessageURL string `json:"messageUrl"`
		PicURL     string `json:"picUrl"`
		Title      string `json:"title"`
		Text       string `json:"text"`
	} `json:"link"`
}

// SendMsg 发送普通通知
func SendMsg(content string) {
	requestData := TextMsg{
		MsgType: "text",
	}

	requestData.Text.Content = content
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://oapi.dingtalk.com/robot/send?access_token="+token, "application/json", bytes.NewBuffer(requestBody))
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
}

// SendLinkMsg 发送链接通知
func SendLinkMsg(word Word) {
	requestData := LinkMsg{
		MsgType: "link",
	}

	requestData.Link.MessageURL = word.ImgURL
	requestData.Link.PicURL = word.ImgURL
	requestData.Link.Title = word.Ciba
	requestData.Link.Text = word.CibaEN

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://oapi.dingtalk.com/robot/send?access_token="+token, "application/json", bytes.NewBuffer(requestBody))
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

// GetWord 获取每日一言
func GetWord() (word Word) {
	url := "https://api.ooopn.com/ciba/api.php?type=json"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result Word
	json.Unmarshal(body, &result)
	fmt.Println(result)
	return result
}
