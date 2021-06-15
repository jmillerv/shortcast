package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const homeDir = "/home/pi" // TODO add if statement to get OS environment. If !pi then other conventions (win, linux, mac)
// instead of a constant this can be set as a default string in the main that uses the os.GetEnv("HOME") to get the home dir.


// installDeps runs updates and installs required programs
func installDeps() error {
	fmt.Println("Installing Dependencies...")
	fmt.Println("running apt update ")
	cmd := exec.Command("apt-get", "update")
	cmd = exec.Command("apt-get","upgrade", "-y")
	cmd = exec.Command("apt-get", "install", "nodejs", "npm", "hugo", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

//TODO use packr to bundle the appImage into the binary.

// moveHokus places the hokus app image to the user's desktop
func moveHokus() error {
	fmt.Println("moving hokus app image")
	err := os.Chdir(homeDir)
	if err != nil {
		log.Println("unable to change directory to", homeDir)
	}
	oldLocation := "hokus-cms_rpi.AppImage"
	newLocation := homeDir+"/Desktop/hokus-cms_rpi.AppImage"
	err = os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Println("unable rename", oldLocation, "to ", newLocation)
	}
	// TODO add switch for architecture
	cmd := exec.Command("cp", "hokus-cms_linux.AppImage ~/Desktop/hokus.AppImage") // Linux amd64
	//cmd := exec.Command("chmod", "+x", homeDir+"/Desktop/hokus-cms_rpi.AppImage") // arm7 architecture
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func setupHugo() error{
	fmt.Println("setting up hugo")
	err := os.Chdir(homeDir)
	if err != nil {
		log.Println("unable to change directory to", homeDir)
	}
	cmd := exec.Command("hugo", "new", "site", "shortcast")
	err = cmd.Run()
	if err != nil {
		log.Fatal("could not create new hugo site ", err)
	}
	err = os.Chdir(homeDir+"/shortcast/themes") // TODO less of this filepath should be hardcoded.
	if err != nil {
		log.Println("unable to change directory to "+ homeDir+"/shortcast/themes")
	}
	cmd = exec.Command("git","clone", "https://github.com/iCyris/hugo-theme-yuki")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// clean deletes unneeded temp files TODO implement
func clean() {
	panic("function not implemented")
	//fmt.Println("Cleaning...")
	//_ = os.RemoveAll("shortcast")
}


func main() {
	err := installDeps()
	if err != nil {
		log.Fatal("failed to install dependencies ", err)
	}
	err = moveHokus()
	if err != nil {
		log.Fatal("Unable to move hokus app image ", err)
	}
	err = setupHugo()
	if err != nil {
		log.Fatal("hugo setup failed", err)
	}
	fmt.Println("running hugo")
	err = os.Chdir(homeDir+"/shortcast")
	if err != nil {
		log.Println("unable to change directory to "+ homeDir+"/shortcast")
	}
	cmd := exec.Command("hugo", "server", "-t", "hugo-theme-yuki")
	err = cmd.Run()
	if err != nil {
		log.Fatal("failed to run hugo ", err)
	}
}