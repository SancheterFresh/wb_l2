package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var siteURL string
	fmt.Printf("Enter URL: ")
	fmt.Scanf("%s\n", &siteURL)
	//siteURL := "https://web.ics.purdue.edu/~gchopra/class/public/pages/webdesign/05_simple.html"
	downloadSite(siteURL)

}

func downloadSite(urlStr string) {
	// Получение имени хоста и адреа
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		log.Println("Error parsing URL:", urlStr)
		return
	}
	hostname := urlObj.Hostname()
	path := urlObj.Path

	// Отправка запроса
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	// Дериктория для сохранения файлов

	basePath := hostname

	err = os.MkdirAll(basePath+path, 0755)
	if err != nil {
		log.Println("Error creating directory:", basePath+path)
		return
	}

	// Запиь ответа в файл
	filename := "index.html"
	if strings.HasSuffix(path, "/") {
		filename = "index.html"
	} else {
		filename = filepath.Base(path)
	}
	filepath := filepath.Join(basePath, path, filename)
	file, err := os.Create(filepath)
	if err != nil {
		log.Println("Error creating file:", filepath)
		return
	}
	defer file.Close()
	_, err = file.Write(respBytes)
	if err != nil {
		log.Println("Error writing file:", filepath)
		return
	}

	// Поис ссылок и скачивание
	links := findLinks(respBytes)
	for _, link := range links {
		if strings.HasPrefix(link, "http") && strings.Contains(link, hostname) {
			downloadSite(link)
		} else if strings.HasPrefix(link, "/") {
			downloadSite(urlStr + link)
		}
	}
}

func findLinks(htmlf []byte) []string {
	var links []string
	doc, err := html.Parse(bytes.NewReader(htmlf))
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return links
	}
	var finder func(*html.Node)
	finder = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			finder(child)
		}
	}
	finder(doc)
	return links
}
