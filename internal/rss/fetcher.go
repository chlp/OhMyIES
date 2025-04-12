package rss

import "ohmyies/pkg/filestore"

type Feed struct {
	name         string
	key          string
	key2         string
	fetchedGuids []string
	fileStore    *filestore.FileStore
}
