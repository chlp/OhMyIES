package rss

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"net/http"
	"ohmyies/internal/model"
	"ohmyies/pkg/logger"
	"sort"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

const fetchRssUrlTemplate = "https://sms.schoolsoft.se/engelska/jsp/public/right_public_parent_rss.jsp?key=%s&key2=%s"

func (f *Feed) fetchNewAndExec(exec func(model.Msg) bool) {
	msgs, err := f.fetch()
	if err != nil {
		logger.Printf("rss::fetchNewAndExec. fetch problem: %v", err)
		return
	}

	for _, msg := range msgs {
		if msg.Guid == "" {
			logger.Printf("rss::fetchNewAndExec. Empty GUID for message: %s", msg.Title)
			continue
		}
		if f.isFetched(msg.Guid) {
			continue
		}

		if exec(msg) {
			f.mu.Lock()
			f.fetchedGuids = append(f.fetchedGuids, msg.Guid)
			f.mu.Unlock()
			f.needSyncFetchedGuids = true
		}
	}
}

func (f *Feed) isFetched(guid string) bool {
	f.mu.RLock()
	defer f.mu.RUnlock()

	for _, g := range f.fetchedGuids {
		if g == guid {
			return true
		}
	}
	return false
}

func (f *Feed) fetch() ([]model.Msg, error) {
	url := fmt.Sprintf(fetchRssUrlTemplate, f.key, f.key2)

	resp, err := http.Get(url)
	if err != nil {
		logger.Printf("rss::fetch. Error fetching feed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Printf("rss::fetch. Unexpected status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var feed feed
	if err = decoder.Decode(&feed); err != nil {
		if err == io.EOF {
			return nil, nil
		}
		logger.Printf("rss::fetch. Error parsing XML: %v", err)
		return nil, err
	}

	if len(feed.Channel.Items) == 0 {
		return nil, nil
	}

	messages := make([]model.Msg, len(feed.Channel.Items))
	for i, item := range feed.Channel.Items {
		messages[i] = model.Msg{
			Guid:        item.Guid,
			Title:       item.Title,
			Description: htmlToText(strings.TrimSpace(item.Description)),
			PubDate:     parsePubDate(item.PubDate),
			Type:        getTypeByTitle(item.Title),
		}
	}

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].PubDate.Before(messages[j].PubDate)
	})

	return messages, nil
}

func htmlToText(html string) string {
	return bluemonday.StrictPolicy().Sanitize(html)
}

func parsePubDate(raw string) time.Time {
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		"Mon, 2 Jan 2006 15:04:05 -0700",
		"Mon, 2 Jan 2006 15:04:05 MST",
		"02 Jan 2006 15:04:05 -0700",
		"2 Jan 2006 15:04:05 -0700",
	}
	for _, layout := range formats {
		if t, err := time.Parse(layout, raw); err == nil {
			return t
		}
	}
	logger.Printf("rss::parsePubDate. Could not parse pub date: %v", raw)
	return time.Time{}
}
