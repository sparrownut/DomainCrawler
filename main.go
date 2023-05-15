package main

import (
	"DomainCrawler/crawler"
	"DomainCrawler/utils"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func run() {
	inputFileName := ""
	threadMax := 10
	app := &cli.App{
		Name:      "DomainCawler",
		Usage:     "内链爬虫 - 找到深度埋藏的资产 \n在输入文件中 一行一个想要爬取的目标资产的url\n爬取后会保存为%ParentDomain%.csv的格式 \nEG:\n--list.txt\n----https://m.example.com\n----https://www.example.com",
		UsageText: "./DC -f list.txt",
		Version:   "0.0.1",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "deepthMax", Aliases: []string{"d"}, Destination: &crawler.MaxDepth, Usage: "爬虫深度 默认5", Value: 5, Required: false},
			&cli.IntFlag{Name: "threadMax", Aliases: []string{"t"}, Destination: &threadMax, Usage: "线程数 默认10", Value: 10, Required: false},
			&cli.StringFlag{Name: "fileInput", Aliases: []string{"f"}, Destination: &inputFileName, Usage: "URLS 文件位置", Value: "list.txt", Required: true},
		}, // 配置信息
		HideHelpCommand: true,
		Action: func(c *cli.Context) error { // 运行位置
			var strList []string
			var tasks []utils.Task
			urlsList, err := utils.ReadURLsFromFile(inputFileName)
			if err != nil {
				return err
			}
			fmt.Println(urlsList)
			for _, it := range urlsList {
				it := it
				tasks = append(tasks, func() {
					crawler.CrawlSingle(it, &strList, 0) // 初始深度0
				})
			}
			utils.RunWithConcurrency(tasks, threadMax)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		utils.Printcritical("PROGRESS PANIC")
		os.Exit(0)
	}
}

func main() {
	run()
}
