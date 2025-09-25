package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Error  error
	Status int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <файл_с_url>")
		os.Exit(1)
	}

	filename := os.Args[1]
	urls, err := readURLs(filename)
	if err != nil {
		fmt.Printf("Ошибка чтения файл %v\n", err)
		os.Exit(1)
	}

	results := make(chan Result, len(urls))

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			status, err := checkUrl(u)
			defer wg.Done()
			results <- Result{
				URL:    u,
				Error:  err,
				Status: status,
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	for res := range results {
		if res.Error != nil {
			fmt.Printf("❌ %s - Ошибка: %v\n", res.URL, res.Error)
		} else if res.Status == http.StatusOK {
			fmt.Printf("✅ %s - Статус: %d\n", res.URL, res.Status)
		} else {
			fmt.Printf("⚠️  %s - Статус: %d\n", res.URL, res.Status)
		}
	}
}

func checkUrl(url string) (int, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Head(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func readURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var urls []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := scanner.Text()
		if url != "" {
			urls = append(urls, url)
		}
	}

	return urls, scanner.Err()
}
