package helpers

import "fmt"

var version string = "v0.1.0"

func PrintVersion() {
	fmt.Printf("CrxBuster %s\n", version)
}
