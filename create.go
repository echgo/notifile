// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package notifile is used to generate a json file that
// is prepared in such a way that it can be read by echgo.
package notifile

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"time"
)

// Set prefix, char, extension is to build the filename.
const (
	prefix    = "notifile-"
	char      = "abcdefghijklmnopqrstuvwxyz1234567890"
	extension = ".json"
)

// Data is to structure the file data.
type Data struct {
	Channels []string `json:"channels"`
	Headline string   `json:"headline"`
	Message  string   `json:"message"`
}

// Create is to create a new notification file.
// First we create a new file with a timestamp
// after that we check the channel & add the data.
func Create(data Data, path string) error {

	if data.Channels == nil {
		return errors.New("unfortunately, no channels are specified")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("unfortunately the directory does not exist")
	}

	name := prefix
	for i := 0; i < 56; i++ {
		rand.Seed(time.Now().UnixNano())
		name += string(char[rand.Intn(len(char))])
	}
	name += extension

	content, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(path+"/"+name, content, 0644)
	if err != nil {
		return err
	}

	return nil

}
