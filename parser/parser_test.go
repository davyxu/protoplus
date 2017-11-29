package parser

import (
	"encoding/json"
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

func TestJson(t *testing.T) {

	fs, err := ParseFile("test.pp")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	data, err := json.Marshal(fs)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(string(data))
}

func TestTest(t *testing.T) {

	a := struct {
		B int `json:"bb"`
	}{B: 1}

	data, _ := json.Marshal(&a)

	t.Log(string(data))

}
