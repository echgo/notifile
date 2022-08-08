package notifile

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

// Set prefix, extension is to build the filename
const (
	prefix    = "notifile-"
	extension = ".json"
)

// Data is to structure the file data
type Data struct {
	Channels []string `json:"channels"`
	Headline string   `json:"headline"`
	Message  string   `json:"message"`
}

// Create is to create a new notification file
// first we create a new file with a timestamp
// after that we check the channel & add the data
func Create(data Data, path string) error {

	if data.Channels == nil {
		return errors.New("unfortunately, no channels are specified")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("unfortunately the directory does not exist")
	}

	name := prefix + time.Now().Format("20060102150405") + extension

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
