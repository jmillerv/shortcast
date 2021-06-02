# Shortcast

## Goals of the project
- [] Provide an user interface to configure a captive portal
- [] Provide an option for users to add content through a GUI 
- [] Lightweight static site generator bundled into the binary
- [] compiles down to a single binary for usage on ARM systems

## Technology Stack  

### Prototype Stack
The content is served using Hugo

The front-end CMS is courtesy the Hokus library. 

The scripts install and captive portal configuration scripts were written by me uisng Golang. 

### Final Product Static 
UI Framework is Fyne.io for configuration and setup
Captive portal is configured through bash scripts that are embedded in the binary
Lightweight static site generator is small enough to be included in the binary

