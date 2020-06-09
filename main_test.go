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

func TestHandler(t *testing.T) {
	testData := []struct {
		base     interface{}
		override interface{}
		want     interface{}
	}{
		{
			Book{BID: "BID0"},
			Book{Title: "Title0"},
			Book{BID: "BID0", Title: "Title0"},
		},
		{
			Book{BID: "BID1"},
			Book{Chapters: []Chapter{{CID: "CID1"}}},
			Book{BID: "BID1", Chapters: []Chapter{{CID: "CID1"}}},
		},
		{
			Book{BID: "BID2"},
			Book{Author: Author{AID: "AID2"}},
			Book{BID: "BID2", Author: Author{AID: "AID2"}},
		},
		{
			Book{Chapters: []Chapter{{CID: "CID3"}}},
			Book{Author: Author{AID: "AID3"}},
			Book{Chapters: []Chapter{{CID: "CID3"}}, Author: Author{AID: "AID3"}},
		},
	}

	for index, d := range testData {
		got := Override(d.base, d.override)
		if !reflect.DeepEqual(d.want, got) {
			t.Errorf("TestCase-%v not match\n%v\n%v\n", index, d.want, got)
		}
	}
}

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}
