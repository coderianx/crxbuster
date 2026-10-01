package commands

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func ScanURL(url string, path string, timeout int) {
	if len(path) == 0 {
		fmt.Println("Please select wordlist path")
		os.Exit(1)
	}

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("[ERROR] Open Wordlist file error")
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.TrimPrefix(line, "/")

		if line == "" {
			continue
		}

		target := strings.TrimRight(url, "/") + "/" + line

		resp, err := client.Get(target)

		if err != nil {
			fmt.Println("[ERROR]", target, err)

			continue
		}

		switch {
		case resp.StatusCode >= 200 && resp.StatusCode < 300:
			fmt.Printf("[+] %d %s\n", resp.StatusCode, target)

		case resp.StatusCode >= 300 && resp.StatusCode < 400:
			fmt.Printf("[>] %d %s\n", resp.StatusCode, target)

		case resp.StatusCode == 404:
			// 404 Not Found - yok say

		case resp.StatusCode >= 400 && resp.StatusCode < 500:
			fmt.Printf("[!] %d %s\n", resp.StatusCode, target)

		case resp.StatusCode >= 500:
			fmt.Printf("[X] %d %s\n", resp.StatusCode, target)
		}

		resp.Body.Close()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[ERROR] Reading Wordlist:", err)
	}
}
