package rss

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"ohmyies/internal/model"
	"ohmyies/pkg/logger"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

const fetchRssUrlTemplate = "https://sms.schoolsoft.se/engelska/jsp/public/right_public_parent_rss.jsp?key=%s&key2=%s"

func fetch(key, key2 string) ([]model.Msg, error) {
	url := fmt.Sprintf(fetchRssUrlTemplate, key, key2)

	resp, err := http.Get(url)
	if err != nil {
		logger.Printf("rss::fetch. Error fetching feed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Printf("rss::fetch. Error reading feed body: %v", err)
	}

	var feed feed
	if err := xml.Unmarshal(body, &feed); err != nil {
		logger.Printf("rss::fetch. Error parsing XML: %v", err)
	}

	if len(feed.Channel.Items) == 0 {
		return nil, nil
	}

	messages := make([]model.Msg, len(feed.Channel.Items))
	for i, item := range feed.Channel.Items {
		pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			pubTime, _ = time.Parse(time.RFC1123, item.PubDate)
		}

		messages[i] = model.Msg{
			Guid:        item.Guid,
			Title:       item.Title,
			Description: htmlToText(strings.TrimSpace(item.Description)),
			PubDate:     pubTime,
			Type:        getTypeByTitle(item.Title),
		}
	}

	return messages, nil
}

func md5String(s string) []byte {
	h := md5.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func htmlToText(html string) string {
	return bluemonday.StrictPolicy().Sanitize(html)
}

func tmp(key, key2 string) ([]model.Msg, error) {
	url := fmt.Sprintf(fetchRssUrlTemplate, key, key2)

	lastReadGuidFile := "./feed_" + fmt.Sprintf("%x", md5String(url))

	resp, err := http.Get(url)
	if err != nil {
		logger.Printf("rss::fetch. Error fetching feed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Printf("rss::fetch. Error reading feed body: %v", err)
	}

	var feed feed
	if err := xml.Unmarshal(body, &feed); err != nil {
		log.Fatalf("rss::fetch. Error parsing XML: %v", err)
	}

	if len(feed.Channel.Items) == 0 {
		return nil, nil
	}

	messages := make([]model.Msg, len(feed.Channel.Items))
	for i, item := range feed.Channel.Items {
		pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			pubTime, _ = time.Parse(time.RFC1123, item.PubDate)
		}

		message := fmt.Sprintf("%s\n*%s*", pubTime.Format("2006-01-02 15:04:05"), item.Title)

		if strings.TrimSpace(item.Description) != "" {
			plainText := htmlToText(item.Description)
			message += "\n\n" + plainText
		}

		// Экранируем спецсимволы для Telegram Markdown
		messages[i] = message
	}

}
