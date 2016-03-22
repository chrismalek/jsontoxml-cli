package main

import (
	"encoding/json"
	"errors"
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

		xmlout, err := convertJSONToXML(inputData)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(string(xmlout))
		os.Exit(0)

	} else {
		os.Exit(1)
	}

}

func convertJSONToXML(inputData []byte) (xmlout []byte, err error) {
	var parsedJson interface{}

	if err := json.Unmarshal(inputData, &parsedJson); err != nil {
		return []byte(""), errors.New(fmt.Sprintf("invalid JSON found %s\n", inputData))
	}

	anyxml.XMLEscapeChars(true)
	xmlout, err2 := anyxml.Xml(parsedJson, "root")
	fmt.Printf("convertJSONToXML-converted xml %s\n\n", string(xmlout))
	if err2 != nil {
		return []byte(""), errors.New("Error converting to XML")
	}

	return xmlout, nil

}
