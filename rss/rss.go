package rss

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/mmcdole/gofeed"
)

// 構造体
type HatenaItem struct {
	Title  string
	Link   string
	Hatebu int
}

// 構造体のスライス
type HatenaItems []HatenaItem

func (h HatenaItems) Len() int {
	return len(h)
}

func (h HatenaItems) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HatenaItems) Less(i, j int) bool {
	return h[j].Hatebu < h[i].Hatebu
}

func RssReader(URL) {
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL(URL)
	items := feed.Items
	var hatenaItems HatenaItems

	fmt.Println(feed.Title)

	for _, item := range items {
		extension := item.Extensions
		hatebu := extension["hatena"]["bookmarkcount"]
		i, _ := strconv.Atoi(hatebu[0].Value)

		var hatenaItem HatenaItem = HatenaItem{item.Title, item.Link, i}
		hatenaItems = append(hatenaItems, hatenaItem)
	}
	sort.Sort(hatenaItems)

	for _, hatenaItem := range hatenaItems {
		fmt.Printf("%d - %s - %s\n", hatenaItem.Hatebu, hatenaItem.Title, hatenaItem.Link)
	}
}
