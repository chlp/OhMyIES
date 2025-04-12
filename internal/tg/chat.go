package tg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"ohmyies/pkg/logger"
	"strings"
	"sync"
)

const sendMessageUrlTemplate = "https://api.telegram.org/bot%s/sendMessage?%s"

type Chat struct {
	botApiToken string
	chatId      string
}

var chats = make(map[string]*Chat)
var mu *sync.Mutex

func NewChat(botApiToken, chatId string) *Chat {
	if mu == nil {
		mu = &sync.Mutex{}
	}

	mu.Lock()
	defer mu.Unlock()

	if chat, ok := chats[chatId]; ok {
		return chat
	}

	c := &Chat{
		botApiToken: botApiToken,
		chatId:      chatId,
	}
	chats[botApiToken+"_"+chatId] = c
	return c
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

	body, err := io.ReadAll(resp.Body)
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
