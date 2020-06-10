package testutil

import (
	"os"
	"reflect"
	"testing"
)

type Book struct {
	BID      string
	Title    string
	Chapters []Chapter
	Author   Author
}

type Chapter struct {
	CID   string
	Title string
}

type Author struct {
	AID  string
	Name string
}

type TestData struct {
	base     interface{}
	fields   map[string]interface{}
	want     interface{}
	testCase string
}

func TestHandler(t *testing.T) {
	testData := []TestData{
		{
			Book{BID: "BID0", Chapters: []Chapter{{CID: "CID4"}}, Author: Author{AID: "AID4"}},
			Fields{"BID": "", "Chapters": []Chapter{}, "Author": Author{}},
			Book{BID: "", Chapters: []Chapter{}, Author: Author{}},
			"Set zero values",
		},
		{
			Book{BID: "BID1"},
			Fields{"Title": "Title0"},
			Book{BID: "BID1", Title: "Title0"},
			"Set scalar value",
		},
		{
			Book{BID: "BID2"},
			Fields{"Chapters": []Chapter{{CID: "CID2"}}},
			Book{BID: "BID2", Chapters: []Chapter{{CID: "CID2"}}},
			"Set slice value",
		},
		{
			Book{BID: "BID3"},
			Fields{"Author": Author{AID: "AID3"}},
			Book{BID: "BID3", Author: Author{AID: "AID3"}},
			"Set struct value",
		},
	}

	for index, d := range testData {
		got := Override(d.base, d.fields)
		if !reflect.DeepEqual(d.want, got) {
			t.Errorf("TestCase-%v not match\n%v\n%v\n", index, d.want, got)
		}
	}
}

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}
