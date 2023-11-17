package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlysevers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:"innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

func main() {
	//------------------解析xml内容--------------------
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error :%v\n", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error :%v\n", err)
		return
	}
	v := Recurlysevers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error :%v\n", err)
		return
	}
	fmt.Println(v)

	//------------------输出xml内容--------------------
	v2 := &Servers{Version: "1"}
	v2.Svs = append(v2.Svs, server{ServerName: "shanghai", ServerIP: "127.0.0.3"})
	v2.Svs = append(v2.Svs, server{ServerName: "beijing", ServerIP: "127.0.0.4"})
	output, err := xml.MarshalIndent(v2, " ", " ")
	if err != nil {
		fmt.Printf("error : %v\n", err)

	}
	os.Exit(1)
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)

}
