package crawler

import (
	"DomainCrawler/data"
	"DomainCrawler/utils"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var (
	domainList = []string{}
	//bodyList = []string{}
	MaxDepth = 5
)

func filter(element string, list []string) bool { // true 为跳过下次爬虫任务 false为执行
	for _, itt := range list { // 如果已经爬过就不爬了
		//fmt.Printf("%s - %s\n", itt, urlStr)
		if itt == element {
			//fmt.Println(*crawlerList)
			return true
		}
	}
	if strings.Count(element, "amp") > 2 {
		return true
	}
	banedKeyword := []string{
		"jpg",
		"png",
		"mp4",
		"mp3",
		"jpeg",
		"gif",
		"exe",
		"apk",
	}
	for _, it := range banedKeyword {
		if strings.HasSuffix(element, it) {
			return true
		}
	}
	return false
}
func CrawlSingle(urlStr string, crawlerList *[]string, depth int) {
	if depth > MaxDepth { // 层数
		return
	}
	tmpDomain, err := data.ExtractDomain(urlStr)
	if err != nil {
	}
	if !utils.IsInList(tmpDomain, domainList) {
		domainList = append(domainList, tmpDomain)
		domainret, err := data.GetParentDomain(urlStr)
		if err != nil {
			return
		}
		data.WriteToFile(domainret, tmpDomain)
		utils.Printsuc(tmpDomain)
	}

	// 获取当前url的域名并写入到列表中

	if filter(urlStr, *crawlerList) {
		return
	}
	resp, err := http.Get(urlStr)
	if err != nil {
		return
	}
	*crawlerList = append(*crawlerList, urlStr) // 加入list
	fmt.Printf("* %s 当前深度 %v/%v\n", urlStr, depth, MaxDepth)
	defer func(Body io.ReadCloser) {

		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	//if strings.Contains(string(body), "37233343-2AA5-42AB-8F81-2D4259C32FF2") {
	//	os.Exit(0)
	//}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return
	}
	filiterHost := ""
	if strings.Contains(baseURL.Host, ":") { // 滤除端口的格式影响
		filiterHost = strings.Split(baseURL.Host, ":")[0]
	} else {
		filiterHost = baseURL.Host
	}
	//println(baseURL.Host)
	s := strings.Split(filiterHost, ".")
	length := len(s)
	parentDomain := s[length-2] + "." + s[length-1]
	re := regexp.MustCompile(`https?://[\w\-]+.` + parentDomain + `(/[\w\-\._~:/\?#\[\]@!\$&'\(\)\*\+,;=]*)?`)

	matches := re.FindAllString(string(body), -1)

	for _, match := range matches {
		if !filter(match, *crawlerList) {
			CrawlSingle(match, crawlerList, depth+1)
		}

	}

	z := html.NewTokenizer(strings.NewReader(string(body)))

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if tt == html.StartTagToken {
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				for _, a := range t.Attr {
					if a.Key == "href" {
						u, err := url.Parse(a.Val)
						if err != nil {
							break
						}
						nextURL := baseURL.ResolveReference(u)

						// Ignore if it goes to a different domain
						if baseURL.Host != nextURL.Host {
							continue
						}
						//fmt.Printf("* %s\n", nextURL.String())
						if !filter(nextURL.String(), *crawlerList) {
							CrawlSingle(nextURL.String(), crawlerList, depth+1)
						}

					}
				}
			}
		}
	}
}
