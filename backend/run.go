package backend

import (
	"fmt"
	"github.com/jmillerv/shortcast/backend/hugo"
)

func StartShortcast() {
	// variadic function to pass commands maybe set theme from the command line?
	// desktop can just use the yaml settings
	fmt.Println("starting shortcast")
	hugo.StartServer()
}

func StopShortcast() {
	fmt.Println("stopping shortcast")
	hugo.StopServer()
}