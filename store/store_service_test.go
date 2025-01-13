package store

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var testStoreService = &StorageService{}

func init() {
    testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
    assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
    initialLink := "https://chess.com"
    userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
    shortURL := "Jsz4k57oAX"

    // save mapping
    SaveUrlMapping(shortURL, initialLink, userUUId)

    // retrieve long url
    retrievedUrl := RetrieveInitialUrl(shortURL)

    assert.Equal(t, initialLink, retrievedUrl)
}
