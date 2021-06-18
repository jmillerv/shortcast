package backend

import "fmt"

func StartShortcast() {
	// variadic function to pass commands maybe set theme from the command line?
	// desktop can just use the yaml settings
	fmt.Println("starting shortcast")
}

func StopShortcast() {
	fmt.Println("stopping shortcast")
}