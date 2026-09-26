package cci

import (
	"errors"

	"github.com/vault-thirteen/auxie/file"
)

const (
	Err_FilePathIsNotSet    = "file path is not set"
	Err_ContentTypeIsNotSet = "content type is not set"
	Err_NoBytes             = "no bytes"
)

type CachedContentItem struct {
	data        []byte
	contentType string
	ttlSec      int
}

func NewCachedContentItemFromFile(filePath string, contentType string, ttl int) (i *CachedContentItem, err error) {
	if len(filePath) == 0 {
		return nil, errors.New(Err_FilePathIsNotSet)
	}
	if len(contentType) == 0 {
		return nil, errors.New(Err_ContentTypeIsNotSet)
	}

	i = new(CachedContentItem)

	i.data, err = file.GetFileContents(filePath)
	if err != nil {
		return nil, err
	}

	i.contentType = contentType
	i.ttlSec = ttl

	return i, nil
}
func NewCachedContentItemFromBytes(bytes []byte, contentType string, ttl int) (i *CachedContentItem, err error) {
	if len(bytes) == 0 {
		return nil, errors.New(Err_NoBytes)
	}
	if len(contentType) == 0 {
		return nil, errors.New(Err_ContentTypeIsNotSet)
	}

	i = new(CachedContentItem)

	i.data = bytes
	i.contentType = contentType
	i.ttlSec = ttl

	return i, nil
}

func (ci *CachedContentItem) Data() []byte        { return ci.data }
func (ci *CachedContentItem) ContentType() string { return ci.contentType }
func (ci *CachedContentItem) TTLSec() int         { return ci.ttlSec }
