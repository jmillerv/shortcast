package backend

import (
	"fmt"
	"os"
	"os/exec"
)

// InstallDependencies runs updates and installs required programs
func InstallDependencies() error {
	fmt.Println("Installing Dependencies...")
	fmt.Println("running apt update ")
	cmd := exec.Command("apt-get", "update")
	cmd = exec.Command("apt-get","upgrade", "-y")
	cmd = exec.Command("apt-get", "install", "go", "git", "hugo", "-y") // We may need nodejs and npm
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// InstallDependenciesWithoutInternet uses the needed embedded binaries these may not be as current as doing it over wifi
func InstallDependenciesWithoutInternet() error {
	panic("implement me")
}

// backupFlatFiles stores a local copy of all of the files that are modified by InstallCaptivePortal so that we can
// restore them just in case.
func backupFlatFiles() error {
	panic("implement me")
}

func InstallCaptivePortal() error {
	// run configCaptive.sh
	panic("implement me")
}

func RestoreFlatFiles() error {
	panic("implement me")
}