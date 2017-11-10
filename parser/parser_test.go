package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	fs, err := ParseFile("test.pp")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(fs.String())

	//v, _ := fileD.StructByName["PhoneNumber"]
	////f, _ := v.FieldByName["number"]
	//
	//tag, _ := v.MatchTag("agent")
	//fmt.Println("tag: ", tag)
	//
	//fmt.Println(fileD.String())
}
