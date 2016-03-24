# JSON to XML Command Line Application.


This a tool that will take JSON over STDIN and convert it to XML to STDOUT.




## Cross Compile

`GOOS=linux GOARCH=amd64 go build -o jsontoxml jsontoxml.go`



## Usage

You can pipe in a file to the application  like this:

`cat ~/Downloads/test.json | go run jsontoxml.go`


`./jsontoxml '{"foo":"bar"}'` 

returns

`<root><foo>bar</foo></root>`


