package helpers

import "fmt"

// PrintHelp prints the CLI help screen.
func PrintHelp() {
	fmt.Printf(`CrxBuster %s - directory & path brute forcer

Usage:
  crxbuster -mode <mode> [options]

Modes:
  1    URL scan (brute force paths on a target URL)
  2    Subdomain scan (brute force subdomains of a target domain)

Options:
  -mode int    Scan mode, see Modes above (default 0)
  -u string    Target URL, e.g. https://example.com
  -w string    Path to the wordlist file
  -t int       HTTP timeout in seconds (default 10)
  -v           Print version
  -h           Show this help screen
  -modes       Show available scan modes

Examples:
  crxbuster -mode 1 -u https://example.com -w wordlist.txt
  crxbuster -mode 1 -u https://example.com -w wordlist.txt -t 5
  crxbuster -v
`, version)
}
