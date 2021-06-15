package main

import (
	"log"
	"os"
	"os/exec"
)

// This script clones and builds Hokus based on the environment.

func cloneHokus() error {
	log.Println("cloning hokus")
	_ = os.Chdir(homeDir)
	cmd := exec.Command("git clone", "https://github.com/julianoappelklein/hokus.git")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func installHokus() error {
	log.Println("installing hokus")
	err := os.Chdir(homeDir+"/hokus")
	if err != nil {
		log.Fatal("failed to change directory ", err)
	}
	cmd := exec.Command("npm", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func buildHokus() error {
	log.Println("building hokus")
	err := os.Chdir(homeDir+"/hokus")
	if err != nil {
		log.Fatal("failed to change directory ", err)
	}
	cmd := exec.Command("npm", "run", "dist-linux")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
