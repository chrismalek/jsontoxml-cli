package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/clbanning/anyxml"
)

func main() {
	stdInbytes, _ := ioutil.ReadAll(os.Stdin)

	if len(stdInbytes) > 0 {
		var parsedJson interface{}

		if err := json.Unmarshal(stdInbytes, &parsedJson); err != nil {
			log.Println("invalid JSON found \n", string(stdInbytes))
			os.Exit(1)
		}

		var xmlout []byte

		xmlout, err2 := anyxml.Xml(parsedJson, "root")

		if err2 != nil {
			log.Println("Error converting to XML")
			os.Exit(1)
		}
		fmt.Print(string(xmlout))
		os.Exit(0)

	} else {
		os.Exit(1)
	}
}
