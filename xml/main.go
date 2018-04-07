package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	// "time"
	// "io/ioutil"
	// "net/http"
)

// type Feed struct {
// 	rss           string   `xml:"rss"`
//  	XMLname       xml.Name `xml:"channel"`
// 	Title         string   `xml:"title"`
// 	Description   string   `xml:"description"`
// 	Link          string   `xml:"link"`
// 	Language      string   `xml:"language"`
// 	LastBuildDate string   `xml:"lastBuildDate"`
// 	Image         image    `xml:"image"`
// 	Items         []item   `xml:"item"`
// }

// type item struct {
// 	Title       string `xml:"title"`
// 	Link        string `xml:"link"`
// 	GUID        string `xml:"guid"`
// 	Comments    string `xml:"comments"`
// 	PubDate     string `xml:"pubDate"`
// 	Description string `xml:"description"`
// 	Category    string `xml:"category,attr"`
// }

// type image struct {
// 	Title  string `xml:"title"`
// 	Width  int    `xml:"windth"`
// 	Height int    `xml:"height"`
// 	Link   string `xml:"link"`
// 	URL    string `xml:"url"`
// }

func main() {
	// tr := &http.Transport{
	// 	MaxIdleConns:       10,
	// 	IdleConnTimeout:    30 * time.Second,
	// 	DisableCompression: true,
	// }
	// client := &http.Client{Transport: tr}
	// resp, err := client.Get("https://rss.packetstormsecurity.com/files/os/osx/")
	// if err != nil {
	// 	panic(err)
	// }

	parser := gofeed.NewParser()
	parsed, _ := parser.ParseURL("https://rss.packetstormsecurity.com/files/os/osx/")

	for _, val := range parsed.Items {
		fmt.Printf("%#v", val)
	}

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// rssExample := &Feed{}

	// err = xml.Unmarshal(body, rssExample)
	// data, err := ioutil.ReadFile("data.xml")
	// if err != nil {
	// 	panic(err)
	// }

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%#v", rssExample)
	// for _,v := range(rssExample.Items){
	//  	fmt.Printf("Seccurity issue: %s \nMore: %s\n\n", v.Description, v.Link)
	// }

}
