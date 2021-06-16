package backend

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	homeDirectory = "HOME"
	themeDirectory = "THEMES"
	defaultTheme  = "https://github.com/iCyris/hugo-theme-yuki"
)

func Hugo() error{
	fmt.Println("setting up hugo")
	err := os.Chdir(getHomeDirectory())
	if err != nil {
		log.Println("unable to change directory to", homeDirectory)
	}
	cmd := exec.Command("hugo", "new", "site", "shortcast")
	err = cmd.Run()
	if err != nil {
		log.Fatal("could not create new hugo site ", err)
	}
	err = os.Chdir(getThemeDirectory())
	if err != nil {
		log.Println("unable to change directory to "+ getThemeDirectory())
	}
	err = GetThemeFromGit(defaultTheme)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getHomeDirectory() string {
	return os.Getenv(homeDirectory)
}

func getThemeDirectory() string {
	path := os.Getenv(homeDirectory) + os.Getenv(themeDirectory)
	return path
}

func GetThemeFromGit(url string) error {
	cmd := exec.Command("git","clone", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}