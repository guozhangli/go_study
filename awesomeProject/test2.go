package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	//len
	tracer := "genji is a adlakdsfl"
	fmt.Println(len(tracer))
	tip2 := "忍者"
	fmt.Println(len(tip2))
	fmt.Println(utf8.RuneCountInString(tip2))

	theme := "狙击 start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c %d\n", theme[i], theme[i])
	}
	for _, s := range theme {
		fmt.Printf("Unicode: %c %d\n", s, s)
	}
	//index
	tracer1 := "夜空中最亮的星,夜byebye"
	comma := strings.Index(tracer1, ",")
	pos := strings.Index(tracer1[comma:], "夜")
	fmt.Println(tracer1[comma:])
	fmt.Println(comma, pos, tracer1[comma+pos:])
	//修改字符串
	angle := "Heros never die"
	angleBytes := []byte(angle)
	for i := 5; i <= 10; i++ {
		angleBytes[i] = ' '
	}
	fmt.Println(string(angleBytes))

	hammer := "吃我一锤"
	sickle := "识别"
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	fmt.Println(stringBuilder.String())

	var progress = 2
	var target = 8
	title := fmt.Sprintf("已采集%d个草药，还需要%d个完成任务", progress, target)
	fmt.Println(title)

	pi := 3.14159
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)

	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}
	fmt.Printf("使用'%%v'%v\n", profile)
	fmt.Printf("使用'%%+v'%+v\n", profile)
	fmt.Printf("使用'%%#v'%#v\n", profile)
	fmt.Printf("使用'%%T'%T\n", profile)

	message := "Away from keyboard. https://golang.org/"
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodedMessage)

	data, err := base64.StdEncoding.DecodeString(encodedMessage)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}

func test(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close() //延迟执行函数
	reader := bufio.NewReader(file)
	var sectionName string
	for {
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		linestr = strings.TrimSpace(linestr)
		if linestr == "" {
			continue
		}
		if linestr[0] == ';' {
			continue
		}
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			sectionName = linestr[1 : len(linestr)-1]
		}
	}
	fmt.Println(sectionName)
	return ""
}

const pi = 3.1415926 //常量pi
const e = 2.718281   //常量e
const ()
