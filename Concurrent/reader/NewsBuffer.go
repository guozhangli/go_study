package reader

import (
	"fmt"
	TestProject "go_study/datastructure"
	"math"
	"sync"
)

type NewsBuffer struct {
	concurrentMap sync.Map
	buffer        *TestProject.LinkedBlockingQueue
}

func NewNewsBuffer() *NewsBuffer {
	return &NewsBuffer{
		buffer: TestProject.NewLinkedBlockingQueue(math.MaxInt32),
	}
}

func (nb *NewsBuffer) add(item *CommonInformationItem) {
	if _, ok := nb.concurrentMap.Load(item.id); !ok {
		nb.buffer.Put(item)
		nb.concurrentMap.Store(item.id, item.source)
	} else {
		fmt.Println("Item " + item.id + " has been processed before")
	}
}

func (nb *NewsBuffer) get() *CommonInformationItem {
	return nb.buffer.Take().(*TestProject.Node).Data.(*CommonInformationItem)
}
