package notifile

import (
	"errors"
	"os"
	"strings"
	"time"
)

// Data is to structure the file data
type Data struct {
	Channel  []string
	Headline string
	Message  string
}

// Create is to create a new notification file
// First we create a new file with a timestamp after that we check the channel & add the data
func Create(data Data, path string) error {

	if len(data.Channel) == 0 {
		return errors.New("unfortunately, no channels are specified. Therefore, the file cannot be created")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("unfortunately the directory does not exist")
	}

	name := "notifile-" + time.Now().Format("20060102150405") + ".txt"

	file, err := os.Create(path + "/" + name)
	if err != nil {
		return err
	}
	defer file.Close()

	channel := strings.Join(data.Channel, ",")

	content := "channel=" + channel + "\nheadline=" + data.Headline + "\nmessage=" + data.Message

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil

}
