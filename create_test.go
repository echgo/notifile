// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package notifile_test

import (
	"fmt"
	"github.com/echgo/notifile"
	"os"
	"testing"
)

// TestCreate is to test the create function
func TestCreate(t *testing.T) {

	tests := []struct {
		path     string
		channels []string
		headline string
		message  string
	}{
		{
			path:     os.TempDir(),
			channels: []string{"gotify", "smtp"},
			headline: "Testing",
			message:  "This is the fist message about test corners.",
		},
		{
			path:     os.TempDir(),
			channels: []string{"smtp"},
			headline: "Testing",
			message:  "This is the second message about test corners.",
		},
		{
			path:     os.TempDir(),
			channels: []string{"trello"},
			headline: "Testing",
			message:  "This is the third message about test corners.",
		},
	}

	for _, value := range tests {

		data := notifile.Data{
			Channels: value.channels,
			Headline: value.headline,
			Message:  value.message,
		}

		fmt.Println(value.path)

		err := notifile.Create(data, value.path)
		if err != nil {
			t.Fatal(err)
		}

	}

}
