package tg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"ohmyies/pkg/logger"
	"strings"
)

const sendMessageUrlTemplate = "https://api.telegram.org/bot%s/sendMessage?%s"

type Chat struct {
	botApiToken string
	chatId      string
}

func NewChat(botApiToken, chatId string) *Chat {
	return &Chat{
		botApiToken: botApiToken,
		chatId:      chatId,
	}
}

func (c *Chat) SendMessage(message string) bool {
	params := url.Values{}
	params.Set("chat_id", c.chatId)
	params.Set("parse_mode", "markdown")
	params.Set("text", escapeMarkdown(message))

	resp, err := http.Get(fmt.Sprintf(sendMessageUrlTemplate, c.botApiToken, params.Encode()))
	if err != nil {
		logger.Printf("tg::SendMessage %s error: %v", c.chatId, err)
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Printf("tg::SendMessage %s read error: %v", c.chatId, err)
		return false
	}

	var result struct {
		Ok bool `json:"ok"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		logger.Printf("tg::SendMessage %s JSON unmarshal error: %v", c.chatId, err)
		return false
	}

	return result.Ok
}

func escapeMarkdown(text string) string {
	replacer := strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"`", "\\`",
		"[", "\\[",
	)
	return replacer.Replace(text)
}
