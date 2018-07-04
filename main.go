package main

import (
	"fmt"
	"log"
	"os"

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

	app.Before = func(c *cli.Context) error {
		fmt.Println("----Before----")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add task list",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		// c.NArg() 		引数の数を返す
		// c.Args() 	    引数を配列で返す
		// c.Args().Get(0)  引数の0番目を取得
		// c.Args()[0] 		上と同じ?

		fmt.Printf("c.GlobalFlagNames() : %+v\n", c.GlobalFlagNames())
		fmt.Printf("c.String() : %+v\n", c.String("lang"))

		return nil
	}

	app.After = func(c *cli.Context) error {
		fmt.Println("----After----")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
