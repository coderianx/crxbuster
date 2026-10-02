package helpers

import "fmt"

// PrintModes prints the available scan modes.
func PrintModes() {
	fmt.Printf(`CrxBuster %s - scan modes

Modes:
  1    URL scan        Brute force paths on a single target URL
  2    Subdomain scan  Brute force subdomains of a target domain

Usage:
  crxbuster -mode <mode> -u <target> -w <wordlist>

Examples:
  crxbuster -mode 1 -u https://example.com -w wordlist.txt
  crxbuster -mode 2 -u example.com -w subdomains.txt
`, version)
}
