package notifile

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

// Data is to structure the file data
type Data struct {
	Channels []string `json:"channels"`
	Headline string   `json:"headline"`
	Message  string   `json:"message"`
}

// Create is to create a new notification file
// First we create a new file with a timestamp after that we check the channel & add the data
func Create(data Data, path string) error {

	if len(data.Channels) == 0 {
		return errors.New("unfortunately, no channels are specified. Therefore, the file cannot be created")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("unfortunately the directory does not exist")
	}

	name := "notifile-" + time.Now().Format("20060102150405") + ".json"

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(path+"/"+name, content, 0644)
	if err != nil {
		return err
	}

	return nil

}
