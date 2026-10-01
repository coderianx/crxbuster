package main

import (
	"flag"
	"fmt"

	"github.com/coderianx/crxbuster/internal/commands"
	"github.com/coderianx/crxbuster/internal/helpers"
)

func main() {
	mode := flag.Int("mode", 0, "Mode (1 = URL scan)")
	wordlistPath := flag.String("w", "", "Wordlist path for scan")
	url := flag.String("u", "", "URL for scan")
	timeout := flag.Int("t", 10, "Timeout")
	showVersion := flag.Bool("v", false, "Print version")
	showHelp := flag.Bool("h", false, "Show help screen")

	flag.Usage = helpers.PrintHelp

	flag.Parse()

	if *showHelp {
		helpers.PrintHelp()
		return
	}

	if *showVersion {
		helpers.PrintVersion()
		return
	}

	switch *mode {
	case 1:
		fmt.Println("[INFO] Scan Starting...")
		commands.ScanURL(*url, *wordlistPath, *timeout)
	default:
		helpers.PrintHelp()
	}

}
