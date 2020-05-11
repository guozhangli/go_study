package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Person struct {
	Age   int32
	Score int32
}

func main() {
	p := &Person{
		Age:   30,
		Score: 99,
	}

	/*var str=`{"Name":"李四","Age":20}`
	json.Unmarshal([]byte(str),p)
	fmt.Printf("%v",p)
	var b =[]byte(str)*/
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, p)
	byteBuffer := bytes.NewReader(buf.Bytes())

	var info Person
	if err := binary.Read(byteBuffer, binary.LittleEndian, &info); err != nil {
		fmt.Println("read byte error", err)
	}
	fmt.Printf("%v", info)
}
