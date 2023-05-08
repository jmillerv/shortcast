# Shortcast
Shortcast is software intended to be run on a single board computer that presents a user-friendly way to configure a captive portal and use the onboard wifi radio to distribute content within a fixed geographic location. 

## Archived 
Development moved to [Codeberg](https://codeberg.org/jmillerv/shortcast)

## Goals of the project
- [ ] Provide an user interface to configure a captive portal
- [ ] Provide an option for users to add content through a GUI 
- [ ] Lightweight static site generator bundled into the binary
- [ ] Compiles down to a single binary for usage on ARM based single board computers

## Technology Stack  

### Prototype Stack
- The content is served using Hugo
- The front-end CMS is courtesy the Hokus library. 
- The scripts install and captive portal configuration scripts were written by me uisng Golang. 

### Final Product Stack 
- UI Framework is Fyne.io for configuration and setup
- Captive portal is configured through bash scripts that are embedded in the binary
- Lightweight static site generator is small enough to be included in the binary

### Status 
This project is in hiatus but not abandoned 
