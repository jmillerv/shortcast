# Install Notes

This is just a page for general notes and problems I encountered while setting this up so that I am sure
to include the solutions when I do the write up for this. 

Install go hugo 

cd into directory 

hugo new site <site name> 

download a theme

cd themes git clone <github address>

hugo new page/

In order to get this working in an automated fashion the user may have to create their pages and posts as markdown and
then the installer will need to read that information somehow and create the structure using a combination of file path
parsing and the `hugo new` commands. 

This should be a go binary. 

https://presstige.io/p/Using-Go-instead-of-bash-for-scripts-6b51885c1f6940aeb40476000d0eb0fc

https://golang.org/pkg/os/exec/

For the demo, I might be able to put the site on netlify. 