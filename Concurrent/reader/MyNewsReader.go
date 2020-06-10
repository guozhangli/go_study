package reader

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	source = "../data/sources.txt"
	output = "../data/output/"
	stoped = false
)

func NewsReader() {
	nb := NewNewsBuffer()
	go NewsWriter(nb)

}

func NewsWriter(nb *NewsBuffer) {
	for !stoped {
		item := nb.get()
		err := ioutil.WriteFile(output+item.getFileName(), []byte(item.toXML()), os.ModePerm)
		if err != nil {
			fmt.Errorf("%s", "write file to disk error")
		}
	}
}
