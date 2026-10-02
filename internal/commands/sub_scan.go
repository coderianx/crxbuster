package commands

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func checkHTTP(client *http.Client, host string) string {
	for _, scheme := range []string{"https", "http"} {
		url := scheme + "://" + host

		resp, err := client.Get(url)
		if err != nil {
			continue
		}

		resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 500 {
			return "OK"
		}
	}

	return "NO"
}

func ScanSubdomain(domain string, path string, timeout int) {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,

		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Wordlist açılamadı: %v\n", err)
		return
	}
	defer file.Close()

	var subs []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sub := strings.TrimSpace(scanner.Text())

		if sub == "" || strings.HasPrefix(sub, "#") {
			continue
		}

		subs = append(subs, sub)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Wordlist okunamadı: %v\n", err)
		return
	}

	var wg sync.WaitGroup

	for _, sub := range subs {
		wg.Add(1)

		go func(sub string) {
			defer wg.Done()

			host := sub + "." + domain

			ips, err := net.LookupHost(host)
			if err != nil {
				return
			}

			httpStatus := checkHTTP(client, host)

			fmt.Printf(
				"[+] %-40s DNS: OK  HTTP: %s  IP: %v\n",
				host,
				httpStatus,
				ips,
			)
		}(sub)
	}

	wg.Wait()
}
