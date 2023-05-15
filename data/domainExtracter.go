package data

import (
	"golang.org/x/net/publicsuffix"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
)

func ExtractDomain(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}

// Function to extract parent domain
func GetParentDomain(u string) (string, error) {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	domain := parsedUrl.Hostname()
	parentDomain, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return "", err
	}

	return parentDomain, nil
}

// Function to write to file
func WriteToFile(parentDomain, domain string) {
	fileName := parentDomain + ".csv"
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	txtBody, _ := io.ReadAll(f)
	if strings.Contains(string(txtBody), domain) {
		return
	}
	if _, err := f.WriteString(domain + "\n"); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func Urls2Domains2List(urls []string) {
	for _, u := range urls {
		parentDomain, err := GetParentDomain(u)
		if err != nil {
			log.Fatal(err)
		}
		WriteToFile(parentDomain, u)
	}
}
