package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	properties map[string]string
}

func (c *Config) Init(filePath string) {
	//var ref string = ""
	c.properties = make(map[string]string)
	file, err := os.Open(filePath)
	defer file.Close()
	reader := bufio.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	for {
		b, isprefix, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err == io.EOF {
			break
		}
		regp, err := regexp.Compile(`#[\s\S]*`)
		if err != nil {
			log.Fatal(err)
		}
		var appendLine, temp string = "", ""
		if isprefix {
			temp = ""
			temp = string(b)
		} else {
			temp += string(b)
		}
		appendLine = temp
		if len(appendLine) > 0 && !regp.MatchString(appendLine) {
			fmt.Println(appendLine)
			//regp, err := regexp.Compile(`[\s\S]`)
			strs := strings.Split(appendLine, "=")
			fmt.Println(len(strs))
			c.properties[strs[0]] = strs[1]
		}
	}
}

func aa() {
	var bb string = "jdbc_url=jdbc:mysql://192.168.1.90:3306/otms?useUnicode\\=true&characterEncoding\\=UTF-8&zeroDateTimeBehavior\\=convertToNull"
	regp, err := regexp.Compile(`[^\\]=`)
	if err != nil {
		log.Fatal(err)
	}
	s := regp.Split(bb, 5)
	for key, value := range s {
		fmt.Println(key, value)
	}
}

func main() {
	conf := &Config{}
	conf.Init("d:\\config.properties")
	for key, value := range conf.properties {
		fmt.Printf("%v=%v\n", key, value)
	}
	fmt.Println("--------------------------------------------")
	aa()
}
