package reader

import (
	"github.com/mmcdole/gofeed"
	"github.com/mmcdole/gofeed/rss"
	TestProject "go_study/datastructure"
)

type RSSDataCapturer struct {
	name   string
	buffer *NewsBuffer
}

func NewRssDataCapturer(name string, buffer *NewsBuffer) *RSSDataCapturer {
	return &RSSDataCapturer{
		name:   name,
		buffer: buffer,
	}
}

func (rss *RSSDataCapturer) load(resource string) *TestProject.LinkedList {

	return nil
}
