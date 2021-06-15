# Install
This script should install all of the dependencies required for the aota project.

Be sure that the const `homeDir` is set to your home directory and that the AppImage exists there.

# Compilation
If building in a development environment different from the target machine be sure to change your environment
variables to match the target machine.

Ex: `env GOOS=linux GOARCH=arm GOARM=5 go build`

Otherwise `go build` is fine