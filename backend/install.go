package install

import (
	"fmt"
	"os"
	"os/exec"
)

const homeDir = "/home/pi" // TODO add if statement to get OS environment. If !pi then other conventions (win, linux, mac)
// instead of a constant this can be set as a default string in the main that uses the os.GetEnv("HOME") to get the home dir.

// Dependencies runs updates and installs required programs
func Dependencies() error {
	fmt.Println("Installing Dependencies...")
	fmt.Println("running apt update ")
	cmd := exec.Command("apt-get", "update")
	cmd = exec.Command("apt-get","upgrade", "-y")
	cmd = exec.Command("apt-get", "install", "nodejs", "npm", "hugo","go", "git", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func DependenciesWithoutInternet() error {
	return nil
}




// clean deletes unneeded temp files TODO implement
func clean() {
	panic("function not implemented")
	//fmt.Println("Cleaning...")
	//_ = os.RemoveAll("shortcast")
}

// DEPRECATED
//func main() {
//	err := Dependencies()
//	if err != nil {
//		log.Fatal("failed to install dependencies ", err)
//	}
//	err = moveHokus()
//	if err != nil {
//		log.Fatal("Unable to move hokus app image ", err)
//	}
//	err = SetupHugo()
//	if err != nil {
//		log.Fatal("hugo setup failed", err)
//	}
//	fmt.Println("running hugo")
//	err = os.Chdir(homeDir+"/shortcast")
//	if err != nil {
//		log.Println("unable to change directory to "+ homeDir+"/shortcast")
//	}
//	cmd := exec.Command("hugo", "server", "-t", "hugo-theme-yuki")
//	err = cmd.Run()
//	if err != nil {
//		log.Fatal("failed to run hugo ", err)
//	}
//}