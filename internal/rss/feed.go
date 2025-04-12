package rss

import (
	"ohmyies/internal/model"
	"ohmyies/pkg/filestore"
	"ohmyies/pkg/logger"
	"os"
	"sync"
	"time"
)

const (
	fetcherPeriod          = 5 * time.Minute
	fetchedGuidsSyncPeriod = 5 * time.Second
)

type Feed struct {
	name string
	key  string
	key2 string

	store                *filestore.FileStore
	fetchedGuids         []string
	mu                   sync.RWMutex
	needSyncFetchedGuids bool
}

func NewFeed(name, key, key2 string, store *filestore.FileStore, newMsgExec func(model.Msg) bool) *Feed {
	if store == nil {
		logger.Fatalf("rss::NewFeed: store is nil")
		return nil
	}

	fetchedGuids := make([]string, 0)
	if err := store.LoadJSON(&fetchedGuids); err != nil {
		if err == os.ErrNotExist {
			err = store.SaveJSON(fetchedGuids)
		}
		if err != nil {
			logger.Fatalf("rss::NewFeed: failed to load %s: %v", name, err)
			return nil
		}
	}

	f := &Feed{
		name:                 name,
		key:                  key,
		key2:                 key2,
		store:                store,
		fetchedGuids:         fetchedGuids,
		mu:                   sync.RWMutex{},
		needSyncFetchedGuids: false,
	}
	f.startFetcher(newMsgExec)
	f.startSyncer()
	return f
}

func (f *Feed) startFetcher(exec func(model.Msg) bool) {
	f.fetchNewAndExec(exec)
	ticker := time.NewTicker(fetcherPeriod)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				f.fetchNewAndExec(exec)
			}
		}
	}()
}

func (f *Feed) startSyncer() {
	ticker := time.NewTicker(fetchedGuidsSyncPeriod)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				f.sync()
			}
		}
	}()
}

func (f *Feed) sync() {
	if !f.needSyncFetchedGuids {
		return
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	if err := f.store.SaveJSON(f.fetchedGuids); err != nil {
		logger.Printf("rss::sync. Error saving fetched guids: %v", err)
	}

	f.needSyncFetchedGuids = false
}
