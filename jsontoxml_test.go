package main

import (
	"errors"
	"testing"
)

var JSONTests = []struct {
	JSONInput     []byte // input
	XMLExpected   []byte // expected result
	ErrorExpected error
}{
	{[]byte("in-am-invalid-json"), []byte(""), errors.New("something")},
	{[]byte("[]"), []byte("<root></root>"), nil},
	{[]byte(`{"key1": "I am fat & happy","key2": "1 < 2","key3": "2 > 1","key4:": "i have have an apostrophe '"}`), []byte(`<root><key1>I am fat &amp; happy</key1><key2>1 &lt; 2</key2><key3>2 &gt; 1</key3><key4:>i have have an apostrophe &apos;</key4:></root>`), nil},
	{[]byte(`{"anElement": {"foo": "bar1"}}`), []byte(`<root><anElement><foo>bar1</foo></anElement></root>`), nil},

	{[]byte(`{"null-element": null}`), []byte(`<root><null-element/></root>`), nil},
}

func TestConversionOfJSON(t *testing.T) {

	for _, testValue := range JSONTests {
		testXML, testErr := convertJSONToXML(testValue.JSONInput)

		if string(testXML) != string(testValue.XMLExpected) {

			t.Log("Json Input: ", string(testValue.JSONInput))
			t.Log("Expected XML ", string(testValue.XMLExpected), "Converted: ", string(testXML))
			t.Log("Test Error ", testErr)

			t.Error("convertJSONToXML: Failed testing")
		}

	}

}
