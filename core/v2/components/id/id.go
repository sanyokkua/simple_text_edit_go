package id

import (
	"simple_text_editor/core/v2/api"
	"time"
)

type idProvider struct {
}

func (receiver idProvider) GetId() int64 {
	return time.Now().UnixNano()
}

func CreateIdProvider() api.IIdProvider {
	return &idProvider{}
}
