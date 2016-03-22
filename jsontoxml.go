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
	var inputData []byte
	if len(os.Args) > 1 {

		inputData = []byte(os.Args[1])

	} else {

		inputData, _ = ioutil.ReadAll(os.Stdin)
	}

	if len(inputData) > 0 {

		var parsedJson interface{}

		if err := json.Unmarshal(inputData, &parsedJson); err != nil {
			log.Println("invalid JSON found \n", inputData)
			os.Exit(1)
		}

		var xmlout []byte
		anyxml.XMLEscapeChars(true)

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
