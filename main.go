package main

import (
	"log"
	"os"

	"github.com/empitsu0813/go-cli-example/rss"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "hatebuViewer"
	app.Usage = "はてなブックマーク人気エントリーのジャンルを選択肢して見れるよ！"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "select",
			Aliases: []string{"a"},
			Usage:   "select genre (all, general, social, economics, life, knwledge, it, fun, entertainment, game)",
			Action: func(c *cli.Context) error {
				selectGenre := c.Args().First()
				var URL string
				if c.NArg() > 0 {
					switch selectGenre {
					case "all":
						URL = "http://b.hatena.ne.jp/hotentry/all.rss"
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
				}

				if URL != "" {
					rss.RssReader(URL)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
