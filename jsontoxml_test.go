package main

import (
	"errors"
	"testing"
)

var JSONTests = []struct {
	JsonInput     []byte // input
	XmlExpected   []byte // expected result
	ErrorExpected error
}{
	{[]byte("in-am-invalid-json"), []byte(""), errors.New("something")},
	{[]byte("[]"), []byte("<root></root>"), nil},
	{[]byte(`{"key1": "I am fat & happy","key2": "1 < 2","key3": "2 > 1","key4:": "i have have an apostrophe '"}`), []byte(`<root><key1>I am fat &amp; happy</key1><key2>1 &lt; 2</key2><key3>2 &gt; 1</key3><key4:>i have have an apostrophe &apos;</key4:></root>`), nil},
	{[]byte(`{"anElement": {"foo": "bar1"}}`), []byte(`<root><anElement><foo>bar1</foo></anElement></root>`), nil},
}

func TestConversionOfJSON(t *testing.T) {

	for tIndex, testValue := range JSONTests {
		testXml, testErr := convertJSONToXML(testValue.JsonInput)
		t.Log("Test Index " + string(tIndex))
		t.Log("Json Input: ", string(testValue.JsonInput))
		t.Log("Test XML ", string(testValue.XmlExpected), "out: ", string(testXml))
		t.Log("Test Error ", testErr)

		if string(testXml) == string(testValue.XmlExpected) {
			t.Log("Passed - ", string(testValue.JsonInput))

		} else {
			t.Error("convertJSONToXML: Failed testing")
		}

	}

}
