## How to create an executable file

1. Go to the main package of your project
2. execute this command in your cmd: `go build`
3. You will see that an executable file is created `[filename].exe`
4. This executable file can be run on your local directory, but to run it as a path command from your command line anywhere please follow these steps
   1. Type `go list -f '{{.Target}}'` to get the target location of the executable file
   2. Follow the steps in the documentation to proceed furthur: https://go.dev/doc/tutorial/compile-install