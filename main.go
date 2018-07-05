package main

import (
	"fmt"
	"log"
	"os"

	"github.com/empitsu0813/go-cli-example/rss"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "hatebu-viewer"
	app.Usage = "はてなブックマーク人気エントリーIT を見れるよ！"
	app.Version = "0.0.1"
	// app.HideHelp helpは引数として見られる

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "japanese",
			Usage: "言語選択",
		},
	}

	//	app.Before = func(c *cli.Context) error {
	//		fmt.Println("----Before----")
	//		return nil
	//	}

	app.Commands = []cli.Command{
		{
			Name:    "select",
			Aliases: []string{"a"},
			Usage:   "select genre",
			Action: func(c *cli.Context) error {
				selectGenre := c.Args().First()
				var URL string
				if selectGenre > 0 {
					switch selectGenre {
					case "general":
						URL = "http://b.hatena.ne.jp/hotentry/general.rss"
					case "social":
						URL = "http://b.hatena.ne.jp/hotentry/social.rss"
					case "economics":
						URL = "http://b.hatena.ne.jp/hotentry/economics.rss"
					case "life":
						URL = "http://b.hatena.ne.jp/hotentry/life.rss"
					case "knowledge":
						URL = "http://b.hatena.ne.jp/hotentry/knowledge.rss"
					case "it":
						URL = "http://b.hatena.ne.jp/hotentry/it.rss"
					case "fun":
						URL = "http://b.hatena.ne.jp/hotentry/fun.rss"
					case "entertainment":
						URL = "http://b.hatena.ne.jp/hotentry/entertainment.rss"
					case "game":
						URL = "http://b.hatena.ne.jp/hotentry/game.rss"
					}
				} else {
					URL = "http://b.hatena.ne.jp/hotentry/all.rss"
				}

				var feed string
				if URL != nil {

					rss.RssReader(URL)
				}

				fmt.Println(feed)
				return nil
			},
		},
	}

	//	app.Action = func(c *cli.Context) error {
	//		// c.NArg() 		引数の数を返す
	//		// c.Args() 	    引数を配列で返す
	//		// c.Args().Get(0)  引数の0番目を取得
	//		// c.Args()[0] 		上と同じ?

	//		fmt.Printf("c.GlobalFlagNames() : %+v\n", c.GlobalFlagNames())
	//		fmt.Printf("c.String() : %+v\n", c.String("lang"))

	//		return nil
	//	}

	//	app.After = func(c *cli.Context) error {
	//		fmt.Println("----After----")
	//		return nil
	//	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
