package main

import (
	"fmt"

	elastic "gopkg.in/olivere/elastic.v2"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200/"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

	fmt.Println("connect es succ")

	for i := 0; i < 10; i++ {
		tweet := Tweet{User: "oliver", Message: "take five"}
		//链式操作
		_, err = client.Index().
			Index("twitter99").
			Type("tweet").
			Id(string(i)).
			BodyJson(tweet).
			Do()

		if err != nil {
			panic(err)
			return
		}

	}
	fmt.Println("insert succ")

}
