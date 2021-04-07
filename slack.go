package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Slack struct {
	Text string `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	MarkdownIn []string `json:"mrkdwn_in"`
	Fields   []Field  `json:"fields"`
}

type Field struct {
	Title string `json:"title"`
	Value int    `json:"value"`
	Short bool   `json:"short"`
}

func Send(webhook string, message string) error {
	requestBody, _ := json.Marshal(Slack{Text: message})
	response, err := http.Post(webhook, "application/json", bytes.NewReader(requestBody))

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	
	if buffer.String() != "ok" {
		return errors.New("Failed to send slack notification")
	}

	return nil
}