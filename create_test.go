package notifile_test

import (
	"github.com/echgo/notifile"
	"testing"
)

// TestCreate is to test the create function
func TestCreate(t *testing.T) {

	data := notifile.Data{
		Channels: nil,
		Headline: "Testing",
		Message:  "This is a message about test corners.",
	}

	err := notifile.Create(data, "/path/to/your/folder")
	if err != nil {
		t.Fatal(err)
	}

}
