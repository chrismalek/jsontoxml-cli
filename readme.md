# JSON to XML Command Line Application.


This a tool that will take JSON over STDIN and convert it to XML to STDOUT. This tool is written in Golang. You can install the golang language from https://golang.org and compile this to your operating system.


You can build it locally using `go build jsontoxml.go` which will build an executable. If you need to cross compile for a different platform that is really easy for golang see below. For example, you have a windows machine and need to compile a Linux executable, golang makes this possible.


## Cross Compile

For a better example of cross compilation in Go, see Dave Cheney's post at http://dave.cheney.net/2015/03/03/cross-compilation-just-got-a-whole-lot-better-in-go-1-5

Here is a command I use to compile from OSX to an 64 bit Linux executable. This will produce a single binary that can be deployed and it includes all dependencies. 

`GOOS=linux GOARCH=amd64 go build -o jsontoxml jsontoxml.go`

There is NO need to install anything else but the artifact produced. The `-o jsontoxml` controls the executable name. So if you want the executable to be called 'superduper-jsontoxml' you would just use that for the `-o` parameters.



## Usage

You can pipe in a file to the application  like this:

In this example we are using the `go run` command which compiles then runs.

`cat ~/Downloads/test.json | go run jsontoxml.go`

In this example we assume we compiled a binary to `jsontoxml` and call it with some parameters.

`./jsontoxml '{"foo":"bar"}'` 

returns

`<root><foo>bar</foo></root>`

