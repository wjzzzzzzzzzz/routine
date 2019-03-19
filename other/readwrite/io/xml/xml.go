package main

import (
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"log"
)

var (
	InputFile = `readwrite/io/xml/device.xml`
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type RootElement struct {
	XMLName        xml.Name       `xml:"UserInformation"`
	ResourceString []ChildElement `xml:"string"`
}
type ChildElement struct {
	XMLName    xml.Name
	StringName string `xml:"age,attr"`
	InnerText  string `xml:",innerxml"`
}

func main() {
	GolangXml, err := ioutil.ReadFile(InputFile) //我们把数据一次性读取到一个切片中。
	if err != nil {
		log.Fatal(err)
	}
	var result RootElement
	fmt.Printf("序列化之前GolangXml = [%v]\n", result)
	err = xml.Unmarshal(GolangXml, &result) //然后对数据进行反序列化。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("序列化之后GolangXml = [%v]\n", result)
	fmt.Println(result.ResourceString)
}
