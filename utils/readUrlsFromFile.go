package utils

import (
	"bufio"
	"net/url"
	"os"
)

func ReadURLsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	var validURLs []string

	for scanner.Scan() {
		u := scanner.Text()
		if _, err := url.ParseRequestURI(u); err == nil {
			validURLs = append(validURLs, u)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return validURLs, nil
}
