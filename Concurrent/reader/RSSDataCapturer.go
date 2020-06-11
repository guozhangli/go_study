package reader

import (
	"crypto/md5"
	TestProject "datastructure"
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
	"time"
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

func (rss *RSSDataCapturer) load(resource string) (*TestProject.LinkedList, error) {
	log.Println("parse url:", resource)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(resource)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	list := TestProject.NewLinkedList()
	for _, v := range feed.Items {
		item := &CommonInformationItem{
			title:       v.Title,
			txtDate:     v.Published,
			date:        time.Now(),
			link:        v.Link,
			description: v.Description,
			author:      v.Author.Name,
			source:      rss.name,
		}
		hashCode := md5.Sum([]byte(v.Description))
		item.id = fmt.Sprintf("%x", hashCode)
		list.HeadAdd(item)
	}
	return list, nil
}
